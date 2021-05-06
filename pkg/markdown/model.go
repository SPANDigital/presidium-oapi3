package markdown

import "github.com/getkin/kin-openapi/openapi3"

type Schema struct {
	Name            string
	PresidiumRefURL string
	*openapi3.SchemaRef
}

type Response struct {
	Name            string
	PresidiumRefURL string
	*openapi3.ResponseRef
}

type Operation struct {
	Name   string
	Method string
	*openapi3.Operation
	MethodTitle bool
}

type Index struct {
	Title string
}
