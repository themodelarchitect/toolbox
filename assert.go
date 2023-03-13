package tools

import (
	"reflect"
)

func Equals(exp, act interface{}) bool {
	return reflect.DeepEqual(exp, act)
}
