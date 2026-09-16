package markdown

import (
	"net/http"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
)

// methodOrder is the conventional REST documentation order used to sequence a
// path's operations. kin-openapi stores a PathItem's operations in a map, so
// the spec's own method order is not recoverable.
var methodOrder = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodHead,
	http.MethodOptions,
	http.MethodTrace,
	http.MethodConnect,
}

// specPathOrder returns the spec's paths in document order by re-reading the
// file: kin-openapi stores paths in a map, so the original key order is lost
// by the time the document is loaded. yaml.v3 parses JSON specs too. Returns
// nil when the order cannot be recovered (caller falls back to routing-match
// order).
func specPathOrder(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return nil
	}
	if len(root.Content) == 0 || root.Content[0].Kind != yaml.MappingNode {
		return nil
	}
	doc := root.Content[0]
	for i := 0; i+1 < len(doc.Content); i += 2 {
		if doc.Content[i].Value != "paths" {
			continue
		}
		pathsNode := doc.Content[i+1]
		if pathsNode.Kind != yaml.MappingNode {
			return nil
		}
		order := make([]string, 0, len(pathsNode.Content)/2)
		for j := 0; j+1 < len(pathsNode.Content); j += 2 {
			order = append(order, pathsNode.Content[j].Value)
		}
		return order
	}
	return nil
}

// orderedPaths returns the swagger's paths in spec document order when
// recoverable, appending any paths the document scan missed — and everything,
// when the spec order is unavailable — in routing-match order.
func orderedPaths(swagger *openapi3.T, filename string) []string {
	seen := map[string]bool{}
	ordered := make([]string, 0, swagger.Paths.Len())
	for _, p := range specPathOrder(filename) {
		if swagger.Paths.Value(p) != nil && !seen[p] {
			ordered = append(ordered, p)
			seen[p] = true
		}
	}
	for _, p := range swagger.Paths.InMatchingOrder() {
		if !seen[p] {
			ordered = append(ordered, p)
			seen[p] = true
		}
	}
	return ordered
}
