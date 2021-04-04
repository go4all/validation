package rules

import (
	"errors"
	"fmt"
	"github.com/go4all/validaiton/utils"
	"reflect"
	"strconv"
)

type Max struct {}

func (rule Max) GetError(kind reflect.Kind, field string, args []string) string {
	switch kind {
	case reflect.Int, reflect.Float64:
		return fmt.Sprintf("%s should not exceed %s", field, args[0])
	case reflect.String:
		return fmt.Sprintf("%s can have max %s characters", field, args[0])
	case reflect.Map, reflect.Slice:
		return fmt.Sprintf("%s can have max %s values", field, args[0])
	}
	return fmt.Sprintf("Max error not defined for %s type", kind)
}

func (rule Max) Check(field string, value interface{}, args []string, message string) error {
	if len(args) == 0 {
		return errors.New("missing args for max validation")
	}
	// Don't check nil value
	if value == nil {
		return nil
	}

	result, convErr := strconv.ParseInt(args[0], 10, 32)

	if convErr != nil {
		return errors.New("invalid args for max validation")
	}

	max := int(result)

	valid := true

	kind := reflect.TypeOf(value).Kind()

	err := utils.ErrorMsg(message, rule.GetError(kind, field, args))

	switch kind {
	case reflect.Int:
		valid = rule.checkInt(value, max)
	case reflect.Float64:
		valid = rule.checkFloat64(value, max)
	case reflect.String:
		valid = rule.checkString(value, max)
	case reflect.Map:
		valid = rule.checkMap(value, max)
	case reflect.Slice:
		valid = rule.checkSlice(value, max)
	default:
		return errors.New("invalid type for max validation")
	}

	if !valid {
		return err
	}

	return nil
}

func (rule Max) checkInt(value interface{}, max int) bool {
	if reflect.TypeOf(value).Kind() == reflect.Int {
		data := value.(int)
		return data <= max
	}
	return false
}

func (rule Max) checkFloat64(value interface{}, max int) bool {
	if reflect.TypeOf(value).Kind() == reflect.Float64 {
		data := value.(float64)
		return data <= float64(max)
	}
	return false
}

func (rule Max) checkString(value interface{}, max int) bool {
	if reflect.TypeOf(value).Kind() == reflect.String {
		data := value.(string)
		return len(data) <= max
	}
	return false
}

func (rule Max) checkSlice(value interface{}, max int) bool {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		data := reflect.ValueOf(value)
		return data.Len() <= max
	}
	return false
}

func (rule Max) checkMap(value interface{}, max int) bool {
	if reflect.TypeOf(value).Kind() == reflect.Map {
		data := reflect.ValueOf(value)
		return data.Len() <= max
	}
	return false
}
