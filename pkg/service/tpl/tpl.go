// Code generated for package tpl by go-bindata DO NOT EDIT. (@generated)
// sources:
// pkg/templates/info.gomd
// pkg/templates/operation.gomd
// pkg/templates/partials/content.gomd
// pkg/templates/partials/headers.gomd
// pkg/templates/partials/parameters.gomd
// pkg/templates/partials/request_body.gomd
// pkg/templates/partials/responses.gomd
// pkg/templates/partials/schema.gomd
// pkg/templates/partials/security.gomd
// pkg/templates/schemas.gomd
// pkg/templates/tag.gomd
package tpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pkgTemplatesInfoGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcf\xcd\x8a\x83\x30\x14\x05\xe0\x7d\x9e\xe2\x80\xbb\x40\x7c\x00\x11\x37\x33\xb3\x93\x19\x70\xb4\x7b\xb1\x11\x02\xfe\x94\x18\xba\x11\xdf\xbd\xdc\x6b\xea\x4f\xda\x45\x5d\xe5\x9e\x7b\xfc\xa2\x4a\x29\xe1\x8c\xeb\x74\x82\x79\x46\x5c\xd2\x11\xcb\x22\x28\x17\xd1\x29\xe3\xe1\xa2\xed\x64\xc6\x81\x2a\x82\xe6\x6f\x3d\x35\xd6\xdc\xdc\x9e\x29\x98\x16\x71\xa9\x6d\x3f\xfd\xb5\xff\xda\xde\x4d\xc3\xa2\x00\x80\x28\x02\x6f\x30\xb6\xf0\xbb\x84\x17\x29\xdf\x14\xbe\x94\xad\xa0\x1e\xae\x24\x3c\xed\xaf\x71\x70\x75\xe3\x8e\xa8\x8f\x78\x0c\x6a\xf1\x6f\xdd\xf3\x07\xc0\x3f\x52\x52\x22\xe5\xfa\xc7\x61\x6b\x23\xfc\xa5\xef\xc4\x9f\xbe\x36\xdd\x99\xe4\x28\x34\xb7\xde\x27\x68\x55\xe4\x67\xb2\x2a\x72\x02\xd3\xa3\xb8\x96\xb2\x17\x70\x3f\x3e\x02\x00\x00\xff\xff\xbc\xf2\x49\x2e\xd1\x01\x00\x00")

func pkgTemplatesInfoGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesInfoGomd,
		"pkg/templates/info.gomd",
	)
}

func pkgTemplatesInfoGomd() (*asset, error) {
	bytes, err := pkgTemplatesInfoGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/info.gomd", size: 465, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgTemplatesOperationGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x3d\x6a\xc4\x30\x10\x85\x7b\x9f\x62\x70\x69\x90\x0f\x90\x22\x45\xd8\x26\x90\x3f\x76\x2f\x60\xe1\x7d\x21\x82\x95\xa5\x8c\xc6\x85\x11\xba\xfb\x32\x46\x6b\x17\xeb\xf2\xbd\xf7\xf1\x49\x63\x8c\x69\xc4\xc9\x0d\x2f\x94\x33\xf5\xdf\x11\x6c\xc5\x85\xe9\xfd\x44\xa5\x34\xba\x0e\xc3\xd0\xe8\xf4\x09\xf9\x0b\x57\x2a\x65\x05\xbf\xac\x87\x12\xba\xea\xec\x7e\xa9\x3f\x21\x32\x46\x2b\x50\xaa\xeb\xf6\xd8\x75\x39\x1b\xc2\xa4\xfd\x03\xbe\xcc\xde\x5b\x5e\xd4\xf7\x9a\xf3\x23\xae\xf6\x0d\xdd\xc5\x69\x64\x17\xf5\x5f\x54\x8a\x3e\xff\xd4\x6c\x72\x43\x02\x1f\x6f\x56\x40\x6d\xb4\x6c\x3d\x04\x9c\x3e\x5c\x92\x96\xfa\x9f\xad\x78\x86\x19\xff\x33\x92\xbc\x85\xeb\xd2\x52\x7f\xde\xd3\x11\x9a\x62\x98\x12\xaa\xf5\x5c\xe3\x81\x34\x61\x9c\xd9\xc9\x52\xc9\x4b\x8d\xeb\x71\xf7\x00\x00\x00\xff\xff\xb3\xdd\x9d\x53\x7b\x01\x00\x00")

func pkgTemplatesOperationGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesOperationGomd,
		"pkg/templates/operation.gomd",
	)
}

func pkgTemplatesOperationGomd() (*asset, error) {
	bytes, err := pkgTemplatesOperationGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/operation.gomd", size: 379, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgTemplatesPartialsContentGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\x6e\x83\x30\x10\x44\xef\x7c\xc5\xc8\xe4\x84\x9a\x70\xaf\xd4\x53\xdb\x63\x15\x29\xcd\x07\xc4\xc5\x43\x83\x54\x96\x14\xdc\x43\x84\xf6\xdf\x2b\xc0\xc1\x56\xaa\x72\x42\x9e\xe7\xdd\x79\x1e\xc7\x2d\xca\x02\xcf\x9d\x78\x8a\x87\x67\x7b\xf9\xb2\x9e\x28\x4a\x6c\x55\xb3\x29\x76\xac\x1b\x21\x4c\xb5\x30\x46\x35\x03\x80\x29\xea\xad\x7c\x12\x1b\xb1\x2d\x1f\xb0\x69\xe9\x1a\x7b\xbc\x5e\x88\xc7\x27\xec\xa0\x9a\xcd\xe0\xf4\xe5\x79\x8e\xd3\x38\xce\xa4\xea\x29\x06\xd3\x94\x75\xa9\x59\x27\x98\x64\x5a\xb2\x8f\xe2\x42\xa9\xf8\x57\x16\x78\x15\x87\xae\xfe\x57\xe2\x86\x1d\x28\x8e\xfd\x00\x8b\xb7\xb5\x69\xe2\x59\x16\x38\xee\x5f\xf6\x01\x83\x3f\x13\x73\x07\xdc\x83\xb7\x07\x89\x6d\x93\x8a\x4d\x8d\xdd\x7b\x75\x66\x6b\x11\x4e\x97\x24\xb1\x1c\xe6\xd8\xc0\x35\x95\x87\x59\xe0\x03\x6b\x73\x7f\x31\x78\x22\x2a\xe3\x8f\x73\xcf\xef\x1f\x0e\x1e\x1f\x9d\xbb\x86\x8e\xbf\x01\x00\x00\xff\xff\xd2\xed\x50\x80\xd3\x01\x00\x00")

func pkgTemplatesPartialsContentGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesPartialsContentGomd,
		"pkg/templates/partials/content.gomd",
	)
}

func pkgTemplatesPartialsContentGomd() (*asset, error) {
	bytes, err := pkgTemplatesPartialsContentGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/partials/content.gomd", size: 467, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgTemplatesPartialsHeadersGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x31\x4f\xf3\x30\x14\xdc\xfb\x2b\x4e\x51\x87\xef\xb3\x68\xcb\x8c\x80\x09\xa4\x0e\x85\xa1\x44\xec\x26\x7e\x6e\x2d\x25\x4e\x70\x5c\x21\xf4\xe2\xff\x8e\x9c\xa6\xc1\xa1\x74\x20\xd3\xd3\xbd\xbb\xe4\xdd\x5d\x98\x17\x58\x09\xac\x49\x2a\x72\xc8\xa9\x6a\x4a\xe9\x09\x62\x85\x45\x08\x33\x66\x45\xda\x58\x42\xb6\xef\x09\x59\x8f\x02\x40\xd4\xcd\x8f\x20\x6e\xee\xb0\x7c\x95\xe5\x81\x90\x2e\x1d\xe9\x7e\xb3\x25\x9d\xe2\x46\x8f\xba\x01\x3d\xdf\x2c\xb7\xf4\x7e\x30\x8e\x14\x42\xc0\xed\x9b\xbb\x87\x70\x03\x22\xc0\x4c\x56\x85\x00\xe6\x54\xf1\x40\x8d\xa3\x42\xfa\x54\xa3\x46\xec\x5b\xd5\x81\xd9\x9f\x5c\x66\x6d\xb1\xa7\x4a\xe6\x9f\x0d\x65\xe3\x9b\x5e\x7a\x0c\x3d\x57\xc4\xbb\x7c\xbd\xce\x9f\x36\xcf\xf4\xb1\x31\x96\xda\xe4\x8b\x6d\xe1\x4c\xe3\x4d\x6d\x63\x2a\x02\xdd\xe8\x91\x6c\xbc\x62\x36\x8c\xc3\xb4\x12\x78\xb4\x0a\xb5\xbe\x14\xf6\x6c\xd2\xc5\xc6\xb4\x1e\xfe\xac\x0f\x4c\x0b\x89\xac\x69\x29\x46\x63\xe7\xf1\xaf\x24\x8b\xe5\x7f\x5c\xff\x0c\xd9\x49\xbb\x23\xcc\xad\xac\xe8\x6a\x52\x60\xc2\x8b\x4f\x07\x21\x98\x8f\x44\x84\x20\xc4\x24\xb7\xd3\xdf\x70\xa1\xc9\xc1\xff\x9f\xe2\xf8\xcd\xef\x57\x00\x00\x00\xff\xff\xd7\x37\x4d\x95\x9d\x02\x00\x00")

func pkgTemplatesPartialsHeadersGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesPartialsHeadersGomd,
		"pkg/templates/partials/headers.gomd",
	)
}

func pkgTemplatesPartialsHeadersGomd() (*asset, error) {
	bytes, err := pkgTemplatesPartialsHeadersGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/partials/headers.gomd", size: 669, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgTemplatesPartialsParametersGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xc1\xae\xd3\x30\x10\xbc\xe7\x2b\x46\x69\x0f\x60\xd1\x94\x33\x02\x4e\x20\x51\xa9\x54\xa8\x54\xdc\x4d\xb2\x69\x2c\x35\x4e\xb0\x5d\x0a\x4a\xfc\xef\x4f\x9b\xba\x89\xd3\x57\xbd\xe7\x43\x62\xed\xec\xec\x7a\x67\xb6\xeb\x56\x58\x0b\xb8\x4a\x59\x38\xaa\xdb\x93\x74\x04\x43\x39\xa9\xbf\x64\x21\xf1\x43\x1a\x59\x93\x23\xb3\xa7\x12\x97\x4a\xe5\x15\xf2\x46\x3b\xa9\x34\xa3\x1c\x94\xba\x88\xf3\x20\xd6\x58\x79\x9f\x70\xe1\x82\x4a\xa5\x09\x69\x7b\x03\x53\x78\x9f\x00\x00\xa3\xcb\x31\x8c\x0f\x9f\x90\xfd\x92\xa7\x33\xcd\x70\x43\xe5\x80\x70\x97\x28\xae\xca\x98\x1a\x00\x3e\x3d\x84\xe8\xba\x08\xcc\x76\xb2\xe6\x92\x42\x3c\xc0\x36\xfa\x8a\x74\xdd\xbc\x62\xb6\xa7\x3f\x67\x65\xa8\x80\xf7\xf8\xf8\xdb\x7c\x86\x30\x21\xc2\xc9\xa4\x0b\xef\x9f\x93\xbe\x50\x6b\x28\x97\x2e\xa6\x15\x63\x6c\x22\xf6\xc3\x08\xa3\xd2\xa9\xcd\x2b\xaa\xe5\xe1\x7f\x4b\x69\x5c\xee\xe7\x10\x46\x20\xc0\x35\xdf\x0e\xdf\xb7\x3b\xba\x6c\x95\x26\x3b\xef\x6b\x73\xa3\x5a\xa7\x1a\xcd\xc9\xa3\x48\xa4\xf9\x21\x49\xb8\x86\xdb\x5a\xe0\xab\x2e\xd0\x94\x68\xef\xed\x4a\x5e\x5b\x84\xad\xb2\x8e\x99\xb3\x85\x78\xc9\x6b\xcb\x8c\xf4\xce\xb8\xa3\xc3\x9b\x13\x69\x64\x6f\xf1\x9e\xa1\xd1\xbb\xc5\x02\xac\x3b\x59\x37\x75\xb0\x49\x64\xed\x60\x65\x8f\x8d\x46\x0f\x96\x0b\x3d\xa2\xd9\xc3\xe4\x43\xea\x6a\x38\xfd\xf8\x19\x7f\xe1\x4c\xa9\xfc\x28\x23\xf5\x91\xb0\x54\xba\xa0\x7f\xef\xee\x77\x32\x5a\xae\x5b\xfe\x64\x5d\xb4\xd5\x8f\xf7\x31\xf2\x21\x68\x30\xd9\x12\x1b\x34\xd9\x32\x8d\x7e\xd5\xfb\xaa\xef\x53\x00\x00\x00\xff\xff\x52\x27\xd1\xd7\xa4\x03\x00\x00")

func pkgTemplatesPartialsParametersGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesPartialsParametersGomd,
		"pkg/templates/partials/parameters.gomd",
	)
}

func pkgTemplatesPartialsParametersGomd() (*asset, error) {
	bytes, err := pkgTemplatesPartialsParametersGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/partials/parameters.gomd", size: 932, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgTemplatesPartialsRequest_bodyGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x53\x4d\x6f\xeb\x20\x10\xbc\xfb\x57\x8c\x9c\x9c\xac\x67\xe7\x1e\x29\x97\xf7\x5e\x8f\xed\xc1\xad\x7a\xa7\x66\x89\x91\x62\x9c\xda\xa4\x55\x84\xf9\xef\x15\xfe\x2a\x50\x97\x13\xbb\x33\xb3\xec\x32\x60\x4c\x8e\x43\x06\x5d\xcb\x1e\x9a\x9a\xeb\x85\x69\x42\x47\x15\xc9\x0f\xea\xc1\x50\xd2\xf9\x46\xbd\xfe\xdb\xf2\x7b\x49\x02\x9f\xb5\xac\x6a\x54\xad\xd2\x4c\xaa\x09\x17\x60\x8a\x83\xe1\x95\x5d\x6e\x84\xec\x80\xdc\xda\xc4\x95\xe5\x24\xa4\x22\xa4\x1d\xbd\x2f\x25\x52\x6b\x13\x00\x70\xb0\x14\x28\xe0\xc5\x7b\x8f\x87\xe3\x09\xc5\x58\x30\x24\x88\x11\x28\x49\xcc\x69\xaf\x56\x20\xf7\x60\xb7\x76\x3b\x94\x13\x0a\x07\x27\x01\x68\x4c\x2c\x2f\x1c\x59\x76\xc4\x61\x6d\x96\x2d\x41\x96\x19\x43\x97\x9e\x5c\xee\xa9\xd5\x08\xf2\x8a\x5b\x1b\x96\x1d\xf0\x48\x5c\xb2\x97\xfb\x95\x30\xe0\x3f\xf5\x55\x27\xaf\x5a\xb6\x0a\x03\x9e\xab\x9a\x1a\x86\x21\x14\xe4\xdf\xcb\xdf\x7b\xd1\x10\x35\x9e\xa3\x63\xea\x4c\xd8\x2b\xd6\xd0\x1f\xec\x9b\xf5\xc4\xe3\x29\x9c\xe8\x5f\xab\x34\x29\x1d\x5f\x8c\x7f\x7f\xab\xb8\x98\xdb\xdb\xe0\xae\x56\xf4\x23\xa5\x9c\x0c\xf9\x21\x8d\x1c\xfa\x45\xbd\x2d\xf5\x5d\x8f\xd7\x00\x63\xc6\x59\xad\x1d\xf7\xe1\x8c\xfe\x1d\x2f\x84\xf5\x49\xa7\xd3\x99\xee\xa4\x74\x73\xd6\xc8\x8d\xa5\xd7\xc9\xd9\xed\x6c\x12\xa5\xfc\xd7\x3c\x87\xd1\xf6\x90\xe1\x41\x71\xb4\x02\x73\xe3\x78\x73\xcf\x75\xfa\x34\x5f\x01\x00\x00\xff\xff\x2c\x10\xba\x6d\x8b\x03\x00\x00")

func pkgTemplatesPartialsRequest_bodyGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesPartialsRequest_bodyGomd,
		"pkg/templates/partials/request_body.gomd",
	)
}

func pkgTemplatesPartialsRequest_bodyGomd() (*asset, error) {
	bytes, err := pkgTemplatesPartialsRequest_bodyGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/partials/request_body.gomd", size: 907, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgTemplatesPartialsResponsesGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x8b\xdb\x30\x10\xbd\xfb\x57\x3c\xec\x1c\xda\xb0\x71\x7a\x2e\xf4\xb4\x2d\xe4\x90\x5d\x4a\x36\xf4\x6e\xe2\xf1\xae\x21\x91\x83\xa5\xa5\x14\xcb\xff\xbd\x58\x92\xf5\x65\x25\xb4\x1b\x08\x48\xd6\x9b\xd1\xcc\x7b\x6f\x34\x0c\x1b\x6c\xd7\x38\x10\xbf\x76\x8c\x13\x04\x5d\xae\xe7\x4a\x10\xd6\x5b\x6c\xc6\x31\x9b\xce\x6b\x6a\x5a\x46\xc8\x7b\x03\xca\x31\x8e\x19\x00\x4c\x87\xab\x9e\x1a\x7c\xfd\x86\x72\x4e\x51\x1e\xa8\x89\x00\x26\x77\x80\xfa\x55\x9d\xdf\x29\xc0\x71\x51\x89\x77\xae\x50\x2f\x7a\xe9\x9d\xb6\x8d\x4b\x54\x3e\x76\x4c\x10\x13\xf3\xf9\x8c\xe9\x2b\xf6\x4a\x58\xb1\xea\x42\x0f\x58\x5d\xa8\x6e\xab\xe3\x9f\xab\xba\xf7\x6e\xac\x7f\x87\x8d\x2a\x5f\x4e\x6f\x74\xa9\x62\x9c\xab\x56\x1d\x1f\x74\xf3\x8b\xb0\x89\x84\xfb\x91\xe9\x30\xc5\x4a\x22\x50\x62\x18\x0c\x41\xe3\xa8\x76\x10\xdd\xee\xf8\xb4\x7f\xa6\xdf\xfb\x96\x11\xf7\x3a\xfc\x4e\xfc\xd4\xb7\x57\xd1\x76\x0c\x33\x58\x91\x62\x77\x56\xe4\x5c\xd7\x32\x55\x90\x27\x7b\x87\x5c\xb0\x44\xac\x8e\x78\xd7\x5f\xac\x54\x74\xe6\xe4\xd3\x96\xa8\xfd\x4e\xb1\x1b\xfd\xcf\xc2\xe4\x66\x05\xb3\xdc\xae\xf1\x83\xd5\xe8\x9a\xdb\xc6\xcd\x22\x67\xef\x5b\x2e\xfe\xc9\xdd\x13\xd0\x73\xf8\xe4\x8a\x57\x81\x4f\x67\x62\x28\x3f\xe3\x8b\xdf\x59\x51\xd8\xf4\x3c\xf3\xfa\xdd\x1d\x8f\x3f\xf1\xd8\xd5\x04\x09\xbf\x41\x89\x27\x6b\x4a\x09\x43\xb2\x63\x58\x6e\xdc\xcf\x5f\x87\x3b\xbb\x96\x29\xf3\x6b\x9e\x1f\xa2\xa9\x4b\xb9\xdd\x99\xc0\xcd\x75\xdd\x9e\x04\xf2\x83\xdd\xbb\x2c\xb9\x1e\xc9\xdc\x8e\xe9\xd2\x03\x98\x4d\xf0\xe1\x92\xe6\xb7\xc4\xd9\xc3\x7b\x4b\x22\xa0\x06\xe8\x57\x24\x08\x09\x1e\x96\x68\xbc\x67\x21\xc3\xf8\x72\x47\x55\x4d\x3d\x37\xea\x2e\x86\xaf\x28\x8a\x42\xb9\xd6\xb6\x0e\x13\xb1\xc4\x4a\x3c\x4f\x93\x26\x61\x44\x0e\xe4\x5f\xa2\x03\x49\x23\xcd\x93\xaf\x87\x53\xed\x4d\x95\xa0\xdd\x9a\xee\x27\x45\x82\xd1\xe9\xc6\xa7\x68\xcc\xfe\x6b\xe2\xfe\x06\x00\x00\xff\xff\x96\xc8\xb0\x2c\x4a\x06\x00\x00")

func pkgTemplatesPartialsResponsesGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesPartialsResponsesGomd,
		"pkg/templates/partials/responses.gomd",
	)
}

func pkgTemplatesPartialsResponsesGomd() (*asset, error) {
	bytes, err := pkgTemplatesPartialsResponsesGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/partials/responses.gomd", size: 1610, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgTemplatesPartialsSchemaGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x5f\x6f\xdb\x36\x10\x7f\xcf\xa7\x38\xa8\x7d\xb0\x85\xd8\xde\xb3\xb1\x0d\x08\xb0\x0e\x2b\x90\xd6\x85\xd3\xed\x65\x28\x06\x56\x3a\xc5\xdc\x64\xca\x95\xe8\x36\x86\xac\xef\x3e\x50\xfc\x23\x92\x22\x2d\x27\xad\x5e\x12\x93\xc7\x1f\xef\x8e\xbf\xfb\x43\xb6\xed\x02\x56\x29\xf0\x1d\x6d\x80\xe3\xfe\x50\x12\x8e\x50\x63\x86\xf4\x2b\x36\x40\xe0\x21\xdb\xe1\x9e\x6c\xb1\x80\x6f\x3b\x9a\xed\x20\xab\x18\x27\x94\x89\x29\x31\x48\x58\x6e\x84\x20\x5d\xc1\xa2\xeb\x6e\x04\x64\x8e\x05\x65\x08\x49\xd3\xcf\x24\x5d\x77\x03\x00\x20\x66\x5e\xd7\x58\xc0\xfa\x17\x58\x1a\xe4\xa5\x00\xb2\x25\xe4\x22\x4f\xe8\x2f\x52\x1e\xd1\x11\xab\xab\x8a\xbf\x27\x7b\xec\x05\xfb\x7f\xac\x59\x5a\xc8\x9d\xba\xee\xa6\x1f\x93\xe3\xab\x14\xb6\xc8\x72\xac\xa1\xc6\x02\x6b\x64\x19\x36\x90\xae\xd4\x3a\xf1\x9d\xa1\x47\x3a\xc3\xc7\xd3\x41\xfc\xf9\x0d\x9b\xac\xa6\x07\x4e\x2b\x06\xe7\x41\x6a\xd1\x7f\xee\x1f\xf5\x59\x52\x40\x58\xc5\x4e\xfb\x63\x03\x67\x68\xdb\xc1\xbf\xca\x2d\x62\x8b\x04\x96\xd0\x75\x72\x5e\x19\xbe\xb4\xf7\x14\x73\x37\xc6\x2a\x2c\x1b\x63\xa5\x65\xa9\x38\x05\xed\xb5\x08\x48\x1c\xfe\xe7\xcf\xf5\xaf\x62\x1a\x59\xee\x79\xcb\x60\xcf\xf0\x8b\x59\xdd\xfb\x25\xa9\x3e\xff\x8b\x19\x4f\xe6\x66\xf8\x43\x5d\x1d\xb0\xe6\x14\x1b\x07\x64\xd2\xa5\xb0\xc5\x86\xd7\x34\x13\xbf\x1a\xcb\xc3\x93\x5e\x5e\x44\x9c\xae\x55\xaf\x09\x7b\x44\x78\xcd\xc8\x1e\x6f\xe1\xf5\x41\xaa\x77\x12\x5c\x09\xab\x0c\xde\x77\x16\x1e\x63\x92\x55\x9e\x5e\x86\x81\x5f\x7b\x4e\x0a\x48\x8d\xaf\x68\xba\x08\x00\xf6\x2b\xf8\x0e\xb7\x92\xff\xc3\x12\x31\x10\x5b\xa0\x0f\x80\x55\x5c\x2f\x9e\xcb\xe3\xe8\xf7\x1e\x9d\x86\x1c\xb5\x2c\x0b\x01\x8b\xef\xef\xb6\xed\xad\xeb\xba\x4f\xb3\xe5\x72\xb5\x5c\xae\xb2\x6a\x7f\xa8\x18\x32\xde\xac\xa4\x83\x9a\xd5\x2b\xa1\xc2\xc0\x5a\x1d\x70\x89\x15\x7b\x5d\xb7\x68\x5b\x28\xab\x6f\x58\x1b\x77\xcd\x83\xa6\xf4\xdc\x8d\xa9\xe3\x6e\x64\x87\xc7\x70\x74\x11\x17\x09\xde\x5e\x70\xdf\x8c\xb2\x07\x5e\x53\xf6\xf8\x50\xd2\x0c\xcd\xe1\x6f\xf1\xcb\x91\xd6\x98\x4b\xa5\xe7\x26\x10\xd2\x54\xcf\xa4\xe9\x35\xf0\xca\xe3\x5b\x24\xf9\x86\x95\x27\x07\x87\xe4\x20\xc6\x9e\x05\xf4\xfe\x58\x96\xe4\x73\x89\x16\x90\x1e\x72\x71\x22\x9c\xd4\xd9\x40\x52\x53\xf2\xc6\x66\xda\x5c\x25\x03\x5e\xfd\xf1\xf1\xdd\xfd\x7b\xfc\x76\x4f\x19\x36\x7a\xf7\x51\xd2\x98\xda\xce\x22\x87\x15\xc7\x89\xde\x7f\xb4\x4e\x23\xfa\x39\x42\x95\x21\x95\x9c\x49\x59\x02\x65\xa5\xa8\x20\x92\xda\x8d\xae\x2d\xfe\x9a\xef\x8d\x72\xa7\x22\x59\x75\x68\x24\x63\x42\xdd\x2d\x44\x11\xf7\x4b\xbf\xd7\xcf\x0d\x56\xdf\x2b\xfa\x7b\xf5\x0a\xac\x6c\x14\x0e\xa0\x35\xbc\x30\x5a\x21\xb2\x69\xa0\x60\x25\x90\xd3\x8c\x43\x62\xaa\xb2\x1d\x9e\x89\xda\x2c\xa6\xa5\x75\xf4\x13\xc3\x26\x5b\xd0\x02\x46\x95\x87\xd4\x35\x39\x25\x21\x94\x81\x3f\x77\x42\x26\xc2\x99\x11\xf3\x83\x55\x71\xcc\x4e\x11\xa2\x4a\xf4\x2d\xc7\x7d\x13\x63\x4b\x9a\x0a\x3d\xd7\x69\x2a\x52\x6c\x34\xa9\xd9\x48\x42\xc9\x4f\x7d\xb0\x47\x73\x83\x49\x5a\x43\x96\x71\x13\x8c\x2a\xe2\xd3\x59\x46\x21\xd9\x69\xc6\xc9\x30\x57\x01\x45\xa3\x5e\xf5\x20\xa1\x23\xf6\x3a\x97\x30\x5c\x9c\x66\xbe\xc3\x42\x1c\xf2\xbb\x97\x28\x87\x9a\xbe\x20\x24\x26\x16\x03\x2d\x8b\xea\x55\xae\xee\x4e\xc2\x8d\xc8\xd9\x6a\xbb\x7a\x44\x95\xd5\x2f\xb1\xeb\x85\x87\xfc\xd2\x03\x0e\x24\xe8\xe9\xe3\xb5\x16\xc5\x0e\x77\x94\xd1\x37\xc5\x2d\x10\x76\x12\x7f\x2a\x86\x9b\x22\x96\xd2\x69\x01\x8f\x1c\x66\x25\x32\x63\xcc\x46\xc8\xcf\xe1\xa7\x5e\x7e\x44\x21\x55\x26\x37\x0c\x61\x53\xac\x95\x89\x57\x11\xed\x9e\x36\xbc\x27\x82\xcd\xb7\xc1\x50\xb9\xaf\x49\x6c\x56\x16\x0d\xd3\x3b\xa0\xf9\x9d\xb0\x5b\x6b\x1e\x51\xfc\xae\x2c\x7f\xb4\xe2\xfd\xb6\xdf\xa7\xb8\x38\xa9\x29\xc5\xd9\xe9\x87\x2b\x2e\xb6\xbd\x4e\xf1\xc0\x75\xc5\x1e\xb3\x7e\xdb\x53\x8a\x96\x0f\xb6\x36\x83\x96\x16\x25\xc7\xb7\x58\xa3\xbb\x57\xaa\x64\xc5\xf3\x2f\xa1\x01\x3a\xab\xdb\x6c\x23\xdc\x1a\xef\x65\x28\xcb\xf1\xe9\x56\x7b\x64\xeb\xdc\x96\xa3\x3d\xcc\x10\xf9\xde\x85\xfa\x9a\x43\x71\x6a\x92\xba\x8d\xdc\xc6\xcf\x54\x47\x7b\x10\x7f\x01\x6f\x65\xe3\x26\xd1\xd6\x61\x98\x98\x26\xf1\xcc\x2f\xd4\x72\x3a\x8c\x58\x6d\x7a\x46\x8b\x31\x0c\x79\x64\x59\xa5\xf0\x86\xe5\x50\x15\x13\x3c\xd1\xc2\x83\x1b\xe1\xe3\x55\x54\x92\x1e\x0f\xbd\x8a\x5c\x78\x0b\x19\xbd\x80\x84\xdf\x38\x16\xa0\xe9\xca\xfe\x73\xe7\xed\x78\x0f\x94\x6a\x05\x19\xeb\xb8\x7c\x87\x3f\xab\xbd\x09\xd6\xe5\xbe\x49\xb6\xb7\x0a\x95\x83\x2b\x28\xb2\x74\x16\x5e\x2a\x48\x7e\x2d\xf6\x8a\xe6\xef\x55\xbd\x27\x5c\x4c\xcc\xda\xd6\x1d\xec\xba\x39\xb4\xad\x6c\x32\x2e\x93\xc8\x4d\x38\x2e\x9d\x2e\x10\x45\x4b\xda\x25\x77\x44\x38\x9b\x49\x6e\x6d\x1e\xb6\xb2\xee\x50\x0e\x96\x47\x48\x5a\xc0\xf2\xdd\xb1\xe4\xf4\x50\x8a\x22\xd7\xb7\x09\xfa\x37\x54\x45\x9a\xae\x45\xb3\xec\x89\x04\x3b\x10\x83\x46\x9e\x14\x0c\x79\xa2\xfb\xe3\xde\x40\xc8\xf1\x8b\x6b\xdf\x3c\x65\xe5\xb1\xa1\x5f\xd1\x80\x98\x11\xd8\x0f\x70\xbc\x3e\xe2\x84\x12\x94\x29\x25\x28\x73\x94\xa0\xec\x19\x4a\x68\x10\x4b\x89\x01\xee\x0a\x25\xc8\xd3\x3d\xb2\x47\xbe\x73\xfc\x01\x65\x3f\x66\xb9\x65\x10\x9a\xb2\xc9\x81\x93\xaa\xf8\x70\xb6\xd0\x45\xb8\x0f\x84\x73\xac\x95\x85\xea\x87\x46\xb1\xe6\xa6\x2c\x94\xc1\xed\x18\x48\xc5\x90\x65\x9f\x11\x99\x32\xcf\xc6\x52\xd6\xb9\x58\x96\xc8\x45\xac\x3f\x19\xfd\x72\x44\x0b\xce\x1a\xb8\xfa\xec\xc4\x8d\xdc\xb3\xec\x60\x2e\xe9\x96\x79\x46\x6e\xca\x3c\x1b\x50\x99\x17\x00\xb4\xe4\x2e\x02\xde\xe5\x39\x15\x21\x4d\xca\xe1\xe9\xe0\xae\x14\x97\xfa\x5c\xee\x31\x08\x58\xdb\x88\x84\x20\x44\xe2\x5e\x88\x66\xad\x60\x4a\xf2\xd2\x96\x7e\x56\xaf\x2a\x0e\x7d\x7d\x0e\x65\x2c\xf3\x22\xe1\x59\x34\xca\xdf\x35\x1e\x4a\x92\x21\xcc\xe4\x53\xc5\x72\x0e\xc9\x3f\x09\x24\x8b\xc4\x71\x86\xaf\xf8\x62\x94\x04\x7d\x75\xfe\x0f\x00\x00\xff\xff\x9b\xef\x39\xe8\xee\x18\x00\x00")

func pkgTemplatesPartialsSchemaGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesPartialsSchemaGomd,
		"pkg/templates/partials/schema.gomd",
	)
}

func pkgTemplatesPartialsSchemaGomd() (*asset, error) {
	bytes, err := pkgTemplatesPartialsSchemaGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/partials/schema.gomd", size: 6382, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgTemplatesPartialsSecurityGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xbf\x4e\xc3\x40\x0c\xc6\xf7\x3c\xc5\xa7\xb4\x03\x44\xbd\x86\x19\x89\x81\x81\x0d\xb1\x94\x17\x38\x25\x4e\x39\xd4\x38\x55\xe3\x4a\xa0\xe4\xde\x1d\xf9\x9a\xff\x04\x09\x0f\x51\xe2\x9f\xed\xf8\xfb\xdc\x34\x06\x69\x82\x03\x65\xd7\x8b\x93\x6f\xbc\x53\x79\x3e\x59\x21\x24\x29\x8c\xf7\x91\xf2\x9c\x0a\xc7\x84\xb8\xee\x8a\xe2\x81\x5c\x2c\x1f\x09\x5b\xb6\x25\xed\xb0\xad\xb3\xea\x4c\x35\x1e\x9f\xb0\x87\xf7\x51\x8b\xa6\xb9\x31\x78\x8f\xf0\xf5\x59\x39\x1e\xea\xe2\x1d\xe2\x40\xc2\x28\xe2\xbc\x1b\x3a\xbe\xa5\x09\x5e\x38\x47\x55\xfc\xbd\x5f\xb4\x10\xf0\xea\x6a\x59\x11\xf1\x4b\x83\xd6\xe9\xdf\x23\x00\xd0\x19\xae\xb8\x6d\x8d\x2e\xba\xdc\x51\x70\x77\x22\xc6\xfe\x1e\x0f\x8a\x07\xae\xb1\xd9\xe0\xf9\x2a\x1f\xc4\xe2\x32\x2b\xae\xe2\x39\x6e\xdf\x6c\x49\xed\x21\xa8\x6d\xe7\xc4\x18\x63\xc2\xc3\x98\x39\x99\xb8\xea\x38\xa7\x2f\xb5\xb5\xd7\xde\x1b\x8b\x45\x68\x8f\xf4\x92\x27\x57\x1a\x3b\x17\x4d\xa3\xcb\xeb\x99\xff\xdf\x63\xcd\xee\x9f\x00\x00\x00\xff\xff\x45\xb9\xf2\x04\x53\x02\x00\x00")

func pkgTemplatesPartialsSecurityGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesPartialsSecurityGomd,
		"pkg/templates/partials/security.gomd",
	)
}

func pkgTemplatesPartialsSecurityGomd() (*asset, error) {
	bytes, err := pkgTemplatesPartialsSecurityGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/partials/security.gomd", size: 595, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgTemplatesSchemasGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\xd5\xe5\x2a\xc9\x2c\xc9\x49\xb5\x52\xa8\xae\x56\xd0\xf3\x4b\xcc\x4d\x55\xa8\xad\xe5\x02\x09\x73\x55\x57\x2b\x94\xa4\xe6\x16\xe4\x24\x96\xa4\x2a\x28\x15\x27\x67\xa4\xe6\x26\x2a\x29\xa4\x64\x26\x97\x28\x28\x05\x83\x79\x41\xa9\x69\x4a\x0a\x7a\x70\xb6\x82\x12\x48\xbb\x12\xdc\x14\x40\x00\x00\x00\xff\xff\x95\x66\xc9\x36\x5c\x00\x00\x00")

func pkgTemplatesSchemasGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesSchemasGomd,
		"pkg/templates/schemas.gomd",
	)
}

func pkgTemplatesSchemasGomd() (*asset, error) {
	bytes, err := pkgTemplatesSchemasGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/schemas.gomd", size: 92, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgTemplatesTagGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\xd5\xe5\x2a\xc9\x2c\xc9\x49\xb5\x52\xa8\xae\x56\x28\xc9\x77\x4e\xcc\x4d\xcd\x51\xd0\xf3\x4b\xcc\x4d\x55\xa8\xad\xe5\x02\x49\x73\x55\x57\x2b\xe8\xb9\xa4\x16\x27\x17\x65\x16\x94\x64\xe6\xe7\x81\xc4\xb9\xaa\xab\x75\x15\x32\xd3\x14\xf4\x5c\x2b\x4a\x52\x8b\xf2\x12\x73\x5c\xf2\x93\x8b\x41\x12\x0a\x0a\x0a\x0a\xd8\xe4\x30\x0c\x50\x80\x82\x68\x90\xe9\x78\x54\xc6\x6a\x60\x28\x08\x0d\xf2\x51\xa8\xad\xd5\x84\xdb\x95\x9a\x53\x9c\x0a\xb3\x1b\x04\x6c\x70\xe8\xb0\x43\xe8\xc8\x4b\x01\x69\x40\x30\x01\x01\x00\x00\xff\xff\xd8\x77\x01\x1a\x07\x01\x00\x00")

func pkgTemplatesTagGomdBytes() ([]byte, error) {
	return bindataRead(
		_pkgTemplatesTagGomd,
		"pkg/templates/tag.gomd",
	)
}

func pkgTemplatesTagGomd() (*asset, error) {
	bytes, err := pkgTemplatesTagGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/templates/tag.gomd", size: 263, mode: os.FileMode(420), modTime: time.Unix(1591022944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
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
	"pkg/templates/info.gomd":                  pkgTemplatesInfoGomd,
	"pkg/templates/operation.gomd":             pkgTemplatesOperationGomd,
	"pkg/templates/partials/content.gomd":      pkgTemplatesPartialsContentGomd,
	"pkg/templates/partials/headers.gomd":      pkgTemplatesPartialsHeadersGomd,
	"pkg/templates/partials/parameters.gomd":   pkgTemplatesPartialsParametersGomd,
	"pkg/templates/partials/request_body.gomd": pkgTemplatesPartialsRequest_bodyGomd,
	"pkg/templates/partials/responses.gomd":    pkgTemplatesPartialsResponsesGomd,
	"pkg/templates/partials/schema.gomd":       pkgTemplatesPartialsSchemaGomd,
	"pkg/templates/partials/security.gomd":     pkgTemplatesPartialsSecurityGomd,
	"pkg/templates/schemas.gomd":               pkgTemplatesSchemasGomd,
	"pkg/templates/tag.gomd":                   pkgTemplatesTagGomd,
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
	"pkg": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"info.gomd":      &bintree{pkgTemplatesInfoGomd, map[string]*bintree{}},
			"operation.gomd": &bintree{pkgTemplatesOperationGomd, map[string]*bintree{}},
			"partials": &bintree{nil, map[string]*bintree{
				"content.gomd":      &bintree{pkgTemplatesPartialsContentGomd, map[string]*bintree{}},
				"headers.gomd":      &bintree{pkgTemplatesPartialsHeadersGomd, map[string]*bintree{}},
				"parameters.gomd":   &bintree{pkgTemplatesPartialsParametersGomd, map[string]*bintree{}},
				"request_body.gomd": &bintree{pkgTemplatesPartialsRequest_bodyGomd, map[string]*bintree{}},
				"responses.gomd":    &bintree{pkgTemplatesPartialsResponsesGomd, map[string]*bintree{}},
				"schema.gomd":       &bintree{pkgTemplatesPartialsSchemaGomd, map[string]*bintree{}},
				"security.gomd":     &bintree{pkgTemplatesPartialsSecurityGomd, map[string]*bintree{}},
			}},
			"schemas.gomd": &bintree{pkgTemplatesSchemasGomd, map[string]*bintree{}},
			"tag.gomd":     &bintree{pkgTemplatesTagGomd, map[string]*bintree{}},
		}},
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