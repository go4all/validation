package rules

import (
	"fmt"
	"github.com/go4all/validation/types"
	"github.com/go4all/validation/utils"
	"reflect"
)

type Required struct {}

func (rule Required) GetError(kind reflect.Kind, field string, args []string) string {
	return fmt.Sprintf("%s is required", field)
}

func (rule Required) Check(config types.RuleConfig) error {
	err := utils.ErrorMsg(config.ErrMsg, rule.GetError(0, config.FieldName, config.RuleArgs))

	if config.FieldValue == nil {
		return err
	}

	valid := true

	kind := reflect.TypeOf(config.FieldValue).Kind()

	switch kind {
	case reflect.String:
		valid = rule.checkString(config.FieldValue)
	case reflect.Map:
		valid = rule.checkMap(config.FieldValue)
	case reflect.Slice:
		valid = rule.checkSlice(config.FieldValue)
	case reflect.Struct:
		valid = rule.checkStruct(config.FieldValue)
	case reflect.Ptr:
		valid = rule.checkPointer(config.FieldValue)
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
