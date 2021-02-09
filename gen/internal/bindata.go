// Code generated by go-bindata.
// sources:
// template/curd.tmpl
// template/curd_common.tmpl
// template/markdown.tmpl
// template/model.tmpl
// template/model_req.tmpl
// DO NOT EDIT!

package internal

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateCurdTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x96\xcf\x6e\xe3\x36\x10\xc6\xcf\xe2\x53\x4c\x75\x08\xa4\x40\xa5\xdb\xab\x01\x1f\x1a\x67\xb3\x08\x10\x64\x91\x75\xda\x1e\x16\x8b\x80\x92\xc6\x0e\x37\x14\x69\x91\x54\xb3\xae\xe1\xb7\xe8\xad\x40\x5f\xa5\x4f\x53\xa0\x8f\x51\x90\x94\xfc\x67\xad\xc4\xce\x6e\x7a\x8a\x4d\x0e\x87\xdf\xfc\x66\xf8\xc5\x73\x56\x3c\xb0\x19\x02\x97\x16\xb5\x64\x82\x10\x5e\xcd\x95\xb6\x90\x90\x28\x9e\x71\x7b\xdf\xe4\xb4\x50\xd5\x20\x47\x69\xb0\x62\x03\xc1\x73\xcd\xf4\x62\x50\x28\x69\xf1\xb3\x8d\x9f\x8f\x32\xb5\x88\x49\x84\xd2\x9a\x5a\x40\x8c\xd2\xce\x14\xe5\x6a\x80\xd2\x0e\x4a\xce\x04\x16\xb6\x0d\x89\x8d\xd5\x5c\xce\x4c\x4c\x52\x42\xa6\x8d\x2c\x60\xac\x91\x59\x5c\x2e\xe9\x2d\xcb\x05\x5e\xb3\x0a\xdf\xaa\xd5\x2a\x29\xe0\xb4\xbd\x9a\x8e\xc3\xdf\x0c\xca\x1c\x4e\x4d\x2d\xe8\xf9\x59\x06\x2a\xff\x04\xa7\x95\x2a\x51\xd0\x2f\xcf\xa6\x90\xa0\xd6\x80\x5a\x2b\x9d\xc2\x92\x44\x79\xc3\x45\x89\x1a\x86\x23\x30\xb5\x38\x0b\xdf\x92\x94\x44\x75\x83\x7a\x91\x01\xd3\x33\xe3\x36\xdb\x38\x7a\x29\x0d\x6a\x9b\xa8\xfc\x53\x48\x9c\xa4\x29\x1d\x2b\xd1\x54\xd2\xf8\xc5\xee\x73\x4a\x29\x4d\xe9\x2f\x4c\x34\x18\x36\x2e\x38\x8a\xb2\x5b\xbf\x71\xc9\xdd\x2d\x77\x99\xd3\x02\x23\x28\x73\xfa\xe6\x33\x16\x6d\x3d\x49\x91\xc1\x96\x00\x77\x86\x44\x1a\x6d\xa3\xa5\x8b\x27\xab\x16\xd0\x39\x0a\xdc\x07\x74\xb6\xb8\x2c\x0f\x42\xe2\x25\x04\xde\x7b\x48\x1c\xbe\xe1\x08\x4e\xfa\x09\x2e\x57\x2f\x87\x16\x64\xee\x42\xfb\xf5\x1e\x35\x26\x61\x2c\xe8\x9b\x9b\x24\xe6\x65\xec\x54\xa5\xaf\x47\xe7\xe7\x79\xd9\x33\x3e\x2f\xa3\x93\xc1\x03\x2e\xd6\x9f\x7f\x73\xfd\x0c\xef\x64\xca\x0a\x5c\xee\xcf\xd3\xeb\xc3\x0b\x55\xec\xc2\x9b\xa0\x4d\x1e\x70\xd1\x0a\xfa\xbf\x60\xae\x41\xbe\x45\xfb\xcd\x33\xf6\xcc\xa3\x0c\xc2\x76\x19\xbe\x2a\xc2\x09\x3a\x9b\xe9\x79\x9f\x17\x5a\x55\x1d\xb6\x00\x77\x1b\xf3\x71\x5c\xd7\x50\xfd\xca\x7b\xf5\xf8\x34\x58\x3a\x29\x98\xdc\x73\x83\x3d\xdc\x17\x5c\x96\x2f\x36\x3d\x8d\x75\xc7\xb7\xef\xfc\x7b\xac\x43\x13\x0c\x7c\xf8\x78\x7c\x1f\x0c\x8c\x9e\x8e\x77\xad\xe8\x23\x3e\xa9\x45\x6f\x09\x1a\xeb\x0c\x26\x37\x57\xb7\x8b\x39\x5e\x71\x63\x5d\xe5\xea\xd1\x84\x6b\x87\x1b\x86\xcf\x4d\x26\x9f\xfa\xe8\xef\x46\x20\xb9\x70\x22\x3b\x78\xd1\x8a\x44\x25\x4e\x51\x83\xcb\x49\xc7\x42\x19\x74\xed\x99\xaa\x76\xe5\xda\xa5\xf4\x75\x1d\xf1\x48\xdb\xae\xfa\x83\x4f\x34\xad\x47\xca\x5a\x8b\x13\xd3\xe1\x63\xf3\x39\xca\xd2\xa3\xf7\xff\x9a\x52\x2f\xf5\x88\x96\xdf\x2a\xcb\xc4\xeb\xf4\xdd\xba\x54\xce\xb7\xbe\x6c\x71\xd8\x18\xc1\x0f\x5f\xdf\xc9\xb1\x6a\xa4\xfd\x9a\x77\x70\xe2\x2f\xdf\x1f\xff\x43\x37\x1f\x28\x37\x73\x9e\xe0\x84\x6d\xbc\xa7\xb3\xef\x0f\x1f\x77\xac\xfb\xdb\xec\xda\x78\x53\x51\x7e\xe7\xa4\xb5\x88\x49\xbb\xe6\x0e\x9a\x47\x6e\x8b\xfb\xb5\x98\x25\x89\x0a\x66\x70\x7b\xfe\x87\x24\xda\x64\x39\xec\x56\xbb\x09\x3c\xf6\xe7\x33\xc4\x3e\x26\x39\x4d\x63\x3f\x73\x64\x3b\xb4\xfb\x78\xc8\x02\x89\x7f\x72\x1a\x6b\x7a\x59\xba\x51\x8f\x63\x3f\xe9\x7d\x99\xfa\xcd\x32\x1c\x4d\x5b\x05\x83\x01\x14\x4e\x14\xfc\xfb\xf7\x1f\xff\xfc\xf9\x97\x4f\xde\x11\x1a\x8d\x76\x6a\xdb\x7a\xdc\x9b\x3b\xd6\xae\xeb\x92\xdd\x19\xf7\x23\x75\xe8\x44\xad\x71\xbb\xeb\x26\x6e\xb9\xe3\x1d\x97\x68\x8a\xd8\x81\x0a\xe1\x23\x68\x15\x9e\xa3\x29\xdc\x38\xd1\x77\xba\x44\x7d\xb6\xe8\xf8\xc6\xac\x3f\xfe\xa7\xbd\xf0\x12\xa7\xac\x11\xf6\xa8\x58\xa7\x57\x85\x6f\xae\xf5\x7e\x9c\xc2\xcf\x5d\x3a\x99\x0b\x6e\xb7\xc3\x33\x88\xb3\xb8\xb5\xae\xbb\x0c\xda\x63\xee\x88\x66\x72\x86\xb0\x95\x67\x19\x7c\xa8\xb8\xc7\xe2\xe1\x52\x4e\xac\xde\x1d\x9b\xf5\xe1\xe0\x7b\xbd\x6d\x6b\x2f\x4d\x7c\x09\x69\xf0\xae\x15\x89\x36\xd3\x42\xdf\x4d\xa7\x06\x6d\xe2\x25\x5e\x37\x15\x7c\x0f\x3f\xa6\x70\x1a\x50\xf3\xdf\x71\xeb\x2d\xd0\x2b\x5e\xb5\xb5\xb4\x3b\x4f\x35\x70\x45\xfe\x0b\x00\x00\xff\xff\x73\x8d\x4e\x6a\x76\x0c\x00\x00")

func templateCurdTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateCurdTmpl,
		"template/curd.tmpl",
	)
}

func templateCurdTmpl() (*asset, error) {
	bytes, err := templateCurdTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/curd.tmpl", size: 3190, mode: os.FileMode(420), modTime: time.Unix(1612870525, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateCurd_commonTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x31\x6b\xc3\x40\x0c\x85\xe7\xd3\xaf\x78\x78\xb2\x4b\x48\xf6\x82\x97\xb6\x4b\x21\x1d\x42\xba\x85\x50\xdc\x8b\xec\x0a\x5f\x75\xc9\x9d\x3c\x84\x90\xff\x5e\x2e\x71\x87\x42\x41\xc3\xe3\x7b\xe8\x93\x8e\x9d\x1f\xbb\x81\x21\x6a\x9c\xb4\x0b\x44\xf2\x7d\x8c\xc9\x50\x13\x00\x54\xac\x36\xc4\xa5\xc4\x15\xab\xad\x0e\xd2\x05\xf6\x56\x91\x63\xb5\x7c\x0a\xff\xd7\xab\x7c\x0a\x15\x35\x44\x3e\x6a\x2e\x22\xb7\xdd\xac\xdf\xcf\x47\x5e\x4b\x36\xa0\x45\x15\x24\x17\xc9\x8c\x9f\xe3\xa4\x56\xb0\x2f\xe1\xb6\x49\xfd\xa4\x1e\xf9\x14\x9e\x26\x09\x07\x4e\x75\x83\x87\xfb\xc9\xe5\xcb\xfd\xc8\x5c\xe0\x42\x2e\xb1\x4d\x49\xf1\xb7\xaf\xe7\x67\x96\x6f\xe7\xed\x66\xdd\xd0\x75\x76\xfa\x2f\xf6\xe3\xab\x6e\x2d\xd5\x07\xc9\xd8\xed\xb3\x25\xd1\x61\x81\x11\xf7\xd4\xe0\x33\xc6\x50\xbc\x7d\x4c\x90\x05\x3e\xf0\xd8\x22\x75\x3a\x30\xca\xc6\x85\x9c\x93\xbe\xc4\x9d\xec\xd1\xb6\x18\x6f\xe8\xf7\x0b\x4b\x13\x93\x73\x57\x2a\x33\xb3\xbe\x0b\x99\xe9\xfa\x13\x00\x00\xff\xff\xba\xb2\x95\x04\x6b\x01\x00\x00")

func templateCurd_commonTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateCurd_commonTmpl,
		"template/curd_common.tmpl",
	)
}

func templateCurd_commonTmpl() (*asset, error) {
	bytes, err := templateCurd_commonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/curd_common.tmpl", size: 363, mode: os.FileMode(420), modTime: time.Unix(1612871348, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMarkdownTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xc1\x6a\xea\x40\x14\x86\xf7\x81\xbc\xc3\x81\xb8\x50\xb8\xc9\x03\x08\xf7\x6e\xae\x1b\xb9\x17\xeb\x42\xba\x11\x17\xb1\x9e\x16\x69\x92\x96\x68\xa0\x61\x66\xc0\x45\xa1\x85\xd2\xd6\x45\xc1\xac\x0a\x05\x11\xbb\xa8\x52\xda\x82\x98\x3e\x8e\x19\xf5\x2d\x4a\x66\x9c\xd4\x56\x9a\xc5\x84\xf9\xff\xc3\xf9\x4f\xbe\x13\xc3\x80\xd5\xc3\x98\xf7\x62\x3e\x3a\xd7\x35\x5d\xa3\xc9\xfc\x36\xb9\x1c\xd0\x54\x4d\xfa\xd7\x40\x61\x35\x7d\xe5\xd1\x0d\x50\x48\x86\x17\xfc\x65\x0c\x54\xd7\x68\xd1\x34\x4d\x71\xa8\x13\xe4\x4b\xd7\x08\xf1\x6d\xef\x08\xc1\xaa\xd9\x4d\x07\xff\xb7\x3b\x5d\xc6\x28\x10\x62\x95\xbd\x16\x9e\x31\x46\xeb\xe9\x45\x98\x15\xdb\x45\xc6\x1a\x79\xe3\xd3\xfd\x6a\x15\x40\x3e\x94\x10\xeb\xef\x89\xeb\xa2\x27\x9a\x89\x14\xf4\x5a\x8c\xa5\x03\xab\xc0\xdc\x31\x86\xbf\x20\xd7\xee\xa2\x0b\xc5\xdf\x60\x95\xb0\x73\x20\xe3\x75\xcd\x30\x0c\x20\x44\x78\x2a\xc9\x52\xf7\xad\x3c\x5d\xfb\x93\x95\x65\x79\x69\x86\x29\x68\x3c\x0d\xf8\xe4\x4d\xd2\x90\xa4\x84\x00\x14\x96\xcf\x71\x72\x7f\x05\x14\x78\x34\x4d\xfa\xa3\xc5\x6c\xbe\x7c\x9c\x53\x58\xc7\xd1\x6a\x32\x4c\x7a\xef\x99\xc3\xa3\xe9\x62\x16\xaf\xef\x26\xbb\x58\x15\x55\x73\x97\xed\x37\x61\x0b\xb2\x1c\x75\x03\x59\x40\x72\x02\xd7\x93\x1f\x03\x20\xb0\x4b\xa9\x16\x9e\x62\x25\x70\x9b\xe8\x33\x26\xf5\x72\xa7\x12\x38\xce\x66\x37\x25\x3c\xb4\x03\xa7\xbb\x6f\x3b\x01\xaa\x82\xaa\xdf\x76\x6d\x3f\xfc\x87\xa1\x52\x64\xab\x1f\x16\x51\xaf\xed\x55\x1b\x79\x23\xfb\x97\x0a\x99\xf7\x11\x00\x00\xff\xff\x54\x04\xaf\xad\x62\x02\x00\x00")

func templateMarkdownTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMarkdownTmpl,
		"template/markdown.tmpl",
	)
}

func templateMarkdownTmpl() (*asset, error) {
	bytes, err := templateMarkdownTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/markdown.tmpl", size: 610, mode: os.FileMode(420), modTime: time.Unix(1599730646, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateModelTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcd\xee\x9b\x30\x10\xc4\xcf\xf1\x53\xac\x2c\x54\x41\x95\xc0\xbd\x52\x4f\x89\x92\x5b\x7b\xe1\x56\xf5\x60\x60\x43\xdd\xf8\x83\xd8\xa6\x2a\xb2\xf6\xdd\x2b\xa0\x28\x24\x8a\x92\xff\x09\xb1\xfe\xcd\x78\x76\xdc\x89\xfa\x22\x5a\x04\x6d\x1b\x54\x8c\x49\xdd\x59\x17\x20\x65\x1b\xde\x88\x20\x2a\xe1\xb1\xf0\x57\xc5\xd9\x86\xb7\x32\xfc\xea\xab\xbc\xb6\xba\x68\xed\xce\x5f\xd5\xae\x71\xf2\x0f\xba\x42\x0f\x33\xf1\x37\x48\x8d\xb0\xe6\x2a\x34\x1e\xb5\x28\x94\xac\x9c\x70\x43\x31\x02\x9c\x65\x8c\xc5\xe8\x84\x69\x11\x92\x0b\x0e\x5b\x48\x4a\xf8\xf2\x15\x72\x22\xc6\xc2\xd0\x21\xc4\x98\x94\x79\x29\x2a\x85\xdf\x84\xc6\x93\x25\x02\x1f\x5c\x5f\x07\x88\x37\xe5\xef\x2d\x24\x32\xa0\x9e\xa4\x13\x7c\x40\x5f\x13\xc1\x28\x1f\x0f\x6e\xc3\xfc\x64\xf7\x56\xf5\xda\x8c\x76\x23\xf1\x9c\x51\xc2\xb4\xe5\xd0\xcd\xc4\x02\x1c\xad\xd3\x22\x1c\x25\xaa\xc6\xcf\xe6\xf2\x0c\x8f\xda\xd9\x7d\x6f\xb5\x46\x13\x46\xaa\x28\x9e\xdc\xf0\x40\xc5\x88\xa6\x21\x62\xcb\x97\x18\x63\xe7\xde\xd4\x90\x6a\xf8\x1c\xe3\x7d\x01\x19\x4c\xbf\x69\x36\x36\x21\x4d\x0b\x91\x39\x0c\xbd\x33\xc0\xd7\x28\x11\x67\xf4\xca\x65\xce\xe0\xd3\x0c\x7e\xfc\x7c\x74\x5a\x26\x11\xde\x97\xcc\x97\xf5\x0e\xd5\x77\x27\xa7\x7a\x88\xf8\xf6\xff\x2e\x40\xaf\x53\xcc\x75\x4e\x21\xa4\x09\xe8\xce\xa2\xc6\x48\xeb\x24\xab\xf1\x47\xe2\x7c\xd2\xf9\xbb\x57\xbf\xcb\xb6\x94\xfe\x2f\x00\x00\xff\xff\x6f\x0b\x2c\xda\xfe\x02\x00\x00")

func templateModelTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateModelTmpl,
		"template/model.tmpl",
	)
}

func templateModelTmpl() (*asset, error) {
	bytes, err := templateModelTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/model.tmpl", size: 766, mode: os.FileMode(420), modTime: time.Unix(1612787993, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateModel_reqTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x31\x6e\x84\x30\x10\x45\x6b\xe6\x14\x23\x44\x19\x71\x00\xa4\x34\x29\x92\x8e\x44\x09\x7d\x30\x78\x64\x39\xe0\x31\xb1\x4d\x61\x2c\xdf\x3d\x72\xb4\x2b\x76\xb5\xed\x9f\x37\x7a\xff\x6f\x62\x5e\x84\x22\x34\x56\xd2\x0a\x10\xe2\x46\xf8\x21\x94\x66\x11\xb4\x65\xf4\xc1\xed\x73\xc0\x04\x55\xbf\x1b\x44\xcd\x01\xc7\x1f\x6f\xb9\xab\x79\x37\xf5\x08\xd5\x97\x3e\xe8\x36\xf6\xfa\xa0\x7a\x84\x0c\x90\x92\x13\xac\x08\x9b\x85\xe2\x13\x36\x03\x76\xcf\xd8\xe6\x7c\x71\xbc\x6a\x96\x29\x35\x43\x3b\x88\x69\xa5\x5e\x18\x7a\xb3\x39\x7f\xd2\xef\xa9\x44\x44\x7c\x44\x00\xaa\x77\x27\xc9\xbd\xc4\x42\x6a\x56\x57\xb3\x2d\xe9\xf7\x14\xff\x5b\x59\x17\xca\xfb\x3d\xe1\xad\x0b\xe5\x7a\xee\x2b\x3d\x21\x25\x62\x99\xf3\x5f\x00\x00\x00\xff\xff\x77\xdd\xd4\x02\x0a\x01\x00\x00")

func templateModel_reqTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateModel_reqTmpl,
		"template/model_req.tmpl",
	)
}

func templateModel_reqTmpl() (*asset, error) {
	bytes, err := templateModel_reqTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/model_req.tmpl", size: 266, mode: os.FileMode(420), modTime: time.Unix(1612886426, 0)}
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
	"template/curd.tmpl":        templateCurdTmpl,
	"template/curd_common.tmpl": templateCurd_commonTmpl,
	"template/markdown.tmpl":    templateMarkdownTmpl,
	"template/model.tmpl":       templateModelTmpl,
	"template/model_req.tmpl":   templateModel_reqTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"curd.tmpl":        &bintree{templateCurdTmpl, map[string]*bintree{}},
		"curd_common.tmpl": &bintree{templateCurd_commonTmpl, map[string]*bintree{}},
		"markdown.tmpl":    &bintree{templateMarkdownTmpl, map[string]*bintree{}},
		"model.tmpl":       &bintree{templateModelTmpl, map[string]*bintree{}},
		"model_req.tmpl":   &bintree{templateModel_reqTmpl, map[string]*bintree{}},
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
