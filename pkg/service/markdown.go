package service

import (
	"fmt"
	"github.com/SPANDigital/presidium-oapi3/pkg/infrastructure/log"
	"github.com/getkin/kin-openapi/openapi3"
	"os"
	"strings"
	"text/template"
)

type Operation struct {
	Name   string
	Method string
	*openapi3.Operation
}

func Join(list []string, sep string) string {
	return strings.Join(list, sep)
}

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"join": Join,
	}
}

func ConvertToMarkdown() {
	log.Info("Converting swagger.yml to markdown")
	swagger, _ := openapi3.NewSwaggerLoader().LoadSwaggerFromFile("swagger.yml")
	GenerateOperations(swagger.Paths)
}

func GenerateOperations(paths openapi3.Paths) error {
	t, err := template.New("operation.gomd").Funcs(FuncMap()).ParseFiles(
		"pkg/templates/operation.gomd",
		"pkg/templates/parameters.gomd",
		"pkg/templates/request_body.gomd",
		"pkg/templates/content.gomd",
		"pkg/templates/responses.gomd",
		"pkg/templates/headers.gomd",
		"pkg/templates/security.gomd",
	)
	if err != nil {
		log.Error(err)
		return err
	}
	for path, item := range paths {
		for method, operation := range item.Operations() {
			tplOperation := Operation{
				Method:    method,
				Name:      path,
				Operation: operation,
			}
			f, err := os.Create(fmt.Sprintf("out/%s.md", tplOperation.OperationID))
			if err != nil {
				log.Error("create file: ", err)
				return err
			}
			err = t.Execute(f, tplOperation)
			if err != nil {
				log.Error("execute: ", err)
				return err
			}
		}
	}
	return nil
}
