package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spandigital/presidium-oapi3/pkg/service"
	"log"
)

func main() {
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile("swagger.yml")
	if err != nil {
		log.Fatal(err)
	}
	operationService, err := service.NewOperationService()
	if err != nil {
		log.Fatal(err)
	}
	for path, item := range swagger.Paths {
		for method, operation := range item.Operations() {
			tplOperation := service.Operation{
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
