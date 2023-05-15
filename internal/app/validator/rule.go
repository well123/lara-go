package validator

import (
	"fmt"
	"reflect"
	"strings"
)

type Errors struct{}

func (*Errors) Required(field, param string) string {
	return fmt.Sprintf("%s字段不能为空", field)
}

func (*Errors) Min(field, param string) string {
	return fmt.Sprintf("%s字段不能小于%s", field, param)
}

func InitValidatorErrors() map[string]func(field, param string) string {
	var errors *Errors
	errorMsgs := make(map[string]func(field, param string) string)

	es := reflect.TypeOf(errors)
	for i := 0; i < es.NumMethod(); i++ {
		tempMethod := es.Method(i)
		errorMsgs[strings.ToLower(tempMethod.Name)] = func(field, param string) string {
			return tempMethod.Func.Call([]reflect.Value{reflect.ValueOf(errors), reflect.ValueOf(field), reflect.ValueOf(param)})[0].String()
		}
	}

	return errorMsgs
}
