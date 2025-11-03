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
					Type: &openapi3.Types{"object"},
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
				Type: &openapi3.Types{"string"},
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
		Type: &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type: &openapi3.Types{"object"},
				Properties: openapi3.Schemas{
					"name": {
						Value: &openapi3.Schema{Type: &openapi3.Types{"string"}},
					},
				},
			},
		},
	}

	assert.True(t, NotEmpty(schemaWithItems.Items), "Expected NotEmpty to return true for SchemaRef with Value and Properties")

	// Test case 2: Array with items that are a reference (should not be rendered as nested)
	schemaWithRef := &openapi3.Schema{
		Type: &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{
			Ref: "#/components/schemas/ExternalSchema",
		},
	}

	assert.True(t, NotEmpty(schemaWithRef.Items), "Expected NotEmpty to return true for SchemaRef with Ref")

	// Test case 3: Array with nil items
	// When Items field is explicitly nil, NotEmpty should return false
	schemaWithNilItems := &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
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

func TestDeref(t *testing.T) {
	tests := []struct {
		name     string
		ptr      *bool
		expected bool
	}{
		{
			name:     "nil pointer",
			ptr:      nil,
			expected: false,
		},
		{
			name:     "pointer to true",
			ptr:      func() *bool { b := true; return &b }(),
			expected: true,
		},
		{
			name:     "pointer to false",
			ptr:      func() *bool { b := false; return &b }(),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Deref(tt.ptr)
			assert.Equal(t, tt.expected, result, "Deref(%v) = %v, expected %v", tt.ptr, result, tt.expected)
		})
	}
}

func TestTypeIs(t *testing.T) {
	tests := []struct {
		name         string
		types        interface{}
		expectedType string
		expected     bool
	}{
		{
			name:         "nil types",
			types:        nil,
			expectedType: "string",
			expected:     false,
		},
		{
			name:         "nil openapi3.Types pointer",
			types:        (*openapi3.Types)(nil),
			expectedType: "string",
			expected:     false,
		},
		{
			name:         "empty openapi3.Types",
			types:        &openapi3.Types{},
			expectedType: "string",
			expected:     false,
		},
		{
			name:         "openapi3.Types with matching type",
			types:        &openapi3.Types{"string"},
			expectedType: "string",
			expected:     true,
		},
		{
			name:         "openapi3.Types with non-matching type",
			types:        &openapi3.Types{"integer"},
			expectedType: "string",
			expected:     false,
		},
		{
			name:         "openapi3.Types with multiple types - match first",
			types:        &openapi3.Types{"string", "null"},
			expectedType: "string",
			expected:     true,
		},
		{
			name:         "openapi3.Types with multiple types - match second",
			types:        &openapi3.Types{"string", "null"},
			expectedType: "null",
			expected:     true,
		},
		{
			name:         "openapi3.Types with multiple types - no match",
			types:        &openapi3.Types{"string", "null"},
			expectedType: "integer",
			expected:     false,
		},
		{
			name:         "backward compatibility - string type matches",
			types:        "string",
			expectedType: "string",
			expected:     true,
		},
		{
			name:         "backward compatibility - string type no match",
			types:        "string",
			expectedType: "integer",
			expected:     false,
		},
		{
			name:         "invalid type - returns false",
			types:        123,
			expectedType: "string",
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TypeIs(tt.types, tt.expectedType)
			assert.Equal(t, tt.expected, result, "TypeIs(%v, %q) = %v, expected %v", tt.types, tt.expectedType, result, tt.expected)
		})
	}
}

func TestFirstType(t *testing.T) {
	tests := []struct {
		name     string
		types    interface{}
		expected string
	}{
		{
			name:     "nil types",
			types:    nil,
			expected: "",
		},
		{
			name:     "nil openapi3.Types pointer",
			types:    (*openapi3.Types)(nil),
			expected: "",
		},
		{
			name:     "empty openapi3.Types",
			types:    &openapi3.Types{},
			expected: "",
		},
		{
			name:     "openapi3.Types with single type",
			types:    &openapi3.Types{"string"},
			expected: "string",
		},
		{
			name:     "openapi3.Types with multiple types",
			types:    &openapi3.Types{"string", "null"},
			expected: "string",
		},
		{
			name:     "openapi3.Types with multiple types - different order",
			types:    &openapi3.Types{"integer", "number", "string"},
			expected: "integer",
		},
		{
			name:     "backward compatibility - string type",
			types:    "string",
			expected: "string",
		},
		{
			name:     "invalid type - returns empty string",
			types:    123,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FirstType(tt.types)
			assert.Equal(t, tt.expected, result, "FirstType(%v) = %q, expected %q", tt.types, result, tt.expected)
		})
	}
}
