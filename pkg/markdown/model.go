package markdown

import "github.com/getkin/kin-openapi/openapi3"

type Schema struct {
	Name            string
	PresidiumRefURL string
	Config          Config
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
	Weight      int
}

type Index struct {
	Title string
}
