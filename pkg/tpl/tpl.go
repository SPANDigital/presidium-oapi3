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

var _templatesInfoGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcf\xbd\x0a\x83\x30\x10\x07\xf0\x3d\x4f\x71\x90\x2d\x90\x3c\x80\x88\x4b\xdb\x4d\x5a\xb0\xda\x5d\x6c\x84\x80\x1f\x25\x06\x17\xf1\xdd\xcb\x9d\xa9\x1f\x69\x87\x3a\xe5\xfe\xf7\xf7\x17\x95\x52\x32\x67\x5c\xa3\x23\x98\x26\x50\x39\x1e\x61\x9e\x19\xe6\x8c\x73\x7e\x48\x61\x54\x38\x3e\xb4\x1d\x4c\xdf\x61\x8d\xe1\x7c\xd6\x43\x65\xcd\xcb\x6d\x99\x04\x53\x83\xca\xb5\x6d\x87\x5b\x7d\xd7\x76\x34\x15\xa9\x0c\x00\x80\x73\xa0\x0d\xf4\x35\xf8\x5d\x44\x8b\x98\xee\x0a\x5f\x4a\x16\x50\x77\x4f\x14\x3e\xf6\xa9\xef\x5c\x59\xb9\x3d\xea\x23\x1a\x83\x9a\xba\x96\x2d\x7d\x00\xf8\x47\x08\x4c\x84\x58\xfe\x3a\x6c\xad\x84\xbf\xf4\x97\x78\x69\x4b\xd3\x1c\x49\x8a\x42\x73\xed\xfd\x83\x16\x59\x7a\x24\x8b\x2c\x45\x30\xde\x8b\x4b\x29\xf9\x02\xb7\xe3\x3b\x00\x00\xff\xff\x36\xe9\x6b\x90\xd5\x01\x00\x00")

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

var _templatesOperationGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcd\x6a\x84\x30\x14\x46\xf7\x3e\xc5\xc5\xa5\x10\x1f\xa0\x8b\x2e\xca\x6c\x0a\xfd\x63\xa6\x0f\x60\xd0\xaf\x34\x60\x34\x4d\xae\x0b\xb9\xe4\xdd\xcb\x95\x8c\x16\x66\xba\x3c\xf9\x4e\x8e\x51\xc4\x90\xfb\xa2\xf6\x15\xfc\x3d\x0f\x9f\x8e\x47\x50\xce\x54\x19\x63\x2a\x56\x7a\x20\x91\xeb\xac\x8b\xd2\x9b\xf5\x6a\x6d\x92\x08\xc6\x84\x02\x7f\x6e\xbc\x07\x44\xcb\x6e\x9e\x9e\x4f\x87\x6a\x08\xd3\x90\x73\xd5\x75\x5d\xf5\x7f\x56\x57\x9d\xf5\x5d\x27\x84\x88\xde\x32\xd4\x6a\x9a\x03\x9b\xa6\xe4\xf4\x46\x91\x2f\x8b\xf7\x36\xae\xda\x7b\x14\xb9\xe2\x56\xdf\xd5\x23\x9c\xfa\xe8\x82\xbe\x90\x72\xd6\xcf\xdf\x9c\xec\x71\x43\x0c\x1f\x46\xcb\xa0\x3a\xd8\x68\x3d\x18\x31\xbd\xb8\xc4\x35\xb5\x1f\xfb\xc1\xad\x1c\xf1\xb3\x20\xf1\xd3\x3c\xac\x35\xb5\xe7\x83\xee\xa9\x29\xcc\x53\x42\xa9\x9e\x0b\xde\x89\x26\xf4\x4b\x74\xbc\x16\xf3\x52\x70\xfb\xb9\xdf\x00\x00\x00\xff\xff\x39\x4e\xed\xbc\xcf\x01\x00\x00")

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

var _templatesPartialsContentGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\x6e\x83\x30\x10\x44\xef\x7c\xc5\xc8\xe4\x84\x9a\x70\xaf\xd4\x53\xdb\x63\x15\x29\xcd\x07\xc4\xc5\x43\x83\x54\x96\x14\xdc\x43\x84\xf6\xdf\x2b\xc0\xc1\x56\xaa\x72\x42\x9e\xe7\xdd\x79\x1e\xc7\x2d\xca\x02\xcf\x9d\x78\x8a\x87\x67\x7b\xf9\xb2\x9e\x28\x4a\x6c\x55\xb3\x29\x76\xac\x1b\x21\x4c\xb5\x30\x46\x35\x03\x80\x29\xea\xad\x7c\x12\x1b\xb1\x2d\x1f\xb0\x69\xe9\x1a\x7b\xbc\x5e\x88\xc7\x27\xec\xa0\x9a\xcd\xe0\xf4\xe5\x79\x8e\xd3\x38\xce\xa4\xea\x29\x06\xd3\x94\x75\xa9\x59\x27\x98\x64\x5a\xb2\x8f\xe2\x42\xa9\xf8\x57\x16\x78\x15\x87\xae\xfe\x57\xe2\x86\x1d\x28\x8e\xfd\x00\x8b\xb7\xb5\x69\xe2\x59\x16\x38\xee\x5f\xf6\x01\x83\x3f\x13\x73\x07\xdc\x83\xb7\x07\x89\x6d\x93\x8a\x4d\x8d\xdd\x7b\x75\x66\x6b\x11\x4e\x97\x24\xb1\x1c\xe6\xd8\xc0\x35\x95\x87\x59\xe0\x03\x6b\x73\x7f\x31\x78\x22\x2a\xe3\x8f\x73\xcf\xef\x1f\x0e\x1e\x1f\x9d\xbb\x86\x8e\xbf\x01\x00\x00\xff\xff\xd2\xed\x50\x80\xd3\x01\x00\x00")

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

var _templatesPartialsHeadersGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x31\x4f\xf3\x30\x14\xdc\xfb\x2b\x4e\x51\x87\xef\xb3\x68\xcb\x8c\x80\x09\xa4\x0e\x85\xa1\x54\xec\x26\x7e\x6e\x2d\x25\x4e\x70\x5c\x21\xf4\xe2\xff\x8e\x9c\xa4\xc1\xa1\x74\x20\x93\x75\xef\xce\x79\x77\x67\xe6\x05\x56\x02\x6b\x92\x8a\x1c\x76\x54\xd6\x85\xf4\x04\xb1\xc2\x22\x84\x19\xb3\x22\x6d\x2c\x21\x3b\x74\x84\xac\x43\x01\x20\xea\xe6\x3d\x88\x9b\x3b\x2c\x5f\x65\x71\x24\xa4\x43\x47\xba\x9b\x6c\x49\xa7\xb8\xd1\xa3\x6e\x40\xcf\x27\xcb\x2d\xbd\x1f\x8d\x23\x85\x10\x70\xfb\xe6\xee\x21\xdc\x80\x08\x30\x93\x55\x21\x80\x39\x55\x3c\x50\xed\x28\x97\x3e\xd5\xa8\x11\xfb\x56\xb5\x60\xf6\x27\x97\x59\x55\x93\x93\xde\x54\xf6\x25\x3f\x50\x29\x77\x9f\x35\x65\xe3\x95\x3d\x86\x4e\x24\xe2\x82\xbe\x5a\xef\x9e\x36\xcf\xf4\xb1\x31\x96\x9a\xe4\xd7\x4d\xee\x4c\x1d\xaf\x89\xf1\x08\xb4\xa3\x59\xb2\x71\x9d\xd9\x70\x1c\x4e\x2b\x81\x47\xab\x50\xe9\x4b\xa9\xcf\x26\xa5\x6c\x4c\xe3\xe1\xcf\x8a\xc1\xb4\x99\xc8\x9a\xb6\x63\x34\xf6\x1e\xff\x0a\xb2\x58\xfe\xc7\xf5\xcf\xb4\x9d\xb4\x7b\xc2\xdc\xca\x92\xae\x26\x4d\x26\xbc\xf8\xb5\x10\x82\xb9\x27\x22\x04\x21\x26\x01\x9e\x9e\xc5\x85\x4a\x07\xff\x7f\x8a\xe3\x57\xbf\x5f\x01\x00\x00\xff\xff\xe6\x6d\x44\x2a\xa7\x02\x00\x00")

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

var _templatesPartialsParametersGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xc1\xae\xd3\x30\x10\xbc\xe7\x2b\x46\x69\x0f\x60\xd1\x94\x33\x02\x4e\x20\x51\xa9\x54\xa8\x54\xdc\x4d\xb2\x69\x2c\x35\x4e\xb0\x5d\x0a\x4a\xfc\xef\x4f\x9b\xba\x89\xd3\x57\xbd\xe7\x43\x62\xed\xec\xac\x77\x77\xa6\xeb\x56\x58\x0b\xb8\x4a\x59\x38\xaa\xdb\x93\x74\x04\x43\x39\xa9\xbf\x64\x21\xf1\x43\x1a\x59\x93\x23\xb3\xa7\x12\x97\x4a\xe5\x15\xf2\x46\x3b\xa9\x34\xa3\x1c\x94\xba\x88\xf3\x20\xd6\x58\x79\x9f\x70\xe1\x82\x4a\xa5\x09\x69\x7b\x03\x53\x78\x9f\x00\x00\xa3\xcb\x31\x8c\x0f\x9f\x90\xfd\x92\xa7\x33\xcd\x70\x43\xe5\x80\xf0\x2b\x51\x5c\x95\x31\x35\x00\x7c\x7a\x08\xd1\x75\x11\x98\xed\x64\xcd\x25\x85\x78\x80\x6d\xf4\x15\xe9\xba\x79\xc5\x6c\x4f\x7f\xce\xca\x50\x01\xef\xf1\xf1\xb7\xf9\x0c\x61\x42\x84\x93\x49\x17\xde\x3f\x27\x7d\xa1\xd6\x50\x2e\x5d\x4c\x2b\xc6\xd8\x44\xec\x87\x11\xc6\x4d\xa7\x4d\x4b\x46\x3a\xd5\xe8\x9f\x79\x45\xb5\x3c\xfc\x6f\x29\x8d\xeb\x5e\xc3\x08\x4c\xb8\xe6\xdb\xe1\xfb\x76\x47\x97\xad\xd2\x64\xe7\x0d\xd8\xdc\xa8\x96\x4b\x71\xf2\xb8\x2d\xd2\xdc\x51\x12\xae\xe1\xb6\x16\xf8\xaa\x0b\x34\x25\xda\x7b\xdd\x92\xd7\x1c\xb1\x55\xd6\x31\x73\xe6\x8c\x97\x44\xb7\xcc\x48\xef\x14\x3c\x3a\xbc\x39\x91\x46\xf6\x16\xef\x19\x1a\x45\x5c\x2c\xc0\x02\x90\x75\xd3\x0b\x36\x89\x34\x1e\x34\xed\xb1\xd1\xe8\xc1\xeb\x42\x8f\x68\xf6\x30\xf9\x90\xba\x1a\x4e\x3f\x7e\xc6\x5f\x38\x53\x2a\x37\x65\xa4\x3e\x12\x96\x4a\x17\xf4\xef\xdd\xbd\x39\x23\x97\xdd\xf2\x27\x0d\x23\x7b\x3f\x36\x66\xa4\x43\xd8\xc1\x24\x4b\x2c\xd0\x24\xcb\x34\xfa\x75\xdf\x61\xbf\x4f\x01\x00\x00\xff\xff\x1b\xaf\x66\x09\xae\x03\x00\x00")

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

var _templatesPartialsResponsesGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x8b\xdb\x30\x10\xbd\xfb\x57\x3c\xec\x1c\xda\xb0\x71\x7a\x2e\xf4\xb4\x2d\xe4\x90\x5d\x4a\x36\xf4\x2e\xe2\xf1\xae\x21\x91\x8d\xa5\xa5\x94\xc8\xff\xbd\xd8\x92\xf5\x65\x25\xb4\x1b\x08\x48\xd6\x9b\x91\xe6\xbd\x37\x73\xbd\x6e\xb0\x5d\xe3\x40\xa2\x6b\xb9\x20\x48\xba\x74\x67\x26\x09\xeb\x2d\x36\xc3\x90\x8d\xe7\x15\xd5\x0d\x27\xe4\xbd\x01\xe5\x18\x86\x0c\x00\xc6\xc3\x55\x4f\x35\xbe\x7e\x43\x39\xa7\x28\x0f\x54\x47\x00\x93\x3b\x40\xfd\x62\xe7\x77\x0a\x70\x42\x32\xf9\x2e\x26\xd4\x8b\x5e\x7a\xa7\x4d\xed\x12\x95\x8f\x2d\x97\xc4\xe5\x7c\x3e\x63\x7a\xc6\x5f\x09\x2b\xce\x2e\xf4\x80\xd5\x85\xaa\x86\x1d\xff\x74\xd3\xbd\x77\x63\xfd\x3b\x6c\x54\xf9\x72\x7a\xa3\x0b\x8b\x71\xee\xb5\xd3\xf1\x41\x17\xbf\x08\x1b\x49\xb8\x1f\x99\x0e\x9b\x58\x49\x04\x2a\x5c\xaf\x86\xa0\x61\x98\x76\x90\xed\xee\xf8\xb4\x7f\xa6\xdf\xfb\x86\x93\xf0\x2a\xfc\x4e\xe2\xd4\x37\x9d\x6c\x5a\x8e\x19\x3c\x91\x62\x77\x56\xe4\xbc\xed\xa8\x67\x23\x52\xdf\x3f\x3e\x25\x4f\x92\x00\xb5\xa0\x8b\x78\x15\x09\xa0\xbf\x58\xcd\xe8\x2c\xc8\xe7\x2f\x51\xc4\x9d\x57\x6f\xf4\x3f\x0b\x93\x9b\x15\xcc\x72\xbb\xc6\x0f\x5e\xa1\xad\x6f\x3b\x38\x8b\x2c\xbe\x6f\x84\xfc\x27\x9b\x8f\x40\xcf\xea\xa3\x3d\x5e\x25\x3e\x9d\x89\xa3\xfc\x8c\x2f\x7e\x65\x45\x61\xd3\x8b\xcc\xab\x77\x77\x3c\xfe\xc4\x63\x5b\x11\x14\xfc\x02\x15\x9e\xac\x3b\x15\x0c\xc9\x8e\x61\xb5\x71\x3f\x7f\x1d\xee\xec\x5a\xa5\xba\x40\xf3\xfc\x10\xb5\x5f\xca\xf6\xce\x0d\xae\xc1\xab\xe6\x24\x91\x1f\xec\xde\x65\xc9\x75\x6f\xe6\xb6\x5f\x97\x1e\xc0\x6c\x82\x0f\x3f\x69\x1e\x2a\xce\x1e\xde\x50\x89\x80\x1a\xa0\xc7\x49\x10\x12\x4c\x98\xa8\xcf\x67\x21\xc3\xf8\x72\x47\xac\xa2\x5e\x18\x75\x17\x5d\x58\x14\x45\x31\xb9\xd6\x96\x0e\x13\xb1\xc4\x2a\x3c\x8f\x2d\xa7\x60\x44\x0e\xe4\x5f\xa2\x03\x49\x23\xcd\x93\x63\xc4\xa9\xf6\x36\x3d\x41\xbb\x35\x5d\x4f\x8a\x04\xa3\xd3\x8d\x4f\x51\x9b\xfd\x5f\xc7\xfd\x0d\x00\x00\xff\xff\xed\xb6\x1c\x3b\x54\x06\x00\x00")

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

var _templatesPartialsSchemaGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xdd\x6e\xdb\xb8\x12\xbe\xcf\x53\x0c\x94\x14\xb0\x85\xca\x39\xbd\x0d\x4e\x0a\x04\xe7\xb4\xd8\x02\x69\x93\x75\xda\xbd\x29\x8a\x85\x62\x51\x09\x77\x69\xca\xd1\x4f\x1b\x43\xd6\xbb\x2f\x48\x8a\xe2\x90\xa2\x2c\xe5\xaf\xc5\xf6\xa6\x26\x39\xfa\x38\xc3\x19\x7e\x33\x24\x53\xd7\x11\x24\x24\xa5\x9c\x40\x50\xac\x6e\xc9\x3a\x0e\x9a\xe6\x00\x00\x40\x8c\x1c\xe5\x24\x85\x93\x53\x58\x5c\xc9\xa1\x25\x49\x17\x4b\x92\x02\x96\x50\x1f\x39\x42\x7f\xc4\xac\x22\x96\xd8\x26\x27\x05\x4d\x68\xb5\x5e\x92\xf4\xcb\xf2\x5c\xca\x5f\x3a\x7d\x58\x9e\xc7\x6b\x22\x85\x3e\x89\x1f\x68\xe4\x38\x84\x0f\x29\x94\xb7\xb4\x80\x76\x6a\x5a\x40\x0c\x39\x49\x49\x4e\xf8\x8a\x40\x78\x0c\x11\x92\xa7\x29\xf0\xac\x54\x96\x34\xcd\x81\xec\x47\x63\x31\x4f\x3a\x1b\xda\xff\x17\xff\x27\xc5\x2a\xa7\x9b\x92\x66\x1c\x9a\x06\xc2\x10\x75\x9c\x84\x21\xfc\xf7\x3a\x7f\x0b\x75\x3d\x24\xaf\x87\x09\x4f\x9c\x19\xc5\x84\xfa\xa3\xcf\xdb\x4d\x67\xd7\x04\x9d\x5a\x71\x08\x43\xf1\x4b\x68\x81\x14\xd0\x83\x78\xe6\xa8\x69\x34\x9e\x16\x5b\x92\x38\xb9\xe0\x6c\xab\x70\x44\x0b\x44\xd3\x58\x14\x0d\x7d\xf9\xa9\x62\x2c\xbe\x66\xad\x06\xba\xe5\xf9\xf0\xc0\xb1\x07\x4a\xb2\xde\xb0\xb8\x24\x10\xe4\xa4\x28\x73\xba\x12\xab\x54\x04\x9d\x7d\x68\x05\xbc\x30\xf6\xa2\x5d\xe6\xd9\x86\xe4\x25\x25\x05\xb8\x73\x85\xa1\x19\x3c\x09\xc3\x3d\x8a\x6c\x3a\xb9\x00\x12\xba\x2a\x21\x50\x81\x6b\x94\x0a\x9c\xc8\x0c\xfa\xf1\x1b\x2c\xb3\xac\x14\xb1\x19\xb4\xb1\xba\xd7\x76\xca\x19\xe5\xe4\xf2\xc5\x27\xf6\x47\x1d\x36\x9d\xf2\xf2\xf7\x2a\xe6\x25\x4d\x29\xc9\x9f\x55\x8f\xc0\xa7\x94\xf6\x69\xd3\x1c\x58\xee\xc5\xb4\x83\xfd\x61\x73\xcf\x5d\x45\x73\x92\x20\x6e\x59\x2c\x75\x1f\x16\xac\x72\x36\xca\x27\x79\xab\xad\x14\xd4\xaa\x77\x6a\xee\x40\x36\x77\x20\x77\xd2\x0e\xf0\x8e\xde\xc1\x12\x45\x2e\xec\xd4\x07\x91\xfc\x67\xff\x17\xf9\x5a\xd1\xae\x53\x22\x8f\xf9\x0d\x51\x0b\xf4\x5a\x2c\xa8\x34\x7b\x8b\xcd\xb3\xe3\x1b\x6f\x8c\x5e\xf8\x6e\xb5\xef\x2e\xbb\xb6\x81\x0c\xde\x57\x8c\x0d\x3a\x07\xbb\xae\x5b\x17\x8f\xcb\xc5\xba\x62\x2a\xd5\xde\x1b\xf3\xe4\xf6\x91\x7e\xb4\x56\x44\x9b\x35\x2d\x2b\x1c\xa5\xad\xc1\x72\x54\x5b\xff\x7c\x51\xd2\x49\x7c\x97\x99\xed\xe4\xd4\x68\xeb\x49\x76\xe5\x2d\x59\xaa\xc4\x69\xa4\x50\xde\xdc\x49\xee\xee\x34\x8e\xba\x6e\x2b\x03\xcc\x64\xda\x52\x50\x73\x98\x91\xbb\x76\x72\xc5\xf6\x41\x76\xfd\x17\x59\x95\xc1\x5c\xf7\xa2\xd0\x89\x50\xec\x7c\xad\x6b\xb9\x6a\x4d\xf3\x6d\x56\xd7\x41\x5d\x17\xb4\x24\x8b\xeb\xb8\x20\x55\xce\x9a\x26\x68\x9a\xba\x3e\x92\x3f\x8f\x57\xd9\x7a\x93\x71\xc2\xcb\xe2\x58\xd1\x40\x71\x7c\x68\x87\x5e\xee\x89\x9b\xa6\x89\xea\x1a\x58\xf6\x83\xe4\xdd\xde\x9f\x9b\x90\x61\x05\x19\x0e\x65\x35\x8d\xb0\x67\x6f\x30\x3f\x35\x76\x51\x6a\x9d\x51\x7e\x55\xe6\x94\xdf\x5c\x31\xba\x22\x28\x36\x25\xfa\xbc\x4b\xa2\x22\x3d\xaa\x11\x99\x69\x4d\xbc\x3b\x78\xed\xe2\xe3\xcc\xda\x7d\x6f\xd2\xeb\x38\x00\x4e\xb0\x2d\x00\xca\xb2\xee\xf7\x6e\x7b\xa7\xd3\xa4\x02\x73\xea\x11\x91\x01\xb2\xdf\x3e\x7f\x3c\xff\x44\x7e\x9c\x53\x4e\x8a\x61\xb9\x1e\xe8\x60\xfa\xfe\xde\x06\x3d\xf4\xe9\x4d\x70\x9b\x8c\xfd\x5e\x60\x3a\xe6\x1f\xa5\x5c\xc8\xc9\xa4\x94\x42\xf0\xaa\x80\xb7\xf0\xaa\xd0\x9e\x3e\xe2\x8f\xe3\x40\x8b\xfc\x52\x6e\xc2\xe7\xe5\x78\xaf\x9f\xde\x4d\xfe\x9b\xca\x7f\x13\x08\x6a\x2c\x8b\x3d\x35\xc7\x74\xd5\x3e\xe2\xaa\x3e\xf1\x19\xbe\xf3\xd5\xd2\x33\x5d\x6c\x3f\x8c\xb2\x9c\x42\xf8\xf0\x50\x9a\xcb\x11\xfb\x76\x73\x9d\xc0\x23\x59\x09\xdc\x62\xbb\xc7\x44\x76\x39\xb4\x24\xe9\x30\x0d\xf9\xa3\x66\xa0\x9a\x75\xdb\xbf\xc4\x5b\xc8\x6e\xed\xab\x2c\xd7\xae\x38\x63\xec\xa2\xe3\x8f\x33\xbe\x35\x8d\x0b\x4e\x2e\x84\x33\x8d\x5f\x7f\x92\xaf\x5c\x67\x3d\x27\xbf\xed\x3b\x1c\x8c\x54\xc8\x6a\x65\x07\xfc\xef\x67\x18\xa7\x30\x46\xf3\xb6\xa5\xf1\x64\xa2\xe9\x6b\x87\xc8\xa1\x77\x12\x47\x63\x0f\x3a\x7e\x0f\xd6\x59\xa3\x05\x12\x3e\xab\xc9\xc0\xd1\x43\x61\x78\xc1\x09\x5c\xa4\xe2\x64\xd6\x5b\xf3\xbb\xa1\xe5\x36\xe7\xc4\x16\xed\x81\xe7\x92\xbd\x2e\x70\xce\x4a\xb6\xf2\x6a\x3f\x74\xca\x9f\x31\xf6\x14\xe5\x15\xda\xcf\x53\x5e\xee\x5f\xa3\x3c\xdf\x3e\x49\x79\x89\xf6\x32\xca\x0f\x06\xfa\x9d\x13\xe3\xcf\x92\x24\xf7\x46\xf7\x63\xcf\x7e\x13\x0f\x7d\x34\xb9\xf7\x33\x7c\x9f\xd7\x55\x35\xf1\xbf\xac\xe2\xa5\x90\x7b\x33\xc8\xed\x03\x47\x10\x98\x78\x0c\x81\x97\x3a\x73\xe8\x7f\x5f\x3f\x48\x4b\x40\x19\x0a\x75\x8d\x4d\x7b\x96\x03\x89\x2f\xe2\x64\x3a\x41\xd9\xa4\x69\x22\x35\x6d\xa4\x40\x22\x57\x8f\x79\x4f\xf1\x9e\x1b\x4e\xa1\xa8\xd6\x76\xd7\x1b\xdf\xe5\x9d\x7b\xdc\xc1\x63\xbf\xe4\xd8\xe3\x66\x16\xe3\xf7\x17\x38\x31\xc0\x63\x4e\x0d\x43\x35\xd1\x33\x6e\x98\x5f\x5a\xd6\x3a\xb5\xd2\x21\xb8\x9b\xc2\xd6\xff\xf1\x15\xd4\x43\x42\xfe\x99\x4a\xe1\xb1\x48\x9c\x58\x0e\xf7\xb2\xc0\x71\x08\x66\x8f\xc0\x67\xad\x5a\x7b\xb3\xdf\x7f\xb2\x50\x5b\xc9\xf7\x6e\x71\xe9\xe1\x3d\xf7\xd9\xe2\x72\xf8\x22\x67\xf8\xbe\x69\xfc\xa6\x68\xc2\x7d\x93\xd8\x80\xbe\x77\x09\xa5\xdc\x39\xe5\x7f\xdb\xe3\x1d\xc5\x98\x64\xef\xae\x31\x4d\x41\xc4\x27\x7e\x18\x08\xe2\x3c\x8f\xb7\x41\x8f\x17\xbe\x3e\x80\x95\x5a\xbc\x0f\x25\x59\x17\x4f\x62\xa6\x6f\x07\x3d\xbe\xd4\x2f\x34\x58\x67\x8f\xae\x3f\xe1\x1a\xcb\xda\x42\xf3\xbe\xa6\x1e\x62\x75\x9f\x60\x9c\x87\x93\xf7\x59\xbe\x8e\xc5\xae\x86\x59\x5d\xdb\x9d\x4d\x33\x87\xba\x26\x3c\xf1\xec\x12\x7b\x93\xa0\x2d\xd2\xfe\x3c\x0e\xe1\x1d\x4f\x20\x4b\xf7\xed\x13\x2d\x89\x39\xd8\xac\x82\x67\x33\xd9\x64\x1d\x99\xb9\x96\x84\x27\x24\x87\x98\x31\x1b\xac\xff\xd4\xb6\xf8\x58\xb1\x92\x6e\x98\x2a\xfd\x21\x0c\x75\x1b\xb2\x34\x0c\x4f\x04\xd3\x38\x22\xbd\x17\x24\x0b\x2d\xbe\x6f\x61\xe2\x7b\xba\xae\xd6\x1d\x84\xea\xdf\xfb\xed\xbb\xfb\x15\xab\x0a\xfa\x9d\x74\x20\x5d\x0f\xac\x0d\x5c\x99\x57\x64\x44\x09\xda\x3e\x03\x7e\xa4\xdc\x52\x82\xf2\x07\x28\xa1\x41\x90\x12\x06\x6e\x82\x12\xf1\xfd\x39\xe1\x37\xe5\xad\xb5\x1e\xc0\x64\x1f\x5a\x16\x23\x34\x66\x93\x05\xa7\x54\x71\xe1\xb0\xd0\x5e\xb8\xcb\xb8\x2c\x49\xde\x5a\xd8\x36\x34\x0a\x1a\x1b\xb3\x50\xd1\x8b\x65\x20\x15\x5d\xc8\xbe\x4e\x64\xcc\x3c\x8c\xd5\x5a\x67\x63\x21\x91\xbd\x58\x5f\x38\xbd\xab\x08\x82\x43\x1d\x93\x7d\x27\xa8\xd4\xb1\xcc\x3c\x7b\x21\xf3\x3a\xb9\x31\xf3\x30\x60\x6b\x9e\x07\x10\xc9\xed\x05\x3c\x4b\x12\x2a\xb6\x74\xcc\x4c\xf1\x72\xc6\x04\x1d\x26\x6a\x0e\x23\x80\xa6\x11\x84\x20\x44\x26\xad\xc2\x4d\x09\x33\x46\x38\x2c\xde\xf1\x6a\x3d\x87\xff\x58\x14\x1f\x86\xa2\x37\x0c\x4f\x04\x82\xc5\x86\x56\x35\x68\x8a\x35\x21\xee\x2d\x74\x95\x4c\xd3\xf4\x80\xa6\x14\x1f\x36\xb5\x7a\x79\xd3\xe1\xd6\x96\x19\x45\x1a\x01\x99\x0e\x7d\xb4\xda\xa5\x1b\x67\xd5\x7b\xef\x01\x39\xd9\xb0\x78\x45\x60\xa6\x12\xd1\x62\x0e\xc1\x9f\x01\x04\x51\x30\x55\xf3\x21\x75\x1e\x58\x55\x09\x07\xc7\xc2\xee\xab\x91\xf2\x6a\x4f\x55\x65\x97\xd3\x4e\xb5\x33\x52\xec\xbc\x78\xad\xe3\x35\xd0\x2e\x72\x1e\x5f\xab\x58\xd5\x34\xa3\x85\x7d\x61\x37\xf4\x87\x0f\xff\xd2\xfa\x42\x87\x8c\x6b\x67\x3f\x2a\x0a\x19\x16\xea\x2a\x0c\x87\xf3\x51\x26\x5f\x8b\x02\x39\x12\xb8\xc4\x84\xaf\xb2\x7a\x80\xa7\x5a\xc0\xdd\x48\x02\x53\x40\x8a\xc1\xc0\x17\x51\xf6\xed\xa4\x0f\x58\x09\x0c\x00\xcb\x41\xff\xae\x6c\x2f\xdd\xc4\xb9\x2e\xdb\x40\xd3\xe8\x3f\x18\xf2\x9f\x6f\xcd\x86\xe9\xe6\x76\xa7\x9c\x16\xb5\xe2\xb3\xd7\xf6\x44\x3e\x9e\xf8\x27\x00\x00\xff\xff\x01\x42\xe9\x41\xef\x25\x00\x00")

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

var _templatesTagGomd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\xd5\xe5\x2a\xc9\x2c\xc9\x49\xb5\x52\xa8\xae\x56\x28\xc9\x77\x4e\xcc\x4d\xcd\x51\xd0\xf3\x4b\xcc\x4d\x55\xa8\xad\xe5\x02\x49\x73\x55\x57\x2b\xe8\xb9\xa4\x16\x27\x17\x65\x16\x94\x64\xe6\xe7\x81\xc4\xb9\xaa\xab\x75\x15\x32\xd3\x14\xf4\x5c\x2b\x4a\x52\x8b\xf2\x12\x73\x5c\xf2\x93\x8b\x41\x12\x0a\x0a\x0a\x0a\xd8\xe4\x30\x0c\x50\x80\x82\x68\x90\xe9\x78\x54\xc6\x6a\x60\x28\x08\x0d\xf2\x51\xa8\xad\xd5\x84\xdb\x95\x9a\x53\x9c\x0a\xb3\x1b\x04\x6c\x70\xe8\xb0\x43\xe8\xc8\x4b\x01\x69\x40\x30\x01\x01\x00\x00\xff\xff\xd8\x77\x01\x1a\x07\x01\x00\x00")

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
		"schemas.gomd": &bintree{templatesSchemasGomd, map[string]*bintree{}},
		"tag.gomd":     &bintree{templatesTagGomd, map[string]*bintree{}},
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
