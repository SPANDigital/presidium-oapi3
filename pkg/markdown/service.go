package markdown

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/SPANDigital/presidium-oapi3/pkg/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/tpl"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
)

type MarkdownService interface {
	ConvertToMarkdown(filename, outputDir string, methodTitle bool) error
}

type markdownService struct {
	templates    *template.Template
	outputDir    string
	referenceURL string
	apiName      string
}

func NewMarkdownService(referenceURL, apiName string) (MarkdownService, error) {
	templates := template.New("").Funcs(tpl.FuncMap(referenceURL))
	for _, path := range tpl.AssetNames() {
		tplResult, err := tpl.Asset(path)
		if err != nil {
			log.PanicWithFields("unable to obtain asset", log.Fields{
				"path":  path,
				"error": err,
			})
		}
		_, err = templates.New(path).Parse(string(tplResult))
		if err != nil {
			log.PanicWithFields("unable to parse template", log.Fields{
				"path":  path,
				"error": err,
			})
		}
	}
	if apiName != "" {
		apiName = fmt.Sprintf("/%s", apiName)
	}

	return &markdownService{
		templates:    templates,
		referenceURL: referenceURL,
		apiName:      apiName,
	}, nil
}

func (ms *markdownService) processSchemas(schemas map[string]*openapi3.SchemaRef) error {
	for name, schema := range schemas {
		log.Infof("Processing schema %s...", name)
		theSchema := Schema{
			Name:            name,
			PresidiumRefURL: ms.referenceURL,
			SchemaRef:       schema,
		}
		dir := fmt.Sprintf("%s/content/_reference%s/components/schemas", ms.outputDir, ms.apiName)
		name := fmt.Sprintf("%s.md", strcase.ToLowerCamel(name))
		err := ms.processTemplate(dir, name, "templates/schemas.gomd", theSchema)
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

func (ms *markdownService) processResponses(responses map[string]*openapi3.ResponseRef) error {
	for name, response := range responses {
		log.Infof("Processing response %s...", name)
		theResponse := Response{
			Name:            name,
			PresidiumRefURL: ms.referenceURL,
			ResponseRef:     response,
		}
		dir := fmt.Sprintf("%s/content/_reference%s/components/responses", ms.outputDir, ms.apiName)
		name := fmt.Sprintf("%s.md", strcase.ToLowerCamel(name))
		err := ms.processTemplate(dir, name, "templates/responses.gomd", theResponse)
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

func (ms markdownService) processOperation(operation Operation, parentFolder string, count int) error {
	if len(operation.Tags) == 0 {
		dir := fmt.Sprintf("%s/content/_reference%s/operations/Default", ms.outputDir, parentFolder)
		name := fmt.Sprintf("%03d-%s.md", count, strcase.ToLowerCamel(operation.OperationID))
		err := ms.processTemplate(dir, name, "templates/operation.gomd", operation)
		if err != nil {
			log.Error(err)
		}
	} else {
		for _, tag := range operation.Tags {
			dir := fmt.Sprintf("%s/content/_reference%s/operations/%s", ms.outputDir, parentFolder, tag)
			name := fmt.Sprintf("%03d-%s.md", count, strcase.ToLowerCamel(operation.OperationID))
			err := ms.processTemplate(dir, name, "templates/operation.gomd", operation)
			if err != nil {
				log.Error(err)
			}
		}
	}
	return nil
}

func (ms markdownService) processOperations(path string, operations map[string]*openapi3.Operation, methodTitle bool, count int) error {
	log.Infof("Processing operations %s...", path)

	for method, operation := range operations {
		tplOperation := Operation{
			Method:      method,
			Name:        path,
			Operation:   operation,
			MethodTitle: methodTitle,
		}
		err := ms.processOperation(tplOperation, ms.apiName, count)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (ms *markdownService) processInfo(info *openapi3.Info) error {
	log.Info("Processing info templates...")
	dir := fmt.Sprintf("%s/content/_reference%s/", ms.outputDir, ms.apiName)
	name := "01_info.md"
	err := ms.processTemplate(dir, name, "templates/info.gomd", info)
	return err
}

func (ms *markdownService) processTags(tags openapi3.Tags) error {
	for _, tag := range tags {
		log.Infof("Processing tag %s...", tag.Name)
		dir := fmt.Sprintf("%s/content/_reference%s/operations/%s", ms.outputDir, ms.apiName, tag.Name)
		name := "index.md"
		err := ms.processTemplate(dir, name, "templates/tag.gomd", tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *markdownService) processTemplate(dir string, name string, tpl string, obj interface{}) error {
	path := fmt.Sprintf("%s/%s", dir, name)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = ms.templates.ExecuteTemplate(&b, tpl, obj)
	if err != nil {
		return err
	}
	cleaned := ms.cleanForMarkdown(b)
	_, err = f.Write(cleaned.Bytes())
	return err
}

func (ms *markdownService) createIndexFiles() error {
	log.Info("Creating index files...")
	dirs := map[string]string{
		"components":           "Components",
		"components/schemas":   "Schemas",
		"components/responses": "Responses",
		"operations":           "Operations",
	}
	for dir, title := range dirs {
		index := Index{Title: title}
		baseDir := fmt.Sprintf("%s/content/_reference%s/%s", ms.outputDir, ms.apiName, dir)
		err := ms.processTemplate(baseDir, "index.md", "templates/index.gomd", index)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *markdownService) ConvertToMarkdown(filename, outputDir string, methodTitle bool) error {
	log.Infof("Loading swagger from file %s...", filename)

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	ms.outputDir = outputDir

	err = ms.createIndexFiles()
	if err != nil {
		return err
	}
	count := 0
	for path, item := range swagger.Paths {
		ms.processOperations(path, item.Operations(), methodTitle, count)
		if err != nil {
			return err
		}
		count++
	}
	err = ms.processSchemas(swagger.Components.Schemas)
	if err != nil {
		return err
	}
	err = ms.processResponses(swagger.Components.Responses)
	if err != nil {
		return err
	}
	err = ms.processInfo(swagger.Info)
	if err != nil {
		return err
	}
	err = ms.processTags(swagger.Tags)
	return err
}
