package service

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"os"
	"text/template"
)

const (
	OutputFormat = "out/%s.md"
)

type Operation struct {
	Name   string
	Method string
	*openapi3.Operation
}

type OperationService interface {
	ProcessOperation(operation Operation) error
}

type operationService struct {
	t *template.Template
}

func NewOperationService() (OperationService, error) {
	t, err := template.New("operation.gomd").Funcs(FuncMap()).ParseFiles(
		"templates/operation.gomd",
		"templates/parameters.gomd",
		"templates/request_body.gomd",
		"templates/content.gomd",
		"templates/responses.gomd",
		"templates/headers.gomd",
		"templates/security.gomd",
		"templates/schema.gomd",
	)
	if err != nil {
		return &operationService{}, err
	}
	return &operationService{t}, nil
}

func (o operationService) ProcessOperation(operation Operation) error {
	f, err := os.Create(fmt.Sprintf(OutputFormat, strcase.ToLowerCamel(operation.OperationID)))
	if err != nil {
		return err
	}
	err = o.t.Execute(f, operation)
	if err != nil {
		return err
	}
	return nil
}
