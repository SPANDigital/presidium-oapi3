package markdown

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func getConfig(t *testing.T) Config {
	return Config{
		ReferenceURL:      "http://example.com",
		ApiName:           "api",
		TitleFormat:       "methodTitle",
		SortFilePath:      false,
		InlineProperties:  false,
		OutputDir:         t.TempDir(),
		AllowExternalRefs: false,
	}
}

func TestNewMarkdownService(t *testing.T) {
	config := getConfig(t)
	expected := config
	expected.ApiName = "/api"

	ms, err := NewMarkdownService(config)

	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)
	assert.NotNil(t, ms.templates, "Expected templates to be non-nil, but got nil")
	assert.Equal(t, ms.cfg, expected, "Config does not match")
}

func TestConvertToMarkdown(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)
	// Success case
	err = ms.ConvertToMarkdown("./testdata/test.yaml")
	assert.NoError(t, err)
	// Fail case
	err = ms.ConvertToMarkdown("./testdata/invalid-file.yaml")
	assert.Error(t, err)
}

func TestBasePath(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	expectedPath := filepath.Clean(fmt.Sprintf("%s%s", config.ReferenceURL, filepath.Join("/", config.ApiName)))
	actualPath := ms.basePath()

	assert.Equal(t, expectedPath, actualPath, "Expected basePath() to return %q, but got %q", expectedPath, expectedPath)
}

func TestRootPath(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	expectedPath := filepath.Join(config.OutputDir, "content", config.ReferenceURL)
	actualPath := ms.rootPath()

	assert.Equal(t, expectedPath, actualPath, "Expected rootPath() to return %q, but got %q", expectedPath, expectedPath)
}

func TestProcessSchemas(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	// Initialize SchemaRef for schema1
	schema1 := openapi3.NewObjectSchema()
	schema1.Type = "object"

	err = ms.processSchemas(openapi3.Schemas{})
	assert.NoError(t, err)

	err = ms.processSchemas(openapi3.Schemas{
		"schema1": {Value: schema1},
	})
	assert.NoError(t, err)
}

func TestProcessResponses(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	err = ms.processResponses(map[string]*openapi3.ResponseRef{})
	assert.NoError(t, err)

	responseWithEmptyDescription := &openapi3.ResponseRef{
		Value: &openapi3.Response{},
	}
	err = ms.processResponses(map[string]*openapi3.ResponseRef{
		"response1": responseWithEmptyDescription,
	})
	assert.NoError(t, err)

	description := "This is a description"
	responseWithDescription := &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Description: &description,
		},
	}
	err = ms.processResponses(map[string]*openapi3.ResponseRef{
		"response2": responseWithDescription,
	})
	assert.NoError(t, err)
}

func TestProcessSecuritySchemas(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	err = ms.processSecuritySchemas(map[string]*openapi3.SecuritySchemeRef{})
	assert.NoError(t, err)

	emptySchema := &openapi3.SecuritySchemeRef{
		Value: &openapi3.SecurityScheme{},
	}
	err = ms.processSecuritySchemas(map[string]*openapi3.SecuritySchemeRef{
		"schema": emptySchema,
	})
	assert.NoError(t, err)

	schema := &openapi3.SecuritySchemeRef{
		Value: &openapi3.SecurityScheme{
			Type:             "http",
			Description:      "auth",
			Name:             "bearer auth",
			In:               "x-API",
			Scheme:           "bearer",
			BearerFormat:     "jwt",
			Flows:            nil,
			OpenIdConnectUrl: "",
		},
	}
	err = ms.processSecuritySchemas(map[string]*openapi3.SecuritySchemeRef{
		"schema": schema,
	})
	assert.NoError(t, err)
}

func TestCleanForMarkdown(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	b := bytes.Buffer{}
	b.WriteString("line1\n\nline2")
	cleaned := ms.cleanForMarkdown(b)
	assert.Equal(t, "line1\n\nline2\n", cleaned.String())
}

func TestProcessOperation(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	operationID := "test"
	operation := Operation{
		Name:   "Test",
		Method: "TestMethod",
		Operation: &openapi3.Operation{
			OperationID: operationID,
			Tags:        []string{"tag1", "tag2"},
		},
		MethodTitle: false,
		Weight:      100,
		GlobalSecurity: openapi3.SecurityRequirements{
			{"bearer": {"scope1", "scope2"}},
		},
	}

	parentFolder := "parent"
	err = ms.processOperation(operation, parentFolder)
	assert.NoError(t, err)
	// Check if file exists
	filePath1 := filepath.Clean(fmt.Sprintf("%s/content/%s/operations/%s/%s.md", ms.cfg.OutputDir, parentFolder, operation.Operation.Tags[0], strcase.ToSnake(operationID)))
	_, err = os.Stat(filePath1)
	assert.NoError(t, err)

	filePath2 := filepath.Clean(fmt.Sprintf("%s/content/%s/operations/%s/%s.md", ms.cfg.OutputDir, parentFolder, operation.Operation.Tags[1], strcase.ToSnake(operationID)))
	_, err = os.Stat(filePath2)
	assert.NoError(t, err)
}

func TestProcessOperations(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	err = ms.processOperations(ms.cfg.OutputDir, map[string]*openapi3.Operation{
		"get":  {},
		"post": {},
	}, 1, nil)
	assert.NoError(t, err)

	err = ms.processOperations(ms.cfg.OutputDir, nil, 0, nil)
	assert.NoError(t, err)
}

func TestProcessInfo(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	info := &openapi3.Info{
		Title: "My API",
	}
	err = ms.processInfo(info)
	assert.NoError(t, err)
	// Check if file exists
	filePath := fmt.Sprintf("%s/content/%s/_index.md", ms.cfg.OutputDir, ms.basePath())
	_, err = os.Stat(filePath)
	assert.NoError(t, err)
}

func TestProcessTags(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	tags := openapi3.Tags{
		{Name: "tag1"},
		{Name: "tag2"},
	}
	err = ms.processTags(tags)
	assert.NoError(t, err)
	// Check if files exist
	filePath1 := fmt.Sprintf("%s/content/%s/operations/tag1/_index.md", ms.cfg.OutputDir, ms.basePath())
	_, err = os.Stat(filePath1)
	assert.NoError(t, err)
	filePath2 := fmt.Sprintf("%s/content/%s/operations/tag2/_index.md", ms.cfg.OutputDir, ms.basePath())
	_, err = os.Stat(filePath2)
	assert.NoError(t, err)
}

func TestProcessTemplate(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	dir := ms.cfg.OutputDir
	name := "testfile.md"
	tplName := "testtpl"
	tplContent := "{{ . }}"
	obj := "Hello World!"

	// Add a test template.
	_, err = ms.templates.New(tplName).Parse(tplContent)
	assert.NoError(t, err)

	// Call process template with test data.
	err = ms.processTemplate(dir, name, tplName, obj)
	assert.NoError(t, err)

	// Ensure that we can read the file and it is cleaned properly.
	b, err := os.ReadFile(fmt.Sprintf("%s/%s", dir, name))
	assert.NoError(t, err)

	// Remove newlines from expected output.
	expected := bytes.NewBufferString(fmt.Sprintf("%s%s", obj, "\n")).Bytes()

	// Assert that the file was created correctly.
	assert.Equal(t, expected, b)
}

func TestCreateIndexFiles(t *testing.T) {
	config := getConfig(t)
	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	loader := &openapi3.Loader{Context: context.Background(), IsExternalRefsAllowed: ms.cfg.AllowExternalRefs}
	swagger, err := loader.LoadFromFile("../../swagger.yml")
	require.NoError(t, err)

	err = ms.createIndexFiles(swagger)
	assert.NoError(t, err)
	// Check if files exist
	dirs := map[string]string{
		"components":           "Components",
		"components/schemas":   "Schemas",
		"components/responses": "Responses",
		"operations":           "Operations",
		"":                     cases.Title(language.English).String(ms.cfg.ReferenceURL),
	}
	if len(swagger.Components.Responses) == 0 {
		delete(dirs, "components/responses")
	}

	if len(swagger.Components.Schemas) == 0 {
		delete(dirs, "components/schemas")
	}
	for dir := range dirs {
		filePath := fmt.Sprintf("%s/content/%s/%s/_index.md", ms.cfg.OutputDir, ms.basePath(), dir)
		_, err = os.Stat(filePath)
		assert.NoError(t, err)
	}
}

func TestCreateSubIndex(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	// create an index file for a non-existing path
	filePath := fmt.Sprintf("%s/test", config.OutputDir)
	err = ms.createSubIndex(filePath)
	assert.NoError(t, err)

	// create an index file for an existing path
	err = ms.createSubIndex(filePath)
	assert.NoError(t, err)
}
