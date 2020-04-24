package main

import (
	"strings"
	"text/template"
)

func Join(list []string, sep string) string {
	return strings.Join(list, sep)
}

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"join": Join,
	}
}
