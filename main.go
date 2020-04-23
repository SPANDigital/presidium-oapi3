package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	"log"
	"os"
	"text/template"
)

type Path struct {
	Name string
	*openapi3.PathItem
}

func main() {
	swagger, _ := openapi3.NewSwaggerLoader().LoadSwaggerFromFile("swagger.yml")
	t, err := template.ParseFiles("templates/operation.gomd")
	if err != nil {
		log.Print(err)
		return
	}

	paths := swagger.Paths
	for key, value := range paths {
		path := Path{
			Name:     key,
			PathItem: value,
		}

		f, err := os.Create("templates/result.md")
		if err != nil {
			log.Println("create file: ", err)
			return
		}
		err = t.Execute(f, path)
		if err != nil {
			log.Print("execute: ", err)
			return
		}
	}
}
