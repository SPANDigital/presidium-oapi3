package tpl

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var referenceURL string

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

func GetRefRootSchema(ref string) string {
	refSplit := strings.Split(ref, "schemas/")
	if len(refSplit) > 0 {
		refSplit = strings.Split(refSplit[1], "/")
	}
	return refSplit[0]
}

func GetSchemaLink(ref string) string {
	idx := strings.LastIndex(ref, "/")
	refName := ref[idx+1:]
	linkPath := ref[:idx]
	linkPath = strings.ReplaceAll(linkPath, "#", fmt.Sprintf("/%s", referenceURL))
	return fmt.Sprintf("[%s]({{%%baseurl%%}}/%s/#%s)", strcase.ToCamel(refName), linkPath, Slugify(refName))
}

func ToHTMLNewLines(str string) string {
	replacement := "<br>"
	replacer := strings.NewReplacer(
		"\n", replacement,
		"\\n", replacement,
		"\\\n", replacement,
	)
	return replacer.Replace(str)
}

func Sum(int1, int2 int) int {
	return int1 + int2
}

func Slugify(s string) string {
	s = strcase.ToKebab(s)
	var nonWordRe = regexp.MustCompile(`(?m)(\W|_)+`)
	slug := nonWordRe.ReplaceAllString(s, "-")
	return strings.Trim(slug, "-")
}

// BreakLine replaces new lines with <br>
func BreakLine(s string) string {
	var newLine = regexp.MustCompile(`\n`)
	return newLine.ReplaceAllString(s, "<br>")
}

func Default(def interface{}, value interface{}) interface{} {
	if value == nil {
		return def
	}
	return value
}

func Marshal(value interface{}) string {
	if value == nil {
		return ""
	}

	v, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return ""
	}

	return fmt.Sprintf("```json\n%s\n```", string(v))
}

func MarshalInline(value interface{}) string {
	if value == nil {
		return ""
	}

	v, err := json.Marshal(value)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("`%s`", string(v))
}

func TableHeader(columns []string) string {
	var header string
	var divider string
	for _, col := range columns {
		header += fmt.Sprintf("| %s ", col)
		divider += fmt.Sprintf("|%s", strings.Repeat("-", len(col)+2))
	}
	header += "|\n"
	header += divider
	header += "|\n"
	return header
}

func Slice(items ...string) []string {
	var s []string
	s = append(s, items...)
	return s
}

func Append(slice []string, value ...string) []string {
	return append(slice, value...)
}

func FuncMap(refUrl string) template.FuncMap {
	referenceURL = refUrl
	return template.FuncMap{
		"join":             Join,
		"dict":             Dict,
		"inStringSlice":    InStringSlice,
		"schemaLink":       GetSchemaLink,
		"toCamel":          strcase.ToCamel,
		"toHTMLNewLines":   ToHTMLNewLines,
		"lower":            strings.ToLower,
		"replace":          strings.ReplaceAll,
		"default":          Default,
		"marshal":          Marshal,
		"marshalInline":    MarshalInline,
		"sum":              Sum,
		"slugify":          Slugify,
		"breakLine":        BreakLine,
		"tableHeader":      TableHeader,
		"getRefRootSchema": GetRefRootSchema,
		"slice":            Slice,
		"append":           Append,
	}
}
