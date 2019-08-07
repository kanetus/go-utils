package tag

import (
  "reflect"
  "unicode"
)

func Anonymous(field reflect.StructField) bool {
  return unicode.IsLower(rune(field.Name[0]))
}
