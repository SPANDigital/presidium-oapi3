package markdown

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
)

const orderedSpec = `openapi: 3.0.2
info:
  title: Order Test
  version: "1.0.0"
paths:
  /zebras:
    get:
      operationId: listZebras
      responses:
        '200':
          description: ok
    post:
      operationId: createZebra
      responses:
        '201':
          description: created
  /apples/{id}/special:
    get:
      operationId: getSpecialApple
      responses:
        '200':
          description: ok
  /apples/{id}:
    get:
      operationId: getApple
      responses:
        '200':
          description: ok
`

func writeTempSpec(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "spec.yaml")
	assert.NoError(t, os.WriteFile(path, []byte(content), 0o600))
	return path
}

func TestSpecPathOrderReturnsDocumentOrder(t *testing.T) {
	path := writeTempSpec(t, orderedSpec)
	// Document order, not alphabetical and not routing-match order.
	assert.Equal(
		t,
		[]string{"/zebras", "/apples/{id}/special", "/apples/{id}"},
		specPathOrder(path),
	)
}

func TestSpecPathOrderUnreadableFile(t *testing.T) {
	assert.Nil(t, specPathOrder(filepath.Join(t.TempDir(), "missing.yaml")))
}

func TestOrderedPathsFollowsSpecWithFallback(t *testing.T) {
	path := writeTempSpec(t, orderedSpec)
	loader := &openapi3.Loader{Context: context.Background()}
	swagger, err := loader.LoadFromFile(path)
	assert.NoError(t, err)

	assert.Equal(
		t,
		[]string{"/zebras", "/apples/{id}/special", "/apples/{id}"},
		orderedPaths(swagger, path),
	)

	// Unreadable source file: every path still appears, via matching order.
	fallback := orderedPaths(swagger, filepath.Join(t.TempDir(), "missing.yaml"))
	assert.ElementsMatch(
		t,
		[]string{"/zebras", "/apples/{id}/special", "/apples/{id}"},
		fallback,
	)
}

func TestGetWeightedFilenamePadsForLexicographicOrder(t *testing.T) {
	assert.Equal(t, "0007-getApple", GetWeightedFilename(7, "getApple"))
	assert.Equal(t, "0123-getApple", GetWeightedFilename(123, "getApple"))
}
