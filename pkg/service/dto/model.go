package dto

import "github.com/getkin/kin-openapi/openapi3"

type Schema struct {
	Name string
	*openapi3.SchemaRef
}

type Operation struct {
	Name   string
	Method string
	*openapi3.Operation
}