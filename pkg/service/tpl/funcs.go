package tpl

import (
	"errors"
	"strings"
	"text/template"
)

func Join(list []string, sep string) string {
	return strings.Join(list, sep)
}

func Dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func InStringSlice(list []string, elem string) bool {
	for _, a := range list {
		if a == elem {
			return true
		}
	}
	return false
}

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"join":          Join,
		"dict":          Dict,
		"inStringSlice": InStringSlice,
	}
}
