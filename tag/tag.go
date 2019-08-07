package tag

import (
  "reflect"
)

type Request struct {
  Data        interface{}
  Field       string
  Tag         string
  ShowHidden  bool
}

type Tag struct {
  Field     string
  Value     string
}

func GetAll(request *Request) ([]*Tag, error) {
  result := []*Tag{}
  if request.Data == nil {return result, NoDataGivenError}
  // if request.Tag == "" {return result, NoTagGivenError}
  t := reflect.TypeOf(request.Data)

  if request.Field != "" && request.Tag != "" {
    field, ok := t.FieldByName(request.Field)
    if !ok {return result, NoSuchFieldError}
    value, ok := field.Tag.Lookup(request.Tag)
    if ok {result = append(result, &Tag{Field: request.Field, Value: value,})}
  } else if request.Tag != "" {
    count := 0
    for i := 0; i < t.NumField(); i++ {
      field := t.Field(i)
      if Anonymous(field) && !request.ShowHidden {continue}
      value, ok := field.Tag.Lookup(request.Tag)
      if ok {
        result = append(result, &Tag{Field: field.Name, Value: value,})
        count++
      }
    }
  }
  return result, nil
}
