package validation

import (
	"github.com/go4all/validation/utils"
	"reflect"
)

func GetValues(data interface{}) map[string]interface{} {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	values := make(map[string]interface{})

	if t.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("json")

		if tag == "" {
			continue
		}

		if field.Type.Kind() == reflect.Struct {
			result := GetValues(value.Interface())
			for key, val := range result {
				values[tag+"."+key] = val
			}
		} else {
			values[tag] = value.Interface()
		}
	}

	return values
}

// ValueByFieldPath will search provided values map and return value which key is matching with fieldPath
func ValueByFieldPath(values map[string]interface{}, fieldPath string) (string, interface{}) {
	value, ok := values[fieldPath]

	if !ok {
		return fieldPath, nil
	}

	name, ok  := utils.GetFieldNameFromPath(fieldPath)

	if ok {
		return name, value
	}

	return fieldPath, nil
}
