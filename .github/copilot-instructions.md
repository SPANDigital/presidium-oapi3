# Copilot Instructions for Presidium OpenAPI 3

## Project Overview

This is a **Go CLI tool** that converts OpenAPI 3 specifications to markdown documentation for integration with Presidium documentation system. The tool processes OpenAPI specs and generates structured markdown files with templates.

## Architecture & Patterns

### Project Structure
- **`main.go`**: Entry point using Cobra CLI framework
- **`cmd/`**: CLI commands (root.go, convert.go) using Cobra patterns
- **`pkg/`**: Core application packages following Go conventions
  - **`pkg/markdown/`**: Core business logic for OpenAPI → Markdown conversion
  - **`pkg/log/`**: Structured logging wrapper around logrus
  - **`pkg/tpl/`**: Template functions and utilities for Go templates

### Key Design Patterns
1. **Service Pattern**: `MarkdownService` encapsulates conversion logic
2. **Configuration Pattern**: `Config` struct for all configurable options
3. **Template Pattern**: Go text/templates with custom function maps
4. **Embedded Resources**: Templates embedded using `//go:embed` directive
5. **Error Handling**: Consistent error wrapping and structured logging

## Code Style & Conventions

### Go Standards
- Follow standard Go project layout
- Use `gofmt` for formatting
- Package names are lowercase, single words
- Interface names end with "-er" when applicable
- Use receiver names that are abbreviations of the type (e.g., `ms *MarkdownService`)

### Error Handling
```go
// Always wrap errors with context
return errors.Wrap(err, "descriptive context")

// Use structured logging for errors
log.ErrorWithFields("operation failed", log.Fields{
    "file": filename,
    "error": err,
})
```

### Logging Patterns
```go
// Use appropriate log levels
log.Info("Starting operation...")
log.Infof("Processing %s...", filename)
log.ErrorWithFields("failed", log.Fields{"key": "value"})

// Include file context in debug mode
log.DebugWithFields("detailed info", log.Fields{
    "component": "markdown",
    "operation": "convert",
})
```

### Testing Conventions
- Test files end with `_test.go`
- Use `testify/assert` for assertions
- Create test configs with `t.TempDir()` for isolation
- Test both success and failure cases
- Use table-driven tests for multiple scenarios

```go
func TestFunctionName(t *testing.T) {
    config := getConfig(t)  // Helper for test configuration
    
    ms, err := NewMarkdownService(config)
    assert.NoError(t, err, "Unexpected error from NewMarkdownService: %v", err)
    
    // Test implementation
}
```

## Core Components

### MarkdownService
The main service that handles OpenAPI to Markdown conversion:

```go
type MarkdownService struct {
    templates *template.Template
    cfg       Config
}
```

**Key Methods:**
- `ConvertToMarkdown(filename string) error` - Main entry point
- `processSchemas()`, `processOperations()`, `processResponses()` - Process different OpenAPI components
- `processTemplate()` - Template rendering with path management
- `cleanForMarkdown()` - Output sanitization

### Template System
- Templates stored in `pkg/markdown/templates/` and embedded
- Custom template functions in `pkg/tpl/funcs.go`
- Consistent template data models in `pkg/markdown/model.go`

### Configuration
All configuration through the `Config` struct:
```go
type Config struct {
    ReferenceURL        string  // Base URL for generated docs
    ApiName             string  // API grouping name
    TitleFormat         string  // "methodTitle" or "operationId"
    SortFilePath        bool    // Weight-based sorting
    InlineProperties    bool    // Schema property handling
    OutputDir           string  // Target directory
    AllowExternalRefs   bool    // External reference handling
    IncludeRestrictions bool    // Schema restrictions
    IncludeExamples     bool    // Schema examples
}
```

## File Organization Patterns

### Generated Output Structure
```
{OutputDir}/content/{sanitized-reference-url}/{api-name}/
├── _index.md                    # API overview
├── components/
│   ├── schemas/                 # Schema definitions
│   ├── responses/               # Response definitions
│   └── securitySchemas/         # Security scheme definitions
└── operations/
    └── {tag}/                   # Operations grouped by tags
        └── {operation}.md       # Individual operations
```

### File Naming Conventions
- Use `strcase.ToSnake()` for file names
- Index files are always `_index.md`
- Optional weight prefixes when `SortFilePath` is enabled
- Path sanitization removes invalid filesystem characters

## Development Guidelines

### Adding New Features
1. **Templates**: Add to `pkg/markdown/templates/` directory
2. **Template Functions**: Add to `pkg/tpl/funcs.go` and register in `FuncMap`
3. **Data Models**: Define in `pkg/markdown/model.go`
4. **Processing Logic**: Add to `MarkdownService` methods
5. **Configuration**: Extend `Config` struct and CLI flags

### Template Development
```go
// Template functions follow this pattern
func CustomFunction(input string) string {
    // Implementation
    return result
}

// Register in FuncMap
func FuncMap(refUrl string) template.FuncMap {
    return template.FuncMap{
        "customFunction": CustomFunction,
        // ... other functions
    }
}
```

### CLI Command Pattern
```go
var customCmd = &cobra.Command{
    Use:   "command-name",
    Short: "Brief description",
    Run: func(cmd *cobra.Command, args []string) {
        // Implementation
    },
}

func init() {
    rootCmd.AddCommand(customCmd)
    // Add flags
    customCmd.Flags().StringVar(&variable, "flag", "default", "description")
}
```

## Testing Strategy

### Unit Tests
- Test each service method independently
- Mock external dependencies when needed
- Use temporary directories for file operations
- Test both happy path and error conditions

### Integration Tests
- Test full conversion workflow with real OpenAPI specs
- Verify generated file structure and content
- Test different configuration combinations

### Test Data
- Store test OpenAPI specs in `pkg/markdown/testdata/`
- Include both valid and invalid specs for error testing
- Use minimal but representative schemas

## Dependencies

### Core Dependencies
- **Cobra**: CLI framework - follow Cobra patterns for commands
- **OpenAPI 3**: Use `github.com/getkin/kin-openapi/openapi3` for parsing
- **Logrus**: Structured logging - use custom wrapper in `pkg/log/`
- **Testify**: Testing assertions and helpers

### Template Functions
- **strcase**: String case conversion (`ToSnake`, `ToCamel`, etc.)
- **Standard library**: `text/template`, `path/filepath`, `strings`

## Release & Distribution

- **GoReleaser**: Automated releases with `.goreleaser.yml`
- **Homebrew**: Distributed via SPANDigital/tap
- **GitHub Packages**: NPM wrapper package for Node.js environments
- **Multi-platform**: Builds for Linux, macOS, Windows

## Common Patterns

### Path Management
```go
// Always use filepath.Clean() for path construction
dir := filepath.Clean(fmt.Sprintf("%s/content/%s", outputDir, basePath))

// Sanitize URLs for filesystem use
func sanitizeReferenceURL() string {
    // Remove protocol and invalid characters
}
```

### Template Processing
```go
// Standard template processing pattern
func (ms *MarkdownService) processTemplate(dir, name, tpl string, obj interface{}) error {
    // 1. Ensure directory exists
    // 2. Create file
    // 3. Execute template
    // 4. Clean output
    // 5. Write to file
}
```

### Error Context
```go
// Always provide meaningful error context
if err != nil {
    return errors.Wrap(err, fmt.Sprintf("failed to process %s", componentName))
}
```

This tool is designed for reliability, extensibility, and integration with the broader Presidium documentation ecosystem. Follow these patterns to maintain consistency and ensure robust functionality.
