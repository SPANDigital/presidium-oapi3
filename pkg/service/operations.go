package service

import (
	"bytes"
	"fmt"
	"github.com/SPANDigital/presidium-oapi3/pkg/infrastructure/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/service/tpl"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"os"
	"text/template"
)

type Operation struct {
	Name   string
	Method string
	*openapi3.Operation
}

type OperationService interface {
	ProcessOperations(path string, operations map[string]*openapi3.Operation) error
}

type operationService struct {
	t *template.Template
}

func NewOperationService() (OperationService, error) {
	t, err := template.New("operation.gomd").Funcs(tpl.FuncMap()).ParseFiles(
		"pkg/templates/operation.gomd",
		"pkg/templates/partials/parameters.gomd",
		"pkg/templates/partials/request_body.gomd",
		"pkg/templates/partials/content.gomd",
		"pkg/templates/partials/responses.gomd",
		"pkg/templates/partials/headers.gomd",
		"pkg/templates/partials/security.gomd",
		"pkg/templates/partials/schema.gomd",
	)
	if err != nil {
		return &operationService{}, err
	}
	return &operationService{t}, nil
}

func (o operationService) processOperation(operation Operation) error {
	for _, tag := range operation.Tags {
		dir := fmt.Sprintf("out/content/_reference/operations/%s", tag)
		path := fmt.Sprintf("%s/%s.md", dir, strcase.ToLowerCamel(operation.OperationID))
		os.MkdirAll(dir, os.ModePerm)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		var b bytes.Buffer
		err = o.t.Execute(&b, operation)
		if err != nil {
			return err
		}
		cleaned := cleanForMarkdown(b)
		f.Write(cleaned.Bytes())
	}
	return nil
}

func (o operationService) ProcessOperations(path string, operations map[string]*openapi3.Operation) error {
	for method, operation := range operations {
		tplOperation := Operation{
			Method:    method,
			Name:      path,
			Operation: operation,
		}
		err := o.processOperation(tplOperation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
