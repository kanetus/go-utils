package suite

import (
  "testing"
  "reflect"
  "fmt"
  "github.com/kylelemons/godebug/pretty"
)

type AnyType  int
type RunTest  func(*testing.T, *Test, int)

const (
  Any AnyType = 0
)

type Test struct {
  Name      string
  Caller    interface{}
  Request   []interface{}
  Response  []interface{}
}

type Table struct {
  Name      string
  Tests     []*Test
  Run       RunTest
}

func Suite(t *testing.T, table *Table) {
  t.Run(table.Name, func(t *testing.T) {
    for i, test := range table.Tests {table.Run(t, test, i)}
  })
}

func Assert(f interface{}) RunTest {
  method := reflect.ValueOf(f)
  return func(t *testing.T, test *Test, count int) {
    defer func() {
      if err := recover(); err != nil {
        t.Error(formatFailure(test.Name, count, err))
      }
    } ()
    response := method.Call(valmap(test.Request))
    assert(t, test.Name, count, response, test.Response)
  }
}

func Params(data ...interface{}) []interface{} {
  result := make([]interface{}, len(data))
  for i, d := range data {
    result[i] = d
  }
  return result
}

func assert(t *testing.T, name string, count int, actual []reflect.Value, expected []interface{}) {
  if len(actual) != len(expected) {
    left, right := showValues(actual, expected)
    t.Error(formatResult(left, right, name, count,
      fmt.Sprintf("Output size mismatch: Got %d, expected %d",
        len(actual), len(expected))))
  } else {
    left, right, match := compare(actual, expected)
    if !match {
      t.Error(formatResult(left, right, name, count, ""))
    }
  }
}

func valmap(data []interface{}) []reflect.Value {
  result := []reflect.Value{}
  for _, d := range data {
    result = append(result, reflect.ValueOf(d))
  }
  return result
}

func compare(actual []reflect.Value, expected []interface{}) (string, string, bool) {
  left := ""
  right := ""
  match := true
  for i, val := range actual {
    rawVal := val.Interface()
    indicator := ""
    if expected[i] != Any && !reflect.DeepEqual(rawVal, expected[i]) {
      indicator = "__"
      match = false
    }
    left = format(left, rawVal, indicator)
    right = format(right, expected[i], indicator)
  }
  return left, right, match
}

func showValues(actual []reflect.Value, expected []interface{}) (string, string) {
  left := ""
  right := ""
  for _, act := range actual {left = format(left, act.Interface(), "")}
  for _, exp := range expected {right = format(right, exp, "")}
  return left, right
}

func format(previous string, val interface{}, indicator string) string {
  cmpt := &pretty.Config{Compact: true}
  return fmt.Sprintf("%s\n\t\t%s%s :: %T%s,", previous, indicator, cmpt.Sprint(val), val, indicator)
}

func formatResult(left string, right string, name string, count int, err string) string {
  failure := formatFailure(name, count, err)
  return fmt.Sprintf("%s%s\n!=%s", failure, left, right)
}

func formatFailure(name string, count int, err interface{}) string {
  if err != "" {
    err = fmt.Sprintf(" :: %s", err)
  }
  return fmt.Sprintf("[%d] %s%s", count, name, err)
}
