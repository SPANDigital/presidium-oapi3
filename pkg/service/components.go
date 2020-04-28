package service

import (
	"bytes"
	"fmt"
	"github.com/SPANDigital/presidium-oapi3/pkg/service/tpl"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"os"
	"text/template"
)

type Schema struct {
	Name string
	*openapi3.SchemaRef
}

type ComponentsService interface {
	ProcessSchemas(schemas map[string]*openapi3.SchemaRef) error
}

type componentsService struct {
	t *template.Template
}

func NewComponentsService() (ComponentsService, error) {
	t, err := template.New("schemas.gomd").Funcs(tpl.FuncMap()).ParseFiles(
		"pkg/templates/schemas.gomd",
		"pkg/templates/partials/schema.gomd",
	)
	if err != nil {
		return &componentsService{}, err
	}
	return &componentsService{t}, nil
}

func (c *componentsService) ProcessSchemas(schemas map[string]*openapi3.SchemaRef) error {
	for name, schema := range schemas {
		theSchema := Schema{
			Name:      name,
			SchemaRef: schema,
		}
		dir := "out/content/_reference/components/schemas"
		path := fmt.Sprintf("%s/%s.md", dir, strcase.ToLowerCamel(name))
		os.MkdirAll(dir, os.ModePerm)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		var b bytes.Buffer
		err = c.t.Execute(&b, theSchema)
		if err != nil {
			return err
		}
		cleaned := cleanForMarkdown(b)
		f.Write(cleaned.Bytes())
	}
	return nil
}
