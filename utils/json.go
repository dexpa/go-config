package utils

import (
	"encoding/json"
	"strings"
)

func JsonParse(data []byte) (map[string]interface{}, error) {
	in := map[string]interface{}{}
	out := map[string]interface{}{}

	err := json.Unmarshal(data, &in)
	if err != nil {
		return out, err
	}

	for f, v := range in {
		processJSON(v, name(f, ""), out)
	}
	return out, nil
}

func processJSON(v interface{}, field string, out map[string]interface{}) {
	switch vv := v.(type) {
	case map[string]interface{}:
		for f, v := range vv {
			processJSON(v, name(f, field), out)
		}
	case string:
		out[field] = vv
	case bool:
		out[field] = bool(vv)
	case int:
		out[field] = int(vv)
	case float64:
		out[field] = float64(vv)
	default:
		// todo processing custom registered types
	}
}

func name(name string, parent string) string {
	if len(parent) == 0 {
		return strings.ToLower(name)
	}
	return strings.ToLower(parent + "." + name)
}
