package go_config

import (
	"fmt"
	"reflect"
	"strconv"
)

// Variable Routines
type Variable struct {
	// field name
	Name string
	// default value
	Def reflect.Value
	// field description
	Description string
	// set value
	Set func(x reflect.Value)
	// field tags
	Tag reflect.StructTag
	// field type
	Type reflect.Type
}

func (v Variable) String() string {
	return fmt.Sprintf("%v[%v] %v", v.Name, v.Type.Kind(), v.Description)
}

func (v *Variable) set(value interface{}) {
	if v.Type.Kind() == reflect.Struct {
		return
	}
	if v.Type.Kind() == reflect.Slice {
		slice, ok := value.([]interface{})
		if !ok {
			return
		}

		switch v.Type.Elem().Kind() {
		case reflect.String:
			resp := make([]string, len(slice))
			for i, v := range slice {
				switch v.(type) {
				case string:
					resp[i] = v.(string)
				default:
					resp[i] = fmt.Sprintf("%v", v)
				}
			}
			v.Set(reflect.ValueOf(resp))
		case reflect.Int:
			resp := make([]int, len(slice))
			for i, v := range slice {
				switch v.(type) {
				case int:
					resp[i] = v.(int)
				default:
					intVal, err := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
					if err != nil {
						continue
					}
					resp[i] = int(intVal)
				}
			}
			v.Set(reflect.ValueOf(resp))
		case reflect.Float64:
			resp := make([]float64, len(slice))
			for i, v := range slice {
				switch v.(type) {
				case float64:
					resp[i] = v.(float64)
				default:
					intVal, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
					if err != nil {
						continue
					}
					resp[i] = float64(intVal)
				}
			}
			v.Set(reflect.ValueOf(resp))
		case reflect.Uint:
			resp := make([]uint, len(slice))
			for i, v := range slice {
				switch v.(type) {
				case uint:
					resp[i] = v.(uint)
				default:
					intVal, err := strconv.ParseUint(fmt.Sprintf("%v", v), 10, 64)
					if err != nil {
						continue
					}
					resp[i] = uint(intVal)
				}
			}
			v.Set(reflect.ValueOf(resp))
		case reflect.Bool:
			resp := make([]bool, len(slice))
			for i, v := range slice {
				switch v.(type) {
				case bool:
					resp[i] = v.(bool)
				default:
					intVal, err := strconv.ParseBool(fmt.Sprintf("%v", v))
					if err != nil {
						continue
					}
					resp[i] = intVal
				}
			}
			v.Set(reflect.ValueOf(resp))
		}
		return
	}
	v.Set(reflect.ValueOf(value).Convert(v.Type))
}
