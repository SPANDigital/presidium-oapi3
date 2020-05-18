package service

import (
	"bytes"
	"fmt"
	"github.com/SPANDigital/presidium-oapi3/pkg/infrastructure/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/service/dto"
	"github.com/SPANDigital/presidium-oapi3/pkg/service/tpl"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"os"
	"strings"
	"text/template"
)

type MarkdownService interface {
	ConvertToMarkdown(filename string) error
}

type markdownService struct {
	operationTemplate *template.Template
	schemaTemplate    *template.Template
	infoTemplate      *template.Template
	tagTemplate       *template.Template
}

func NewMarkdownService() (MarkdownService, error) {
	operationTemplate, err := template.New("operation.gomd").Funcs(tpl.FuncMap()).ParseFiles(
		"pkg/templates/operation.gomd",
		"pkg/templates/partials/parameters.gomd",
		"pkg/templates/partials/request_body.gomd",
		"pkg/templates/partials/content.gomd",
		"pkg/templates/partials/responses.gomd",
		"pkg/templates/partials/headers.gomd",
		"pkg/templates/partials/security.gomd",
		"pkg/templates/partials/schema.gomd",
	)
	schemaTemplate, err := template.New("schemas.gomd").Funcs(tpl.FuncMap()).ParseFiles(
		"pkg/templates/schemas.gomd",
		"pkg/templates/partials/schema.gomd",
	)
	infoTemplate, err := template.New("info.gomd").Funcs(tpl.FuncMap()).ParseFiles(
		"pkg/templates/info.gomd",
	)

	tagTemplate, err := template.New("tag.gomd").Funcs(tpl.FuncMap()).ParseFiles(
		"pkg/templates/tag.gomd",
	)
	if err != nil {
		return &markdownService{}, err
	}
	return &markdownService{
		operationTemplate: operationTemplate,
		schemaTemplate:    schemaTemplate,
		infoTemplate:      infoTemplate,
		tagTemplate:       tagTemplate,
	}, nil
}

func (ms *markdownService) processSchemas(schemas map[string]*openapi3.SchemaRef) error {
	for name, schema := range schemas {
		theSchema := dto.Schema{
			Name:      name,
			SchemaRef: schema,
		}
		dir := "out/content/_reference/components/schemas"
		name := fmt.Sprintf("%s.md", strcase.ToLowerCamel(name))
		err := ms.processTemplate(dir, name, ms.schemaTemplate, theSchema)
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

// cleanForMarkdown ensures what is written to the markdown file is clean:
// - trimmed line spaces
// - empty lines
func (ms markdownService) cleanForMarkdown(b bytes.Buffer) bytes.Buffer {
	var result bytes.Buffer
	var previousLine = ""
	for _, line := range strings.Split(b.String(), "\n") {
		// Keep just one empty line if exists
		line = strings.TrimSpace(line)
		if len(line) == 0 && len(previousLine) != 0 {
			result.WriteString("\n")
		}
		previousLine = line
		if len(line) > 0 {
			result.WriteString(line + "\n")
		}
	}
	return result
}

func (ms markdownService) processOperation(operation dto.Operation) error {
	for _, tag := range operation.Tags {
		dir := fmt.Sprintf("out/content/_reference/operations/%s", tag)
		name := fmt.Sprintf("%s.md", strcase.ToLowerCamel(operation.OperationID))
		err := ms.processTemplate(dir, name, ms.operationTemplate, operation)
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

func (ms markdownService) processOperations(path string, operations map[string]*openapi3.Operation) error {
	for method, operation := range operations {
		tplOperation := dto.Operation{
			Method:    method,
			Name:      path,
			Operation: operation,
		}
		err := ms.processOperation(tplOperation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (ms *markdownService) processInfo(info *openapi3.Info) error {
	dir := "out/content/_reference/"
	name := "01_info.md"
	err := ms.processTemplate(dir, name, ms.infoTemplate, info)
	return err
}

func (ms *markdownService) processTags(tags openapi3.Tags) {
	for _, tag := range tags {
		dir := fmt.Sprintf("out/content/_reference/operations/%s", tag.Name)
		name := "index.md"
		ms.processTemplate(dir, name, ms.tagTemplate, tag)
	}
}

func (ms *markdownService) processTemplate(dir string, name string, tpl *template.Template, obj interface{}) error {
	path := fmt.Sprintf("%s/%s", dir, name)
	os.MkdirAll(dir, os.ModePerm)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, obj)
	if err != nil {
		return err
	}
	cleaned := ms.cleanForMarkdown(b)
	f.Write(cleaned.Bytes())
	return nil
}

func (ms *markdownService) ConvertToMarkdown(filename string) error {
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	for path, item := range swagger.Paths {
		ms.processOperations(path, item.Operations())
	}
	ms.processSchemas(swagger.Components.Schemas)
	ms.processInfo(swagger.Info)
	ms.processTags(swagger.Tags)
	return err
}
