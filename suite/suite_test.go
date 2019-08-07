package suite_test

import (
  "github.com/kanetus/go-utils/suite"
  "testing"
  "errors"
)


var DivideByZeroError error = errors.New("Divide by Zero")

func divide(x int, y int) (int, error) {
  if y == 0 {return 0, DivideByZeroError}
  if x == 3 {return 4, nil}
  return x/y, nil
}

func TestExample(t *testing.T) {
  divisionTests := []*suite.Test{
    &suite.Test{Request: suite.Params(4,2), Response: suite.Params(2,nil),},
    &suite.Test{Request: suite.Params(10,0), Response: suite.Params(suite.Any, DivideByZeroError),},
    &suite.Test{Request: suite.Params(3,1), Response: suite.Params(3, nil), Name: "Should fail"},
    &suite.Test{Request: suite.Params(2,2), Response: suite.Params(1), Name: "Incorrect output size"},
    &suite.Test{Request: suite.Params(2), Response: suite.Params(1, nil), Name: "Incorrect input size"},
  }

  suite.Suite(t, &suite.Table{
    Name: "Division Tests",
    Tests: divisionTests,
    Run: suite.Assert(divide),
  })
}
