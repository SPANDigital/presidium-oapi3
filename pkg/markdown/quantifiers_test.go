package markdown

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const oneOfSpec = `openapi: 3.0.2
info:
  title: Quantifier Test
  version: "1.0.0"
paths:
  /results:
    get:
      operationId: getResult
      tags:
        - Results
      responses:
        '200':
          description: ok
          content:
            application/json:
              schema:
                oneOf:
                  - $ref: '#/components/schemas/AlphaResult'
                  - $ref: '#/components/schemas/BetaResult'
                  - $ref: '#/components/schemas/GammaResult'
components:
  schemas:
    AlphaResult:
      type: object
      properties:
        value:
          type: string
    BetaResult:
      type: object
      properties:
        value:
          type: integer
    GammaResult:
      type: object
      properties:
        value:
          type: boolean
`

// OneOf/AnyOf/AllOf alternatives render inside table cells; joining them with
// bare commas produces one long unbreakable run that forces horizontal
// scrolling. They must be <br>-separated, with no trailing separator.
func TestListQuantifiersJoinAlternativesWithLineBreaks(t *testing.T) {
	config := getConfig(t)
	ms, err := NewMarkdownService(config)
	require.NoError(t, err)
	require.NoError(t, ms.ConvertToMarkdown(writeTempSpec(t, oneOfSpec)))

	matches, err := filepath.Glob(filepath.Join(
		config.OutputDir, "content", "*", "*", "operations", "results", "get_result.md",
	))
	require.NoError(t, err)
	require.Len(t, matches, 1)
	raw, err := os.ReadFile(matches[0])
	require.NoError(t, err)
	content := string(raw)

	assert.Contains(t, content, "AlphaResult")
	assert.Contains(t, content, ")<br>[", "alternatives must be <br>-separated")
	assert.NotContains(t, content, "),", "alternatives must not be comma-joined (unbreakable in table cells)")
}
