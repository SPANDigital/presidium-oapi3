package dto

import "github.com/getkin/kin-openapi/openapi3"

type Schema struct {
	Name            string
	PresidiumRefURL string
	*openapi3.SchemaRef
}

type Operation struct {
	Name   string
	Method string
	*openapi3.Operation
}

type Index struct {
	Title string
}
