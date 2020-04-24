package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"log"
	"os"
	"text/template"
)

type Operation struct {
	Name   string
	Method string
	*openapi3.Operation
}

func main() {
	swagger, _ := openapi3.NewSwaggerLoader().LoadSwaggerFromFile("swagger.yml")
	GenerateOperations(swagger.Paths)
}

func GenerateOperations(paths openapi3.Paths) error {
	t, err := template.New("operation.gomd").Funcs(FuncMap()).ParseFiles(
		"templates/operation.gomd",
		"templates/parameters.gomd",
		"templates/request_body.gomd",
		"templates/content.gomd",
		"templates/responses.gomd",
		"templates/headers.gomd",
		"templates/security.gomd",
	)
	if err != nil {
		log.Print(err)
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
				log.Println("create file: ", err)
				return err
			}
			err = t.Execute(f, tplOperation)
			if err != nil {
				log.Print("execute: ", err)
				return err
			}
		}
	}
	return nil
}
