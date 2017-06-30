// Code generated by go-bindata.
// sources:
// plugins/plugins.toml
// plugins/templates/python/.dockerignore
// plugins/templates/python/.gitignore.template
// plugins/templates/python/CHANGELOG.md
// plugins/templates/python/Dockerfile
// plugins/templates/python/LICENSE.md
// plugins/templates/python/README-short.txt
// plugins/templates/python/README.md
// plugins/templates/python/circle.yml
// plugins/templates/python/plugin.toml
// plugins/templates/python/requirements.txt
// plugins/templates/python/scan.py
// DO NOT EDIT!

package plugins

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

var _pluginsPluginsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x5d\x73\xda\x38\x14\x7d\xe7\x57\xdc\x31\x2f\xbb\x3b\x21\x10\x0a\xa1\xed\x4c\x1e\x58\x42\xda\x74\xf3\x35\x98\x64\xa7\x93\xc9\xa4\xb2\x7d\x6d\x6b\x63\x5b\x1e\x49\x86\x90\x5f\xbf\x23\xd9\x60\x3e\xdc\x0d\x22\x29\x9b\x87\x18\x74\x7d\xcf\xbd\xe7\xe8\x48\x96\xa9\xc3\x80\xa5\x33\x4e\x83\x50\xc2\x6f\xee\xef\xd0\x6e\x1d\x7d\x80\x86\xba\x1c\x83\x13\x11\xf7\x49\xb2\x14\xbe\x31\x11\x66\x04\x2e\x09\x4d\xf0\x00\xfa\x51\x04\x23\x95\x20\x60\x84\x02\xf9\x04\xbd\xc3\x5a\x1d\x6c\x44\xb8\x38\x1f\x0c\xaf\xec\x21\xf8\x8c\x43\x44\x5d\x4c\x04\x02\x4d\x7c\xc6\x63\x22\x29\x4b\x0e\x6b\xb5\xfa\xfb\xfc\xd5\xea\x70\x73\x71\xfb\xe5\xfc\x0a\x06\x2c\xf1\x69\x90\x71\x5d\x00\xcc\x71\xde\xa9\x9f\x9a\xa4\x32\x42\x38\x01\xeb\x92\x28\xe6\x70\x13\x65\x01\x4d\x56\xdb\x13\x56\xad\x76\x7f\x9f\xea\xc8\xc3\x43\x0d\x00\x13\xe2\x44\xe8\xc1\x09\x48\x9e\x61\x0d\x20\x21\xb1\x06\x49\x04\x8f\xac\x1a\x80\x87\xc2\xe5\x34\xd5\xdc\x4e\xc0\xba\xb2\x47\x17\x70\x4a\x24\x71\x88\x40\xf8\x4a\x44\x08\x36\x12\xee\x86\xea\x5e\x97\x48\x0c\x18\x9f\xa9\x1b\x69\x22\x51\x03\xd0\x98\x04\x1a\x31\xd6\x6d\x35\x15\xf0\x67\x11\x92\x23\x15\xe4\x98\x32\x41\x65\x91\x13\x4a\x99\x8a\xcf\xcd\x66\x40\x65\x98\x39\x87\x2e\x8b\x9b\x79\x12\x65\xc5\x87\x86\xca\x3e\x0c\xa8\x54\xc9\x4e\x46\x23\xd5\xb9\x4f\x22\xa1\x5a\x77\x63\xf5\xcd\x8a\x18\x7b\xca\x52\x75\x43\x4c\x73\x2e\x21\x11\xba\x3f\x75\x95\xb3\x14\x05\x9c\xc0\x3d\x58\xba\x09\xc8\x55\x98\xa8\x21\xeb\xb2\xaf\xfc\xf3\x38\x3e\xbf\x1c\x5e\xdf\x8e\xad\x87\x6d\xc5\x9a\x50\x9e\x09\xc9\x24\xa9\x92\xec\x4e\x05\xc7\x2a\x08\x0d\xf0\x69\x84\x02\x84\x4b\x12\x20\x89\xa7\x3b\x82\xb2\xe1\xad\x04\x5c\x2d\x66\xac\x60\x99\xfe\x33\x1d\xb3\x34\x62\x64\x31\x00\x6a\x61\xc9\xf9\xa0\x64\xf0\x43\x91\xff\x01\xd4\x87\x19\xcb\x60\x4a\x12\xa9\x46\x8b\xb8\x20\x71\x1a\xa1\x1a\x58\xaa\xe3\xb2\xb8\x06\x40\x52\xfa\x84\xba\xcb\x76\xf7\xc3\xa7\xee\xd1\xb1\xd7\xe9\x1d\x79\x3d\x07\x9d\x63\xa7\xfd\x91\xf4\xda\x2d\xaf\x47\xda\xdd\x56\xbb\x83\x9e\xdb\xf2\x7b\xdd\x4f\x2d\xef\x43\xa7\xeb\xbb\xbd\x4e\xaf\x73\xf4\xf1\xb8\xd3\x6d\xb5\xda\xc4\xed\xf4\x1c\x0b\xea\x30\x0e\xa9\x00\x2a\x80\x80\x44\x21\xe1\x09\x67\x07\x90\x46\xa8\x5c\xc9\x31\x8d\x88\x8b\x30\xa5\x32\x54\x3d\x72\x60\xd3\x64\x07\x7f\xc4\x5e\xd7\x3a\x28\x6c\x92\x5f\xdb\xdd\xe3\x4a\xc3\xdc\x8d\x1f\xfb\x37\xe7\xea\xa6\xad\x1d\x34\x17\x7b\x6e\x21\xad\xd4\xbc\x91\x35\x07\xd5\x17\x41\x68\xec\x64\x99\x15\x70\x63\xc7\x2c\xb2\x7f\xad\x61\xca\x32\xee\x2c\xe6\x59\xe1\x9a\x4c\x20\x57\x7d\xaa\xba\x85\x7d\xde\xb4\xb4\xbf\x3e\xde\xda\xc3\xd1\xf2\x54\x7d\x7d\xfc\x6b\xf8\xdd\x68\xee\xd6\x56\xbf\x08\x89\xc7\xa6\x0d\xfd\x20\xe2\x15\xd3\x67\xeb\xb8\xad\xc3\x3b\x4e\xe0\x46\x09\xe3\x49\x5c\x41\xf8\xd9\x44\x1a\x2c\x88\xdd\xf6\xcd\x0d\xd7\x23\x89\x1b\x7a\xba\x2b\x74\x1b\x23\x89\x07\x2a\xb6\xab\xeb\x57\xc0\xb7\x54\xac\x91\xf7\x2d\x96\xb2\xff\x37\xb5\xd6\x7c\xa6\x9e\x1d\xea\x2c\x53\x21\x95\x10\x1e\x62\xda\x1c\x8f\xe8\x69\x13\x9f\xa9\x2f\x19\x8b\xd6\x75\x8a\x51\x12\x8f\x48\x52\x25\xd5\x32\xb4\xb1\xb5\xe6\xc9\xaf\xe9\xf4\x87\xf5\x0e\x22\xcc\x08\x27\x15\x02\x7c\xef\x8f\xfa\x60\xbb\x24\x59\x67\x4d\x26\x55\x7c\xe7\x28\xc6\x5c\x55\xe2\x5e\x78\x92\x09\x11\xb2\x82\x68\x5f\x8d\x43\x3f\x91\x54\x9f\x2b\xb6\xa3\xbb\x00\x33\xe6\xab\x33\xf7\x44\x38\xa8\xa2\x7b\xf7\xc5\x9c\x6c\xb0\x23\xd5\x60\x2f\x44\x1d\x2a\x3d\xf4\x31\xf1\x2a\x1f\x16\x7f\x96\x51\x53\xe2\x6b\xc0\xc6\x02\x2c\xe5\xef\x45\x08\x37\x22\x71\xce\x64\x4d\x83\x41\x44\xe2\xfe\xdd\x76\xa4\x4b\x10\x63\xbe\x79\xea\x7e\xa8\xb2\x98\x79\x55\x1b\xf7\x40\x07\x4c\x67\xba\x84\x33\x27\xad\x53\xf7\x42\xda\x4f\x39\xab\xda\xc2\xce\x1a\x37\xa3\xeb\xb1\x29\xe7\x05\x9a\xf9\xf3\x49\x65\xee\x87\xb1\x40\x37\xe3\x58\xc9\xd9\xd6\x21\x63\xd6\x25\xa2\x39\xef\x3c\x77\x2f\xcc\x05\x4b\x43\x26\xaa\x0e\xbf\x3a\x60\x4a\xbb\x84\x33\x3f\xe8\xea\xd4\x5f\x43\x7a\xfd\xf8\x3a\xa5\x89\xc7\xa6\xa2\xf1\x1f\xfb\xf9\xdf\xf9\x2d\x70\xba\xe3\xa6\x5e\x55\xc2\xf4\x28\x5b\x7c\x5d\x87\xda\x8f\x46\xf8\x5c\xb5\x1e\x6e\x86\xd0\x00\x75\x4a\xd5\x6f\x7e\x9c\x2a\xd2\x29\xe3\x52\xa1\x00\x3e\xa3\x9b\xe9\x8f\x1b\x22\x15\x68\xeb\x2a\x15\xc3\xa6\xc2\xe0\xf3\xab\xab\x83\xa4\x69\x44\x5d\xfd\xe3\x59\xf3\xb9\xe1\x31\xa1\x9a\x7b\x97\xbd\x22\x62\xa2\x6a\xc1\x9c\x51\x8e\xc3\x19\xc2\x05\x71\x04\x5c\x3b\x7e\x26\x14\x7f\x0f\x6c\xc9\x69\x12\x80\xcd\xa2\xe2\x15\x70\x0b\x59\x16\x35\xcc\xf7\x0e\x95\xb9\x3f\x6d\xd6\x4d\xc3\x7c\x9f\xba\x55\xbe\xb9\xd6\x81\x4d\xef\x5c\x5f\x0c\x9b\xa3\xf1\x19\x78\xcc\xcd\x62\x4c\xe4\x86\x71\xe6\x81\x2a\x99\xca\x6a\xa6\x06\xca\x33\xf7\xb3\x90\x52\xcf\xaf\x5a\x48\xa7\x67\x9b\x6a\xa8\xc1\x9d\x94\x28\x6a\x98\xca\x90\x7a\xbe\x89\x57\x8a\x2a\x6f\x56\xe4\x1f\xf5\x76\xa2\xc5\xa8\x10\xe6\xdb\x22\xb8\xa9\xcf\x37\x1b\xf2\x90\x91\x38\xab\xe5\x4c\x35\x2a\xb3\x4d\xa4\x5a\xad\xf9\x66\xc5\x08\x77\x43\x3a\xa9\x5a\x58\xfd\x3c\xb2\xa4\x55\x96\x14\x77\x43\x71\xdd\x7c\x62\x95\x68\x1b\x2f\x61\x65\xc8\x54\xa8\x22\xf5\x55\x95\xca\x0a\x2f\x34\x5d\xfa\xdd\x45\x12\x6e\x1d\x80\x15\xbc\xa8\xff\x69\xef\xe5\x95\x1f\x5f\xfe\x0d\x00\x00\xff\xff\x21\x0b\xfd\xad\x80\x1a\x00\x00")

func pluginsPluginsTomlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsPluginsToml,
		"plugins/plugins.toml",
	)
}

func pluginsPluginsToml() (*asset, error) {
	bytes, err := pluginsPluginsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/plugins.toml", size: 6784, mode: os.FileMode(420), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonDockerignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xf0\x4c\xcf\xcb\x2f\x4a\x55\xd0\x4b\xcf\x2c\x51\x48\xcb\xcf\x49\x49\x2d\xe2\x02\xb1\xb5\xb8\xc0\x54\x26\x58\x96\x2b\xc8\xd5\xd1\xc5\xd7\x55\x2f\x37\x85\xcb\xc7\xd3\xd9\xd5\x2f\x18\xcc\x74\xf6\x70\xf4\x73\x77\xf5\xf1\x77\x07\x71\x92\x4a\x33\x73\x52\xb8\x8a\x52\x73\x52\x13\x8b\x53\xb9\x7c\x13\xb3\x53\xd3\x32\x73\x52\xb9\x92\x33\x8b\x92\x73\x52\xf5\x2a\x73\x73\x00\x01\x00\x00\xff\xff\xba\xc0\x2a\xa8\x6a\x00\x00\x00")

func pluginsTemplatesPythonDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonDockerignore,
		"plugins/templates/python/.dockerignore",
	)
}

func pluginsTemplatesPythonDockerignore() (*asset, error) {
	bytes, err := pluginsTemplatesPythonDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/.dockerignore", size: 106, mode: os.FileMode(420), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonGitignoreTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x52\xb1\x6e\x1b\x31\x0c\xdd\xf9\x15\x04\xb2\x19\xb1\xb4\x14\x1d\x3a\xa6\x5e\x0a\x64\x08\x90\x66\x2a\x0a\x43\x96\x78\x77\x4c\x75\xa2\x20\xd2\x8e\xaf\x5f\x5f\x48\x4e\x9b\x2e\x77\x8f\x7c\xe2\xd3\xa3\x48\x77\x78\x3e\x3e\x9b\x34\x02\xb8\xc3\x87\xcd\x68\x1f\x65\xad\x9c\x29\xa1\x47\xa9\xc6\x2b\xff\x1e\xf8\xf0\xf8\x88\x13\x67\x52\x38\x1e\xeb\x16\x43\x5c\xe8\x78\xf4\xb0\x73\x75\xfb\x11\x25\xfd\xec\xf5\x5f\x91\xae\x46\x45\x59\x8a\xc2\xce\xa9\xf4\xe4\x81\xd5\x1a\x9f\xce\xc6\x52\xd0\x63\x0d\xf1\x57\x98\xb9\xcc\xe0\x9e\x36\x5b\xa4\x00\x95\x8b\x87\xd3\x99\x73\xf2\x90\xe8\x42\x59\xea\x9e\xe6\x59\x3d\x24\x56\xf3\x90\xe4\xad\x64\x09\x49\x3d\xdc\xd2\xee\xf6\xcb\x7c\x1a\x9f\xcf\x9f\x3c\xd4\xd0\x4c\x3d\xe8\xad\xe2\x12\x5a\x77\x46\xf3\xbc\xe7\x32\x89\x07\xc7\x45\x2d\xe4\x4c\xc9\xc5\x69\xbe\x51\xdd\xdb\xd3\xf6\xed\x9d\x68\x70\x87\xf8\xa2\xe7\x90\xf3\x86\xb6\x90\xd2\xad\x5b\x0c\x8d\xf0\xad\xb1\x19\x15\x3c\x6d\x18\xb0\x0e\xd7\xa8\xb1\x71\x35\x9c\x9a\xac\x18\xd0\x68\xad\x39\x18\x75\x95\x13\x4d\xd2\xe8\x7f\x6d\x1c\xdd\x69\xd7\x45\xba\xd2\x3d\xaa\x60\x50\x34\x41\x2e\xaf\x14\x0d\x53\x30\xf2\x62\x0b\x35\xec\x86\x15\xb9\x74\xd2\x1c\xec\xdc\x1a\x0a\x4f\xa4\xd6\x1f\xb4\x52\xec\xb6\x3f\x84\xb3\xcc\x0a\x95\xeb\x3e\xcb\xec\xec\x6a\x03\x27\xca\x64\xb4\xb7\x85\x75\x9f\xb8\x51\x34\x69\xdb\x60\xe1\x0e\x5f\x0a\x1b\x1a\xa9\xa1\xc7\x28\x17\x6a\x61\x26\x6c\x54\xa5\x99\xc2\x62\x6b\x8e\x72\xf1\xe0\x4c\xae\x1e\xdc\xdf\x03\x1f\xc8\xed\xc0\x8d\xe1\x43\x11\xa5\xae\xa3\xee\xba\x66\xf8\xc7\xf7\x60\x77\x3f\xc2\x7e\xdd\xf7\x16\x8a\xe6\x60\xef\x2b\xb1\x4a\xdf\x18\x19\x4e\x0e\xaf\xa1\xcc\x82\x6a\xe7\x69\xfa\x02\x3b\x97\x65\xcc\xe4\xb9\x2e\x5c\xae\x98\x24\x9e\x57\x2a\x36\x4a\x21\x49\x54\x7f\x7c\x5f\x92\x31\xb8\x87\x8e\xa9\x81\x85\x36\x93\xf9\x3f\x01\x00\x00\xff\xff\x27\xc3\x36\xdb\xc8\x02\x00\x00")

func pluginsTemplatesPythonGitignoreTemplateBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonGitignoreTemplate,
		"plugins/templates/python/.gitignore.template",
	)
}

func pluginsTemplatesPythonGitignoreTemplate() (*asset, error) {
	bytes, err := pluginsTemplatesPythonGitignoreTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/.gitignore.template", size: 712, mode: os.FileMode(493), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonChangelogMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x56\x30\xd0\x33\xd4\x33\x50\xd0\x55\x70\xcb\x2c\x2a\x2e\x51\x08\x4a\xcd\x49\x4d\x2c\x4e\xe5\xd2\x52\x70\x2d\x4b\x2d\xaa\x54\x48\x4b\x4d\x2c\x29\x2d\x4a\x55\x48\x4c\x49\x49\x4d\x81\x8b\x26\x95\xa6\x2b\xa4\x65\x56\xa4\xa6\x70\x01\x02\x00\x00\xff\xff\x82\x96\x39\xb5\x41\x00\x00\x00")

func pluginsTemplatesPythonChangelogMdBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonChangelogMd,
		"plugins/templates/python/CHANGELOG.md",
	)
}

func pluginsTemplatesPythonChangelogMd() (*asset, error) {
	bytes, err := pluginsTemplatesPythonChangelogMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/CHANGELOG.md", size: 65, mode: os.FileMode(493), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\xab\xe2\x30\x14\x85\xf7\xf9\x15\x97\x2e\x5c\x0c\xa4\xdd\x0f\xb8\xe8\xb4\x1d\x0c\x8e\x6d\x89\xce\x0c\xe2\x48\x89\xed\xb5\x06\xd3\x34\x93\xa6\xa8\x14\xff\xfb\x50\xe7\x55\x1e\x3c\x79\xbb\xe4\xdc\xef\x72\xcf\x39\xdf\x79\xb6\x82\x46\x28\x59\x62\x20\x94\x91\x1a\xbf\x3a\xa9\x25\x21\xab\x90\xa5\x9b\x90\xa5\x09\x87\x61\x80\xd2\xa2\x70\xad\x85\xfb\x9d\x90\x28\xcb\xb7\xe0\x43\xe0\x1a\x13\x74\xb6\x0c\x86\x01\x8c\xea\x6b\xa9\x0b\x2d\x1a\x1c\x11\xfe\x33\x05\x61\xce\x54\xea\xce\x09\xa5\xc0\xdc\xdc\xa9\xd5\x1f\x64\xea\xc0\x3f\xf4\x52\x55\xb4\x42\xd3\x41\x2d\x1d\x34\x68\xcb\xde\x4a\x31\xee\x50\x23\x0d\xfc\x21\x00\xb3\x19\x74\xe8\x80\x5e\xa7\x5f\x59\x7d\x72\x7d\x82\xf0\x6a\x5a\xeb\x20\x67\x79\x91\x66\x45\x14\x46\x8b\xa4\x88\x19\x9f\xb7\xc7\xe3\x0b\x24\x66\xeb\xf0\xdb\x8f\xa4\x18\xdf\xbf\x12\xbe\x66\x59\x5a\x44\x8b\x24\x5a\xce\x5b\x3d\xe1\xa3\x9f\xa7\x77\xda\x9b\xda\x8a\x0a\x1f\xea\xe5\x84\xa8\x5e\x62\x16\x2c\xfe\xed\xa5\xc5\x06\xb5\xeb\x7c\x77\x75\x4f\xec\x51\xca\x98\xac\x37\xbe\xb9\x3d\x57\xde\xc6\xb6\x01\x6a\x8f\xff\x73\x7e\x99\x44\x61\xce\x50\xe1\x78\xdd\xf4\xb6\xc6\xf7\xf5\x11\xf2\x3b\xe3\xcb\x98\x71\x08\x1a\xa1\x2e\xc2\x22\x21\x49\xba\xe1\xdb\x3c\x63\xe9\x06\x76\x5e\x70\x90\x3a\xe8\x4a\xa1\xbd\x3d\x21\xd1\x2a\x86\x9d\x47\xe9\x09\x95\xf1\xf6\xff\x02\x00\x00\xff\xff\x22\x74\x40\xdc\x05\x02\x00\x00")

func pluginsTemplatesPythonDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonDockerfile,
		"plugins/templates/python/Dockerfile",
	)
}

func pluginsTemplatesPythonDockerfile() (*asset, error) {
	bytes, err := pluginsTemplatesPythonDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/Dockerfile", size: 517, mode: os.FileMode(420), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonLicenseMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x51\x4f\x8f\xe3\x26\x14\xbf\xfb\x53\xfc\x34\xa7\x5d\xc9\x9a\xde\x7b\x63\x62\x1c\xa3\x3a\x10\x61\xb2\x69\x4e\x15\xb1\x49\x4c\xe5\x40\x04\xa4\xa3\x68\xb5\xdf\xbd\x82\x64\x3b\x3b\x3d\x59\xe6\xbd\xdf\xdf\xb7\xf2\xd7\x7b\xb0\xe7\x39\xe1\xcb\xf8\x15\xdf\xbf\x63\xbc\x85\x60\x5c\xfa\xeb\x6e\x74\xc0\x8f\x1f\xe5\x29\x18\x9d\x7c\xfe\xab\xaa\xad\x09\x17\x1b\xa3\xf5\x0e\x36\x62\x36\xc1\x1c\xef\x38\x07\xed\x92\x99\x6a\x9c\x82\x31\xf0\x27\x8c\xb3\x0e\x67\x53\x23\x79\x68\x77\xc7\xd5\x84\xe8\x1d\xfc\x31\x69\xeb\xac\x3b\x57\x1a\xa3\xbf\xde\xf3\x66\x9a\x6d\x44\xf4\xa7\xf4\xae\x83\x81\x76\x13\x74\x8c\x7e\xb4\x3a\x99\x09\x93\x1f\x6f\x17\xe3\x92\x4e\x59\xef\x64\x17\x13\xf1\x25\xcd\xa6\x7a\x19\x9e\x88\x97\xaf\x45\x64\x32\x7a\x81\x75\x48\xb3\xc1\xcf\x11\xde\x6d\x9a\xfd\x2d\x21\x98\x98\x82\x1d\x33\x47\x0d\xeb\xc6\xe5\x36\x65\x0f\x3f\xc7\x8b\xbd\xd8\xa7\x42\x86\x97\x32\x62\x26\xbd\x45\x53\x17\x9f\x35\x2e\x7e\xb2\xa7\xfc\x35\x25\xd6\xf5\x76\x5c\x6c\x9c\xeb\x6a\xb2\x99\xfa\x78\x4b\xa6\x46\xcc\x8f\xa3\x71\x19\xa5\xdd\xf4\x9b\x0f\x88\x66\x59\x32\x83\x35\xf1\x91\xf5\xc3\x5d\xd9\x41\xf2\xd5\x35\x17\x9a\x9e\x15\x15\xdd\xf7\xd9\x5f\x3e\x27\xb1\x11\xa7\x5b\x70\x36\xce\x66\x2a\x71\x3d\xa2\x2f\x8a\x7f\x9b\x31\x65\x96\xbc\x7e\xf2\xcb\xe2\xdf\xad\x3b\x63\xf4\x6e\xb2\x39\x51\xfc\xbd\xaa\xd4\x6c\xa0\x8f\xfe\x1f\x53\xb2\x3c\x6e\xed\x7c\xb2\xe3\xa3\xee\x72\x80\xeb\xc7\x55\x9f\xa3\x38\xeb\x65\xc1\xd1\x54\x8f\xc2\xcc\x94\xeb\xd5\xbf\xc4\x09\x59\x3e\x26\xed\x92\xd5\x0b\xae\x3e\x14\xbd\xff\xc7\x7c\xad\x2a\xd5\x51\x0c\xa2\x55\x7b\x22\x29\xd8\x80\xad\x14\xdf\x58\x43\x1b\xbc\x90\x01\x6c\x78\xa9\xb1\x67\xaa\x13\x3b\x85\x3d\x91\x92\x70\x75\x80\x68\x41\xf8\x01\x7f\x30\xde\xd4\x15\xfd\x73\x2b\xe9\x30\x40\x48\xb0\xcd\xb6\x67\xb4\xa9\xc1\xf8\xaa\xdf\x35\x8c\xaf\xf1\xb6\x53\xe0\x42\xa1\x67\x1b\xa6\x68\x03\x25\x90\x05\x9f\x54\x8c\x0e\x10\x6d\xb5\xa1\x72\xd5\x11\xae\xc8\x1b\xeb\x99\x3a\xd4\x68\x99\xe2\x99\xb3\x15\x12\x04\x5b\x22\x15\x5b\xed\x7a\x22\xb1\xdd\xc9\xad\x18\x28\x08\x6f\x2a\x2e\x38\xe3\xad\x64\x7c\x4d\x37\x94\xab\x57\x30\x0e\x2e\x40\xbf\x51\xae\x30\x74\xa4\xef\x8b\x14\xd9\xa9\x4e\xc8\xe2\x6f\x25\xb6\x07\xc9\xd6\x9d\x42\x27\xfa\x86\xca\x01\x6f\xb4\xea\x19\x79\xeb\xe9\x43\x8a\x1f\xb0\xea\x09\xdb\xd4\x68\xc8\x86\xac\x69\x41\x09\xd5\x51\x89\xbc\xf6\x74\xb7\xef\x68\x79\x62\x1c\x84\x83\xac\x14\x13\xbc\x12\x2d\x56\x82\x2b\x49\x56\xaa\x86\x12\x52\xfd\x07\xdd\xb3\x81\xd6\x20\x92\x0d\xb9\x90\x56\x8a\x4d\x8d\x5c\xa7\x68\x4b\x67\x3c\xe3\x38\x7d\xb0\xe4\xaa\xf1\xe9\x22\x42\x96\xff\xdd\x40\x3f\xbc\x34\x94\xf4\x8c\xaf\x87\x0c\xfe\x75\xf9\xb5\xfa\x37\x00\x00\xff\xff\xe4\x3f\x0e\xa2\x2f\x04\x00\x00")

func pluginsTemplatesPythonLicenseMdBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonLicenseMd,
		"plugins/templates/python/LICENSE.md",
	)
}

func pluginsTemplatesPythonLicenseMd() (*asset, error) {
	bytes, err := pluginsTemplatesPythonLicenseMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/LICENSE.md", size: 1071, mode: os.FileMode(493), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonReadmeShortTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x28\xc8\x29\x4d\xcf\xcc\x8b\x4f\x49\x2d\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x53\xa8\xad\x05\x04\x00\x00\xff\xff\xf7\x36\xbe\xfa\x18\x00\x00\x00")

func pluginsTemplatesPythonReadmeShortTxtBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonReadmeShortTxt,
		"plugins/templates/python/README-short.txt",
	)
}

func pluginsTemplatesPythonReadmeShortTxt() (*asset, error) {
	bytes, err := pluginsTemplatesPythonReadmeShortTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/README-short.txt", size: 24, mode: os.FileMode(420), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcb\xb1\x4e\xc5\x20\x14\x80\xe1\x9d\xa7\x38\xc6\x45\x07\xa0\xb4\x29\xbd\xb8\xf9\x00\x26\xee\xc6\x18\x38\x1c\x5a\x92\x5e\x68\x0a\x0c\xa6\xe1\xdd\x4d\x8c\x83\xeb\xff\xe7\x7b\x84\xeb\x82\x63\x6f\x6b\x4c\x5f\xc9\xde\x09\x7a\x67\xec\xcd\xee\x11\xe9\xdf\xf1\x54\xf0\x8c\x47\x8d\x39\x41\xef\xf0\xfe\x5b\x05\x63\x0f\x1f\xaf\x50\xf0\x24\x4a\x65\xcb\x15\x72\x80\xef\xdc\xce\x3f\xf5\xf9\xb4\xd5\x7a\x94\x17\x29\x83\xc0\x3d\x37\x2f\xd6\x58\xb7\xe6\x04\xe6\xbb\xb4\xa5\x50\x2d\x52\x1b\xa5\x8d\x1c\x47\x33\x8c\xf3\x20\x71\x9a\xfd\x4d\x2f\x96\xdb\x41\x2d\x5c\x29\x9a\xf8\x4d\x3b\xe2\xe8\x17\x9c\x5d\x98\x42\x30\x4e\xac\x31\x3c\xb3\x9f\x00\x00\x00\xff\xff\x6d\x30\xd9\xc7\xb7\x00\x00\x00")

func pluginsTemplatesPythonReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonReadmeMd,
		"plugins/templates/python/README.md",
	)
}

func pluginsTemplatesPythonReadmeMd() (*asset, error) {
	bytes, err := pluginsTemplatesPythonReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/README.md", size: 183, mode: os.FileMode(493), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonCircleYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xcd\x6e\xc2\x30\x10\x84\xef\x79\x8a\x15\xed\xa1\x3d\x84\xdc\x41\xea\x8b\x20\x84\x16\x7b\x48\x56\xf8\x27\x5a\x3b\x91\x2a\x44\x9f\xbd\x22\x38\x14\x24\xca\xcd\x1a\xef\x7c\x33\x1a\xcf\xa6\x93\x80\x55\x45\x94\xa0\xa3\x18\xa4\xcb\x9b\xa8\x26\x1b\xcd\x11\x5a\x55\x16\x3d\x82\x45\x30\x72\xfd\x33\x6c\x3a\xec\xac\x28\x4c\x8e\x2a\x7f\x86\xc5\x4f\x73\xf5\x2c\x2a\xa2\x38\x42\x55\x2c\x1e\x69\x24\xe1\x10\x8b\x22\x07\xda\x6c\xa8\x06\xcd\xb6\x46\x3c\xb7\x58\x66\x56\xda\x6e\xd7\x94\x3b\x84\xd9\xe6\x22\x5b\xaa\x6b\x09\xfd\x90\x9f\xdc\xaf\xe9\x20\x8f\x39\xfb\x41\x9c\xa5\x3a\x93\x11\x35\x0e\x46\x9a\xd3\x89\x7a\x37\xb4\x12\x76\x81\x3d\xe8\x7c\xa6\x65\xb1\xf8\xa3\x15\xa5\xba\xbf\x81\xd7\x33\x25\xf1\x88\x57\x84\xaf\x27\x5d\xaa\x2a\x23\xe5\xd5\xff\x13\xe8\x10\x5e\x31\x2f\xee\x69\x74\x17\xbf\x3d\xc2\x44\x7a\x23\xcf\x29\x43\xaf\x6f\xa2\xbd\x72\x30\xdd\xaa\xa8\x45\x34\xd1\x7b\x0e\x36\xcd\x47\x97\x48\x8c\xec\xe8\xfd\xe3\x2e\xb9\x75\x62\xa1\x8e\xf7\xa9\xe9\xb9\xc5\xb4\x13\xf4\x56\xa8\x36\xde\x7e\x56\x44\x0a\x07\x4e\xa5\xfa\x1c\x57\xc4\x49\xbb\x4f\x2b\x33\xf2\x11\xb7\x93\xdf\x00\x00\x00\xff\xff\x5b\x80\xf9\x05\x58\x02\x00\x00")

func pluginsTemplatesPythonCircleYmlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonCircleYml,
		"plugins/templates/python/circle.yml",
	)
}

func pluginsTemplatesPythonCircleYml() (*asset, error) {
	bytes, err := pluginsTemplatesPythonCircleYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/circle.yml", size: 600, mode: os.FileMode(420), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonPluginToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xcb\x6a\xeb\x30\x10\x86\xf7\x7e\x8a\x41\xd9\x1e\x0e\xb4\x90\xee\xb2\x28\xa1\x94\x40\x5a\xba\xe9\xca\x04\xa3\xc4\x13\x67\x88\x6e\x48\xb2\x4b\x08\x7e\xf7\xa2\x8b\xd3\xd8\xe9\xa6\x10\x6f\x0c\xdf\x7c\x1a\xfd\x9a\x99\xdd\xe7\x2b\x66\xf0\xb1\xfe\x7c\x5d\xbd\xc3\x52\xab\x3d\x35\xad\xe5\x9e\xb4\x82\xbf\xf7\xb9\x53\x9e\xa2\x2c\x8d\x68\x1b\x52\x9b\x4d\x01\xa0\xb8\x44\x58\x00\x3b\x9f\x21\xd1\x2a\x92\xbe\x67\x05\x40\x87\xd6\x85\xac\xa3\xfa\x00\x93\x52\xa3\xdb\x59\x32\xfe\x46\xbb\x2e\x24\x75\xc7\x3d\x36\xda\x9e\xc6\xde\x85\x26\xe9\x88\xa7\x2f\x6d\x6b\x07\x0b\x28\xaf\xac\x0b\xee\x7b\x16\x62\x93\xe4\xcd\x24\x77\x42\xa9\x8b\x45\xa3\x1d\xf9\x9b\xcb\xae\x78\x12\x05\xed\x50\xb9\x49\xa7\x01\x26\x05\x55\x43\x0a\x5d\x56\x24\x0f\xd5\x2a\xc3\xca\xb5\xc6\x68\xeb\xb1\xce\x72\x1e\xed\x7f\x43\x46\x04\x21\x64\x05\xf0\x7a\x7c\x81\x21\x83\xae\xf2\x3a\x1f\x02\xd8\x5b\x2d\x7f\x53\x22\x4f\xd2\xb6\x25\x51\xc3\x02\xf6\x5c\x38\x2c\x00\xb8\xa1\x23\xc6\xe7\xc5\xd9\xca\x7a\x32\x56\x39\x44\x92\x34\xdd\x70\x24\xa9\x78\xe0\xee\xe0\x4f\x26\x3e\xaf\x04\x26\xeb\x39\xfb\x07\xcc\x1d\xf8\x43\xfe\x3f\xce\x9f\x18\x6c\xe2\x18\xba\xb8\x93\xb7\xe7\xf5\x6a\xf9\x52\xfd\x74\x43\xd5\x55\x1d\xb7\xc3\x66\x50\xf1\xad\xc0\x90\xc6\xdb\x16\xbf\x03\x00\x00\xff\xff\xab\x02\x2d\x63\x45\x03\x00\x00")

func pluginsTemplatesPythonPluginTomlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonPluginToml,
		"plugins/templates/python/plugin.toml",
	)
}

func pluginsTemplatesPythonPluginToml() (*asset, error) {
	bytes, err := pluginsTemplatesPythonPluginTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/plugin.toml", size: 837, mode: os.FileMode(420), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonRequirementsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x28\xc8\x2c\x50\xc8\xcc\x2b\x2e\x49\xcc\xc9\x51\xd0\x2d\x52\x28\x4a\x2d\x2c\xcd\x2c\x4a\xcd\x4d\xcd\x2b\x29\xd6\x2b\xa9\x28\x01\x04\x00\x00\xff\xff\x5d\x24\x95\x7b\x21\x00\x00\x00")

func pluginsTemplatesPythonRequirementsTxtBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonRequirementsTxt,
		"plugins/templates/python/requirements.txt",
	)
}

func pluginsTemplatesPythonRequirementsTxt() (*asset, error) {
	bytes, err := pluginsTemplatesPythonRequirementsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/requirements.txt", size: 33, mode: os.FileMode(420), modTime: time.Unix(1498789947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonScanPy = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\x4f\x8b\xdb\x3e\x10\xbd\xeb\x53\xbc\x9f\x96\xc5\xc9\x6f\x63\xa7\x2d\x6c\x29\x06\x1d\x0a\xd9\x42\x0f\x0b\x3d\xf4\x6e\x14\x7b\x62\x0b\x62\x49\x1d\xc9\xd9\x1a\x93\x7e\xf6\x22\xc5\xed\x16\x4a\x7d\x11\xf3\xe6\xbd\x79\xf3\xc7\x77\xff\xed\xa7\xc0\xfb\xa3\xb1\x7b\xb2\x17\xf8\x39\x0e\xce\x8a\x3b\x94\xff\x97\x68\x5d\x67\x6c\x5f\x63\x8a\xa7\xf2\x43\x42\x84\x94\x52\x84\x56\xdb\xca\xcf\xe2\xc7\xfa\x89\x65\x41\x47\xa1\x65\xe3\xa3\x71\x16\xd7\x6b\x25\xea\xd6\xf9\x99\x4d\x3f\xc4\x1a\x9b\x76\x8b\x77\x6f\xde\xbe\xc7\x71\xc6\xb2\xa0\x65\xd2\xd1\xf1\x8d\x76\x36\x2d\xd9\x40\x35\x9e\x3f\x7f\xcd\xc5\x85\x19\xbd\xe3\x08\xcd\xbd\xd7\x1c\x48\x08\xd1\xd1\x09\xa3\x36\x76\xb3\xad\x05\x00\x64\x9c\xa1\x7e\x73\xaa\x8f\xdc\x4f\x23\xd9\xf8\x25\x67\x36\x9e\x5d\xaf\x8a\x65\x81\x3f\x4f\xbd\xb1\x8d\xd5\x23\xe1\x7a\x2d\xb6\x7f\xc8\x2b\xdd\x75\x8d\x5e\x75\x1b\x59\x5e\xe4\x0e\xb2\x2c\x2f\xc4\x47\x17\x48\xee\x30\xd0\xd9\x2b\x79\x30\xc1\x9f\xf5\x8c\x15\x87\x9b\xa2\x9f\x22\x46\x0a\x41\xf7\x89\xa6\xdb\x34\xb3\x92\x21\x3a\xa6\x26\xf2\x94\x40\xa6\x6f\x93\x61\xea\xd4\x27\x7d\x0e\xf4\x6f\xdb\x62\xd0\x61\x28\x76\x18\x29\xea\x8b\x66\x55\x3c\x1f\x1e\x8b\x1d\xe2\xec\x49\x85\xc8\x3b\x58\xcd\x7d\x50\xc5\x43\xb1\xf6\x53\x68\x8c\xdd\x23\x92\x0c\xd1\x21\x90\xe6\x76\xc0\xc9\x71\xb5\x0e\x97\xf8\x50\xbf\xcc\xf2\x93\xec\xc2\x66\x2b\x72\xde\x9c\x32\xa5\x4a\x15\x6e\xdb\xcc\xbd\xb1\xb1\x11\xf2\xe0\x10\xdc\x48\x71\x30\xb6\x87\x7e\xa1\x14\xe0\xc5\xc4\x21\x1b\xd6\x90\x78\x78\x55\x67\x31\x53\x9c\xd8\x0a\x21\xcc\x09\x4d\xde\x73\xd3\x40\x29\xc8\xa6\x49\x17\x6b\x1a\x79\x33\x89\x3c\xbf\xba\xdd\x6e\x99\x43\xfa\xde\x92\x8f\x78\xca\x4f\xfa\x77\x74\x00\xfd\xd5\xd7\x13\xb3\xe3\x1a\xf7\x41\xe2\x1e\x24\x7e\x06\x00\x00\xff\xff\xe2\x82\x12\x89\xb2\x02\x00\x00")

func pluginsTemplatesPythonScanPyBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonScanPy,
		"plugins/templates/python/scan.py",
	)
}

func pluginsTemplatesPythonScanPy() (*asset, error) {
	bytes, err := pluginsTemplatesPythonScanPyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/scan.py", size: 690, mode: os.FileMode(420), modTime: time.Unix(1498789947, 0)}
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
	"plugins/plugins.toml": pluginsPluginsToml,
	"plugins/templates/python/.dockerignore": pluginsTemplatesPythonDockerignore,
	"plugins/templates/python/.gitignore.template": pluginsTemplatesPythonGitignoreTemplate,
	"plugins/templates/python/CHANGELOG.md": pluginsTemplatesPythonChangelogMd,
	"plugins/templates/python/Dockerfile": pluginsTemplatesPythonDockerfile,
	"plugins/templates/python/LICENSE.md": pluginsTemplatesPythonLicenseMd,
	"plugins/templates/python/README-short.txt": pluginsTemplatesPythonReadmeShortTxt,
	"plugins/templates/python/README.md": pluginsTemplatesPythonReadmeMd,
	"plugins/templates/python/circle.yml": pluginsTemplatesPythonCircleYml,
	"plugins/templates/python/plugin.toml": pluginsTemplatesPythonPluginToml,
	"plugins/templates/python/requirements.txt": pluginsTemplatesPythonRequirementsTxt,
	"plugins/templates/python/scan.py": pluginsTemplatesPythonScanPy,
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
	"plugins": &bintree{nil, map[string]*bintree{
		"plugins.toml": &bintree{pluginsPluginsToml, map[string]*bintree{}},
		"templates": &bintree{nil, map[string]*bintree{
			"python": &bintree{nil, map[string]*bintree{
				".dockerignore": &bintree{pluginsTemplatesPythonDockerignore, map[string]*bintree{}},
				".gitignore.template": &bintree{pluginsTemplatesPythonGitignoreTemplate, map[string]*bintree{}},
				"CHANGELOG.md": &bintree{pluginsTemplatesPythonChangelogMd, map[string]*bintree{}},
				"Dockerfile": &bintree{pluginsTemplatesPythonDockerfile, map[string]*bintree{}},
				"LICENSE.md": &bintree{pluginsTemplatesPythonLicenseMd, map[string]*bintree{}},
				"README-short.txt": &bintree{pluginsTemplatesPythonReadmeShortTxt, map[string]*bintree{}},
				"README.md": &bintree{pluginsTemplatesPythonReadmeMd, map[string]*bintree{}},
				"circle.yml": &bintree{pluginsTemplatesPythonCircleYml, map[string]*bintree{}},
				"plugin.toml": &bintree{pluginsTemplatesPythonPluginToml, map[string]*bintree{}},
				"requirements.txt": &bintree{pluginsTemplatesPythonRequirementsTxt, map[string]*bintree{}},
				"scan.py": &bintree{pluginsTemplatesPythonScanPy, map[string]*bintree{}},
			}},
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

