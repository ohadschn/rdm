// Code generated by go-bindata.
// sources:
// templates/conduct.md
// templates/license/mit.md
// templates/license/unlicense.md
// DO NOT EDIT!

package file

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

var _templatesConductMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x56\x4d\x8f\xdc\x46\x0e\xbd\xd7\xaf\x20\xe0\x43\x80\x46\xbb\x07\x01\x72\xf2\xcd\x6b\x24\x8b\x1c\x16\x31\x62\xef\xc9\xf0\x81\x5d\x45\xb5\xb8\x2e\x55\x09\x2c\xaa\xdb\x4a\x90\xff\xbe\x20\x4b\xfd\x31\x8e\x07\xb9\xcc\xa8\x25\x16\x3f\x1e\x1f\x1f\xeb\x15\xbc\xab\x45\x85\x8f\x8b\x56\x81\x77\xf5\x4c\x05\x8b\xc2\xbb\x9a\x08\xea\x60\x1f\xd3\x12\x35\x84\x57\xaf\xe0\xb7\x45\xe0\x7d\xa6\x74\xa2\x10\x7e\x2d\xa0\x23\x01\x17\x25\xa1\xa6\x66\x3a\xd4\xa6\x24\x5c\x4e\x80\x05\xea\x4c\x05\xb0\x24\xb8\x50\x8e\x75\xb2\xb7\x54\xce\x2c\xb5\x4c\x54\x74\x0f\x17\x02\x6c\x21\xde\x43\x37\xb7\x9e\x90\x8b\x22\x17\x92\x06\xb3\x87\x02\xad\x30\xe1\x17\x73\x30\xa3\x28\x47\x9e\x51\xb9\x16\xe0\x02\x75\x11\x98\xa5\xfe\x8f\xa2\xda\xe9\x60\xbf\x63\x9d\xa6\xa5\xb0\xae\x80\x30\xa2\x60\x6b\x16\xf0\xf5\x20\x44\x40\x5f\x67\x12\xa6\x12\x09\x86\x2a\x40\x67\x92\xb5\x16\xda\x83\xd0\x09\x25\x65\x6a\xcd\x0a\xc1\x13\xed\xe1\x58\xd3\x1a\x1a\xff\x41\x7b\x48\xdc\xf0\xc8\x99\x75\xdd\x03\xe9\x58\x38\xfa\xe3\x89\x4a\x22\x01\x4e\x54\xd4\xe3\x95\x64\x11\x84\x5a\xe3\x5a\xf6\x90\xe9\x4c\xd9\xfc\xdd\xc3\xee\x43\xf1\xe4\xb1\x3b\x9b\x49\x9a\xfd\x00\x9c\x67\x42\x41\xb3\x00\x41\xff\x4b\x99\x4f\xee\xa6\x0a\x34\xfa\xba\x60\x7e\x16\x29\x54\xf3\xa8\xee\xed\x70\xeb\xce\x07\xc5\x92\x50\x52\x0b\xe1\xe7\xaf\x38\xcd\x99\xbc\xa0\x23\x8d\x78\xe6\x2a\xa0\x23\x2a\xdc\x40\xa7\x66\xe0\x46\x21\x54\xef\x1a\xcc\xb5\xb1\xf2\x99\x1e\x5b\x15\xb8\xc4\xbc\x24\x7a\x13\xc2\x0e\xfe\xdb\xcc\xf0\xde\x52\x2b\xd9\xbf\x37\x3b\x95\xb1\x9c\x16\x3c\x51\xd8\xc1\xbf\xc8\x3e\x0b\xb5\x99\xa2\x0e\x8b\xc3\x90\x78\x18\x3a\x3f\xce\x4c\x97\xb9\x72\xd1\x76\x05\x6d\xc3\xa7\x85\x1d\xfc\xdb\x00\x18\x96\x9c\x57\xc0\x18\x69\xf6\xdc\x62\x2d\x4d\x65\x89\x9e\x5d\x14\x36\x1e\xb4\x29\xec\xe0\x97\x1a\x17\x4f\xaa\x16\xb8\x58\x75\xdc\xe0\x68\x84\x1c\xbc\x5c\xba\xf3\x21\xec\xe0\xc3\x58\x2f\xce\xc4\x69\x46\x1d\x57\xd0\x7a\x31\xb0\xa0\xea\x48\x8f\xcc\x99\x68\x3a\x92\x7c\x03\xe2\x52\x7a\x3a\x78\xcc\x74\x47\xf4\xb8\xde\x69\x69\xf5\x3c\xa2\xf5\x71\x24\x58\x9a\x8f\x51\xef\x20\xff\x41\xe9\x86\x92\x35\x96\x27\x3c\x91\x74\xea\x2c\xa5\xe3\x4a\xd7\x76\xa3\xaa\xf5\xbb\x16\xa8\x12\x30\x9d\x71\x03\xe8\xa3\xd4\x9c\xb9\x9c\xf6\xc0\xa5\x2d\xd9\xf0\x79\x4a\x24\xf5\x84\x5a\x65\xf5\x32\xa8\x68\xdb\xbb\xd7\x1b\xc5\xaa\xc0\x5c\xb3\x01\xd7\x5d\x63\xfc\x62\xce\xde\x2f\xc7\xcc\xd1\xbf\x0a\x9f\x51\xe9\x61\x68\xae\x9f\xdb\xe8\x00\x1b\x48\xed\x87\x9b\x1d\x97\xa1\xca\xe4\x04\xdc\x43\x5b\xe2\x08\xd8\x8c\x43\xe3\xda\x3c\x86\x0d\x58\xa6\xa8\x52\x0b\xc7\x00\x80\x29\xd9\x68\xec\xe1\xc2\x3a\xd6\x45\xad\xed\xd9\x46\xc9\x72\x9c\xd8\x87\x26\xec\xe0\xb7\xad\x17\x2e\x3b\x70\x19\x39\x8e\x10\xeb\x92\x13\x08\xa1\x95\x72\xcc\x2b\x1c\xc9\x19\xc1\x89\x84\x8c\x81\x38\xcf\x52\x67\xe1\x9e\x17\x60\x00\xd3\x85\xa1\x4f\x22\x66\x68\xa4\x06\xd3\x6d\x4e\x7e\xa7\x36\xdb\x79\x9f\x6a\xa6\x16\xc2\xfb\x4d\x46\x1e\x05\x08\x85\x9c\xc2\x6e\x99\xbb\x66\xc4\x8c\xc2\xc3\x6a\x80\x18\xbb\xda\x75\xe0\x5c\x37\x6e\xfc\x08\x37\x7e\x58\x0f\xcc\x8f\x71\x3c\x2a\x25\x9b\x38\xc5\x2f\x04\x8f\x29\x9b\xd1\x80\x6c\x65\x8b\x50\x27\x39\xc6\x4d\xe5\xc2\x96\x82\x2b\x21\x96\xd5\xba\xae\xce\x85\x17\x59\x79\xf8\x7e\x3d\x23\x9e\xc9\xb3\x16\x3e\x8d\xae\x98\xf7\xf2\x5c\xde\x2c\x82\xd0\x54\xcf\xb4\x07\x4a\xac\x26\x3e\x41\xc8\x1d\xdd\x69\x65\x4f\xdc\x1f\x12\x59\x3b\xbf\xb0\x5b\xb7\x3d\x70\x6b\x0b\x6d\xcc\xbb\x0e\xd5\x26\x36\x5c\x4b\x0b\xae\x3f\x06\x47\xa9\x0a\x98\xf9\x54\x36\x44\x46\x6e\xdf\xae\x1c\x57\x3e\xad\x70\xc4\x02\x4a\xd3\x5c\x05\x85\xf3\x6a\x29\x19\x61\xb0\x50\x51\xd3\x88\xb2\xc2\xc3\x1a\xf1\x26\xf5\xd0\x57\x34\x5a\x97\x3d\x1d\x69\x85\x44\x34\x3d\xe7\xcb\x3e\xe8\x68\x0a\x48\xc5\x47\xaa\x0e\x03\x15\x53\x33\x0f\x3f\xa2\x4c\xc3\x92\xbb\xc0\x7e\x88\x75\xa6\x10\x3e\x7e\x27\x57\xeb\x66\x66\x6a\x70\xac\x3a\x3a\xc1\xb9\xdc\x36\x53\x9b\xd1\x9a\xd5\xb5\x12\xe6\x3e\x70\xfd\x65\xb8\x8c\xbe\x26\x81\x4b\xe2\x33\x27\x57\xf9\x06\x42\xb6\x45\x6c\xf8\x37\x9a\x5d\x5d\x99\x64\x68\xbb\x2b\xd5\x01\x1e\x14\x2a\x3c\x3b\x86\x8f\x87\xee\xd2\xb6\xe9\x13\x74\xd1\xb4\x45\x3d\x0c\x1c\x19\xf3\xcd\x9c\x5e\x4f\xc8\x39\xdc\xc6\x75\xae\x4d\xbb\x68\xe3\x33\xfb\x56\xfd\xdf\x44\xc9\x3e\xc4\x58\x97\xd2\x7b\x66\xcc\x35\xdf\x56\xb2\x01\x63\x42\x4f\xe9\x9e\x1d\x76\x7e\xab\x7b\x2b\x99\x8b\x4b\x61\x1d\x06\x7f\xa4\x33\x15\x3d\xc0\xef\x0f\xd6\xa6\x80\xc3\x43\x41\x13\x9a\x02\x84\x61\x11\xef\x73\xa2\x81\x8d\x47\x06\x70\x1f\x4f\xa6\xe4\xc2\xfc\xf7\x19\xe8\x9d\xfc\xd9\xa4\x2b\x92\x6b\x5c\xf8\xf5\x71\x9c\xf0\xb8\xf4\xe6\x77\x15\xec\x94\xd8\x08\x75\xe1\x46\x2f\x6c\x81\x2d\x23\xa1\xb9\x8a\xf6\xe0\xc6\xc9\x0d\x09\xbb\x96\x28\xe1\x64\x35\xef\x76\x7f\xfe\x09\x87\xff\x20\x67\xf8\xeb\xaf\xdd\xee\x00\x6f\x3b\x4c\x59\x49\x8a\x23\x13\x06\xa2\x0c\x7e\x5f\xf1\x71\xc4\x38\x82\xc9\xe5\x36\xfd\x36\xf3\xed\x26\xeb\x79\x3d\xc0\xdb\x9c\x43\xac\xd3\x9c\xd1\x37\xea\x85\x73\x36\x81\x14\xb2\x35\xbb\xe1\xc2\xe5\x4c\x4d\xf9\x84\xba\xbd\x70\x2b\x21\xdb\x1f\xae\x98\x70\x57\x99\x11\x35\x70\xf3\x51\xa1\x04\x85\x22\xb5\x86\xdb\x92\x7a\x54\x2d\x9f\x5b\x82\xc8\x12\x97\x69\xc3\xf0\xe0\x4b\xef\x0a\xbc\x17\xcd\x2d\xd4\x63\xde\x42\xfb\x65\xae\x77\xc3\x10\x1a\xfa\x95\xc6\xaf\x43\x3e\x36\xdb\x4d\xec\xea\x7b\x03\x54\xbc\x35\x36\x25\xd1\x0f\x1c\xc2\x2f\xb7\xde\x2b\x72\xf6\xd6\xd9\x55\x83\x07\x8e\x40\xf7\xee\xfa\xca\x8b\x36\x98\xbd\x43\x4e\x66\x4a\xd0\x68\x46\x41\xa5\xbc\xbe\x20\x95\x97\xb1\x42\xaa\x2e\x53\x43\xcd\xb9\x5e\x7c\x9d\x75\xc7\x9e\xd9\xb7\x02\xc0\x05\x4e\xb5\xa6\x30\xa0\x55\x61\xd1\x06\x34\xd3\x4d\xb8\x56\xdf\xb0\x57\xd9\xb2\xba\x48\xe2\xe2\xeb\xa9\xd9\xa4\x24\x52\x5b\x82\xa5\x73\xc7\xe9\x16\xb6\x5b\x88\x45\x79\xd0\x80\x1f\x1a\x64\xc2\x44\xd2\x46\x9e\x3b\x9d\xdf\xea\x4d\x62\x5f\x90\x27\x6e\x80\x09\x67\x2b\x7d\x90\x3a\xb9\xbf\x4f\xdf\xbb\xf7\x7f\xfe\x34\xd6\x89\x66\x3c\xd1\xe7\x3d\x9c\x49\x2c\x41\xf8\xf1\xf0\xd3\x3e\xe0\x19\x39\x3b\xeb\x51\xe1\xd3\xa8\x3a\xbf\x79\x7a\x7a\x10\xde\xd7\x71\x73\x71\xa8\x72\x7a\xda\x8e\x3e\xfd\xf8\xf4\xd3\xe7\x4f\xdb\x8f\xcf\x21\xdc\xbd\xbf\x81\x7f\x70\x11\x6e\xc7\xfe\xd1\xf4\x31\xda\x53\xf8\x7f\x00\x00\x00\xff\xff\xf9\xe8\x7f\x0d\xd1\x0c\x00\x00")

func templatesConductMdBytes() ([]byte, error) {
	return bindataRead(
		_templatesConductMd,
		"templates/conduct.md",
	)
}

func templatesConductMd() (*asset, error) {
	bytes, err := templatesConductMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/conduct.md", size: 3281, mode: os.FileMode(420), modTime: time.Unix(1497174582, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLicenseMitMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x51\x41\x8e\xe3\x36\x10\xbc\xf3\x15\x05\xe4\xb2\x0b\x08\x93\x7b\x6e\x1c\x89\xb6\x88\x48\xa4\x41\xd1\xeb\xf8\x28\x4b\xb4\xc5\x40\x26\x0d\x91\xce\x60\xb0\x58\x60\x1e\x92\x7d\x40\xbe\x91\xa7\xcc\x4b\x02\x52\xde\xdd\x24\x27\x89\xdd\x5d\x55\x5d\xd5\x3f\x41\x4f\x06\x2d\xd7\x68\xec\x60\x5c\x30\xf8\xd0\x72\xfd\x91\x90\xd2\xdf\x5e\x17\x7b\x99\x22\xfe\xfe\x0b\x9f\x3f\xe3\xe9\x68\xfa\x05\x5f\xbe\xe4\x7f\x7a\x8f\x93\x4f\x2f\x42\x76\x66\xb9\xda\x10\xac\x77\xb0\x01\x93\x59\xcc\xe9\x15\x97\xa5\x77\xd1\x8c\x05\xce\x8b\x31\xf0\x67\x0c\x53\xbf\x5c\x4c\x81\xe8\xd1\xbb\x57\xdc\xcc\x12\xbc\x23\xfe\x14\x7b\xeb\xac\xbb\xa0\xc7\xe0\x6f\xaf\x69\x32\x4e\x36\x20\xf8\x73\x7c\xe9\x17\x83\xde\x8d\xe8\x43\xf0\x83\xed\xa3\x19\x31\xfa\xe1\x7e\x35\x2e\xf6\xd1\x7a\x47\xce\x76\x36\x01\x1f\xe2\x64\xf0\xfe\xf6\x67\xf7\xc0\xbc\xbf\x7d\xfd\x98\x85\x46\xd3\xcf\xb0\x0e\xa9\xff\xad\x89\x17\x1b\x27\x7f\x8f\x64\x31\x21\x2e\x76\x48\x3c\x05\xac\x1b\xe6\xfb\x98\xf6\x78\xb4\x31\xdb\xab\x5d\x55\x32\x3c\x07\x11\x12\xe9\x3d\x98\x82\xa4\x5d\x0b\x5c\xfd\x68\xcf\xe9\x6b\xb2\xb5\xdb\xfd\x34\xdb\x30\x15\x18\x6d\xa2\x3e\xdd\xa3\x29\x10\x52\x31\xe7\x5a\x24\x2f\x3f\xfb\x05\xc1\xcc\x73\x62\xb0\x26\xac\x7e\x7f\x6c\x97\x67\x92\xca\x2d\x85\x1a\x1f\x31\x65\xdd\x97\xc9\x5f\xd3\x2c\xf9\xee\xc4\x06\x9c\xef\x8b\xb3\x61\x32\x19\x33\x7a\x04\x9f\x15\x7f\x37\x43\x4c\x95\x44\x7d\xf6\xf3\xec\x5f\xac\xbb\x90\xc1\xbb\xd1\x26\x47\xe1\x17\x42\xd2\xd1\xfb\x93\xff\xc3\xe4\xdc\xd7\x3b\x3b\x1f\xed\xb0\x46\x9e\x8f\x70\xfb\x71\xd9\x47\x2b\x4c\xfd\x3c\xe3\x64\xc8\x1a\x98\x19\x53\xbc\xa9\xf4\xcd\xce\x92\xe4\x43\xec\x5d\xb4\xfd\x8c\x9b\x5f\xb2\xde\xff\x6d\x3e\x11\xa2\x6b\x86\x4e\x6e\xf4\x81\x2a\x06\xde\x61\xa7\xe4\x27\x5e\xb1\x2a\x1d\x92\x76\xe0\xdd\xfb\xdb\xd7\x02\x07\xae\x6b\xb9\xd7\x38\x50\xa5\xa8\xd0\x47\xc8\x0d\xa8\x38\xe2\x57\x2e\xaa\x82\xb0\xdf\x76\x8a\x75\x1d\xa4\x02\x6f\x77\x0d\x67\x55\x01\x2e\xca\x66\x5f\x71\xb1\xc5\xf3\x5e\x43\x48\x8d\x86\xb7\x5c\xb3\x0a\x5a\x22\x89\x3e\xa8\x38\xeb\x88\xdc\xa0\x65\xaa\xac\xa9\xd0\xf4\x99\x37\x5c\x1f\x0b\x6c\xb8\x16\x89\x73\x23\x15\x28\x76\x54\x69\x5e\xee\x1b\xaa\xb0\xdb\xab\x9d\xec\x18\xa8\xa8\x88\x90\x82\x8b\x8d\xe2\x62\xcb\x5a\x26\xf4\x13\xb8\x80\x90\x60\x9f\x98\xd0\xe8\x6a\xda\x34\x59\x8a\xee\x75\x2d\x55\xde\xaf\x94\xbb\xa3\xe2\xdb\x5a\x93\x5a\x36\x15\x53\x1d\x9e\x19\x1a\x4e\x9f\x1b\xb6\x4a\x89\x23\xca\x86\xf2\xb6\x40\x45\x5b\xba\x65\x19\x25\x75\xcd\x54\x1e\x5b\xb7\x23\x87\x9a\xe5\x12\x17\xa0\x02\xb4\xd4\x5c\x8a\x94\x49\x29\x85\x56\xb4\xd4\x05\xb4\x54\xfa\x3b\xf4\xc0\x3b\x56\x80\x2a\xde\x71\xb1\x25\x1b\x25\xdb\x02\x29\x4e\xb9\xc9\x99\x89\x84\x13\x6c\x65\x49\x51\xe3\x3f\x57\x91\x2a\xbf\xf7\x5d\xfa\x25\xeb\x2e\x15\xa3\x0d\x17\xdb\x2e\x81\xff\x3d\xfc\xf4\x4f\x00\x00\x00\xff\xff\xe4\x1b\xfe\x15\x47\x04\x00\x00")

func templatesLicenseMitMdBytes() ([]byte, error) {
	return bindataRead(
		_templatesLicenseMitMd,
		"templates/license/mit.md",
	)
}

func templatesLicenseMitMd() (*asset, error) {
	bytes, err := templatesLicenseMitMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/license/mit.md", size: 1095, mode: os.FileMode(420), modTime: time.Unix(1497168993, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLicenseUnlicenseMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x54\x4f\x8f\xe2\x36\x14\xbf\xfb\x53\xfc\xb4\xbd\xec\x4a\x29\x73\xef\xcd\x03\x66\xb0\x1a\x6c\xe4\x98\xa5\xa3\xaa\x87\x90\x3c\x06\xb7\xc1\x4e\x6d\xa7\x2b\xfa\xe9\x2b\x07\x66\x66\xa7\xda\x5b\x70\xde\xef\x2f\x2f\xfe\x09\x7b\x3f\xb8\x8e\x7c\x22\x7c\xde\x4d\xc7\xc1\x75\x58\x85\x4b\xeb\xfc\x17\xc6\xec\xd9\x25\xb8\x84\x53\x24\x42\xeb\x7b\x4c\x9e\x7c\x37\x5d\x8e\x14\xa9\x47\x0a\xa7\xfc\xad\x8d\x84\x48\x03\xb5\x89\x7a\x38\x9f\x03\xf2\x99\x30\xde\x88\xfa\x99\x68\xc1\x18\xf7\xd7\xe0\xe9\x8d\x2a\x07\x74\x61\xbc\x56\xb8\x84\xde\x9d\xae\xd5\x6d\x3e\x9d\x2b\x4c\x89\x2a\x74\xe1\x32\xba\x81\x2a\x24\x1a\x86\x0a\x21\xb2\xde\xa5\x1c\xdd\x71\xca\x84\x5c\x3c\xbd\x4a\x57\x20\x97\xcf\x14\xe1\x3c\x52\x98\x62\x47\xe8\x42\x4f\x38\x85\x78\x41\x88\x68\x13\xda\x57\xba\x9e\x1d\x9d\x6f\xe3\xb5\x2a\x6f\xd1\xfa\x2b\xc6\x29\x8e\xe1\x2e\x78\xa1\xd8\xb9\x76\x28\x20\x1f\xfc\xcf\xef\x27\xd5\x1c\xfc\x78\x2d\x08\x76\xa1\xd6\xa7\x05\x63\xd2\xe3\xcf\x29\xba\xd4\xbb\x2e\xbb\xe0\x13\xf2\xb9\xcd\x88\xd4\x85\x17\xef\xfe\xa5\x39\x5d\x74\x2f\xe7\x8c\xa1\xfd\x96\xaa\xb9\x93\x76\xca\xe7\x10\x67\x57\xf3\x53\x62\xe1\xf4\x31\x0d\x7a\xea\x5d\xd7\x66\x9a\xdd\x15\xd9\x76\x18\xbe\xe3\x72\x3e\x53\xa4\x54\x1e\x0a\x23\x7b\xc3\xfd\xb0\x75\x1c\x08\x97\xf6\xaf\x7b\x63\x77\x6a\x17\xfc\x1c\xbf\x8c\x1f\xc9\xd3\xc9\xe5\x9b\x8d\x37\x74\x5b\x3c\xc7\x97\xdb\xff\x7d\x27\xee\x29\x47\x77\x21\x9f\x11\x4e\x08\x53\xc4\x99\x5c\x4c\x65\x82\xa5\xa9\xeb\x28\xa5\x10\xd3\x2c\x58\x2c\x16\xdc\xff\x24\x73\xc0\xb1\x30\x22\xfc\x43\x31\xa3\xed\x0a\x13\x8b\x34\x38\xff\xf7\xe4\xd2\x79\xe6\x76\x1e\x23\xc5\x91\xf2\xe4\xf2\xb5\x28\x95\xf4\x63\xa4\x54\x5e\x16\x37\xa7\x29\x4f\x65\xdd\x4a\x19\xe9\xe6\xcd\xa5\xf7\x16\x26\xdf\x53\xfc\x58\xfd\x82\x31\xbb\x11\x68\xf4\xda\x1e\xb8\x11\x90\x0d\x76\x46\x7f\x95\x2b\xb1\xc2\x27\xde\x40\x36\x9f\x2a\x1c\xa4\xdd\xe8\xbd\xc5\x81\x1b\xc3\x95\x7d\x86\x5e\x83\xab\x67\xfc\x2a\xd5\xaa\x62\xe2\xb7\x9d\x11\x4d\x03\x6d\x20\xb7\xbb\x5a\x8a\x55\x05\xa9\x96\xf5\x7e\x25\xd5\x13\x1e\xf7\x16\x4a\x5b\xd4\x72\x2b\xad\x58\xc1\x6a\x14\xc1\x3b\x95\x14\x0d\xf4\x9a\x6d\x85\x59\x6e\xb8\xb2\xfc\x51\xd6\xd2\x3e\x57\x58\x4b\xab\x0a\xe7\x5a\x1b\x70\xec\xb8\xb1\x72\xb9\xaf\xb9\xc1\x6e\x6f\x76\xba\x11\xe0\x6a\x05\xa5\x95\x54\x6b\x23\xd5\x93\xd8\x0a\x65\x17\x4c\x2a\x28\x0d\xf1\x55\x28\x8b\x66\xc3\xeb\x7a\x96\xe2\x7b\xbb\xd1\xa6\xc1\xa3\x40\x2d\xf9\x63\x2d\x6e\xac\xea\x19\xcb\x9a\xcb\x6d\x85\x15\xdf\xf2\xa7\x62\xc4\x30\x6d\x37\xc2\xcc\x63\x77\x23\x87\x8d\x98\x8f\xa4\x02\x57\xe0\x4b\x2b\xb5\x2a\xf1\x97\x5a\x59\xc3\x97\xb6\x82\xd5\xc6\x96\xec\x33\xf4\x20\x1b\x51\x31\x6e\x64\x53\xb2\xaf\x8d\xde\x56\x28\xcd\xe9\xf5\x5c\x8f\x2a\x38\x25\x6e\x2c\xa5\x55\x7c\x28\x5f\x9b\xf9\xf7\xbe\x11\xef\x5e\x56\x82\xd7\x52\x3d\x35\x05\xfc\xfd\xf0\x82\xb1\x75\x88\xb8\x84\x58\x76\xaa\x7c\xce\xf3\x22\x55\x18\xe7\xab\x06\x91\x4e\x14\xcb\x0a\xfc\x3e\xbd\x5e\x5f\x8b\x10\x5f\xfe\xf8\x7c\xce\x79\xfc\xe5\xe1\xe1\xc3\xe9\xc3\x17\xf6\x5f\x00\x00\x00\xff\xff\x42\x13\x92\x2d\xe7\x04\x00\x00")

func templatesLicenseUnlicenseMdBytes() ([]byte, error) {
	return bindataRead(
		_templatesLicenseUnlicenseMd,
		"templates/license/unlicense.md",
	)
}

func templatesLicenseUnlicenseMd() (*asset, error) {
	bytes, err := templatesLicenseUnlicenseMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/license/unlicense.md", size: 1255, mode: os.FileMode(420), modTime: time.Unix(1497167702, 0)}
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
	"templates/conduct.md": templatesConductMd,
	"templates/license/mit.md": templatesLicenseMitMd,
	"templates/license/unlicense.md": templatesLicenseUnlicenseMd,
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
		"conduct.md": &bintree{templatesConductMd, map[string]*bintree{}},
		"license": &bintree{nil, map[string]*bintree{
			"mit.md": &bintree{templatesLicenseMitMd, map[string]*bintree{}},
			"unlicense.md": &bintree{templatesLicenseUnlicenseMd, map[string]*bintree{}},
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

