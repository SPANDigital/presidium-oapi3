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

var _templatesIndexGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\xd5\xe5\x2a\xc9\x2c\xc9\x49\xb5\x52\xa8\xae\x56\xd0\x0b\x01\x31\x15\x6a\x6b\xb9\x8a\x73\x4a\xd3\xc1\x42\x20\x46\x66\x5a\x25\x92\x14\x48\x0b\x20\x00\x00\xff\xff\xb1\x7c\xd0\xe9\x37\x00\x00\x00")

func templatesIndexGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexGomd,
		"templates/index.gomd",
	)
}

func templatesIndexGomd() (*asset, error) {
	bytes, err := templatesIndexGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInfoGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xcd\x6a\x84\x30\x10\xc7\xef\x79\x8a\x81\xdc\x02\xc9\x03\x88\x78\x69\x7b\x93\x16\xac\xf6\x2e\x36\x29\x01\x3f\x4a\x92\x15\x16\xf1\xdd\x97\x49\xb2\x7e\x64\xf7\xb0\x9e\x66\xfe\x1f\xbf\x91\x70\xce\x89\xd3\xae\x97\x19\x2c\x0b\x88\x1a\x47\x58\x57\x62\xfb\xcb\x9f\x97\x70\xd0\xea\x7a\xb0\xb0\x42\x28\xa5\xa7\x02\xcc\x02\xd7\x1f\x69\xac\x9e\x46\x8c\x11\xdc\xdf\xa5\xed\x8c\xfe\x77\xbb\xc6\x41\x2b\x10\xb5\x34\x83\xfd\x52\xdf\xd2\xcc\xba\xf3\x54\x02\x00\x40\x29\x78\x07\x26\x05\xd1\xcb\x82\x93\xfb\x63\x69\xab\x08\x44\x39\xfe\x22\xe2\x0e\x7f\x9b\x46\xd7\x76\xee\x48\x8d\x52\xd8\x93\x9c\xf8\x6c\x07\xff\x0b\x10\x3f\xc6\x50\x61\x2c\x3c\x49\x9a\xda\x10\xf1\xea\x33\xe2\xc7\xd0\xea\xfe\x8c\xf4\x52\xca\xdc\x72\xaf\x40\x9b\xaa\x3c\x23\x9b\xaa\x44\x60\x7e\x24\x86\x50\xf1\x00\xdc\xc7\x5b\x00\x00\x00\xff\xff\x00\x8c\x66\x58\xf2\x01\x00\x00")

func templatesInfoGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesInfoGomd,
		"templates/info.gomd",
	)
}

func templatesInfoGomd() (*asset, error) {
	bytes, err := templatesInfoGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/info.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOperationGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xcf\x4a\xc4\x30\x10\xc6\xef\x79\x8a\xa1\xc7\x42\xfa\x00\x1e\x3c\x48\x2f\x82\xff\xd8\x15\xbc\x36\x74\x3f\xdd\x40\xd3\xc6\x64\x8a\x94\x90\x77\x97\xe9\xb6\x56\xdc\x7a\xcb\x6f\xe6\x37\xdf\x0c\xd1\x5a\xab\x94\x34\xd9\x77\xaa\x1e\xc1\xe7\xe1\xf4\x6a\xb9\x03\xe5\xac\x58\x1e\x37\x94\xd2\xda\xa1\x9c\x67\x7a\x32\x6e\x16\x52\x22\x74\xf1\xaf\xfb\xec\x11\x0c\xdb\xa1\xbf\xaf\x2f\x92\x26\xf4\x32\xab\xbe\x60\x3f\xce\x7c\xb1\xde\xe6\xb7\x54\xe5\x02\xd5\x34\x8d\xfa\x7f\x91\x74\xa5\x2d\x47\xd6\xf0\x01\xad\x61\x88\x55\x96\x1b\x96\xe5\xaf\x55\xab\x7d\x1c\x9d\x33\x61\x92\xc0\xdb\x94\x56\x9c\xe3\xaf\xdc\x1a\xb1\x0d\xd6\xcb\xe9\x94\xb3\xec\xbf\xaa\x6c\x13\x9a\x18\xce\x77\x86\x41\x85\x37\xc1\x38\x30\x42\x7c\xb0\x91\x0b\xaa\x5e\x7e\x0a\xeb\x0f\x6c\x72\xc0\xe7\x88\xc8\x77\xc3\x69\x2a\xa8\x3a\x6c\xb4\xa7\x46\x3f\xf4\x11\x4b\xea\x61\xc1\x9d\xd0\x88\x76\x0c\x96\xa7\xc5\x3c\x2e\x48\x39\x7f\x07\x00\x00\xff\xff\x34\x7c\xcc\x01\xe0\x01\x00\x00")

func templatesOperationGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesOperationGomd,
		"templates/operation.gomd",
	)
}

func templatesOperationGomd() (*asset, error) {
	bytes, err := templatesOperationGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/operation.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsContentGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x3f\x6f\xb3\x30\x10\xc6\x77\x3e\xc5\x23\x8b\xc9\x7a\x03\x7b\xa4\x77\xea\x9f\xa9\x55\x2a\x9a\xee\x71\xf1\xd1\x58\xaa\x0d\x05\x67\x88\x2c\x7f\xf7\x0a\x02\xb6\x4b\xda\xa1\x65\x42\xbe\xdf\xd9\xcf\xfd\xce\xb9\x0d\x4a\x8e\x9b\xd6\x58\x32\x16\x96\x74\xf7\x2e\x2c\x81\x97\xd8\x78\x9f\x8d\x65\x49\x8d\x32\x04\x56\x5f\x18\xe6\x7d\x06\x00\x63\x29\xef\x7a\x1a\x94\x54\x27\x5d\x51\xf3\x52\x3d\x60\xfb\x1f\xc5\xd3\xea\x2c\xe1\x7b\x61\xde\x08\xb9\x11\x9a\xfe\x21\xd7\x24\x95\xd8\x9f\x3b\x9a\xda\x96\x08\x33\x3e\x7e\x9c\xdf\xb7\xbd\x16\x76\xcb\x39\x0e\xce\x4d\x7d\xde\x1f\xb2\x00\x38\x17\x13\xb3\x70\x1d\x83\x54\xb5\x05\x7b\x8c\x07\xc9\x5b\x6c\x95\x8f\x5d\x4f\x91\x24\x26\x23\x67\x0d\xf1\xaf\xe4\xb8\x33\x12\x6d\xf3\xa3\xb6\x05\xab\xc8\x48\xea\x07\x08\x84\x2c\xa9\xd9\x92\x63\xbf\xbb\xdd\xcd\x18\xec\x91\x30\xe5\xc4\x1a\x5c\x56\x10\x47\xfc\xc3\x12\x42\x83\x6a\x50\x84\x3c\xc5\x73\x7d\x24\x2d\x52\xed\x5f\xac\x0e\x53\x79\x51\x7a\x81\x2b\x6a\xd8\x37\x57\xfc\xda\x2c\xa2\x5a\x5c\xb9\xed\xe9\xe3\x44\x83\xc5\x6b\x2b\xcf\xb3\x8b\xcf\x00\x00\x00\xff\xff\x8e\x9d\xe9\x63\xad\x02\x00\x00")

func templatesPartialsContentGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsContentGomd,
		"templates/partials/content.gomd",
	)
}

func templatesPartialsContentGomd() (*asset, error) {
	bytes, err := templatesPartialsContentGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/content.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsHeadersGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x3d\x4f\xc3\x30\x14\xdc\xfb\x2b\x4e\x51\x07\xb0\x68\xca\x8c\xc4\x56\x24\x06\xa6\x52\xb1\x5b\xf1\xb9\xb5\x94\x38\xc1\x75\x07\xe4\xe4\xbf\xa3\x7c\x34\xb2\x1b\x3a\x90\xe9\xe9\xde\xdd\x29\xef\xce\x21\x6c\xb0\x15\x78\xa7\x54\x74\x38\xb0\x6a\x4a\xe9\x09\xb1\xc5\xa6\xeb\x56\x21\x28\x6a\x63\x89\xec\x34\x10\xb2\x01\x05\x80\x5e\xb7\x1e\x41\xbc\xbc\x22\xff\x92\xe5\x85\x88\x97\x8e\x7a\xd8\xec\xa9\x63\xdc\xe8\x59\x37\xa1\xcb\x4d\xbe\xe7\xf7\xc5\x38\xaa\x98\x22\xdc\x04\x8a\x48\x45\xab\x12\x97\xd8\x64\xc7\xc6\xb1\x90\xfe\xc6\x46\xcd\xf0\xc2\x08\x2d\x42\xf0\xd7\x0c\xb2\xba\xa1\x93\xde\xd4\xf6\xb3\x38\xb1\x92\x87\x9f\x86\xd9\xec\x3e\x62\x18\x44\x22\x4a\x23\xdf\xf1\x5c\x38\xd3\xf4\xba\x3e\x2d\x81\x76\xbe\x9d\x76\xf8\x95\x69\x9c\xa6\xad\xc0\x9b\x55\xa8\xf5\xbd\x12\x56\x49\x47\x1f\xe6\xec\xe1\x17\x3d\x21\x2d\xaa\x67\xa5\x65\x19\x8d\xa3\xc7\x43\x49\x8b\xfc\x11\xcf\xb7\xe1\x3b\x69\x8f\xc4\xda\xca\x8a\x4f\x49\xb1\x11\xaf\xff\x5a\x08\x11\xc2\x48\x44\xd7\x09\x91\x24\x76\x7d\x25\x77\x1a\x9e\xee\xff\x57\x1c\x7f\xde\xfb\x1b\x00\x00\xff\xff\x1b\x02\x9f\x55\xb6\x02\x00\x00")

func templatesPartialsHeadersGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsHeadersGomd,
		"templates/partials/headers.gomd",
	)
}

func templatesPartialsHeadersGomd() (*asset, error) {
	bytes, err := templatesPartialsHeadersGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/headers.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsParametersGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x4d\x6f\xd4\x30\x10\xbd\xe7\x57\x3c\xa5\x7b\x68\x2d\x36\xcb\x19\x89\x1b\x20\xad\x54\x21\x28\x88\xbb\x95\x4c\xba\x16\x89\x93\xda\x5e\xa0\x4a\xfc\xdf\x91\x1d\x27\xb1\xf7\xa3\xad\x0f\xbb\xb1\x9f\xdf\xcc\xf8\xcd\xc7\x30\x6c\xb1\x63\x30\x07\xa1\x61\xa8\xed\x1b\x6e\x08\x8a\x4a\x12\x7f\x48\x83\xe3\x1b\x57\xbc\x25\x43\xea\x81\x6a\xfc\x3d\x88\xf2\x80\xb2\x93\x86\x0b\xe9\x50\x77\xc8\x65\x15\xdf\x03\xdb\x61\x6b\x6d\xe6\x0c\x57\x54\x0b\x49\xc8\xfb\x19\xcc\x61\x6d\x06\x00\x0e\xdd\x2c\xc7\xf8\xf0\x11\xc5\x2f\xde\x1c\x29\xc1\x15\xd5\x1e\x71\x5e\xa2\x73\x51\xc7\xd4\x00\xb8\x35\x82\xb1\x61\x88\xc0\xe2\x2b\x6f\x9d\x49\xc6\x2e\x60\x7b\x39\x21\xc3\x90\x5a\x2c\x1e\xe8\xe9\x28\x14\x55\xb0\x16\x4c\x85\x8d\xbb\x47\xb2\xb2\xf6\xfc\xfe\x27\xea\x15\x95\xdc\x4c\x0c\xb0\x6a\xd9\xaf\xa4\xd1\x47\xbe\x08\xbc\x2a\xf2\xf3\xb9\xa7\x3c\xb6\xf6\xa3\x3c\x50\xcb\x11\x38\x48\x1d\xe9\x52\x89\xde\x88\x4e\x3a\x78\x11\x84\xa4\xf3\x9c\x85\xcf\xf0\xb5\x63\xf8\x2c\x2b\x74\x35\xfa\xd3\xd4\x64\xd7\x92\x2e\x68\x4a\x7a\x88\xa1\xee\x14\xf8\x39\xfd\x62\x66\xa7\x77\xbc\x21\x7b\x1b\x3d\x19\x4f\x52\x9e\xe4\x56\xad\x84\xf9\x78\xe2\xdc\x0b\xf9\x3b\x81\xfd\x8b\x1b\x4d\x9e\xa6\x67\xe1\x12\xa6\xa8\x41\x4f\x33\x58\xb8\x28\x91\x73\xa5\xf8\x73\xee\xdf\x82\x68\xbd\x9c\xa1\x60\x61\x6f\xa8\xd5\x09\x35\x0e\x42\x76\x26\xf5\x75\xee\x23\x72\xd1\x08\x6d\xbe\x1f\xb9\x34\xa2\x16\xa4\x74\x7e\xed\x0d\xde\xfa\x85\x60\x13\x4f\xbe\x32\xb7\x91\x12\xc5\x97\x4e\xb5\xdc\x38\xe0\x76\x18\xd2\x43\x6b\xef\xe6\xd2\x4c\x3d\x4d\x95\x74\xb9\xb0\xb0\x56\x16\x12\x75\x5e\xad\xab\x79\x98\xdc\x0b\x6d\x5c\x45\x26\x43\xe5\xa5\xaa\xd2\x8e\x91\x9f\x34\xff\xa3\xc1\x6d\x43\x12\xc5\x1d\xde\x2f\xc5\xe3\xd6\xcd\x0d\x5c\xef\x92\x36\xab\x07\x9d\x45\xe3\xc1\x8f\x83\x11\x7b\x89\x11\x3e\xf0\x11\x51\x4f\x85\x8e\xf2\x57\xb7\x7e\x8d\xcb\xcf\xf2\x17\xd6\x98\xc8\xa6\xb8\x7c\x24\x6c\x84\xac\xe8\xdf\xbb\xd3\xb9\xf6\xb6\x2a\xcb\xaf\xcc\xb4\x48\xfb\xa0\xc1\x9a\x8a\x93\xa4\x84\x76\x5f\x9f\x3e\xe9\x1d\xf4\xfd\x1f\x00\x00\xff\xff\x2f\x0e\x01\x21\xe9\x05\x00\x00")

func templatesPartialsParametersGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsParametersGomd,
		"templates/partials/parameters.gomd",
	)
}

func templatesPartialsParametersGomd() (*asset, error) {
	bytes, err := templatesPartialsParametersGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/parameters.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsRequest_bodyGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x53\xcb\x6e\x83\x30\x10\xbc\xf3\x15\x23\x92\x13\x2a\xe4\x1e\x29\x97\x3e\x8e\xed\x81\x56\xbd\xbb\xb0\x04\x4b\xc1\xa6\xe0\xb4\x8a\x8c\xff\xbd\x32\xaf\xda\x2e\xf5\xc9\xbb\x33\xb3\xf6\x7a\xd6\x5a\xa7\x38\x24\x50\x35\xef\xa1\xa8\x69\x2f\x4c\x11\x3a\x2a\x88\x7f\x51\x0f\x86\x9c\xce\x57\xea\xd5\xbd\x2c\x6f\x39\x55\xf8\xae\x79\x51\xa3\x90\x42\x31\x2e\x26\xbc\x02\x13\x25\x18\xde\xd9\xe5\x4a\x48\x0e\x48\x8d\x89\x6c\xd9\x92\x2a\x2e\x08\x71\x47\x9f\x4b\x89\xd8\x98\x08\x00\x2c\xcc\x2b\x64\x70\xe2\xbd\xc3\xc3\xf1\x84\x6c\x2c\xe8\x13\xaa\x11\xc8\xa9\x9a\xd3\x4e\x2d\x4f\xee\xc0\x76\xed\x76\xc8\x27\x14\x16\x8e\x3c\x50\xeb\x50\x9e\x59\x32\xef\xa8\x84\x31\x49\xb2\x04\x49\xa2\x35\x5d\x7a\xb2\xb9\x17\xa9\xe0\xe5\x45\x69\x8c\x5f\x76\xc0\x33\x95\x9c\xbd\xdd\x5a\xc2\x80\x47\xea\x8b\x8e\xb7\x8a\x4b\x81\x01\xaf\x45\x4d\x0d\xc3\xe0\x0b\xd2\xdf\xe5\xee\x9d\x68\x08\x2e\x9e\xa2\x63\xe2\x4c\xd8\x0b\xd6\xd0\x1d\xf6\xcd\x7a\xe2\xf1\xe4\x77\xf4\x20\x85\x22\xa1\xc2\x87\x71\xdf\x6f\x15\x67\xf3\xf5\x36\xb8\xab\x15\xfd\x48\xc9\x27\x43\xfe\x48\x03\x87\xfe\x51\x6f\x4b\x5d\xd7\xc3\x35\x40\xeb\xb1\x57\x63\xc6\xbd\xdf\xa3\xfb\xc6\x0b\x61\x1d\xe9\x58\xb6\xd4\x31\x8b\x4d\xe7\xd8\x23\xe3\xcd\xa6\x03\x5b\x96\x4b\x4f\x16\x6f\x67\xa3\x20\xe5\x8e\xf5\x1c\x06\xdb\x43\x82\x27\x51\x42\x56\x98\x3b\xc0\x87\x9d\xdb\xf9\xf7\xfc\x04\x00\x00\xff\xff\x24\xd9\x72\x61\x95\x03\x00\x00")

func templatesPartialsRequest_bodyGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsRequest_bodyGomd,
		"templates/partials/request_body.gomd",
	)
}

func templatesPartialsRequest_bodyGomd() (*asset, error) {
	bytes, err := templatesPartialsRequest_bodyGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/request_body.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsResponsesGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4b\x6f\x9c\x30\x10\xbe\xf3\x2b\x46\xb0\x87\x36\xca\x92\x9e\x2b\xb5\x97\x6d\xa5\x1c\xb6\x51\x44\xb6\xbd\x56\x68\x19\x12\xab\x8b\x41\xd8\x7b\xa8\x80\xff\x5e\xf9\x81\x5f\x38\x5b\xa2\x0d\x27\x6c\x7f\xe3\x79\x7d\xf3\x79\x18\xb6\x70\x77\x03\x05\xb2\xae\xa5\x0c\x81\x63\xd3\x9d\x4a\x8e\x70\x73\x07\xdb\x69\x4a\xc4\x79\x85\x35\xa1\x08\x69\xaf\x41\x29\x4c\x53\x02\x00\x20\x0e\x37\x3d\xd6\xf0\xf9\x0b\xe4\xf3\x15\x79\x81\x75\x00\xd0\x77\x7b\xa8\x5f\xe5\xe9\x8c\x1e\x8e\xf1\x92\x9f\x99\x44\x3d\xa9\x5f\xe7\x94\xd4\x50\xd2\x4a\x79\x33\x37\xe6\xbb\x96\x72\xa4\x5c\xe3\x66\x6c\x5f\xd2\x67\x84\x0d\x2d\x1b\xbc\x85\x4d\x83\x15\x29\x0f\x7f\x3b\xe9\x7f\x61\x0a\x8e\xad\xf8\x46\x18\x06\x1d\xc9\x34\xc9\x95\x63\xf3\x0d\xd9\xb1\x27\x1d\x27\x2d\x05\x73\x2a\xdc\x98\x15\x3b\xbe\x60\x53\xee\x09\xfd\xa3\x42\x15\xfb\x5e\x6c\x48\x2b\x37\x2d\x3c\x31\x14\xb9\x49\xb4\x13\xca\xdb\xc2\xd8\xfe\xc7\xb9\xeb\xe9\x62\x05\xae\xad\x9e\xee\x94\xb5\xca\x9f\x64\x4c\x21\xce\xf6\x5c\x1e\x17\x8a\x42\x0b\x33\x41\xa5\xcb\x96\x71\x33\xc9\xad\x88\x61\xa4\xaa\x6f\xea\xae\x19\x8e\xb4\xed\xb0\x2f\x05\x52\x79\x14\xce\xd3\x68\xda\x4e\xfb\x1d\x0a\x04\x25\xb7\x3b\xa6\x53\x57\x71\x61\xeb\xf6\x5d\xde\xed\x50\x4f\x4f\xfc\x77\x5a\x41\x5b\xbf\x3e\xf8\x49\xa0\x0c\x7b\xc2\xf8\x2a\x75\x10\x40\x47\x21\x04\x1f\x9e\x39\x7c\x38\x21\x85\xfc\x23\x7c\x72\x33\xcb\x32\x73\x3d\x4b\x9c\x7c\xef\x0f\x87\x47\xd8\xb5\x15\xc2\x08\x6e\x82\x23\xfc\x30\x74\x1c\x41\xd7\xd8\x16\x78\xdc\xda\xcf\xfd\xf7\x57\xe6\x7f\x8c\xd1\x5e\xd5\xf9\x36\x50\xad\x18\xcf\x2d\x19\xac\x2e\x56\xe4\xc8\x21\x2d\xcc\xda\xde\x92\x2a\x49\x4b\x8d\xcc\x2d\x29\x20\xb6\xae\x0b\x69\xd6\x62\x4b\x0f\x47\x8b\x03\xa0\x02\x28\x15\xf6\x4c\x3c\x61\x0e\x06\x7b\x6e\xa4\x6f\x9f\xdf\x63\x59\x61\xcf\x74\x77\x17\x63\x97\x65\x59\x26\x59\x6b\x52\x07\x6d\xb1\xc4\x8e\xf0\x20\x26\x6e\x04\xdd\x64\xaf\xfd\x4b\xb4\xd7\xd2\xa0\xe7\x51\xdd\xb0\x5d\x7b\x91\x21\x28\xb6\xc6\xf3\x89\x15\xc1\xd1\xef\xc8\x56\x30\x66\xd7\x4d\xdc\x6f\x2d\x71\x6b\x86\x4e\xcd\x42\x7a\xe9\x5d\x2e\xb0\x5e\xf7\x34\x0b\xe0\xf2\x75\xee\x7a\x64\xa4\x22\xe7\xa6\xc0\xfa\x67\xb1\x97\x16\x8f\xc1\x9e\x8b\x97\xc2\x29\x40\x0f\x4a\x41\x13\xf7\x29\xa7\x2d\x9f\x9f\x28\x53\xca\xaf\xc3\x10\x95\x35\x7f\x2a\xa2\x2c\xbc\x44\xc0\xf5\x84\x5a\x4f\xa6\x15\x44\x8a\x71\x28\x3e\xe7\xf1\x8c\xf4\x03\x1b\x08\xa6\xb6\xb0\xbe\x8f\x0a\x36\x2b\xcf\x6e\x5e\x2e\x5f\xea\x34\x68\x56\xfa\x6a\xfb\xde\x9d\xd7\xc9\xbf\x00\x00\x00\xff\xff\x69\x5e\x64\x23\x65\x0a\x00\x00")

func templatesPartialsResponsesGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsResponsesGomd,
		"templates/partials/responses.gomd",
	)
}

func templatesPartialsResponsesGomd() (*asset, error) {
	bytes, err := templatesPartialsResponsesGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/responses.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsSchemaGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\x5f\x6f\xdb\x38\x12\x7f\xcf\xa7\x18\x28\x29\x60\x0b\x95\x73\xbd\x47\x03\x39\x20\xe8\xb5\x40\x81\xb6\xf1\x39\xed\xbd\x14\xc5\x42\xb5\xa9\x84\xbb\x34\xe5\xe8\x4f\x36\x81\xac\xef\xbe\x20\x29\x4a\x43\x8a\x94\x95\xd8\x49\xf3\x12\x8b\x1c\x0e\x67\x38\xc3\x1f\x7f\x43\x56\x55\x04\x6b\x92\x50\x4e\x20\xc8\x57\xb7\x64\x13\x07\x75\x7d\x02\x00\x20\x7a\xce\x32\x92\xc0\xfc\x02\x66\xd7\xb2\x6b\x49\x92\xd9\x92\x24\x80\x25\xd4\x20\x4b\xe8\xff\x31\x2b\x89\x21\xb6\xcd\x48\x4e\xd7\xb4\xdc\x2c\x49\xf2\x7d\xf9\x59\xca\x2f\xac\x36\x2c\xcf\xe3\x0d\x91\x42\x5f\xc5\x0f\xd4\x73\x1e\xc2\xa7\x04\x8a\x5b\x9a\x43\x33\x35\xcd\x21\x86\x8c\x24\x24\x23\x7c\x45\x20\x3c\x87\x08\xc9\xd3\x04\x78\x5a\x28\x4f\xea\xfa\x44\xb6\xa3\xbe\x98\xaf\x5b\x1f\x9a\xff\xb3\xff\x92\x7c\x95\xd1\x6d\x41\x53\x0e\x75\x0d\x61\x88\x1a\xe6\x61\x28\xc6\xfa\x64\x45\x17\xe1\x6b\x6b\x26\x31\x91\x1e\xf0\xed\x71\xdb\xfa\x33\xc2\x96\x46\x1c\xc2\x50\xfc\x12\xb3\xa3\xc9\x75\xa7\x9e\x35\xaa\x6b\xad\x4b\x8b\x2c\x49\xbc\xbe\xe2\xec\x51\xe9\x10\x5f\x20\x3e\x95\x17\x91\x6f\xd4\xd7\x92\xb1\xf8\x17\x6b\x66\xd6\x5f\xd6\xa0\x13\xcb\x07\x28\xc8\x66\xcb\xe2\x82\x40\x90\x91\xbc\xc8\xe8\x4a\xac\x4a\x1e\xb4\x3e\x21\xaf\x9d\x6a\xcc\x85\x5a\x64\xe9\x96\x64\x05\x25\x39\xd8\x73\x85\x61\xd7\x39\x0f\xc3\x01\x43\xb6\xad\x5c\x00\x6b\xba\x2a\x20\x50\x49\xda\x19\x15\x58\x59\x18\xf4\x73\x35\x58\xa6\x69\x21\xf2\x30\x68\xf2\x72\xd0\x77\xca\x19\xe5\x64\xf1\xe2\x13\xbb\x33\x0d\xbb\x4e\x79\xf1\xbf\x32\xe6\x05\x4d\x28\xc9\x8e\x6a\x47\xe0\x32\x4a\xc7\xb4\xae\x4f\x8c\xf0\x62\x88\xc1\xf1\x30\x71\xe6\xae\xa4\x19\x59\x23\x1c\x99\x2d\x75\x1b\x16\x2c\x33\xb6\x17\x3b\xb2\xc6\x5a\x29\xa8\x4d\x6f\xcd\xdc\x81\xfc\xdc\x81\xdc\x3d\x3b\xc0\x3b\x78\x07\x4b\x94\xb9\xb0\x53\x03\x22\xf9\x67\xfe\x8b\x5c\x5f\xd1\xae\x35\x22\x8b\xf9\x0d\x51\x0b\xf4\x56\x2c\xa8\x74\xfb\x11\xbb\x67\xe6\x37\xde\x18\xbd\xf4\x7d\xd4\xb1\x5b\xb4\xdf\x9d\xca\xe0\x63\xc9\x98\x37\x38\x38\x74\xed\xba\x38\x42\x2e\xd6\x15\xc3\xa6\x8e\x9e\x37\x92\xb4\x20\x9b\x7c\xe1\x09\xe7\x41\x51\x6a\x25\x56\x65\x96\x11\xde\x09\xbd\x47\xdf\x38\xe9\xf6\xaf\xb5\x3c\x8f\x06\x56\xfc\x2c\xe1\x42\x5c\x6e\x99\x04\x82\x37\x39\xfc\x07\xde\x08\xe4\xc2\x26\xb4\xd9\x7e\x9c\x58\x25\xfc\x95\x22\xd5\x19\xf6\xac\x1d\x67\xac\xa7\x76\x6a\xdc\x59\x7d\x96\x34\xee\xca\x5e\xed\xfb\x0b\x64\xca\xbd\xe4\x1b\xf3\x8b\xce\x5a\x07\x05\x29\x6e\xc9\x52\xd1\x99\x4e\x0a\xb1\x99\x9d\x3c\x59\x5b\x8b\xa3\xb6\xd9\x38\x9f\x27\x92\x4c\x28\x55\x53\x98\x90\xbb\x66\x72\x75\x16\x07\xe9\xaf\x3f\xc9\xaa\x08\xa6\xba\x15\xa5\x5c\x84\x32\xe7\x47\x55\xc9\x55\xab\xeb\x9f\x93\xaa\x0a\xaa\x2a\xa7\x05\x99\xfd\x8a\x73\x52\x66\xac\xae\x83\xba\xae\xaa\x33\xf9\xf3\x7c\x95\x6e\xb6\x29\x27\xbc\xc8\xcf\x15\x60\xe7\xe7\xa7\x66\xe2\x65\x8e\xbc\xa9\xeb\xa8\xaa\x80\xa5\x7f\x93\xac\xcd\xdb\x69\x97\x32\x2c\x1f\x48\x64\x35\x8d\xf0\x67\x30\x95\x0f\xcd\x5d\x44\x7c\x26\x94\x5f\x17\x19\xe5\x37\xd7\x8c\xae\x08\xca\x4d\xa9\x7d\x2a\x29\x8e\x20\x2f\xaa\x55\x72\xa0\x2e\xd7\x2d\x5d\xcd\xc2\x63\xde\x63\x11\x9f\xfd\x83\x31\xfd\x31\xf9\x8f\x3d\xd6\xfe\xde\x69\x12\xa3\x14\x59\xec\xb0\xaa\xc0\xdf\xd3\x53\xe3\xa5\x53\xf7\x4d\x6a\xc3\xce\x66\xb9\x93\x83\x52\x74\x6a\xe7\x84\x06\x56\x81\xaa\x72\xdf\xf4\x92\xda\x41\x64\xbd\x58\xca\x15\x88\xba\xc6\x8c\x43\x51\x2f\x7c\x1e\x94\x7f\xbe\x70\xfa\x4f\xbe\x1e\xc1\xeb\x0e\xa3\xb1\xb8\x3a\x02\xf8\xf6\xf1\x98\x43\x59\x46\x5b\xdb\x21\x0c\xec\x03\x6a\x87\xa3\xae\xca\x69\xa2\x4b\xab\xa7\x41\xa1\x95\x01\xa7\xa7\xd2\x5d\xfb\x78\x95\x73\xcd\xe1\x99\x68\x07\x76\x96\xf5\x10\xce\x24\xc4\x4b\x92\xf8\xe1\xcd\x9d\x49\x23\x13\xe9\xb7\x44\x0b\xf9\xad\x63\x95\x66\x3a\x14\x97\x8c\x5d\xb5\x18\x75\xc9\x1f\xbb\x8f\x2b\x4e\xae\x44\x30\xbb\xb8\xbe\x52\xac\xec\x60\x1d\x86\xa1\x43\x05\xe1\x9e\xaa\x48\xad\xa5\x27\xe2\x6e\x9c\xb1\x8a\x21\x34\x6f\x53\x0e\x8d\x86\x96\xbe\x75\x08\x0e\x7a\x37\x2d\xa8\xef\x49\xd7\x2b\x5e\xc6\xb6\x97\x6a\xe1\xfa\x5c\xa6\x8a\xee\x0a\xc3\x2b\x4e\xe0\x2a\x11\xd5\x78\x6f\xcd\xef\x7c\xcb\xdd\xdd\x0d\x34\xda\x9e\x58\x8b\x0e\x86\xc0\xaa\x8f\x4d\xe3\xd5\x0e\x68\x8d\xbf\x64\xec\x10\xe3\x95\xb6\xd7\x33\x5e\xee\xd8\xce\x78\xfe\x78\x90\xf1\x52\xdb\xcb\x18\xef\x4d\xf4\x3b\x2b\xc7\x8f\x72\x2c\x0e\x66\xf7\x73\xeb\xfd\x91\x85\x3e\x5d\x3f\xb8\x31\xbd\x8f\xe4\x8a\x3f\xbc\x4f\x4b\x5e\x08\xb9\x77\x5e\x34\xf7\x14\x33\x30\xb2\xa0\x81\x97\xaa\x5e\x10\xc0\x69\x96\x77\xbd\x4d\x0b\xf8\xf7\x1c\xde\x9c\xde\xf7\x0f\x47\xfd\xf7\xe3\x93\xf4\x1c\xd4\xc2\x40\x55\xe1\xa5\x38\x4a\x29\xe4\xca\x50\x79\xe0\xa0\xf3\xa6\xae\x23\x35\x6d\xa4\x94\x44\xb6\x1d\x53\x87\xa3\x56\xd8\x2e\x20\x2f\x37\x66\xd3\x3b\x17\xaf\xb5\x0b\x2d\xdc\xf7\x5b\x0a\x2e\xfb\x24\xea\xf2\xe4\x28\x75\x0b\x3c\xa7\x76\xf1\xf1\xa4\x23\x6e\xa9\xdf\x4a\x75\x2d\xfe\x74\x0a\xf6\x36\x30\xed\x7f\x3e\xab\x7a\x4a\x92\x1f\x89\x1e\xef\xcb\xbd\xe7\xd6\x5a\xe7\x21\x74\xbb\x02\xbe\x69\xd3\x9a\xb7\x9d\xfe\xa3\x95\xda\x3c\xae\x97\xab\x85\x03\x19\xed\x87\xab\x85\xff\xd2\xc8\x7f\xb7\xb5\xff\x56\x6a\xc4\xdd\x96\xd8\x72\xae\x97\x29\x65\xdc\x67\xca\xff\x32\xfb\x5b\x50\xe9\xe8\x80\xbd\xc6\x34\x01\x91\x9f\xf8\x89\x28\x88\xb3\x2c\x7e\x0c\x9c\x48\x70\x26\x2f\x72\xe5\x21\xd2\x0c\xf9\x24\x1b\x5c\xb2\x34\x01\x51\x44\xc8\x11\xcd\x9d\x6a\x53\x3d\x18\x4d\xaa\xba\xc0\x4d\x8a\xdd\xb9\xcf\x10\x9c\xda\x8c\xe6\x26\xfd\xc5\x5a\xdc\x40\x26\x16\xc3\xa5\xf8\xc7\x13\x30\xd6\x70\xfc\x10\x9c\xfd\xb9\x17\x68\x71\x00\x7b\x61\x6a\x70\xc4\x1d\x27\xd7\xed\xb3\xd6\x24\xb1\x09\xab\xb2\x35\xbc\xce\x85\xa3\x01\x40\xd3\xbe\xa5\x2e\xb7\xac\xa7\x4c\xeb\x21\xf2\x63\x9a\x6d\x62\x81\x89\x30\xa9\x2a\xb3\xb1\xae\xa7\x50\x55\x84\xaf\x1d\x18\x63\x42\x0c\x02\x98\xe6\xe7\x79\x08\x1f\xf8\x1a\xd2\x64\x08\x65\xb4\x24\x3e\xc1\xba\x55\x70\x40\x91\x79\xd4\x45\xdd\x5c\x4b\xc2\xd7\x24\x83\x98\x31\x53\x59\xff\xa9\x7a\xf6\xa5\x64\x05\xdd\x32\x55\x5a\x41\x18\xea\x6f\x48\x93\x30\x9c\x0b\x9c\xb6\x44\x9c\xb8\x2a\x35\xc5\x0f\x8d\x8a\xf8\x81\x6e\xca\x4d\x3b\x5c\xb5\x7b\xc7\x7d\x78\x58\xb1\x32\xa7\xf7\xa4\x55\xd0\xb6\xc0\xa6\x53\x55\x64\x25\x19\x98\x9c\x36\x4f\xe7\x5f\x28\x37\x26\xa7\x7c\xe4\xe4\x5a\x01\x9a\xbc\x53\xb5\x67\xf2\xf8\xe1\x33\xe1\x37\xc5\xad\xe1\x3f\x30\xd9\x86\x96\xa1\x13\x1a\xf2\xc3\x50\xa5\x4c\xb0\x55\x61\x21\xaf\xaa\x45\x5c\x14\x24\x6b\xbc\x6a\x3e\xb4\x06\xd4\x37\xe4\x95\x02\x28\xc3\x29\x89\x8f\xc8\xa7\x56\x64\xc8\x25\xac\xa7\xf1\xc8\xd4\x83\x44\xbc\x7a\xbe\x73\x7a\x57\x12\xa4\x0a\x35\x8c\x8a\x91\x00\x60\xcb\x9b\xee\x79\x18\xb9\xd4\xca\x0d\xb9\x84\x95\x35\x2e\x39\x94\x21\x39\xaf\xb2\xcb\xf5\x9a\x8a\xed\x19\xb3\x8e\xc6\x5d\x32\x01\x6d\x6b\xa5\xbf\x13\x40\x53\x88\xcd\x2d\x44\xf6\x7a\x7e\x53\xc0\x84\x11\x0e\xb3\x0f\xbc\xdc\x4c\xe1\x5f\x06\x54\x87\xa1\x68\x0d\xc3\xb9\xeb\x1a\x5e\xf1\xe1\x8e\xae\x0a\x51\xf7\x51\x21\x65\x9e\x41\xbd\x4c\x68\x74\xe2\x9e\x85\x8d\x0d\xb2\x89\x63\x00\xe4\x01\xe5\x82\xc5\xf6\xb8\xb0\x56\xba\x77\x30\x66\x64\xcb\xe2\x15\x81\x89\x3a\x48\x66\x53\x08\xfe\x08\x20\x88\x82\xb1\x96\xfb\xcc\xd1\xfd\x22\x5a\xb1\x70\xe8\x7a\x1c\xb9\x74\xc8\xbb\x59\xe6\x00\xb9\x34\xab\x0a\x8b\xf4\x61\xef\x1d\x9c\xcf\x3e\xe7\xfd\x94\xaf\x3d\x47\x69\xc1\xdc\x4c\xc9\x2f\xe1\xe5\x51\x4f\xa4\x93\xd0\xa7\x5d\xce\x05\xec\x11\x4d\x07\x6d\x1a\x4b\x6c\xc0\x2e\x5e\xfa\x14\xb2\xbf\x66\x30\xa2\x3c\x3f\x2e\x31\x81\x3e\x39\x71\x34\x8d\xdb\x97\x63\x92\xd8\xc8\x62\x7b\x49\xfa\x89\x2a\xf9\x7f\x43\xdc\xf1\x56\x3b\x4b\xe5\xbb\x5f\x20\x7b\x02\x1b\x28\xf1\x35\x64\x4f\xe1\x85\x16\xb0\x33\x4d\xe8\x14\x2a\x45\x67\xe0\x4a\x72\xf3\x66\xd9\xa5\xb8\x5f\x4e\x20\xc5\xb2\xd3\x8d\x18\xcd\x85\xa9\xa8\xb8\xd3\x2d\xd4\xf5\xbc\x79\xd1\x35\x00\xb6\xdb\xba\xed\x94\x91\xa0\xc7\x63\xd3\x5a\x48\xbf\xf5\x86\xf2\x9f\x00\x00\x00\xff\xff\x0d\x70\x91\x94\x70\x29\x00\x00")

func templatesPartialsSchemaGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsSchemaGomd,
		"templates/partials/schema.gomd",
	)
}

func templatesPartialsSchemaGomd() (*asset, error) {
	bytes, err := templatesPartialsSchemaGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/schema.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsSecurityGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xbf\x4e\xc3\x40\x0c\xc6\xf7\x3c\xc5\xa7\xb4\x03\x44\xbd\x86\x19\x89\x81\x81\x0d\xb1\x94\x17\x38\x25\x4e\x39\xd4\x38\x55\xe3\x4a\xa0\xe4\xde\x1d\xf9\x9a\xff\x04\x09\x0f\x51\xe2\x9f\xed\xf8\xfb\xdc\x34\x06\x69\x82\x03\x65\xd7\x8b\x93\x6f\xbc\x53\x79\x3e\x59\x21\x24\x29\x8c\xf7\x91\xf2\x9c\x0a\xc7\x84\xb8\xee\x8a\xe2\x81\x5c\x2c\x1f\x09\x5b\xb6\x25\xed\xb0\xad\xb3\xea\x4c\x35\x1e\x9f\xb0\x87\xf7\x51\x8b\xa6\xb9\x31\x78\x8f\xf0\xf5\x59\x39\x1e\xea\xe2\x1d\xe2\x40\xc2\x28\xe2\xbc\x1b\x3a\xbe\xa5\x09\x5e\x38\x47\x55\xfc\xbd\x5f\xb4\x10\xf0\xea\x6a\x59\x11\xf1\x4b\x83\xd6\xe9\xdf\x23\x00\xd0\x19\xae\xb8\x6d\x8d\x2e\xba\xdc\x51\x70\x77\x22\xc6\xfe\x1e\x0f\x8a\x07\xae\xb1\xd9\xe0\xf9\x2a\x1f\xc4\xe2\x32\x2b\xae\xe2\x39\x6e\xdf\x6c\x49\xed\x21\xa8\x6d\xe7\xc4\x18\x63\xc2\xc3\x98\x39\x99\xb8\xea\x38\xa7\x2f\xb5\xb5\xd7\xde\x1b\x8b\x45\x68\x8f\xf4\x92\x27\x57\x1a\x3b\x17\x4d\xa3\xcb\xeb\x99\xff\xdf\x63\xcd\xee\x9f\x00\x00\x00\xff\xff\x45\xb9\xf2\x04\x53\x02\x00\x00")

func templatesPartialsSecurityGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsSecurityGomd,
		"templates/partials/security.gomd",
	)
}

func templatesPartialsSecurityGomd() (*asset, error) {
	bytes, err := templatesPartialsSecurityGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/security.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResponsesGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcd\x41\x0a\xc2\x30\x14\x84\xe1\x7d\x4e\x31\xcc\x3e\x39\x80\x67\x10\x91\x88\x07\x08\xed\x14\x03\x8d\x96\xe6\xb9\x0a\xbd\xbb\x04\xb2\x90\x2e\xe7\x5b\xfc\xe3\xbd\x77\x96\x6d\xd5\x05\xad\x21\xdc\x52\x11\x8e\xc3\x75\x76\xad\xc1\x54\xb6\x35\x99\xc0\x5d\x75\xfb\xbc\xab\x1e\xd3\x4b\x25\x11\x73\x9e\x0c\x8c\x43\xa3\x16\x22\xfc\x2d\xb0\xa7\x38\x8a\xbc\xef\xaa\x79\xce\xdf\x12\xb5\x3c\xe3\x95\x08\x27\xe9\xa7\xbf\x00\x00\x00\xff\xff\x56\x01\xc5\xa0\x8c\x00\x00\x00")

func templatesResponsesGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesResponsesGomd,
		"templates/responses.gomd",
	)
}

func templatesResponsesGomd() (*asset, error) {
	bytes, err := templatesResponsesGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/responses.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSchemasGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\xd5\xe5\x2a\xc9\x2c\xc9\x49\xb5\x52\xa8\xae\x56\xd0\xf3\x4b\xcc\x4d\x55\xa8\xad\xe5\x02\x09\x73\x55\x57\x2b\x94\xa4\xe6\x16\xe4\x24\x96\xa4\x2a\x28\x15\x27\x67\xa4\xe6\x26\x2a\x29\xa4\x64\x26\x97\x28\x28\x05\x83\x79\x41\xa9\x69\x4a\x0a\x7a\x70\xb6\x82\x12\x48\xbb\x12\xd4\x14\xa5\x80\xa2\xd4\xe2\xcc\x94\xcc\xd2\xdc\xa0\xd4\xb4\xd0\x20\x1f\x25\x05\x3d\x34\x11\x90\x45\x80\x00\x00\x00\xff\xff\xf5\x76\x34\x32\x80\x00\x00\x00")

func templatesSchemasGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesSchemasGomd,
		"templates/schemas.gomd",
	)
}

func templatesSchemasGomd() (*asset, error) {
	bytes, err := templatesSchemasGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/schemas.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTagGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\xd5\xe5\x2a\xc9\x2c\xc9\x49\xb5\x52\xa8\xae\x56\x28\xc9\x77\x4e\xcc\x4d\xcd\x51\xd0\xf3\x4b\xcc\x4d\x55\xa8\xad\xe5\x02\x49\x73\x55\x57\x2b\xe8\xb9\xa4\x16\x27\x17\x65\x16\x94\x64\xe6\xe7\x81\xc4\xb9\xaa\xab\x75\x15\x32\xd3\x14\xf4\x5c\x2b\x4a\x52\x8b\xf2\x12\x73\x5c\xf2\x93\x8b\x41\x12\x0a\x0a\x0a\x0a\xd8\xe4\xd0\x0d\x50\x80\x82\x68\x90\xe1\x78\x14\xc6\x6a\x60\x28\x08\x0d\xf2\x51\xa8\xad\xd5\x84\x5b\x95\x9a\x53\x9c\x8a\x6c\xa4\x0d\x0e\x1d\x76\x08\x1d\x79\x29\x20\x0d\x08\x26\x20\x00\x00\xff\xff\x21\xf8\x46\x74\x06\x01\x00\x00")

func templatesTagGomdBytes() ([]byte, error) {
	return bindataRead(
		_templatesTagGomd,
		"templates/tag.gomd",
	)
}

func templatesTagGomd() (*asset, error) {
	bytes, err := templatesTagGomdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tag.gomd", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
