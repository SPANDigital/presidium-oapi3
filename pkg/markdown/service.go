package markdown

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/SPANDigital/presidium-oapi3/pkg/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/tpl"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
)

const (
	methodTitle = "methodTitle"
	operationId = "operationId"
)

type MarkdownService struct {
	templates *template.Template
	cfg       Config
}

func NewMarkdownService(cfg Config) (MarkdownService, error) {
	if len(cfg.ApiName) > 0 {
		cfg.ApiName = fmt.Sprintf("/%s", cfg.ApiName)
	}

	templates := template.New("").Funcs(tpl.FuncMap(fmt.Sprintf("%s%s", cfg.ReferenceURL, cfg.ApiName)))
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

	return MarkdownService{
		templates: templates,
		cfg:       cfg,
	}, nil
}

func (ms *MarkdownService) ConvertToMarkdown(filename string) error {
	log.Infof("Loading swagger from file %s...", filename)

	loader := &openapi3.Loader{Context: context.Background(), IsExternalRefsAllowed: ms.cfg.AllowExternalRefs}
	swagger, err := loader.LoadFromFile(filename)
	if err != nil {
		return err
	}

	err = ms.createIndexFiles()
	if err != nil {
		return err
	}

	sequence := 0
	for path, item := range swagger.Paths {
		_ = ms.processOperations(path, item.Operations(), sequence)
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

	return ms.processTags(swagger.Tags)
}

func (ms *MarkdownService) basePath() string {
	return filepath.Clean(fmt.Sprintf("%s%s", ms.cfg.ReferenceURL, ms.cfg.ApiName))
}

func (ms *MarkdownService) rootPath() string {
	return filepath.Join(ms.cfg.OutputDir, "content", ms.cfg.ReferenceURL)
}

func (ms *MarkdownService) processSchemas(schemas openapi3.Schemas) error {
	for name, schema := range schemas {
		log.Infof("Processing schema %s...", name)
		if !strings.HasPrefix(ms.cfg.ReferenceURL, "/") {
			ms.cfg.ReferenceURL = fmt.Sprintf("/%s", ms.cfg.ReferenceURL)
		}

		theSchema := Schema{
			Name:            name,
			PresidiumRefURL: filepath.Clean(ms.cfg.ReferenceURL),
			SchemaRef:       schema,
			Config:          ms.cfg,
		}
		dir := filepath.Clean(fmt.Sprintf("%s/content/%s/components/schemas", ms.cfg.OutputDir, ms.basePath()))
		name := fmt.Sprintf("%s.md", strcase.ToSnake(name))
		err := ms.processTemplate(dir, name, "templates/schemas.gomd", theSchema)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *MarkdownService) processResponses(responses map[string]*openapi3.ResponseRef) error {
	for name, response := range responses {
		log.Infof("Processing response %s...", name)
		theResponse := Response{
			Name:            name,
			PresidiumRefURL: ms.cfg.ReferenceURL,
			ResponseRef:     response,
		}
		dir := filepath.Clean(fmt.Sprintf("%s/content/%s/components/responses", ms.cfg.OutputDir, ms.basePath()))
		name := fmt.Sprintf("%s.md", strcase.ToSnake(name))
		err := ms.processTemplate(dir, name, "templates/responses.gomd", theResponse)
		if err != nil {
			return err
		}
	}
	return nil
}

// cleanForMarkdown ensures what is written to the markdown file is clean:
// - trimmed line spaces
// - empty lines
func (ms MarkdownService) cleanForMarkdown(b bytes.Buffer) bytes.Buffer {
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

func (ms MarkdownService) processOperation(operation Operation, parentFolder string) error {
	// Create filename with weight as prefix if flag selected
	var name string
	if ms.cfg.SortFilePath {
		name = GetWeightedFilename(operation.Weight, operation.OperationID)
		name = fmt.Sprintf("%s.md", strcase.ToSnake(name))
	} else {
		name = fmt.Sprintf("%s.md", strcase.ToSnake(operation.OperationID))
	}

	if len(operation.Tags) == 0 {
		dir := filepath.Clean(fmt.Sprintf("%s/content/%s/operations/Default", ms.cfg.OutputDir, parentFolder))
		err := ms.processTemplate(dir, name, "templates/operation.gomd", operation)
		if err != nil {
			return err
		}
	} else {
		for _, tag := range operation.Tags {
			dir := filepath.Clean(fmt.Sprintf("%s/content/%s/operations/%s", ms.cfg.OutputDir, parentFolder, tag))
			err := ms.processTemplate(dir, name, "templates/operation.gomd", operation)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (ms MarkdownService) processOperations(path string, operations map[string]*openapi3.Operation, count int) error {
	log.Infof("Processing operations %s...", path)

	for method, operation := range operations {
		tplOperation := Operation{
			Method:      method,
			Name:        path,
			Operation:   operation,
			MethodTitle: ms.cfg.TitleFormat == methodTitle,
			Weight:      count + 1,
		}
		err := ms.processOperation(tplOperation, ms.basePath())
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *MarkdownService) processInfo(info *openapi3.Info) error {
	log.Info("Processing info templates...")
	dir := filepath.Clean(fmt.Sprintf("%s/content/%s/", ms.cfg.OutputDir, ms.basePath()))
	name := "info.md"
	err := ms.processTemplate(dir, name, "templates/info.gomd", info)
	return err
}

func (ms *MarkdownService) processTags(tags openapi3.Tags) error {
	for _, tag := range tags {
		log.Infof("Processing tag %s...", tag.Name)
		dir := filepath.Clean(fmt.Sprintf("%s/content/%s/operations/%s", ms.cfg.OutputDir, ms.basePath(), tag.Name))
		name := "_index.md"
		err := ms.processTemplate(dir, name, "templates/tag.gomd", tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *MarkdownService) processTemplate(dir string, name string, tpl string, obj interface{}) error {
	path := fmt.Sprintf("%s/%s", strings.ToLower(dir), name)
	err := ms.createSubIndex(dir)
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

func (ms *MarkdownService) createIndexFiles() error {
	log.Info("Creating index files...")
	dirs := map[string]string{
		"components":           "Components",
		"components/schemas":   "Schemas",
		"components/responses": "Responses",
		"operations":           "Operations",
		"":                     cases.Title(language.English).String(ms.cfg.ReferenceURL),
	}

	for dir, title := range dirs {
		index := Index{Title: title}
		baseDir := filepath.Clean(fmt.Sprintf("%s/content/%s/%s", ms.cfg.OutputDir, ms.basePath(), dir))
		err := ms.processTemplate(baseDir, "_index.md", "templates/index.gomd", index)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *MarkdownService) createSubIndex(path string) error {
	path = filepath.Clean(path)
	log.Debugf("creating index: %s", path)
	if ms.rootPath() == path {
		return nil
	}

	indexPath := filepath.Join(path, "_index.md")
	if _, err := os.Stat(indexPath); !os.IsNotExist(err) {
		return nil
	}

	if err := os.MkdirAll(path, fs.ModePerm); err != nil {
		return err
	}

	dirName := filepath.Base(path)
	fm := fmt.Sprintf("---\ntitle: %s\n---", dirName)
	if err := os.WriteFile(indexPath, []byte(fm), os.ModePerm); err != nil {
		return err
	}

	return ms.createSubIndex(filepath.Dir(path))
}
