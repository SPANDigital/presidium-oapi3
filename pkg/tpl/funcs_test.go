package tpl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{
			name:     "nil value",
			value:    nil,
			expected: false,
		},
		{
			name:     "empty string",
			value:    "",
			expected: false,
		},
		{
			name:     "non-empty string",
			value:    "hello",
			expected: true,
		},
		{
			name:     "nil slice",
			value:    ([]interface{})(nil),
			expected: false,
		},
		{
			name:     "empty slice",
			value:    []interface{}{},
			expected: false,
		},
		{
			name:     "non-empty slice",
			value:    []interface{}{"item1", "item2"},
			expected: true,
		},
		{
			name:     "nil map",
			value:    (map[interface{}]interface{})(nil),
			expected: false,
		},
		{
			name:     "empty map",
			value:    map[interface{}]interface{}{},
			expected: false,
		},
		{
			name:     "non-empty map",
			value:    map[interface{}]interface{}{"key": "value"},
			expected: true,
		},
		{
			name:     "integer",
			value:    42,
			expected: true,
		},
		{
			name:     "boolean true",
			value:    true,
			expected: true,
		},
		{
			name:     "boolean false",
			value:    false,
			expected: true,
		},
		{
			name: "non-nil SchemaRef pointer with Value",
			value: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "object",
				},
			},
			expected: true,
		},
		{
			name: "non-nil SchemaRef pointer without Value (reference)",
			value: &openapi3.SchemaRef{
				Ref: "#/components/schemas/SomeSchema",
			},
			expected: true,
		},
		{
			name: "non-nil Schema pointer",
			value: &openapi3.Schema{
				Type: "string",
			},
			expected: true,
		},
		{
			name:     "empty struct",
			value:    struct{}{},
			expected: true,
		},
		{
			name: "struct with fields",
			value: struct {
				Name string
			}{Name: "test"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NotEmpty(tt.value)
			assert.Equal(t, tt.expected, result, "NotEmpty(%v) = %v, expected %v", tt.value, result, tt.expected)
		})
	}
}

func TestNotEmptyWithSchemaRefInTemplate(t *testing.T) {
	// This test simulates the real-world scenario where NotEmpty is used
	// in templates to check if an array items SchemaRef should be rendered

	// Test case 1: Array with items that have a Value (should be rendered)
	schemaWithItems := &openapi3.Schema{
		Type: "array",
		Items: &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type: "object",
				Properties: openapi3.Schemas{
					"name": {
						Value: &openapi3.Schema{Type: "string"},
					},
				},
			},
		},
	}

	assert.True(t, NotEmpty(schemaWithItems.Items), "Expected NotEmpty to return true for SchemaRef with Value and Properties")

	// Test case 2: Array with items that are a reference (should not be rendered as nested)
	schemaWithRef := &openapi3.Schema{
		Type: "array",
		Items: &openapi3.SchemaRef{
			Ref: "#/components/schemas/ExternalSchema",
		},
	}

	assert.True(t, NotEmpty(schemaWithRef.Items), "Expected NotEmpty to return true for SchemaRef with Ref")

	// Test case 3: Array with nil items
	// When Items field is explicitly nil, NotEmpty should return false
	schemaWithNilItems := &openapi3.Schema{
		Type:  "array",
		Items: nil,
	}

	// Verify Items is actually nil
	if schemaWithNilItems.Items == nil {
		assert.False(t, NotEmpty(schemaWithNilItems.Items), "Expected NotEmpty to return false for nil Items")
	} else {
		// If Items is not nil (due to zero value initialization), skip this assertion
		t.Log("Items is not nil, skipping nil check")
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple string",
			input:    "Hello World",
			expected: "hello-world",
		},
		{
			name:     "string with special characters",
			input:    "Hello_World!",
			expected: "hello-world",
		},
		{
			name:     "camelCase",
			input:    "helloWorld",
			expected: "hello-world",
		},
		{
			name:     "PascalCase",
			input:    "HelloWorld",
			expected: "hello-world",
		},
		{
			name:     "with numbers",
			input:    "Hello123World",
			expected: "hello-123-world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Slugify(tt.input)
			assert.Equal(t, tt.expected, result, "Slugify(%q) = %q, expected %q", tt.input, result, tt.expected)
		})
	}
}

func TestBreakLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single line",
			input:    "Hello World",
			expected: "Hello World",
		},
		{
			name:     "multiple lines",
			input:    "Hello\nWorld",
			expected: "Hello<br>World",
		},
		{
			name:     "multiple newlines",
			input:    "Hello\n\nWorld",
			expected: "Hello<br><br>World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BreakLine(tt.input)
			assert.Equal(t, tt.expected, result, "BreakLine(%q) = %q, expected %q", tt.input, result, tt.expected)
		})
	}
}

func TestInStringSlice(t *testing.T) {
	tests := []struct {
		name     string
		list     []string
		elem     string
		expected bool
	}{
		{
			name:     "element exists",
			list:     []string{"apple", "banana", "cherry"},
			elem:     "banana",
			expected: true,
		},
		{
			name:     "element does not exist",
			list:     []string{"apple", "banana", "cherry"},
			elem:     "date",
			expected: false,
		},
		{
			name:     "empty list",
			list:     []string{},
			elem:     "apple",
			expected: false,
		},
		{
			name:     "nil list",
			list:     nil,
			elem:     "apple",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InStringSlice(tt.list, tt.elem)
			assert.Equal(t, tt.expected, result, "InStringSlice(%v, %q) = %v, expected %v", tt.list, tt.elem, result, tt.expected)
		})
	}
}
