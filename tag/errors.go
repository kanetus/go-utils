package tag

import "errors"

var NoSuchFieldError  error = errors.New("No such field in structure")
var NoTagFoundError   error = errors.New("No such tag found in structure")
var NoDataGivenError  error = errors.New("No data given")
var NoTagGivenError   error = errors.New("No tag given")
