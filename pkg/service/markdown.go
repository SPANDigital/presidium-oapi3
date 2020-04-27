package service

import (
	"fmt"
	"github.com/SPANDigital/presidium-oapi3/pkg/infrastructure/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/service/tpl"
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
	t, err := template.New("operation.gomd").Funcs(tpl.FuncMap()).ParseFiles(
		"pkg/templates/operation.gomd",
		"pkg/templates/parameters.gomd",
		"pkg/templates/request_body.gomd",
		"pkg/templates/content.gomd",
		"pkg/templates/responses.gomd",
		"pkg/templates/headers.gomd",
		"pkg/templates/security.gomd",
		"pkg/templates/schema.gomd",
	)
	if err != nil {
		return &operationService{}, err
	}
	return &operationService{t}, nil
}

func (o operationService) ProcessOperation(operation Operation) error {
	os.MkdirAll("out", os.ModePerm)
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

func ConvertToMarkdown() {
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile("swagger.yml")
	if err != nil {
		log.Fatal(err)
	}
	operationService, err := NewOperationService()
	if err != nil {
		log.Fatal(err)
	}
	for path, item := range swagger.Paths {
		for method, operation := range item.Operations() {
			tplOperation := Operation{
				Method:    method,
				Name:      path,
				Operation: operation,
			}
			err := operationService.ProcessOperation(tplOperation)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}