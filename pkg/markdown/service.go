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
	if apiName != "" {
		apiName = fmt.Sprintf("/%s", apiName)
	}

	templates := template.New("").Funcs(tpl.FuncMap(fmt.Sprintf("%s%s", referenceURL, apiName)))
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

	return &markdownService{
		templates:    templates,
		referenceURL: referenceURL,
		apiName:      apiName,
	}, nil
}

func (ms *markdownService) basePath() string {
	return fmt.Sprintf("%s%s", ms.referenceURL, ms.apiName)
}

func (ms *markdownService) processSchemas(schemas map[string]*openapi3.SchemaRef) error {
	for name, schema := range schemas {
		log.Infof("Processing schema %s...", name)
		theSchema := Schema{
			Name:            name,
			PresidiumRefURL: ms.referenceURL,
			SchemaRef:       schema,
		}
		dir := fmt.Sprintf("%s/content/%s/components/schemas", ms.outputDir, ms.basePath())
		name := fmt.Sprintf("%s.md", strcase.ToSnake(name))
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
		dir := fmt.Sprintf("%s/content/%s/components/responses", ms.outputDir, ms.basePath())
		name := fmt.Sprintf("%s.md", strcase.ToSnake(name))
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

func (ms markdownService) processOperation(operation Operation, parentFolder string) error {
	if len(operation.Tags) == 0 {
		dir := fmt.Sprintf("%s/content/%s/operations/Default", ms.outputDir, parentFolder)
		name := fmt.Sprintf("%s.md", strcase.ToSnake(operation.OperationID))
		err := ms.processTemplate(dir, name, "templates/operation.gomd", operation)
		if err != nil {
			log.Error(err)
		}
	} else {
		for _, tag := range operation.Tags {
			dir := fmt.Sprintf("%s/content/%s/operations/%s", ms.outputDir, parentFolder, tag)
			name := fmt.Sprintf("%s.md", strcase.ToSnake(operation.OperationID))
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
			Weight:      count + 1,
		}
		err := ms.processOperation(tplOperation, ms.basePath())
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (ms *markdownService) processInfo(info *openapi3.Info) error {
	log.Info("Processing info templates...")
	dir := fmt.Sprintf("%s/content/%s/", ms.outputDir, ms.basePath())
	name := "info.md"
	err := ms.processTemplate(dir, name, "templates/info.gomd", info)
	return err
}

func (ms *markdownService) processTags(tags openapi3.Tags) error {
	for _, tag := range tags {
		log.Infof("Processing tag %s...", tag.Name)
		dir := fmt.Sprintf("%s/content/%s/operations/%s", ms.outputDir, ms.basePath(), tag.Name)
		name := "index.md"
		err := ms.processTemplate(dir, name, "templates/tag.gomd", tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *markdownService) processTemplate(dir string, name string, tpl string, obj interface{}) error {
	path := fmt.Sprintf("%s/%s", strings.ToLower(dir), name)
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
		baseDir := fmt.Sprintf("%s/content/%s/%s", ms.outputDir, ms.basePath(), dir)
		err := ms.processTemplate(baseDir, "_index.md", "templates/index.gomd", index)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *markdownService) ConvertToMarkdown(filename, outputDir string, methodTitle bool) error {
	log.Infof("Loading swagger from file %s...", filename)

	swagger, err := openapi3.NewLoader().LoadFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	ms.outputDir = outputDir

	err = ms.createIndexFiles()
	if err != nil {
		return err
	}
	sequence := 0
	for path, item := range swagger.Paths {
		_ = ms.processOperations(path, item.Operations(), methodTitle, sequence)
		if err != nil {
			return err
		}
		sequence++
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
