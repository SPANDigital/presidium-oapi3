package service

import (
	"bytes"
	"github.com/SPANDigital/presidium-oapi3/pkg/infrastructure/log"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

// cleanForMarkdown ensures what is written to the markdown file is clean:
// - trimmed line spaces
// - empty lines
func cleanForMarkdown(b bytes.Buffer) bytes.Buffer {
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

func ConvertToMarkdown() {
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile("swagger.yml")
	if err != nil {
		log.Fatal(err)
	}
	operationService, err := NewOperationService()
	componentsService, err := NewComponentsService()
	infoService, err := NewInfoService()
	if err != nil {
		log.Fatal(err)
	}
	for path, item := range swagger.Paths {
		operationService.ProcessOperations(path, item.Operations())
	}
	componentsService.ProcessSchemas(swagger.Components.Schemas)
	infoService.ProcessInfo(swagger.Info)
}
