package rules

import (
	"fmt"
	"github.com/go4all/validaiton/utils"
	"reflect"
)

type Required struct {}

func (rule Required) GetError(kind reflect.Kind, field string, args []string) string {
	return fmt.Sprintf("%s is required", field)
}

func (rule Required) Check(field string, value interface{}, args []string, message string) error {
	err := utils.ErrorMsg(message, rule.GetError(0, field, args))

	if value == nil {
		return err
	}

	valid := true

	kind := reflect.TypeOf(value).Kind()

	switch kind {
	case reflect.String:
		valid = rule.checkString(value)
	case reflect.Map:
		valid = rule.checkMap(value)
	case reflect.Slice:
		valid = rule.checkSlice(value)
	case reflect.Struct:
		valid = rule.checkStruct(value)
	case reflect.Ptr:
		valid = rule.checkPointer(value)
	}

	if !valid {
		return err
	}

	return nil
}

func (rule Required) checkString(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.String {
		data := value.(string)
		return data != ""
	}
	return false
}

func (rule Required) checkSlice(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		data := reflect.ValueOf(value)
		return data.Len() > 0
	}
	return false
}

func (rule Required) checkMap(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Map {
		data := reflect.ValueOf(value)
		return data.Len() > 0
	}
	return false
}

func (rule Required) checkStruct(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Struct {
		data := reflect.ValueOf(value)
		return data.NumField() > 0
	}
	return false
}

func (rule Required) checkPointer(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		data := reflect.ValueOf(value)
		return !data.IsNil()
	}
	return false
}
