package markdown

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Pinned so the test guards a known contract between generated links and the
// theme's rendered URLs/anchors. Bump deliberately when the theme moves.
const (
	layoutsModule = "github.com/spandigital/presidium-layouts-base@v0.23.2"
	stylingModule = "github.com/spandigital/presidium-styling-base@v0.11.0"
)

const renderTestEnv = "PRESIDIUM_RENDER_TEST"

// TestRenderedSiteHasNoBrokenLinks converts the test spec, renders it with
// Hugo and the pinned Presidium theme, and asserts every internal link and
// anchor in the rendered HTML resolves. It needs the hugo binary and network
// access for module downloads, so it only runs when PRESIDIUM_RENDER_TEST is
// set (see the render-test Makefile target).
func TestRenderedSiteHasNoBrokenLinks(t *testing.T) {
	if os.Getenv(renderTestEnv) == "" {
		t.Skipf("set %s=1 to run (requires hugo and network access)", renderTestEnv)
	}
	if _, err := exec.LookPath("hugo"); err != nil {
		t.Skip("hugo binary not found in PATH")
	}

	site := t.TempDir()
	cfg := Config{
		ReferenceURL: "reference",
		ApiName:      "Aggregate API", // exercises URL normalization end to end
		OutputDir:    site,
		SortFilePath: true,
	}
	ms, err := NewMarkdownService(cfg)
	require.NoError(t, err)
	require.NoError(t, ms.ConvertToMarkdown("./testdata/linkcheck.yaml"))

	writeSiteScaffolding(t, site)

	runHugo(t, site, "mod", "get", layoutsModule)
	runHugo(t, site, "mod", "get", stylingModule)
	runHugo(t, site, "--logLevel", "error")

	broken := checkLinks(t, filepath.Join(site, "public"))
	require.Empty(t, broken, "broken links in rendered site:\n%s", strings.Join(broken, "\n"))
}

func writeSiteScaffolding(t *testing.T, site string) {
	t.Helper()
	files := map[string]string{
		"go.mod": "module example.com/presidium-oapi3-render-test\n\ngo 1.22\n",
		"config.yml": `title: "Render Test"
pluralizelisttitles: false
params:
  sortByFilePath: true
markup:
  goldmark:
    renderer:
      Unsafe: true
menu:
  main:
    - identifier: reference
      name: Reference
      url: /reference/
      weight: 1
module:
  imports:
    - path: github.com/spandigital/presidium-styling-base
    - path: github.com/spandigital/presidium-layouts-base
enableInlineShortcodes: true
`,
		"content/_index.md":           "---\ntitle: Home\n---\n",
		"content/reference/_index.md": "---\ntitle: API Reference\nurl: /reference/\n---\n",
	}
	for name, content := range files {
		path := filepath.Join(site, name)
		require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
		require.NoError(t, os.WriteFile(path, []byte(content), 0o644))
	}
}

func runHugo(t *testing.T, dir string, args ...string) {
	t.Helper()
	cmd := exec.Command("hugo", args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, "hugo %s failed:\n%s", strings.Join(args, " "), out)
}

var (
	hrefRe = regexp.MustCompile(`href="([^"]+)"`)
	idRe   = regexp.MustCompile(`id="([^"]+)"`)
)

var assetExtensions = []string{".css", ".js", ".xml", ".json", ".svg", ".png", ".ico"}

// checkLinks walks the rendered site and returns every site-absolute href
// whose target page or anchor does not exist.
func checkLinks(t *testing.T, public string) []string {
	t.Helper()

	served := map[string]bool{}   // URL path -> exists
	anchors := map[string]map[string]bool{} // page URL -> ids
	type page struct {
		url  string
		html string
	}
	var pages []page

	err := filepath.Walk(public, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		rel, err := filepath.Rel(public, path)
		if err != nil {
			return err
		}
		url := "/" + filepath.ToSlash(rel)
		served[url] = true
		if filepath.Base(path) != "index.html" {
			return nil
		}
		dirURL := strings.TrimSuffix(url, "index.html")
		served[dirURL] = true
		served[strings.TrimSuffix(dirURL, "/")] = true
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		html := string(content)
		ids := map[string]bool{}
		for _, m := range idRe.FindAllStringSubmatch(html, -1) {
			ids[m[1]] = true
		}
		anchors[dirURL] = ids
		pages = append(pages, page{url: dirURL, html: html})
		return nil
	})
	require.NoError(t, err)

	var broken []string
	for _, p := range pages {
		for _, m := range hrefRe.FindAllStringSubmatch(p.html, -1) {
			href := m[1]
			if !strings.HasPrefix(href, "/") || strings.HasPrefix(href, "//") {
				continue
			}
			target, fragment, _ := strings.Cut(href, "#")
			if isAsset(target) {
				continue
			}
			if !served[target] && !served[strings.TrimSuffix(target, "/")] {
				broken = append(broken, fmt.Sprintf("%s -> %s (missing page)", p.url, href))
				continue
			}
			if fragment == "" {
				continue
			}
			key := target
			if !strings.HasSuffix(key, "/") {
				key += "/"
			}
			if ids, ok := anchors[key]; ok && !ids[fragment] {
				broken = append(broken, fmt.Sprintf("%s -> %s (missing anchor)", p.url, href))
			}
		}
	}
	return broken
}

func isAsset(path string) bool {
	for _, ext := range assetExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}
