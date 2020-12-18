// Code generated by go-bindata.
// sources:
// ../../api/swagger/v1/hiro.swagger.yaml
// DO NOT EDIT!

package hiro

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

var _ApiSwaggerV1HiroSwaggerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x94\xcd\x6e\xe3\x36\x10\xc7\xef\x7a\x8a\x01\x7d\x69\x81\x58\x4e\x73\xd4\xcd\xd8\xa4\x4d\x0a\xec\x26\x48\xd2\xf6\x1a\x9a\x1c\x49\xdc\x50\x1c\x81\x43\x45\x51\xb7\xfb\xee\xc5\x50\xb2\x0d\x37\x09\x8a\xbd\x58\xf2\x70\xe6\xa7\xf9\xf8\x0f\x57\xb0\x85\x87\x51\x37\x0d\x46\xb8\x28\xcf\xe1\x27\x5d\x3e\x97\xba\x84\xdb\x1e\xc3\xf6\xee\xe6\x67\xb0\x58\xbb\xe0\x92\xa3\x00\x54\x43\x6a\x11\x5a\x17\x09\xb6\x77\x37\x65\xb1\x2a\x56\xf0\xd8\x3a\x06\xc7\x30\x30\x5a\xa8\x29\x42\x83\x01\xa3\x4e\x2e\x34\xe2\x04\x96\xcc\xd0\x61\x48\x3a\x23\x74\xb0\x99\x91\xa6\x1e\x97\x98\xdd\x24\x96\x62\x05\xc6\x3b\x0c\x69\xc3\x18\x5f\x30\x96\xf0\x80\x08\xba\x77\x9b\xfb\xab\xed\xe5\xe7\xab\xb2\x9b\xe9\x1d\x45\x04\x17\x6a\x8a\x5d\x26\xce\x49\x3c\x50\x87\xc0\x69\xf2\x08\x81\x12\x72\x55\xac\x60\x3d\xa7\x56\x3b\x8f\x87\xfc\x76\x13\xdc\xe3\x25\x99\x33\x18\x5b\x67\x5a\xd0\xde\xd3\xc8\xf0\x9b\x4b\xd7\xc3\x0e\x7e\xf5\xfa\x85\x22\x5a\xf8\xac\xe3\xb3\xa5\x31\x80\x0b\xc5\x0a\x00\x2c\xb2\x89\xae\x97\xef\x71\xb9\xa0\x31\x66\x6c\x20\xe8\xf4\xab\xeb\x86\x0e\xbc\x0b\x08\x1e\x43\x93\xda\xb3\x9c\x2b\x6a\x46\x69\x1a\x5a\x97\xdb\x21\xc5\xf7\x11\x53\x9a\xc0\xba\xba\x5e\x50\xd4\xe7\x76\x51\xb8\xb1\x0c\x3a\x57\x97\x5b\x34\x97\x08\xea\x0b\x0d\xe1\x4f\x8c\x3b\x75\x06\xa3\x4b\x2d\x68\x60\x17\x9a\xc1\xeb\x08\x81\x86\x50\x16\x05\xcf\xf3\xab\x40\x5d\x94\xe7\xaa\x60\xd3\x62\x27\x3d\x00\x58\x83\x6a\x53\xea\xd5\xf1\x95\x55\xd1\x47\xb2\x83\x39\x38\xe8\xbe\xf7\xce\xe4\x0c\x36\x5f\x99\x82\x7a\x6b\x7e\xed\xbc\x2a\x0c\x05\x1e\xba\x1f\x0a\x5b\x8f\xe3\xb8\x96\x3a\xd6\x43\xf4\x18\x0c\x59\xb4\xaa\x28\x64\x7c\x42\x49\x2e\x79\xac\x40\x5d\x2f\x82\x12\xc4\x0b\x46\x76\x14\x2a\x50\xbf\x94\xe7\x52\xcd\x49\xf7\x2b\xf8\xa7\x00\x80\x79\xb2\xa2\xae\x3e\xd2\x8b\xb3\xc8\xb9\x63\x1d\xa6\x96\x2c\xcf\x3a\xd1\x41\x37\xd2\xf4\xac\x56\xda\x7d\x45\x93\xb8\x3c\xa8\x35\x50\x02\x2d\x9a\x88\xa2\xb1\xcc\x94\xe9\x70\x4b\x83\xb7\x40\xc1\x4f\xb0\xc3\xa3\xa6\x77\xda\x3c\x03\xd5\xb5\x33\x38\x93\x51\x24\x5d\x42\x91\x23\x57\xb0\x1d\x52\x8b\x21\x2d\x85\xcf\xd6\x5c\xd5\xc0\xc8\x70\x2b\xc7\x79\xbb\x84\xa5\x4f\x7c\xf3\x67\xc5\x44\xd1\xfd\xbd\x28\x7a\x81\x5e\xc5\x48\x91\x8b\xa5\x60\xcc\xf5\x66\x1e\x27\x1d\xac\x8e\x16\xae\x1f\x1f\xef\xe4\x5f\x1a\x18\xa4\xb9\x0c\x89\xc0\x05\x2b\x6c\xcc\x2d\xe1\xc1\x18\x64\x06\x8a\x50\x6b\xe7\x87\x88\x19\xb7\x6c\xb2\x10\x8d\xf6\xbe\xcc\xfc\x1d\xd9\x69\x7f\x12\x91\x7b\x0a\x8c\x30\x3a\xef\xa5\x15\xbf\x3f\xdc\x7e\x39\x2a\x53\xd6\xc6\x85\x26\xb3\x66\x9d\x56\xf9\xfd\xe9\xe9\x29\x3f\xbf\xe5\x5f\x00\xd5\x21\xb3\x6e\x50\x55\xa0\xe6\x19\xe4\xce\xd7\x34\x04\xab\xce\xf6\x4e\x16\x93\x76\x5e\x7c\xf2\x40\x2c\xe1\x3c\x20\x7c\x75\x9c\x16\xb7\xef\x07\x7e\x91\x74\xb3\x88\x30\xe8\x6e\xaf\x1f\x95\xcf\xdf\x91\xca\x32\x08\x29\xf5\x76\xbf\x69\x2c\x4b\x83\x66\x88\x2e\x4d\x97\x87\xfb\x2d\x43\xf3\xac\xe6\x62\xe4\x8e\x92\xbc\x65\x3a\x17\x33\xbf\xf6\x34\x56\xa0\x74\x6e\xea\x27\xb2\x38\x9b\x4f\xe6\xf7\x47\xf4\xd5\xb2\x6b\xd5\x66\x23\x47\xe5\x72\xa5\x79\x32\xda\x6f\x32\x6f\xb3\x0f\x59\x08\x89\x9e\xf1\x7f\x23\xe5\x36\xcc\x8e\x73\x0c\x1b\xea\xe7\x75\x9c\x73\xb0\x0e\x83\xc1\x2a\xa2\xb6\x15\xa8\x7b\xd4\xf6\x60\x64\xf5\x5f\xaf\x31\xba\x24\xd5\x7d\x8a\x28\x5a\xc9\x37\xab\x75\xf5\xf4\x4e\xc8\x71\xa3\x4f\xd9\x47\xfb\xbb\xbe\x1f\x7f\xe1\x9d\xc0\x48\xfe\x34\xf3\xac\x04\xb1\x9e\xba\x7c\x08\x7d\xeb\x2f\x96\x37\xc8\xd3\xd3\x3d\xed\x2f\x79\xee\x61\x69\xa1\xb1\x2a\x0a\x7b\x2a\x8e\x6f\xdf\xff\x0d\x00\x00\xff\xff\x08\x15\x16\x1c\x2f\x07\x00\x00")

func ApiSwaggerV1HiroSwaggerYamlBytes() ([]byte, error) {
	return bindataRead(
		_ApiSwaggerV1HiroSwaggerYaml,
		"../../api/swagger/v1/hiro.swagger.yaml",
	)
}

func ApiSwaggerV1HiroSwaggerYaml() (*asset, error) {
	bytes, err := ApiSwaggerV1HiroSwaggerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../api/swagger/v1/hiro.swagger.yaml", size: 1839, mode: os.FileMode(420), modTime: time.Unix(1608248946, 0)}
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
	"../../api/swagger/v1/hiro.swagger.yaml": ApiSwaggerV1HiroSwaggerYaml,
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"api": &bintree{nil, map[string]*bintree{
				"swagger": &bintree{nil, map[string]*bintree{
					"v1": &bintree{nil, map[string]*bintree{
						"hiro.swagger.yaml": &bintree{ApiSwaggerV1HiroSwaggerYaml, map[string]*bintree{}},
					}},
				}},
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

