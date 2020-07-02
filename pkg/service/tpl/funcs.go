package tpl

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	strings "strings"
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

func GetSchemaLink(ref string) string {
	idx := strings.LastIndex(ref, "/")
	refName := ref[idx+1:]
	linkPath := ref[:idx]
	linkPath = strings.ReplaceAll(linkPath, "#", "../..")
	return fmt.Sprintf("[%s](%s/#%s)", strcase.ToCamel(refName), linkPath, strings.ToLower(refName))
}

func ToHTMLNewLines(str string) string {
	return strings.ReplaceAll(str, "\n", "<br>")
}

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"join":           Join,
		"dict":           Dict,
		"inStringSlice":  InStringSlice,
		"schemaLink":     GetSchemaLink,
		"toCamel":        strcase.ToCamel,
		"toHTMLNewLines": ToHTMLNewLines,
		"lower":          strings.ToLower,
		"replace":        strings.ReplaceAll,
	}
}
