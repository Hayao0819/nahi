package rfutils

import "reflect"

func GetFields(i interface{}) map[string]interface{} {
	fields := make(map[string]interface{})
	v := reflect.ValueOf(i)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fields[t.Field(i).Name] = v.Field(i).Interface()
	}
	return fields
}
