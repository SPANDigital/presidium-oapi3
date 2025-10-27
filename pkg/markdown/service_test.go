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

	// Expect sanitized ReferenceURL (http:// removed) + API name
	expectedPath := "example.com/api"
	actualPath := ms.basePath()

	assert.Equal(t, expectedPath, actualPath, "Expected basePath() to return %q, but got %q", expectedPath, actualPath)
}

func TestRootPath(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	// Expect sanitized ReferenceURL (http:// removed)
	expectedPath := filepath.Join(config.OutputDir, "content", "example.com")
	actualPath := ms.rootPath()

	assert.Equal(t, expectedPath, actualPath, "Expected rootPath() to return %q, but got %q", expectedPath, actualPath)
}

func TestProcessSchemas(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	// Initialize SchemaRef for schema1
	schema1 := openapi3.NewObjectSchema()
	schema1.Type = &openapi3.Types{"object"}

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
		var filePath string
		if dir == "" {
			filePath = fmt.Sprintf("%s/content/%s/_index.md", ms.cfg.OutputDir, ms.basePath())
		} else {
			filePath = fmt.Sprintf("%s/content/%s/%s/_index.md", ms.cfg.OutputDir, ms.basePath(), dir)
		}
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

func TestProcessSchemasWithArrayOfObjectsWithProperties(t *testing.T) {
	config := getConfig(t)
	config.IncludeRestrictions = true

	ms, err := NewMarkdownService(config)
	require.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)

	// Create a schema similar to EditorialItem with canvases array
	// This tests the fix for arrays of objects with nested properties
	editorialItemSchema := openapi3.NewObjectSchema()
	editorialItemSchema.Type = "object"
	editorialItemSchema.Description = "Editorial Item Dataset - Complete Editorial Item information from Apple App Store."

	// Create the canvases array property with items containing nested object properties
	canvasItemSchema := openapi3.NewObjectSchema()
	canvasItemSchema.Type = "object"
	canvasItemSchema.Description = "Display canvas for editorial item."
	canvasItemSchema.Properties = openapi3.Schemas{
		"storefronts": {
			Value: &openapi3.Schema{
				Type:        "array",
				Description: "(Required) Array of storefronts where this canvas is displayed.",
				Items: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: "string",
					},
				},
			},
		},
		"shelves": {
			Value: &openapi3.Schema{
				Type:        "object",
				Description: "(Required) Canvas shelves, listed in order of display.",
				Properties: openapi3.Schemas{
					"textBlockShelf": {
						Value: &openapi3.Schema{
							Type:        "object",
							Description: "A text block shelf to display.",
							Properties: openapi3.Schemas{
								"editorialCopy": {
									Value: &openapi3.Schema{
										Type:        "string",
										Description: "Map of locale to text.",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	editorialItemSchema.Properties = openapi3.Schemas{
		"id": {
			Value: &openapi3.Schema{
				Type:        "string",
				Description: "(Required) The ID for the editorial item.",
			},
		},
		"canvases": {
			Value: &openapi3.Schema{
				Type:        "array",
				Description: "(Required) Displayable canvases for this editorial item.",
				Items:       &openapi3.SchemaRef{Value: canvasItemSchema},
			},
		},
	}

	// Process the schema
	err = ms.processSchemas(openapi3.Schemas{
		"EditorialItem": {Value: editorialItemSchema},
	})
	assert.NoError(t, err)

	// Verify the markdown file was created
	schemaFilePath := filepath.Clean(fmt.Sprintf("%s/content/%s/components/schemas/editorial_item.md", ms.cfg.OutputDir, ms.basePath()))
	_, err = os.Stat(schemaFilePath)
	require.NoError(t, err, "Expected schema markdown file to exist at %s", schemaFilePath)

	// Read the generated markdown file
	content, err := os.ReadFile(schemaFilePath)
	require.NoError(t, err, "Failed to read generated markdown file")
	markdownContent := string(content)

	// Verify key sections exist in the markdown
	t.Run("verify main properties table", func(t *testing.T) {
		assert.Contains(t, markdownContent, "| id|string|", "Expected 'id' property in main properties table")
		assert.Contains(t, markdownContent, "| canvases|[canvases](#editorialitem-canvases)|", "Expected 'canvases' property with link in main properties table")
	})

	t.Run("verify canvases section is rendered", func(t *testing.T) {
		assert.Contains(t, markdownContent, "## canvases { #editorialitem-canvases }", "Expected canvases section header")
		assert.Contains(t, markdownContent, "Display canvas for editorial item.", "Expected canvases description")
	})

	t.Run("verify nested properties in canvases", func(t *testing.T) {
		assert.Contains(t, markdownContent, "| storefronts|string|", "Expected 'storefronts' property in canvases")
		assert.Contains(t, markdownContent, "| shelves|[shelves](#editorialitem-shelves)|", "Expected 'shelves' property with link in canvases")
	})

	t.Run("verify shelves nested section", func(t *testing.T) {
		assert.Contains(t, markdownContent, "## shelves { #editorialitem-shelves }", "Expected shelves section header")
		assert.Contains(t, markdownContent, "Canvas shelves, listed in order of display.", "Expected shelves description")
	})

	t.Run("verify deeply nested properties", func(t *testing.T) {
		assert.Contains(t, markdownContent, "## textBlockShelf { #editorialitem-textblockshelf }", "Expected textBlockShelf section header")
		assert.Contains(t, markdownContent, "A text block shelf to display.", "Expected textBlockShelf description")
		assert.Contains(t, markdownContent, "| editorialCopy|string|", "Expected 'editorialCopy' property in textBlockShelf")
	})
}

func TestNotEmptyWithSchemaRef(t *testing.T) {
	config := getConfig(t)

	ms, err := NewMarkdownService(config)
	require.NoError(t, err)

	// Create a schema with an array that has a non-nil Items SchemaRef
	schema := openapi3.NewObjectSchema()
	schema.Properties = openapi3.Schemas{
		"items": {
			Value: &openapi3.Schema{
				Type: "array",
				Items: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: "object",
						Properties: openapi3.Schemas{
							"name": {
								Value: &openapi3.Schema{
									Type: "string",
								},
							},
						},
					},
				},
			},
		},
	}

	// Process the schema - this should not panic or error
	err = ms.processSchemas(openapi3.Schemas{
		"TestSchema": {Value: schema},
	})
	assert.NoError(t, err)

	// Verify the markdown file contains the nested items section
	schemaFilePath := filepath.Clean(fmt.Sprintf("%s/content/%s/components/schemas/test_schema.md", ms.cfg.OutputDir, ms.basePath()))
	content, err := os.ReadFile(schemaFilePath)
	require.NoError(t, err)
	markdownContent := string(content)

	// Verify that the items section was rendered (proving notEmpty works with SchemaRef)
	assert.Contains(t, markdownContent, "## items { #testschema-items }", "Expected items section to be rendered when Items SchemaRef is non-nil with properties")
}
