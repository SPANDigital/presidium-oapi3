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

var _templatesIndexGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\xd5\xe5\x2a\xc9\x2c\xc9\x49\xb5\x52\xa8\xae\x56\xd0\x0b\x01\x31\x15\x6a\x6b\xb9\x40\xe2\x80\x00\x00\x00\xff\xff\xff\xd8\xd2\xd8\x1c\x00\x00\x00")

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

var _templatesInfoGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x3b\xab\x83\x30\x14\xc7\xf7\x7c\x8a\x03\xd9\x02\x09\xdc\x55\xc4\xe5\xde\xbb\x49\x0b\x56\xbb\x8b\x8d\x6d\xc0\x47\x89\xc1\x0e\xe2\x77\x2f\x27\x49\x7d\xa4\x1d\xea\x74\xfe\x8f\xf3\x3b\x12\xce\x39\x31\xca\x34\x32\x82\x69\x02\x91\xe3\x08\xf3\x4c\x1e\x52\x5d\x6f\x26\x82\x1f\x82\x0d\x42\x29\xdd\xe5\x30\x0a\x94\x67\xa9\x07\xd5\x77\xb8\x40\x50\xff\xc9\xa1\xd2\xea\x6e\x56\x8f\x83\xaa\x41\xe4\x52\xb7\xc3\xb1\x3e\x49\x3d\xaa\xca\xf2\x09\x00\x00\xa5\x60\x13\xe8\x6b\xf0\x59\xe4\x92\xd8\x1e\x0b\xb7\x12\x47\x94\xdd\x05\x11\x2f\xf8\x6f\xdf\x99\xb2\x32\x5b\xaa\xb7\x9c\x0e\x7a\xe2\x50\xb6\xf6\x17\xc0\x7f\x8c\xa1\xc3\x98\x7b\x81\xb0\xb5\x20\xfc\xd5\x4f\xc4\xff\xb6\x54\xcd\x1e\x69\xad\x90\xb9\xf4\xbe\x81\x16\x59\xba\x47\x16\x59\x8a\xc0\x78\x4b\x74\xa5\xe4\x0d\xb8\x8e\xcf\x00\x00\x00\xff\xff\x83\x10\xc4\x02\xe1\x01\x00\x00")

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

var _templatesPartialsSchemaGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\xdb\x6e\xdb\x38\x13\xbe\xcf\x53\x0c\x94\x14\xb0\x85\x2a\xf9\xfb\x5f\x1a\xc8\x02\x41\xb7\x05\x0a\xb4\x8d\xd7\x69\xf7\xa6\x28\x16\xaa\x4d\x25\xdc\x95\x29\x47\x87\x6c\x02\x59\xef\xbe\xe0\x49\x1a\x52\xa4\x2c\xc7\x4e\x9b\x9b\x58\xe4\x68\x38\x27\x7e\xfc\x86\xaa\xeb\x08\x56\x24\xa1\x8c\x40\x50\x2c\xef\xc8\x3a\x0e\x9a\xe6\x04\x00\x80\xcf\x9c\xe5\x24\x81\xd9\x25\x9c\xdf\x88\xa9\x05\x49\xce\x17\x24\x01\x2c\x21\x5f\xb2\x84\xfe\x8c\xd3\x8a\x18\x62\x9b\x9c\x14\x74\x45\xab\xf5\x82\x24\x5f\x17\x1f\x85\xfc\xdc\x1a\xc3\xf2\x2c\x5e\x13\x21\xf4\x99\xff\x40\x33\x17\x21\x7c\x48\xa0\xbc\xa3\x05\xa8\xa5\x69\x01\x31\xe4\x24\x21\x39\x61\x4b\x02\xe1\x05\x44\x48\x9e\x26\xc0\xb2\x52\x7a\xd2\x34\x27\x62\x1c\xcd\xc5\x6c\xd5\xfa\xa0\xfe\x9f\xff\x4e\x8a\x65\x4e\x37\x25\xcd\x18\x34\x0d\x84\x21\x1a\x98\x85\x21\x7f\xd7\x27\xcb\xa7\x08\x5b\x59\x2b\xf1\x85\xf4\x0b\x5f\x9e\x36\xad\x3f\x23\x6c\x51\xe2\x10\x86\xfc\x17\x5f\x1d\x2d\xae\x27\xf5\xaa\x51\xd3\x68\x5d\x5a\x64\x41\xe2\xd5\x35\x4b\x9f\xa4\x0e\xfe\x04\xfc\x51\x7a\x11\xf9\xde\xfa\x5c\xa5\x69\xfc\x23\x55\x2b\xeb\x27\xeb\xa5\x13\xcb\x07\x28\xc9\x7a\x93\xc6\x25\x81\x20\x27\x45\x99\xd3\x25\x8f\x4a\x11\xb4\x3e\x21\xaf\x9d\x6a\xcc\x40\xcd\xf3\x6c\x43\xf2\x92\x92\x02\xec\xb5\xc2\xb0\x9b\x9c\x85\xe1\x80\x21\x9b\x56\x2e\x80\x15\x5d\x96\x10\xc8\x22\xed\x8c\x0a\xac\x2a\x0c\xfa\xb5\x1a\x2c\xb2\xac\xe4\x75\x18\xa8\xba\x1c\xf4\x9d\xb2\x94\x32\x32\x7f\xf1\x85\xdd\x95\x86\x5d\xa7\xac\xfc\xa3\x8a\x59\x49\x13\x4a\xf2\xa3\xda\x11\xb8\x8c\xd2\x39\x6d\x9a\x13\x23\xbd\x18\x62\x70\x3e\x4c\x9c\xb9\xaf\x68\x4e\x56\x08\x47\xce\x17\x7a\x0c\x0b\x56\x79\xba\x13\x3b\x72\x65\xad\x10\xd4\xa6\xb7\x66\x6e\x41\x3c\x6e\x41\xec\x9e\x2d\xe0\x1d\xbc\x85\x05\xaa\x5c\xd8\xca\x17\x22\xf1\x67\xfe\x8b\x5c\x4f\xd1\xb6\x35\x22\x8f\xd9\x2d\x91\x01\x7a\xcd\x03\x2a\xdc\x7e\xc2\xee\x99\xf5\x8d\x37\x46\xaf\x7c\x9f\x74\xee\xe6\xed\x73\xa7\x32\x78\x5f\xa5\xa9\x37\x39\x38\x75\x6d\x5c\x1c\x29\xe7\x71\xc5\xb0\xa9\xb3\xe7\xcd\x24\x2d\xc9\xba\x98\x7b\xd2\x79\x50\x96\x5a\x89\x65\x95\xe7\x84\x75\x42\x6f\xd1\x33\x2e\xba\xdd\xb1\x16\xe7\xd1\x40\xc4\xcf\x12\xc6\xc5\xc5\x96\x49\x20\x78\x55\xc0\x6f\xf0\x8a\x23\x17\x36\xa1\xad\xf6\xe3\xe4\x2a\x61\x3f\x29\x53\x9d\x61\xcf\xda\x71\x46\x3c\xb5\x53\xe3\xce\xea\xb3\x44\xb9\x2b\x66\xb5\xef\x2f\x50\x29\x0f\x82\x6f\xcc\x2e\x3b\x6b\x1d\x14\xa4\xbc\x23\x0b\x49\x67\x3a\x29\xc4\x66\xb6\xe2\x64\x6d\x2d\x8e\xda\x61\xe3\x7c\x9e\x08\x32\x21\x55\x4d\x61\x42\xee\xd5\xe2\xf2\x2c\x0e\xb2\x1f\x7f\x93\x65\x19\x4c\xf5\x28\x2a\xb9\x08\x55\xce\xb7\xba\x16\x51\x6b\x9a\xef\x93\xba\xe6\x41\x68\x9a\x8b\x65\xb6\xde\x64\x8c\xb0\xb2\xb8\x90\xd0\x5c\x5c\x9c\x9a\x25\x96\x3b\x2a\xa4\x69\xa2\xba\x86\x34\xfb\x97\xe4\x6d\x85\x4e\xbb\xe2\x48\x8b\x81\x92\x95\xcb\x70\xcb\x07\x8b\xf6\xd0\x2a\x45\x14\x67\x42\xd9\x4d\x99\x53\x76\x7b\x93\xd2\x25\x41\x55\x28\xb4\x4f\x05\x99\xe1\x34\x45\x8e\x0a\xb6\xd3\x55\xb5\xa5\x4b\x85\x18\x33\x1c\x8b\xe2\xec\x7e\x19\x13\x1d\x93\xe9\xd8\xef\xda\xcf\x5b\x4d\x57\xa4\x22\x8b\x07\xd6\x35\xf8\x67\x7a\x6a\xbc\xc4\xe9\x41\x15\x31\x6c\x6d\x3e\x3b\x39\xa8\x18\xa7\x76\x4d\x68\x08\xe5\xf8\x29\x76\x48\xaf\x7c\x1d\x94\xd5\x8b\x9a\x4c\xc2\xa5\xeb\x9d\x71\x78\xe9\x05\xca\x83\xea\xcf\x97\x4e\xff\x19\xd7\xa3\x72\xdd\xb1\x33\x16\x41\x47\x40\xdc\x2e\xc6\x72\x28\x9f\x68\xbb\x38\x84\x76\x7d\xe8\xec\x10\xd3\xd5\x23\x4d\x74\x13\xb5\x1f\xe8\x59\x15\x70\x7a\x2a\xdc\x55\x30\x05\x35\x3c\x13\xe0\xc0\x2e\xac\x1e\xa8\x99\x6c\x77\x41\x12\x3f\xa2\xb9\x8b\x67\x64\xed\xfc\x92\x04\x21\xbf\x75\x7a\xb2\x5c\x47\xff\x2a\x4d\xaf\x5b\x58\xba\x62\x4f\xdd\xc3\x35\x23\xd7\x3c\x7f\x5d\x2a\x7f\x5a\x7a\x0e\x03\xca\xa1\xfe\x6e\x47\x93\x23\xa3\xe7\xc9\xb1\x1b\x4c\xac\xde\x06\xad\xab\xba\x9b\xd1\xf8\xd1\xb7\x0e\xed\xf9\xde\xc5\x09\x9a\xdb\xeb\xb6\xc4\x4b\xc0\x76\x32\x27\xdc\x6e\x8b\xe2\xd0\x53\x61\x78\xcd\x08\x5c\x27\xbc\xb9\xee\xc5\xfc\xde\x17\xee\xae\xd5\x57\xda\xf6\x6c\x2d\x07\x53\x60\xb5\xbb\xa6\xf1\xb2\xe6\x5b\xe3\xaf\xd2\xf4\x10\xe3\xa5\xb6\x9f\x67\xbc\xd8\xa3\x9d\xf1\xec\xe9\x20\xe3\x85\xb6\x97\x31\xde\x5b\xe8\xf7\x56\x8d\x1f\xe5\xec\x1b\xac\xee\xe7\xb6\xef\x23\xfb\x76\xba\x7a\x74\xa3\x78\x1f\xbb\x25\x49\x78\x9b\x55\xac\xe4\x72\x6f\xbc\xf8\xed\xe9\x4d\x60\x64\x7f\x02\x2f\xd5\x8c\x20\x80\xd3\x54\xee\x66\x93\x95\xf0\xff\x19\xbc\x3a\x7d\xe8\x1f\x87\xfa\xef\xdb\x07\xe1\x39\xc8\xc0\x40\x5d\xe3\x50\x1c\xab\xb3\x41\x07\x4b\xd3\x44\x72\x81\x48\x2a\x89\xec\x15\xa7\x0e\x97\xac\x04\x5d\x42\x51\xad\xcd\xa1\x37\x2e\x9a\x6a\xf7\x4d\x78\xee\x97\xf4\x4f\xf6\x99\xd3\x55\xc4\x51\xda\x10\x78\x4e\x2b\xe2\xe3\x40\x47\xdc\x3c\xbf\x94\xb9\x5a\xdc\xe8\x14\xec\x82\x37\xed\xdf\x8b\x31\xed\x53\xd7\x47\x62\xbb\xbb\xca\xed\xb9\xdd\xd2\x45\x08\xdd\x46\x80\x2f\xda\x34\xf5\x1d\xa6\xff\x81\x49\xee\x17\xd7\x57\xa6\xb9\x03\xf6\xec\x8f\x4c\x73\xff\x05\x8f\xff\x1e\x6a\xf7\x0d\xd2\x88\x7b\x28\xbe\xcb\x5c\x5f\x91\xa4\x71\x1f\x29\xfb\xc7\x9c\x6f\x71\xa4\x3b\xeb\xed\x18\xd3\x04\x78\x49\xe2\xcf\x39\x41\x9c\xe7\xf1\x53\xe0\xdc\xfc\x67\xe2\xd2\x55\x9c\x10\xea\x95\x0f\x62\xc0\x25\x4b\x13\xe0\x3d\x81\x78\x43\xdd\x7f\xaa\x66\xc0\x18\x92\xcd\x02\x1e\x92\xd4\xcd\x7d\x40\xe0\xd2\x4e\x69\x61\x72\x5b\xac\xc5\x8d\x5d\x3c\x18\x2e\xc5\xdf\xf6\x80\x55\xc3\xf1\x43\xa0\xf5\xfb\x4e\x6c\xc5\x09\xec\xa5\x49\x41\x87\x3b\x4f\xae\x9b\x62\xad\x49\xc0\x11\x56\x65\x6b\x38\xf6\xe5\xa0\x01\x35\xd3\xbe\x4d\x2e\x07\xac\x0f\x8c\xd6\xe7\xc1\xf7\x59\xbe\x8e\x05\xe0\x71\x03\x8d\xc1\xa6\x99\x42\x5d\x13\xb6\x72\xa0\x89\x09\x26\x08\x4a\xd4\xcf\x8b\x10\xde\xb1\x15\x64\xc9\x10\x9e\x68\x49\x7c\x3c\x75\x51\x70\x80\x8e\x79\x8e\x45\xdd\x5a\x0b\xc2\x56\x24\x87\x38\x4d\x4d\x65\xfd\x0f\xc8\xe7\x9f\xaa\xb4\xa4\x9b\x54\x76\x48\x10\x86\xfa\x19\xb2\x24\x0c\x67\x1c\x91\x2d\x11\x27\x82\x0a\x4d\xf1\xa3\x52\x11\x3f\xd2\x75\xb5\x6e\x5f\x97\xe3\xde\xf7\xde\x3d\x2e\xd3\xaa\xa0\x0f\xa4\x55\xd0\x8e\xc0\xba\x53\x55\xe6\x15\x19\x58\x9c\xaa\x0f\xda\x9f\x28\x33\x16\xa7\x6c\xe4\xe2\x5a\x01\x5a\xbc\x53\xb5\x63\xf1\xf8\xf1\x23\x61\xb7\xe5\x9d\xe1\x3f\xa4\x62\x0c\x85\xa1\x13\x1a\xf2\xc3\x50\x25\x4d\xb0\x55\x61\x21\xaf\xaa\x79\x5c\x96\x24\x57\x5e\xa9\x07\xad\x01\xcd\x0d\x79\x25\xa1\xc8\x70\x4a\x20\x21\xf2\xa9\x15\x19\x72\x09\xeb\x51\x1e\x99\x7a\x90\x88\x57\xcf\x57\x46\xef\x2b\x82\x54\xa1\x81\x51\x39\xe2\x50\x6b\x79\xd3\x7d\xb4\x45\x2e\xb5\x72\x43\x2e\x61\x65\xca\x25\x87\x32\x24\xe7\x55\x76\xb5\x5a\x51\xbe\x3d\xe3\xb4\xe3\x68\x57\x29\x87\xb6\x95\xd4\xdf\x09\xa0\x25\xf8\xe6\xe6\x22\x3b\x3d\xbf\x2d\x61\x92\x12\x06\xe7\xef\x58\xb5\x9e\xc2\xff\x0c\x50\x0e\x43\x3e\x1a\x86\x33\xd7\x95\xb9\x24\xbb\x1d\x17\xe5\xa2\xee\x43\x41\xc8\x3c\x83\x64\x99\xd0\xe8\xc4\x3d\x0b\x1b\x15\xb2\xf1\x63\x00\xc4\x51\xe4\x82\xc5\xf6\xb8\xb0\x22\xdd\x3b\x02\x73\xb2\x49\xe3\x25\x81\x89\x3c\x48\xce\xa7\x10\xfc\x15\x40\x10\x05\x63\x2d\xf7\x99\xa3\xe7\x79\xb6\x62\xee\xd0\xcd\x38\x1a\xe9\x90\x77\xf3\xc9\x01\x1a\x69\xb6\x0c\x16\xbd\xc3\xde\x3b\xd8\x9d\x7d\xa2\xfb\xc9\x5d\x7b\x8e\xd2\x32\x75\x73\x22\xbf\x84\x97\x31\xed\x49\x1c\xa1\x4f\xb0\x9c\x01\xec\x51\x4a\x07\x41\x1a\x4b\x61\xc0\x6e\x53\xfa\x64\xb1\x1f\x33\x18\xd1\x7b\x1f\x97\x98\x40\x9f\x9c\x38\x86\xc6\xed\xcb\x31\x45\x6c\x54\xb1\x1d\x92\x7e\xa1\x0a\xa6\xaf\x28\x3a\xde\x6a\x67\x99\xf8\x46\x17\x88\x99\xc0\x06\x4a\x7c\x9b\xd8\x53\x78\xa9\x05\xec\x4a\xe3\x3a\xb9\x4a\x3e\x19\xb8\x8a\xdc\xbc\x20\x76\x29\xee\x37\x0e\x48\xb1\x98\x74\x23\x86\xba\xf7\xe4\xed\x74\xb6\x81\xa6\x99\xa9\xaf\xaf\x06\xc0\x76\x5b\xb7\x5d\x32\x6a\x9a\xf1\x65\xcd\xa5\x5f\x7b\x53\xf9\x5f\x00\x00\x00\xff\xff\x14\x50\x57\x5b\x06\x29\x00\x00")

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

var _templatesSchemasGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\xd5\xe5\x2a\xc9\x2c\xc9\x49\xb5\x52\xa8\xae\x56\xd0\xf3\x4b\xcc\x4d\x55\xa8\xad\xe5\x02\x09\x73\x55\x57\x2b\x94\xa4\xe6\x16\xe4\x24\x96\xa4\x2a\x28\x15\x27\x67\xa4\xe6\x26\x2a\x29\xa4\x64\x26\x97\x28\x28\x05\x83\x79\x41\xa9\x69\x4a\x0a\x7a\x70\xb6\x82\x12\x48\xbb\x12\xd4\x14\xa5\x80\xa2\xd4\xe2\xcc\x94\xcc\xd2\xdc\xa0\xd4\xb4\xd0\x20\x1f\x25\x05\x3d\x34\x11\x85\xda\x5a\x40\x00\x00\x00\xff\xff\x00\x66\x70\xe3\x7f\x00\x00\x00")

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
