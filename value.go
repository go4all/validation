package validation

import (
	"reflect"
	"strings"
)

// ValueByJsonTag will check for json tag on struct fields and return value of matching tag with provided name as argument
func ValueByJsonTag(data interface{}, fieldPath string) (string, interface{}) {
	path := strings.Split(fieldPath, ".")

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	if t.Kind() != reflect.Struct {
		return path[0], nil
	}

	var value interface{} = nil

	for i := 0; i < t.NumField(); i ++ {
		fieldTag := t.Field(i).Tag.Get("json")
		if path[0] == fieldTag {
			value = v.Field(i).Interface()
		}
	}

	if value == nil || len(path) == 1 {
		return path[0], value
	} else {
		return ValueByJsonTag(value, strings.Join(path[1:], "."))
	}
}



