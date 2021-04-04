package rules

import (
	"errors"
	"fmt"
	"github.com/go4all/validaiton/utils"
	"reflect"
	"strconv"
)

type Min struct {}

func (rule Min) GetError(kind reflect.Kind, field string, args []string) string {
	switch kind {
	case reflect.Int, reflect.Float64:
		return fmt.Sprintf("%s should be atleast %s", field, args[0])
	case reflect.String:
		return fmt.Sprintf("%s should have atleast %s characters", field, args[0])
	case reflect.Map, reflect.Slice:
		return fmt.Sprintf("%s should have atleast %s values", field, args[0])
	}
	return fmt.Sprintf("Min error not defined for %s type", kind)
}

func (rule Min) Check(field string, value interface{}, args []string, message string) error {
	if len(args) == 0 {
		return errors.New("missing args for min validation")
	}

	if value == nil {
		return nil
	}

	result, convErr := strconv.ParseInt(args[0], 10, 32)

	if convErr != nil {
		return errors.New("invalid args for min validation")
	}

	min := int(result)

	valid := true

	kind := reflect.TypeOf(value).Kind()

	err := utils.ErrorMsg(message, rule.GetError(kind, field, args))

	switch kind {
	case reflect.Int:
		valid = rule.checkInt(value, min)
	case reflect.Float64:
		valid = rule.checkFloat64(value, min)
	case reflect.String:
		valid = rule.checkString(value, min)
	case reflect.Map:
		valid = rule.checkMap(value, min)
	case reflect.Slice:
		valid = rule.checkSlice(value, min)
	default:
		return errors.New("invalid type for min validation")
	}

	if !valid {
		return err
	}

	return nil
}

func (rule Min) checkInt(value interface{}, min int) bool {
	if reflect.TypeOf(value).Kind() == reflect.Int {
		data := value.(int)
		return data >= min
	}
	return false
}

func (rule Min) checkFloat64(value interface{}, min int) bool {
	if reflect.TypeOf(value).Kind() == reflect.Float64 {
		data := value.(float64)
		return data >= float64(min)
	}
	return false
}

func (rule Min) checkString(value interface{}, min int) bool {
	if reflect.TypeOf(value).Kind() == reflect.String {
		data := value.(string)
		return len(data) >= min
	}
	return false
}

func (rule Min) checkSlice(value interface{}, min int) bool {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		data := reflect.ValueOf(value)
		return data.Len() >= min
	}
	return false
}

func (rule Min) checkMap(value interface{}, min int) bool {
	if reflect.TypeOf(value).Kind() == reflect.Map {
		data := reflect.ValueOf(value)
		return data.Len() >= min
	}
	return false
}
