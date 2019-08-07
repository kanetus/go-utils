package tag_test

import (
  "github.com/kanetus/go-utils/tag"
  "github.com/kanetus/go-utils/suite"
  "testing"

)

type TestStruct struct {
  FieldA      string    `val:"A", anon:"false"`
  FieldB      string    `val:"B", anon:"false"`
  FieldC      string    `val:"C", anon:"false"`
  Field       string    `anon:"false"`
  anonFieldD  string    `val:"D", anon:"true"`
}

func TestGetAll(t *testing.T) {
  tagTests := []*suite.Test{
    &suite.Test{
      Request: suite.Params(&tag.Request{
        Data: TestStruct{}, Tag: "val", Field: "FieldA",}),
      Response: suite.Params([]*tag.Tag{&tag.Tag{Field: "FieldA", Value: "A"}}, nil),
    },
    &suite.Test{
      Request: suite.Params(&tag.Request{Data: TestStruct{}, Tag: "val", ShowHidden: true}),
      Response: suite.Params([]*tag.Tag{
        &tag.Tag{Field: "FieldA", Value: "A"},
        &tag.Tag{Field: "FieldB", Value: "B"},
        &tag.Tag{Field: "FieldC", Value: "C"},
        &tag.Tag{Field: "anonFieldD", Value: "D"},
      }, nil),
    },
    &suite.Test{
      Request: suite.Params(&tag.Request{Data: TestStruct{}, Tag: "val"}),
      Response: suite.Params([]*tag.Tag{
        &tag.Tag{Field: "FieldA", Value: "A"},
        &tag.Tag{Field: "FieldB", Value: "B"},
        &tag.Tag{Field: "FieldC", Value: "C"},
      }, nil),
    },
  }

  suite.Suite(t, &suite.Table{
    Name: "Get Specified Tag",
    Tests: tagTests,
    Run: suite.Assert(tag.GetAll),
  })
}
