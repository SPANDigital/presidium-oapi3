package markdown

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var markdownLinkRe = regexp.MustCompile(`\]\(([^)]*)\)`)

// collectLinks returns every markdown link destination found in .md files
// under root, keyed by the file it appears in.
func collectLinks(t *testing.T, root string) map[string][]string {
	t.Helper()
	links := map[string][]string{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".md") {
			return err
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for _, m := range markdownLinkRe.FindAllStringSubmatch(string(content), -1) {
			links[path] = append(links[path], m[1])
		}
		return nil
	})
	require.NoError(t, err)
	return links
}

// Generated links must survive CommonMark parsing (no spaces) and match the
// URLs Hugo renders for the generated directories (lowercase, hyphenated).
func TestGeneratedLinksAreHugoCompatible(t *testing.T) {
	cfg := Config{
		ReferenceURL: "reference",
		ApiName:      "Aggregate API", // deliberately uppercase with a space
		OutputDir:    t.TempDir(),
	}
	ms, err := NewMarkdownService(cfg)
	require.NoError(t, err)
	require.NoError(t, ms.ConvertToMarkdown("./testdata/linkcheck.yaml"))

	// The API name is normalized into the output path, matching link URLs.
	base := filepath.Join(cfg.OutputDir, "content", "reference", "aggregate-api")
	_, err = os.Stat(base)
	require.NoError(t, err, "expected normalized output directory %s", base)

	links := collectLinks(t, base)
	require.NotEmpty(t, links, "expected generated markdown to contain links")

	for file, dests := range links {
		for _, dest := range dests {
			url := strings.TrimPrefix(dest, "{{%baseurl%}}")
			assert.NotContains(t, url, " ", "link with space (breaks CommonMark) in %s: %s", file, dest)
			assert.Equal(t, strings.ToLower(url), url, "uppercase link (Hugo renders lowercase URLs) in %s: %s", file, dest)
		}
	}
}

func TestSecuritySchemeLinks(t *testing.T) {
	cfg := Config{
		ReferenceURL: "reference",
		ApiName:      "linkcheck",
		OutputDir:    t.TempDir(),
	}
	ms, err := NewMarkdownService(cfg)
	require.NoError(t, err)
	require.NoError(t, ms.ConvertToMarkdown("./testdata/linkcheck.yaml"))

	opPath := filepath.Join(cfg.OutputDir, "content", "reference", "linkcheck",
		"operations", "data query - service", "get_users.md")
	content, err := os.ReadFile(opPath)
	require.NoError(t, err)

	// Path must include the API name segment and the lowercase directory name;
	// the anchor must be kebab-case to match Presidium's rendered heading ids.
	assert.Contains(t, string(content),
		"[bearerAuth]({{%baseurl%}}/reference/linkcheck/components/securityschemas/#bearer-auth)")
	assert.Contains(t, string(content),
		"[legacyAccessToken]({{%baseurl%}}/reference/linkcheck/components/securityschemas/#legacy-access-token)")
}

func TestSchemaLinks(t *testing.T) {
	cfg := Config{
		ReferenceURL: "reference",
		ApiName:      "linkcheck",
		OutputDir:    t.TempDir(),
	}
	ms, err := NewMarkdownService(cfg)
	require.NoError(t, err)
	require.NoError(t, ms.ConvertToMarkdown("./testdata/linkcheck.yaml"))

	opPath := filepath.Join(cfg.OutputDir, "content", "reference", "linkcheck",
		"operations", "data query - service", "get_users.md")
	content, err := os.ReadFile(opPath)
	require.NoError(t, err)
	assert.Contains(t, string(content),
		"[User]({{%baseurl%}}/reference/linkcheck/components/schemas/#user)")

	schemaPath := filepath.Join(cfg.OutputDir, "content", "reference", "linkcheck",
		"components", "schemas", "user.md")
	content, err = os.ReadFile(schemaPath)
	require.NoError(t, err)
	assert.Contains(t, string(content),
		"[HomeAddress]({{%baseurl%}}/reference/linkcheck/components/schemas/#home-address)")
}
