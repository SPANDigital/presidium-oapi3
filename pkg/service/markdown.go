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
	schemaTemplate *template.Template
	infoTemplate *template.Template
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
	if err != nil {
		return &markdownService{}, err
	}
	return &markdownService{
		operationTemplate: operationTemplate,
		schemaTemplate: schemaTemplate,
		infoTemplate: infoTemplate,
	}, nil
}

func (ms *markdownService) processSchemas(schemas map[string]*openapi3.SchemaRef) error {
	for name, schema := range schemas {
		theSchema := dto.Schema{
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
		err = ms.schemaTemplate.Execute(&b, theSchema)
		if err != nil {
			return err
		}
		cleaned := ms.cleanForMarkdown(b)
		f.Write(cleaned.Bytes())
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
		path := fmt.Sprintf("%s/%s.md", dir, strcase.ToLowerCamel(operation.OperationID))
		os.MkdirAll(dir, os.ModePerm)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		var b bytes.Buffer
		err = ms.operationTemplate.Execute(&b, operation)
		if err != nil {
			return err
		}
		cleaned := ms.cleanForMarkdown(b)
		f.Write(cleaned.Bytes())
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
	path := fmt.Sprintf("%s/01_%s.md", dir, "info")
	os.MkdirAll(dir, os.ModePerm)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = ms.infoTemplate.Execute(&b, info)
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
	return err
}
