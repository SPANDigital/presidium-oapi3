// Code generated for package tpl by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/index.gomd
// templates/info.gomd
// templates/operation.gomd
// templates/partials/content.gomd
// templates/partials/headers.gomd
// templates/partials/parameters.gomd
// templates/partials/request_body.gomd
// templates/partials/responses.gomd
// templates/partials/schema.gomd
// templates/partials/security.gomd
// templates/responses.gomd
// templates/schemas.gomd
// templates/tag.gomd
package tpl

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// templatesIndexGomd reads file data from disk. It returns an error on failure.
func templatesIndexGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/index.gomd"
	name := "templates/index.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesInfoGomd reads file data from disk. It returns an error on failure.
func templatesInfoGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/info.gomd"
	name := "templates/info.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesOperationGomd reads file data from disk. It returns an error on failure.
func templatesOperationGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/operation.gomd"
	name := "templates/operation.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesPartialsContentGomd reads file data from disk. It returns an error on failure.
func templatesPartialsContentGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/partials/content.gomd"
	name := "templates/partials/content.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesPartialsHeadersGomd reads file data from disk. It returns an error on failure.
func templatesPartialsHeadersGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/partials/headers.gomd"
	name := "templates/partials/headers.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesPartialsParametersGomd reads file data from disk. It returns an error on failure.
func templatesPartialsParametersGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/partials/parameters.gomd"
	name := "templates/partials/parameters.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesPartialsRequest_bodyGomd reads file data from disk. It returns an error on failure.
func templatesPartialsRequest_bodyGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/partials/request_body.gomd"
	name := "templates/partials/request_body.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesPartialsResponsesGomd reads file data from disk. It returns an error on failure.
func templatesPartialsResponsesGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/partials/responses.gomd"
	name := "templates/partials/responses.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesPartialsSchemaGomd reads file data from disk. It returns an error on failure.
func templatesPartialsSchemaGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/partials/schema.gomd"
	name := "templates/partials/schema.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesPartialsSecurityGomd reads file data from disk. It returns an error on failure.
func templatesPartialsSecurityGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/partials/security.gomd"
	name := "templates/partials/security.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesResponsesGomd reads file data from disk. It returns an error on failure.
func templatesResponsesGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/responses.gomd"
	name := "templates/responses.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesSchemasGomd reads file data from disk. It returns an error on failure.
func templatesSchemasGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/schemas.gomd"
	name := "templates/schemas.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesTagGomd reads file data from disk. It returns an error on failure.
func templatesTagGomd() (*asset, error) {
	path := "/Users/meyer/Documents/Workspace/Span/presidium-oapi3/templates/tag.gomd"
	name := "templates/tag.gomd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/index.gomd":                 templatesIndexGomd,
	"templates/info.gomd":                  templatesInfoGomd,
	"templates/operation.gomd":             templatesOperationGomd,
	"templates/partials/content.gomd":      templatesPartialsContentGomd,
	"templates/partials/headers.gomd":      templatesPartialsHeadersGomd,
	"templates/partials/parameters.gomd":   templatesPartialsParametersGomd,
	"templates/partials/request_body.gomd": templatesPartialsRequest_bodyGomd,
	"templates/partials/responses.gomd":    templatesPartialsResponsesGomd,
	"templates/partials/schema.gomd":       templatesPartialsSchemaGomd,
	"templates/partials/security.gomd":     templatesPartialsSecurityGomd,
	"templates/responses.gomd":             templatesResponsesGomd,
	"templates/schemas.gomd":               templatesSchemasGomd,
	"templates/tag.gomd":                   templatesTagGomd,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"index.gomd":     &bintree{templatesIndexGomd, map[string]*bintree{}},
		"info.gomd":      &bintree{templatesInfoGomd, map[string]*bintree{}},
		"operation.gomd": &bintree{templatesOperationGomd, map[string]*bintree{}},
		"partials": &bintree{nil, map[string]*bintree{
			"content.gomd":      &bintree{templatesPartialsContentGomd, map[string]*bintree{}},
			"headers.gomd":      &bintree{templatesPartialsHeadersGomd, map[string]*bintree{}},
			"parameters.gomd":   &bintree{templatesPartialsParametersGomd, map[string]*bintree{}},
			"request_body.gomd": &bintree{templatesPartialsRequest_bodyGomd, map[string]*bintree{}},
			"responses.gomd":    &bintree{templatesPartialsResponsesGomd, map[string]*bintree{}},
			"schema.gomd":       &bintree{templatesPartialsSchemaGomd, map[string]*bintree{}},
			"security.gomd":     &bintree{templatesPartialsSecurityGomd, map[string]*bintree{}},
		}},
		"responses.gomd": &bintree{templatesResponsesGomd, map[string]*bintree{}},
		"schemas.gomd":   &bintree{templatesSchemasGomd, map[string]*bintree{}},
		"tag.gomd":       &bintree{templatesTagGomd, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
