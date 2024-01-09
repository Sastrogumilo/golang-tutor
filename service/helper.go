package services

import (
	"reflect"
)

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Convert struct to map
	structValue := reflect.ValueOf(obj).Elem()
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structType.Field(i).Tag.Get("json")
		result[fieldName] = structValue.Field(i).Interface()
	}

	return result
}
