package utils

import (
	"reflect"
	"strconv"
	"strings"
)

func ParseSlice(src, delimiter string, kind reflect.Kind) ([]interface{}, error) {
	if len(src) == 0 {
		return []interface{}{}, nil
	}

	values := strings.Split(src, delimiter)
	slice := make([]interface{}, len(values))

	for i, v := range values {
		switch kind {
		case reflect.String:
			slice[i] = v
		case reflect.Int:
			val, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				continue
			}
			slice[i] = val
		case reflect.Float64:
			val, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue
			}
			slice[i] = val
		case reflect.Uint:
			val, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				continue
			}
			slice[i] = val
		case reflect.Bool:
			val, err := strconv.ParseBool(v)
			if err != nil {
				continue
			}
			slice[i] = val
		}
		slice = append(slice)
	}

	return slice, nil
}
