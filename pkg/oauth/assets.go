// Code generated by go-bindata.
// sources:
// ../../api/swagger/v1/oauth.swagger.yaml
// DO NOT EDIT!

package oauth

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

var _ApiSwaggerV1OauthSwaggerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7d\xeb\x72\x1b\xb9\xb1\xf0\x7f\x3d\x45\x17\xfd\x55\xd9\xae\xa2\x28\xc9\xeb\x75\xb2\xca\x77\x52\x47\x2b\x5f\xa2\xc4\x5e\xab\x24\x7b\xb7\x36\x8e\x4b\x04\x67\x9a\x24\x2c\x10\x98\x05\x30\xa2\x98\x64\xab\xf2\x0a\x79\xc5\x3c\xc9\xa9\x6e\x00\x73\xe1\x45\xa4\x2e\xf6\x71\x72\xb4\x95\x8a\xa9\x19\xa0\xd1\x68\xf4\x1d\x0d\xcc\x03\x38\x80\xd3\xa9\x18\x8d\xd0\xc2\x93\xde\x2e\x3c\x12\xbd\xf3\x9e\xe8\xc1\xdb\x02\xf5\xc1\xf1\xd1\x63\xc8\x71\x28\xb5\xf4\xd2\x68\x30\x43\xf0\x63\x84\xb1\xb4\x06\x0e\x8e\x8f\x7a\x5b\x0f\xb6\x1e\xc0\xbb\xb1\x74\x20\x1d\x94\x0e\x73\x18\x1a\x0b\x23\xd4\x68\x85\x97\x7a\x44\x8d\x20\x37\x59\x39\x41\xed\x05\x83\x10\x3a\x67\x18\x7e\x56\x60\xec\x33\x98\xd1\x93\xad\x07\x90\x29\x89\xda\xef\x38\xb4\x17\x68\x7b\x70\x8a\x08\xa2\x90\x3b\x27\x2f\x0e\x9e\xbf\x79\xd1\x9b\x04\xe8\x13\x63\x11\xa4\x1e\x1a\x3b\x61\x88\x01\x89\x53\x33\x41\x70\x7e\xa6\x10\xb4\xf1\xe8\xf6\xb7\x1e\xc0\x76\x40\x6d\x28\x15\x56\xf8\x0d\x66\x70\x82\xcf\x4d\xd6\x85\xe9\x58\x66\x63\x10\x4a\x99\xa9\x83\x57\xd2\xff\xa1\x1c\xc0\x4b\x25\x2e\x8c\xc5\x1c\xde\x08\x7b\x9e\x9b\xa9\x06\xa9\xb7\x1e\x00\x40\x8e\x2e\xb3\xb2\xa0\xf1\x5c\x2f\x82\x46\xcb\x60\xb5\x81\x89\xb8\x94\x93\x72\x02\x4a\x6a\x04\x85\x7a\xe4\xc7\x5d\xc6\x15\x85\x43\x22\x1a\xe6\x92\xc9\x41\x93\x2f\x2c\x7a\x3f\x83\x5c\x0e\x87\x11\x94\x29\x98\x5c\x46\x1f\xe5\x0e\x04\xcf\x8e\x49\x14\xa6\x08\x9d\x1f\x4c\xa9\x7f\x44\x3b\xe8\x74\x61\x2a\xfd\x18\x04\x38\xa9\x47\xa5\x12\x16\xb4\x29\x75\x6f\x6b\xcb\x85\xf5\xdb\x87\xce\x93\xde\x6e\x67\xcb\x65\x63\x9c\x10\x0d\x00\xb6\xa1\x33\xf6\xbe\xe8\xd4\x3f\x5d\x67\xab\xb0\x26\x2f\xb3\xaa\x81\x28\x0a\x25\x33\xc6\x60\xe7\x93\x33\xba\xb3\xf8\xf8\x72\xa2\x3a\x5b\x99\xd1\xae\x9c\x5c\xab\xdb\xf6\x74\x3a\xdd\xa6\x79\x6c\x97\x56\xa1\xce\x4c\x8e\x79\x67\x6b\x8b\x96\x8f\xa0\x78\xe9\x15\xee\x43\xe7\xed\x41\xe9\xc7\xcc\x7d\x07\xc7\x47\x04\xe7\x02\xad\x93\x46\xef\x43\x67\xaf\xb7\x4b\x53\x6a\x2d\xc1\x3e\xfc\x7d\x0b\x00\xc2\xf2\x12\x8b\x15\xd6\x5c\xc8\x1c\x1d\x93\x4d\x63\x86\xce\x09\x3b\x03\xe7\x6d\x99\xf9\xd2\xa2\x63\xca\x8f\x85\xce\x15\x5a\xc7\x4b\x53\x0f\x59\x3a\xb4\x20\x4a\x3f\x46\xed\x23\xe2\xbd\x2d\x86\xff\x00\x0e\x5a\x8f\xc3\xd3\x9f\x68\x0d\x68\x20\xbc\xcc\xb0\x68\xca\x45\x1b\x48\xbd\xae\x8c\xd7\xc4\xa1\xba\x40\x07\x8f\xfc\xd8\x38\x04\x4f\x2b\x96\x43\x9f\x46\xe8\x3f\xee\x12\x23\x12\xb3\xf3\x08\x99\x50\xca\x81\xc5\x5f\x4a\x69\x11\x04\x5c\x08\x25\x73\xe8\x7f\x8f\xc2\xa2\xed\x83\x37\xe7\xa8\x13\x93\xfc\xe1\xdd\xbb\xe3\x00\xc5\x58\xf9\x57\x1e\xae\x0f\x63\x14\x39\x49\xd0\xbb\x31\xd2\x50\xd4\x3e\x70\x96\x93\x23\x8d\x39\x0f\x42\x14\x89\x92\x8a\x39\x94\xc4\x52\x0c\xb0\x6f\x68\x1a\xbd\x43\xa3\xbd\x35\x4a\xd1\x80\x52\x7b\xb4\x43\x91\x61\x45\x97\xb7\x83\x4f\x98\x79\x38\xca\x69\xba\x43\x89\x36\xbc\x38\xa2\x86\x5a\x28\x35\xe3\xf9\x18\x6e\x15\x79\x3a\x47\x3d\x94\xd5\x48\x02\xca\x52\xe6\x3d\x78\x71\x99\x7a\x74\x69\x74\x47\x0d\x43\x87\x81\x70\xf8\xed\x6f\x21\x32\x4d\x8f\xe1\x1f\x32\x61\x3c\x2b\x9f\x26\x79\xdd\xd8\x94\x2a\x07\xa1\xa6\x62\xc6\x82\xce\x33\x89\x10\x2e\x84\x2a\xd1\x55\xa8\xbf\xb0\xd6\x58\xb7\x15\x19\x08\x19\x54\xe9\xd0\x81\xf3\x42\xe7\xc2\xe6\x81\xa4\xce\x0b\x5f\x3a\xa0\xc1\x79\x44\xa9\x73\x5a\xd7\x00\xd9\x95\x19\xf1\x18\x18\x0b\x43\x21\x55\x69\x91\xc1\x45\x36\x20\x88\xb4\x84\x4c\x7f\x18\x98\x7c\x96\xde\x58\x74\x85\xd1\x0e\x61\x2a\x95\x82\x01\xc2\x1f\x4f\xdf\xfe\x50\x8b\x3b\xe9\x22\xa9\x47\x0c\x2b\x08\xff\x3e\xff\xee\xf7\xfb\xfc\xef\xdf\xf8\xff\x01\x3a\x13\x62\xf0\x11\x76\xf6\xa1\x13\x68\x4c\x4a\x0f\x86\xa6\xd4\x79\xa7\x9b\x1a\xe5\xe8\x85\x54\xd4\x86\x39\x3c\x37\xe8\xb8\x19\x5e\x4a\xe7\x63\xb3\x5f\x2b\xf8\x5b\x5e\x8c\xa2\x64\x6b\x31\x21\xa1\x7c\xef\xd0\x76\xf8\xfd\x12\xd1\x03\xa0\xd7\xcd\x35\x88\x32\x08\xa4\xbd\x65\x16\x88\x46\xd3\xe2\xc1\x95\x3c\x47\x18\xa1\x67\x2d\x58\x58\xc3\x3a\x39\x17\x5e\x30\x1b\xba\xf4\x42\x38\x37\x35\x36\x77\xbd\x26\x1e\xc4\xdb\x57\xe0\xc1\x72\xbc\x0c\x0f\x39\xd2\x65\xd1\x05\x65\x46\x32\x98\x1d\xd1\x14\x92\x0a\xcf\xf6\x58\x41\xfd\xac\x1a\xea\xf8\x88\xcc\x62\x1c\x69\x6b\xcb\x61\x56\x5a\xe9\x67\xcf\x2b\x0b\xc9\x14\x64\xd5\x12\x56\x8e\xac\x1c\x2d\x12\x0d\xfd\x24\x40\x1e\x2a\x33\xdd\x87\x8e\x60\x0e\x3a\x34\x39\x86\xc7\x2d\xe4\xde\x5b\xb5\x1f\xb5\xf5\xfe\xce\x0e\xcb\x63\x34\x8a\xca\x64\x42\xed\x30\xbc\x9d\xd4\x25\x42\x60\x31\x5f\xd7\x93\xec\x29\x37\x0c\x7d\x5c\x66\x8a\xa0\xd0\x99\x7f\x0b\xd4\x32\x27\x32\x04\xf6\x16\xbc\x78\x0f\x1d\x0b\xaf\x97\x7e\xd6\x89\x0d\xe3\x0a\x2e\x69\x19\xdf\xa4\x86\x66\x38\x24\xa3\x78\x16\x66\x1b\x98\x8a\x8d\x70\x30\xf5\xcc\x24\xa4\xe7\xd0\x79\x28\xd0\x4e\xa4\x23\xb5\x4f\x4f\x49\x88\x2d\x0e\x2d\xba\x71\x54\x60\x09\xa6\xc8\x73\x1b\x80\x9d\x04\x0d\x99\x07\x28\x22\xaf\xf8\x8d\x10\x19\xcf\x9c\xcc\x84\x4a\xcd\x53\x6f\x9c\x08\xa9\xd6\xf4\xe5\x36\xf3\x1d\x8b\xb1\xd1\xb8\x76\x50\xa3\x11\x74\x39\x19\x24\xd1\x49\x03\x5e\xa0\x95\xc3\xd9\x5c\xf7\xf0\xb0\x26\xde\x15\xe3\x6e\xd4\x7f\xd9\xf0\x4c\xbb\x7d\x42\x94\x3b\x13\xbe\xd1\x7c\x54\x1e\xd4\x7c\xd3\x0b\x73\x1e\x26\x4a\x3f\xe6\x88\xef\x90\x57\x68\x3f\xb3\x28\x3c\xb5\x62\x1b\xcc\x6c\x14\x94\x5a\x78\x01\x02\x06\xd6\x4c\x49\xf4\x63\x0f\x98\x8e\x51\xb3\x0f\x56\x4d\x2c\x4a\x7b\x1b\x48\x74\xc8\xc2\x9c\x68\x92\xd9\x58\xe8\x11\xab\x5d\x69\xab\x3e\x9d\xad\xad\xbc\x2d\x74\xc7\x91\x25\xaf\x90\xde\x4a\xf1\x44\x9d\x49\xf6\x21\x07\x13\xb4\xef\x87\xc0\xfc\x90\x19\xad\xe9\x65\xb2\x06\x1f\x1f\x25\x61\x0a\x0d\x7a\x1a\xfd\x8e\x2b\x30\x73\xf1\xc1\x76\xec\xb1\x9d\x19\x8b\xdb\x7b\x67\xbb\xbd\xb1\x9f\xa8\x07\xa7\xb1\xff\xa1\x12\x72\xe2\x1e\xf7\x5a\xea\x80\xc7\x0f\x74\xb8\xdc\xd6\xa5\x52\x62\x40\xd2\xe4\x6d\x19\x0c\x49\x61\x49\xa1\x79\x59\x8b\xa6\x2b\x07\xe9\x67\x05\xc6\x79\x2b\xf5\xa8\x53\x3d\x5e\x3a\x6b\xfa\xef\xb4\x0c\x13\xde\x6e\x18\x6c\x76\x82\x68\xe2\x2f\x74\xbe\xcd\xba\x5c\x78\xfe\xfb\xc8\xb9\x12\x6d\xaf\xea\x7e\xb9\x3d\x32\xdb\x51\x3d\x46\x40\x69\x48\x7e\x7a\x63\xac\xd2\xc0\x0f\x1d\x0c\x4b\xa5\x18\x1a\xd9\xc2\x5c\xba\x42\x89\x19\x91\x84\xad\x20\x48\x9d\xa9\x32\x67\xaf\x21\x35\x2b\x84\xf5\xae\x0b\x85\x71\x4e\x0e\xd4\xac\xd1\x84\x5d\xca\xe0\xef\xb9\x72\x38\x94\x97\xe8\xba\xd0\x18\xd5\xd8\x1c\x49\x7c\x44\x96\x19\x1b\x7a\x98\x16\x1d\x1e\x3a\x60\x55\x89\xc9\x5b\x1f\xa2\x45\x1d\x0d\x05\xfd\x37\x92\x17\xa8\xcf\x6e\x37\xf7\x57\x04\x83\xa7\xf2\xc8\x3d\x66\x37\x42\x5a\xe7\xeb\x07\xc3\x16\x4a\x3d\xf8\xc1\xb0\xe7\x21\x3c\x11\xc8\x51\xac\x93\x95\x8a\x3d\xdb\x2e\x14\x68\x0a\x85\x90\x09\x0d\x63\x71\x81\x30\x29\x95\x97\xf4\x64\x54\x0d\xe2\x7e\xd7\x24\x01\x51\x91\x5a\x0f\x90\xa6\xe7\x50\xfb\x18\x58\xb0\xf7\x4c\xcd\x61\x80\x44\x19\x87\x85\x08\xee\xe1\x60\x06\xae\x10\x19\x92\x34\x5a\x91\x79\xb4\x15\x39\x86\x62\x22\xd5\xec\x96\xf4\x38\x2d\x6d\x83\x18\x4a\xdc\x39\x2d\x02\x96\x71\x76\x86\x02\xa7\xe6\xa3\xcf\x48\x9e\x89\xcc\x73\x85\xb7\x24\xcf\x1b\x06\x72\xc7\x24\x99\xd4\x40\x3f\x0b\x7f\xc0\x81\x72\x86\xa3\xf1\x55\xa8\x35\x31\x60\x9f\x9f\xdc\x53\x32\x11\x89\x78\x5a\x66\xe7\xb7\xa3\xdc\xa1\x70\xa5\x88\x4a\x63\x8e\x6c\x01\xab\x89\x98\x11\x3f\xd0\x3f\x34\xfa\x20\x3a\xf8\xd4\x5e\x84\x78\xb2\x96\xf7\x1e\xbc\x34\x16\xa4\x26\xdb\x90\x61\x4b\xad\x88\x0a\xd7\x10\x6f\xd0\x60\x6f\xc8\xe9\x9d\xc8\xd1\x98\xc1\x5a\xf4\xa5\xd5\xa4\x79\x94\xd1\x23\x47\x2e\xaa\x68\xc0\x6e\x76\xcb\xc6\x02\x55\xaf\xf2\xb5\x48\x03\x59\xcc\xcf\xc8\x22\xde\x52\xcc\xc6\xc6\x7a\x0a\x83\x03\x41\x06\xb3\x98\x05\x69\x91\x65\x2a\xdd\x38\xb8\xef\x8c\x75\x18\x9c\xfe\x8c\x06\xe2\xe4\xb8\x4b\x01\xd0\x98\xe8\xf3\x49\x68\xcc\x0d\x12\x05\x3f\xf5\x72\x83\xbd\x10\x93\x87\xb9\xbc\x39\xf8\x99\x40\x08\x3d\x8b\x11\x6c\x03\x13\x0e\x7c\x02\xde\x0d\xe5\x4d\x76\x55\x0a\xd5\xe0\xa2\x6a\xa4\xff\xee\xc2\x4e\x97\xc6\x99\x8e\xa5\x47\x66\xb6\x10\x5f\x9d\x1c\xc3\x9b\xf7\xa7\xef\xe0\x87\xb7\xef\xc0\xa2\x9a\x41\x59\xb0\x49\xaf\xd0\x08\xbc\x5a\x6a\xf9\x4b\x89\xbd\x39\x07\xf6\xc6\x94\x7c\x7f\xf2\x7a\x9e\x9f\x6a\xef\x17\x0a\x31\x8a\xd8\x65\x46\x7b\x76\x74\xb9\xb1\x74\xf0\x13\x0e\xf8\x35\x9c\xfe\xe1\xed\xfb\xd7\xcf\x99\x40\x03\x53\xfa\xb6\x44\x57\x23\xc5\x30\x10\x3a\xa5\x95\x95\xe3\x24\x39\xa9\xf1\xd9\x90\x0f\xe0\xe3\x52\x52\x5b\xa6\x2f\x73\x02\xb3\x81\x06\x39\xa1\x19\x70\xeb\x47\x9c\xe7\xba\x14\x93\x42\x61\x17\x04\x1c\xff\xf0\xaa\x0b\x7f\x3c\x7e\xf1\x8a\x17\xeb\xd5\xd1\xcb\x46\xe3\xc7\x2d\x91\xb1\xc2\x8f\x83\x14\xb2\xc3\x2f\x6a\xd2\x10\xd1\x84\xd4\x21\x63\x16\xfa\x37\x75\x9c\x4f\x78\x45\x12\x32\xd7\x0c\xc9\xd9\x57\x33\xa8\xec\x35\x88\x7a\x46\x63\xe3\x4d\x9a\xf1\x12\x0f\x04\x5c\x29\x7d\xf2\x37\x92\xff\x41\xa3\xb3\xcf\x1a\xc8\x38\x48\xf9\x91\xd4\xa9\xdb\x9a\x00\xfd\xcf\x0e\xa4\xb7\xc2\xce\xe2\x78\x5e\x90\x9f\x1d\x92\x9b\x0b\x2b\x3b\xc5\x81\x93\xfe\xce\xd7\xb0\x22\xa1\xb1\x30\x50\x66\x14\x17\x71\x9e\xe9\x22\x81\x9b\x41\x00\x14\xe5\x40\x91\xe4\xe7\xf3\x28\x13\x2c\xa1\xdb\x3e\xd4\x48\xe8\x14\x47\xc7\x25\x69\xb4\x97\x0e\xc4\x70\x28\x95\x64\xd3\x40\x96\xa3\x72\x9d\x50\xe7\x68\xf7\xaf\xe5\x1b\x86\x3e\x3d\xf8\x91\xd3\x38\x21\x1b\x9d\x90\xa4\xa8\x21\x2d\x7e\xc8\x2e\x5b\x84\x21\x4e\x92\xf7\x46\x3f\x7a\xf0\x96\x97\x29\xa4\x81\x92\x52\xe2\x64\x30\x2f\xaf\x46\xc9\xef\x9b\xf3\x0b\x74\x4d\x43\xc5\x9e\x04\x3b\x26\x37\x07\x0a\x7b\x57\xaf\xdc\x40\x5a\x3f\xce\x45\x73\x85\x37\x99\x6c\xec\x36\xeb\x82\xc5\x68\x7b\xc9\x62\x90\x3f\x0b\x47\xa7\x6f\xe1\xb7\xcf\x76\xf7\xf6\x9f\xec\xee\x3e\x85\x0f\x47\xa7\x6f\xe9\xaf\x7f\xfd\xe3\x9f\xf4\xf7\x47\xf8\xf9\xe7\x9f\x7f\xde\x7e\xf3\x66\xfb\xf9\xf3\xa8\x34\x82\x02\x9a\xa1\xb0\x69\xd2\xbb\xbb\xbb\xbb\xdd\x94\xc9\x0a\xfc\x4c\x76\xd9\xd3\x8a\x99\x89\xf4\x1e\xf3\x5e\x93\x0e\xef\x4c\x8d\x06\x18\xad\x02\x5f\x10\xc4\x2e\x0f\x97\x32\xd4\xb4\xe0\x14\xb6\x51\xf7\x5a\x4c\x73\x2c\x50\xb3\x5a\x8f\x21\x56\x49\x0b\xa9\x58\xb0\x0a\x25\x3c\x75\x7e\xe8\x80\x88\x44\x8a\x9b\x99\x65\x58\xea\x8c\x08\xd4\x8d\x69\x1c\x6a\xfb\xa9\x74\xbe\x89\x15\xcf\x88\xfc\x13\x8b\xae\x54\xec\x57\x5c\x08\xcb\x60\x27\x46\xfb\x31\xaf\x3c\xd3\xd0\x05\xcf\x5e\x92\x6e\x9a\x10\x29\xad\x03\x8d\xc1\x98\x91\x74\x06\x0e\x1a\x8a\xcc\xb3\x59\x27\x25\x94\x65\xa6\xd4\x9e\xe3\x4e\x63\x2d\x66\x5e\xcd\x08\x15\x4e\x72\x30\x4f\x08\x8f\xae\xd7\x59\x27\xb6\x95\xda\xa6\xf6\x9d\x46\x2c\xb5\x24\xd8\x03\xf8\xab\xd1\x98\x52\xe2\x6b\x78\xe5\x34\x98\xcc\xa1\x35\x93\xaa\x1b\x7c\x48\xbf\x3e\x82\x97\x13\xe4\x17\x9c\x5d\xa3\x08\xb7\x5e\xc2\x79\x0d\xf6\xd0\xd5\xcd\x83\x6b\x53\xe9\xf1\x17\x25\x45\xa0\x3b\xc7\xc2\x4a\x76\x97\x0f\x26\x68\x65\x26\x76\x5e\x1b\x77\x76\xa0\x47\xa8\xea\x70\x68\x05\x0d\x42\x0c\x75\x3d\xee\x0f\x7d\x16\x79\x1f\xbe\x3f\x3c\x7e\xfa\x1b\xf8\x70\xf2\xf2\xf0\xdb\x67\x4f\x9f\x7d\x04\x25\xf4\xa8\x24\x7d\xe6\x45\x52\x71\xd2\x11\x1e\xd1\x08\x44\x69\x79\xf6\xcd\x77\xdb\x7b\x70\xa0\x8a\xb1\xd8\x7e\xc2\x12\xf3\xec\x9b\xef\xfe\xf5\x8f\x7f\xee\x35\x20\x64\x26\xe7\xa8\x93\xd8\xd7\x66\x44\xaf\xa6\x57\xa7\xf3\x04\xeb\x9b\xbd\x67\xcf\xe6\x80\xd1\xa3\x00\x8d\x99\xc6\xce\x2a\x60\x65\x51\x04\x60\xdd\xb6\x8f\x2c\x20\x17\x6e\x3c\x47\x6b\xd4\xdb\xef\x4f\x39\x06\xb4\xdb\x87\x07\x3d\x38\xa0\x19\x67\x66\x52\x08\x2f\x07\x52\x49\xcf\x9e\x69\xdb\xdd\x64\x47\xba\xe2\xeb\x98\xf9\x64\xf7\x9e\x15\x9b\xd0\x41\xd8\x5c\x66\x6c\xe5\xc7\x46\x4c\x8c\x6d\x1b\x2e\x46\x29\x6e\x57\xd5\x28\x9d\xbd\x3f\xfd\x1d\x9c\x60\x90\xd6\x63\xc1\xb9\x08\x56\x23\xd9\xd8\xf0\xf6\x05\xcb\x0a\x16\x2d\xc9\x64\x71\x8a\xc1\xb3\x9b\x69\x2f\x2e\x69\xec\x29\x2a\xb5\x86\x59\x88\x3d\x86\xa2\x54\x24\x31\x4c\x8e\x76\xf2\x6c\x39\x13\x75\xde\xad\xc9\xd8\xad\x97\x4e\xee\xd6\x1a\xeb\x8c\x73\x6a\x12\xf3\x0d\x38\xf7\x9d\x2d\x11\xe4\x82\x0d\xc6\xed\x26\x32\x30\x16\x14\x2e\xa1\x86\x04\xf9\x77\x60\x88\xfc\x53\xe9\x28\x2c\x55\x0e\x7b\xf0\x13\x19\x22\xa6\x1e\xa7\x8b\x82\xad\x63\x96\xb6\x25\x76\xc3\x9b\x09\x0a\xde\x48\x8a\xd6\xf6\xed\x71\x8b\xf2\xc6\x9c\xb3\xd1\x65\x6b\x7e\x81\xe0\x3c\x16\xec\xc3\xa3\x76\xa5\x6d\x3a\x4e\x73\xe8\x4d\x85\x63\x77\x80\xf7\x7a\x16\x6d\x7f\x1c\x8d\xf5\x04\xfd\x08\x93\x88\xd6\x96\xfa\x16\x68\x89\x9c\xf3\x96\x63\x8c\x11\xe1\x2a\xc6\x10\x7a\x7e\x68\xf2\xd0\x23\x4d\xe8\x37\xbb\xca\x97\x7e\x3b\x59\xf4\x6e\x50\xe4\x6c\x45\xc8\x00\x45\xbf\x1e\x89\x28\xce\xc3\xd0\x8a\x09\x4e\x8d\x3d\x27\xd1\xe1\x19\x88\xcc\x53\xac\x27\x46\x16\x59\x2c\x5c\x13\x23\x72\x43\xa4\x6e\xc4\x3b\x45\x64\x6a\x32\xeb\x71\xfb\x40\x8f\xe6\xf9\x74\x60\x8c\x42\xa1\x9b\xa9\xd8\xb3\x90\x5f\xdd\x80\x29\x9b\xe9\x58\xd2\x0a\x2f\x7a\x7b\xcf\x9e\x46\xe6\x5b\xc3\xa3\xcd\x91\xee\x80\x25\x5b\x98\x7c\x6d\x0c\xd9\x42\xee\xcb\xb1\xe3\xdc\xf2\x7c\x85\xbc\x18\x57\x21\x52\x1c\x5b\x4c\x11\xd7\x85\x23\xb3\x01\xce\x73\x57\x34\x5c\x33\xc0\x4b\x8f\xda\xb1\x75\x48\x4d\x9b\x11\x58\xc3\xce\x4a\x0d\x27\x2f\x0f\xe1\x9b\xef\x9e\x3d\x4b\xfe\xe3\x3c\x8f\xb6\x85\x21\x6d\xc3\x54\x8d\xfe\x9f\xc5\xe1\x3e\x74\x1e\xec\x34\x32\xf3\x3b\x07\x6d\x95\x5c\x16\xe4\x10\xe5\x67\x69\x67\xb3\x21\x00\x52\x7b\x1c\xd5\xfb\x16\x57\x31\x79\x5a\xfd\x06\x87\x37\x23\x19\xe2\x03\x4e\x21\xc6\xd1\x7a\x70\xe4\x53\x32\x80\x5c\xd5\x90\x7c\x88\x74\x5c\xf0\x8e\xe2\x73\x33\x04\x87\x99\xd1\xb9\x0b\xce\xd6\xde\x77\xbf\xd9\xdd\xde\xdd\xdb\xde\xdd\x7b\xb7\xbb\xbf\xbb\xbf\xfb\xe7\xa6\x9b\xc0\xc2\x40\x6c\xcd\x74\x7c\xff\xee\x10\x4a\xed\xa5\xaa\x3c\xc6\x1d\x62\xd8\x25\x41\xbd\xd4\xfe\xd9\xd3\xce\xd6\x16\xc0\x41\x93\x9a\x6d\x8d\xf2\xb6\x40\x7d\xf4\xbc\xd2\x9a\x19\x2f\xbc\xa8\x23\x21\xa9\x09\x53\x9e\xf9\xb7\xbd\xbd\xde\x5e\x0a\x5f\xd2\x4e\x06\xfb\x00\x7b\xbd\xdd\x76\xa8\xd4\x59\xb5\x21\xb1\xb8\xf7\x10\xd0\xf5\x4d\x05\x74\xdd\x68\xf5\x65\xa9\x14\x90\xf2\xe7\x98\x3e\xcc\xa4\x5b\x03\x6e\xc6\xdd\x24\x42\xa5\x43\x0a\x1a\x44\xd5\x45\x89\x01\xaa\x5e\xaa\xe7\x41\x95\x07\x5f\x24\x86\xb1\x55\x4a\x53\x49\x8d\x6e\xce\xe7\xd2\x38\xe5\xc7\x2d\x85\xf0\x43\x7c\x98\xf2\x9c\x4d\x49\x88\xf1\x20\x3b\x9e\x99\xb0\x56\x92\x9b\x18\xb2\x77\x3b\x5c\xe1\x33\xa4\x10\xa2\x10\xd2\xc2\xa3\xce\x5f\xec\x5f\x74\x87\x93\xd6\xdc\xde\x49\x3d\x8a\x78\x84\x66\x55\x26\x8b\xda\xea\xce\xe3\xde\xba\x58\xc0\x79\x8b\xe8\xcf\x16\xa4\xeb\x46\x14\x0f\xc0\x6a\xd6\x31\x93\xc2\xe8\x90\xcf\x65\xdd\x43\x44\x0c\xd9\x37\x84\xb1\x21\xaa\x07\xee\xef\xa6\x9e\x5a\x4c\xb0\x0b\xc7\xc6\x79\x78\x3b\x1c\xca\x0c\xe1\x7b\x73\x19\x94\x21\x13\x7d\x9b\x67\xca\x5a\x26\xc7\x7c\x7e\xbc\x06\x3a\xcd\x82\xad\x5b\x2e\xe3\xad\xd6\xae\x81\xd2\x67\x5d\x46\xf6\x82\xa5\x9f\x5d\x6f\x01\x3b\x87\xe4\xed\x1b\x5b\x75\xaf\x97\xac\xb7\x36\x8a\xb4\x38\x22\x20\xd7\x1b\xf0\xd4\x0b\x0a\x2d\x38\xda\xe6\x9c\x36\xa7\x9b\x39\x01\xc8\x09\xbc\x00\xf4\x3a\x68\x14\xc6\x79\xa1\xce\x28\x14\xba\x26\x2e\x7f\x96\x45\x88\xa0\x8c\x8d\x50\xc2\x9f\xd7\x18\x3c\x06\x62\xd7\xa5\x7a\x0c\xdf\x38\x29\xbe\xd1\x68\x5b\x00\xa1\x1a\xeb\x1d\xef\x99\x2f\x6a\xed\x24\x8a\x8d\x56\xc1\xbc\x57\xdb\x00\xd1\xb5\xe9\xc7\x62\x0e\xde\x66\xef\xc3\x04\xfd\xd8\xe4\xad\xca\xad\x54\xd6\x14\x05\x35\x99\x5f\x0d\xfd\x50\x55\x71\x16\xbb\x06\x91\x8e\x42\xc1\x61\x60\x0c\xd2\x9c\x03\x51\xc8\x08\x3a\x9a\x33\x51\xa5\x5c\x62\x1d\x4c\xe2\xe8\x30\xb0\xb0\xc1\xc2\x72\xd2\x27\x95\x71\x4d\x84\x16\xa3\x1a\x73\x82\x29\xbd\x43\x35\x0c\x99\x26\x36\xe0\xa1\x36\x86\xcc\x2c\x5e\x16\x98\xf9\x58\x6c\x06\x5c\x42\x95\xe2\xd5\x0d\xea\xcc\x60\xb9\x71\x8a\x35\x71\xf9\x3e\x7c\xe8\xf0\xb4\xcf\xa8\x51\xa7\x9b\xea\x69\x02\x2d\xe8\x6f\xbc\x2c\xa4\x45\x77\x26\x75\xe7\xe3\x0a\xbb\x26\xf3\xd0\xfc\xe6\x4a\x96\xa7\x9d\xa3\x26\x49\x0d\x6b\x15\x75\x59\x34\xd3\xae\x91\xdb\xa7\xa0\xa0\xae\xa4\x89\x4a\x30\xd6\x62\xb8\x96\x53\x16\x19\x24\xba\xf1\xfd\x50\x66\xd0\x4f\x45\x16\xc2\xc1\xc8\x0a\xd2\x76\xbd\x46\xaf\xa3\xd8\x38\xa6\xbc\x97\xb4\x8e\x5e\x7b\x28\xd2\x88\x1a\xb7\xaa\xcf\x8a\xdd\x7a\xed\xb9\x39\x8c\x05\x42\xc1\x31\x2d\x3d\xef\x5a\xd5\xee\x05\xb9\xc8\xd5\xc6\x86\x92\x03\x2b\xec\xac\x1b\xf3\x89\x65\x11\x37\xd6\x1b\x20\x63\x69\x16\xef\x22\x4c\x84\xac\x70\xa8\xc6\x49\xe3\x37\x17\xf3\xaa\x40\x2b\x90\x2f\xec\x54\x55\xe5\xc7\xad\x4a\xaa\xb5\xb9\x80\xa6\x68\x73\xf0\x53\x69\x53\xae\x3d\xba\x0b\x0e\x69\x95\x31\xc1\x44\xcc\x1a\xf2\x99\x2a\x30\x41\x90\x85\x4b\xe2\x1a\x5a\x3a\x13\xcb\xa4\xe6\x32\x52\xbc\x62\x99\x45\x66\x25\xa1\x1c\xe4\xbc\xcd\x1a\xf2\x3f\x44\x74\x2b\x2e\xd0\xc6\x22\xc8\x29\x97\x90\x8e\x84\xd4\xbd\x79\xc6\xad\x6a\xe5\x1a\xbc\xd6\xaa\xd7\x4a\x5c\x24\x5d\x2a\xd2\x6a\xc3\x48\x45\x3b\x5f\x9c\x3f\x6a\xd1\xdf\x80\x3b\x66\x05\x76\x93\x0e\x0d\xea\x78\x2d\x53\xa0\x2e\x27\xa4\x62\x06\xa1\xf9\xc7\xb5\xdc\x52\x6b\x9b\x2b\x11\xa2\xc0\x85\xd5\x6f\x5f\x9b\x69\xbf\x8e\x9d\x03\xa6\x11\xc8\x02\x76\x0b\x91\xd1\x7c\xfc\x70\x35\x76\x22\xcf\x39\x18\x13\xea\x78\x41\x05\x2e\x51\xb1\xf3\x88\x1f\x54\xdd\x1b\x2a\x94\x80\xd6\x96\x20\x65\xf3\x1b\x15\x7a\x2b\x47\x5c\x18\x73\x0b\x42\x49\xee\x49\xac\x8d\x5d\x16\x02\x1d\x90\x55\x9e\x18\x0d\x48\x2d\xab\x32\xda\x95\x31\x4c\xd3\x4c\xa4\x5a\xd9\x55\x56\x20\xbe\xbf\xa6\xcb\x40\xcb\x19\x90\x49\xf0\xd7\xf2\x48\xa8\xc7\xbd\xf1\x38\xb1\x9c\x77\x6b\xab\x10\x7e\xcc\xd8\xef\x70\x69\x6b\x00\x38\xc2\x2a\xa2\x5e\xa1\x8b\xde\xbb\x58\x1f\x5f\x29\xcd\xaa\x54\x2f\x95\xe8\xa5\x92\xb0\x50\x11\xaf\x73\xb0\x98\x4b\x1b\x4a\xb8\x53\xb9\x94\x9a\x45\x88\x8d\x53\x0b\xe4\x47\x06\x10\x87\x0c\xb2\x2a\x28\x14\x23\x47\x4b\xc0\x15\xbc\x49\x88\xc8\xa7\x9f\xa0\x47\xdb\x60\x89\xaa\xfe\x36\x0d\x78\xd6\xd8\xee\x5e\x4e\x94\xd4\x12\xe6\x5a\x4a\x6a\xf0\x4b\x89\x76\xd6\x7c\xba\x82\xd6\x4d\x4e\x69\x78\x91\x0d\x84\x9c\x6f\x6d\xe0\x2c\x75\x9f\x89\x96\x85\x70\x55\xc5\xf3\x1d\xa0\x96\x38\xbc\x41\xa3\x6f\x76\x9f\xec\xaf\xc4\x23\x94\x7e\x37\x5e\x87\x13\x00\xae\xd9\x03\xe0\xb5\x09\xb1\x7f\xfb\xe9\x95\xd4\x7d\x7f\x72\x54\xb1\x05\x19\xd6\x64\x08\x3a\x73\x20\x56\x12\x18\x56\x14\x31\x00\x3c\xdd\xdd\x5d\x3d\xa1\x81\xc8\x6b\x4e\x69\xc2\xe3\xb3\x2d\xa2\x3d\x83\xa5\x99\xa7\x96\x52\x69\x8e\xba\xb7\x7a\xd4\x52\x57\x35\xd5\xf9\x5d\x0e\xfa\xed\x55\x53\x95\xf1\xd0\x04\x84\x42\xed\x20\xee\x77\x37\x7a\x2a\x52\x6f\x4a\x5b\xa8\x50\x87\x0f\x9d\x76\x39\x6f\xe7\x63\x38\x16\x41\xc1\xd7\x1a\x6d\xd2\x38\x18\x83\x75\xb5\xee\x46\x6e\xfd\xbc\xea\x78\x4d\x5a\xec\x4a\x95\x51\x1d\x3c\x82\x0f\x1b\x1c\x2f\xda\x48\xd1\xa8\xe6\xa0\x70\x15\x03\x2f\x4f\xef\xc3\x02\x00\x19\xe4\xd0\x4e\x9e\x0b\x2f\xae\xa3\x62\xea\x9a\xe6\x9b\x63\xb3\x0c\xc6\x8d\x11\x8a\x22\x7e\xd6\x38\x27\x70\x1d\xac\xfe\xde\x62\x54\xc6\xb1\x7d\xf2\x22\x15\xfc\x2f\xc0\xbf\x31\xc6\xb4\xf0\x69\x9b\xc4\x5e\x4d\xae\xe3\x3f\x1d\xbe\x08\x69\x85\x65\xed\x57\x62\x70\x3d\x03\xb2\x44\x7d\x3f\xd9\x7d\xba\x5a\x05\x68\x93\x0a\xc5\x3a\xf7\xfa\xfe\x5e\xdf\xdf\x70\xf4\xe8\x12\x9a\xd2\x6f\xea\x13\xbe\x36\x23\x0a\xdb\xe2\xe6\x96\xf3\xd6\x04\x87\x3e\x2b\xad\x45\xed\xe7\xcf\x70\xac\x52\xdf\xa6\xf4\x77\xef\xf2\x5d\x47\x05\x5a\x19\x4e\xe3\x44\x26\xf6\x06\xc4\xd0\x07\x25\xdd\xc0\x0d\x56\xf8\x60\xb5\x1a\xe1\x68\xfb\x4c\xde\x54\x15\x87\xfe\x20\x17\xd5\xf0\x82\xd7\xb7\x4e\xa3\x89\x32\x97\xa8\x33\xfc\x92\xae\xed\x86\x33\x0d\x0b\x0e\x0b\xdd\x17\xd1\xbb\xf7\x62\x3f\x97\x56\xfb\xdf\x53\x30\xe1\x58\xe5\xfe\xa6\x7e\xe2\x09\x8e\xa4\x23\x51\xdc\xd8\x45\xac\xba\xd2\x0a\xc6\xa4\x11\x4c\x4a\xe7\xc3\x3e\xb7\xf4\x52\x28\xd2\xd7\xcd\x23\xcb\x3f\x49\x3f\x3e\x50\xca\x4c\x4f\x19\xbb\x47\xc4\xf0\x8f\xfb\x60\x8a\x70\x98\x3b\xc0\x9b\x0b\x58\xb9\xe5\xbf\x9b\xdb\xc9\x45\x72\x5f\x91\xe3\xf9\xd0\xdd\xad\xeb\x29\xf5\x85\xf4\x78\x37\x9e\xe7\x91\x96\x17\x32\xde\x3a\x91\x8e\xc0\x13\x8f\xb0\x26\x08\xd9\x78\x6f\xf8\x98\x30\x94\x45\xa8\x05\xe6\x1a\xe8\xac\x7a\x26\x1d\xe4\xd2\x89\x81\x6a\xe7\xda\x57\xcc\xee\xde\x7f\xfe\x9c\xfe\xf3\xbd\xe1\xe0\xff\xee\xdd\xe1\x8d\x47\x27\x6b\x15\x8f\x4d\x6f\xe8\x0e\xff\x18\xce\x53\xa7\xbb\x62\xea\x13\xd6\x6c\xba\x5a\x35\xa6\xf5\x5e\x44\xdb\xac\x04\x10\xf3\x66\x85\xaf\x52\xb8\xe3\x04\xe8\xbc\xd3\x1b\x3c\xde\x66\x41\xdc\x8d\x7d\xc7\x4d\xb3\x9f\x2d\x51\xd9\xc0\x1f\x5c\x39\xe4\x7d\xbc\x7c\xaf\x20\xbe\xbc\x82\x88\xbd\xaf\xca\x8f\x86\xbd\xf7\x4e\x17\x3a\xe9\x62\x8b\x8f\x9b\x7a\xbe\xa7\xa8\xf3\x87\x0e\x44\xbb\x46\xb5\xda\x7c\x69\x86\xd9\xac\x5e\x12\x17\xc5\x2a\x3d\xcc\xc3\x2e\x7d\xe6\x63\xe1\xc6\x15\xea\x86\x86\xfa\xd2\x9e\x6c\x40\xea\xa6\x81\x72\x6b\x66\xb5\xa2\x6d\xf6\x4a\x7b\xb0\xa1\x46\x9f\x56\x60\x6c\x34\x36\xf6\x62\xaf\xeb\xef\xdc\x85\x92\xb9\x97\xc9\xaf\x53\x26\xc9\xd6\x57\xf7\x89\x5c\x69\xed\x3b\xaf\xd0\x83\xa8\xa2\x96\xd6\x9e\x67\x67\xa9\x8c\x1d\xc7\xa6\xaf\x70\x21\xcb\xb5\xb9\x5d\x5f\x88\xda\xae\x19\xdf\xdd\x5d\xda\x87\x0b\xa6\xae\x44\xa4\xa2\x0d\x35\x6d\x16\xa9\xdc\x02\xab\x24\xcd\x4a\xea\x73\x5a\x3a\xe3\x0b\xfa\xc7\xa2\x43\xdf\x92\xe9\xfb\xcd\x89\xfb\xe0\xea\xde\x77\xfa\xda\xf5\x74\x6a\x7a\xb5\x33\xd4\x39\x65\x5d\xbb\x22\x51\xb4\x5c\xcb\x9e\xae\xd1\xb2\x9f\xc1\x95\xb9\x65\x1e\x4c\xe3\xf4\xae\xf7\x5f\x1d\x7a\x2e\xa0\xde\x50\x4f\x73\x07\x98\xef\x70\x47\xea\xe2\x0b\x6d\xd5\xb4\x27\xb3\xd1\x3c\x96\x89\xfe\xbd\x72\xbb\x57\x6e\x77\x34\xfa\x95\x4e\x68\x25\xf0\xc1\xf7\xac\x66\xbf\x69\xaa\xe9\x15\x5a\x76\xa9\xf8\x66\x07\xde\x0f\x59\xe1\x87\x0c\x8d\x6d\x5c\x61\x59\x17\xe4\xf1\x05\xa7\x71\x7b\x36\x86\x93\xa1\x19\xf1\x4c\x6b\x17\x85\x8f\x6a\x72\x9c\xc5\xb5\xdf\xb5\xe8\xb5\x21\xf4\x45\x51\x90\x60\xf7\xe3\x25\xc3\x3a\x5d\xbd\x0a\xfd\x96\xcf\xd5\xaf\xd9\xab\xd7\x1e\x85\x64\x3b\xde\xc4\x9a\x2e\x8d\xa8\xaa\x09\xe3\xc5\xc5\x01\x45\x63\xe3\xed\x9c\xf3\xc8\x02\x7c\x6f\xfc\xb8\x81\x0a\xcd\xb9\xdf\x54\x3b\xfd\x6a\x2b\xc8\xc6\x9d\xa5\x78\xc5\x48\x75\x24\x21\x6c\xc2\xce\xed\x23\xf1\x3d\x72\x59\x40\xcd\x50\x7f\x91\x9d\xf3\xa4\x43\x0d\x65\x9f\x50\xea\xc7\xbd\xad\x74\x92\x82\xa9\x15\xcf\x93\x86\x0b\x5a\x25\x31\x5b\x53\x75\xd5\xa8\x89\x22\x11\xae\x1a\xd7\x2f\xf8\x96\x44\xce\x06\x62\xe1\x26\xda\xea\x2e\xd8\x1c\x95\xbc\xc0\x78\xd7\xd2\xdc\x6a\xf8\xa9\x01\x76\xb6\x97\x5a\x32\xe2\xcc\x3e\x73\xff\x7f\xfd\xff\x50\x2c\x4a\x0a\xfb\xf7\x7d\x78\x24\x7b\xd8\x83\xfe\x40\xe4\x67\x71\x0d\xfb\x5d\x40\x9f\x3d\x86\x25\x7d\xcf\x1a\xac\x9a\xe0\x84\xa2\xd3\xdf\xf7\xb7\x96\xda\xec\x83\xf6\x05\xa4\x37\xdb\xfd\x0f\x42\x77\x36\x1f\x96\x2c\x8f\x2f\x16\x15\xee\x22\x85\xeb\x5b\x76\xe7\x61\xae\x8d\x4e\xd8\x84\x5e\x27\x1a\x59\x5a\x31\xb0\x29\xea\x4b\xcb\x05\x6e\x6c\x9c\xbf\x40\xc5\x40\xe4\xc9\xf9\x11\x78\x80\x2b\x67\xfa\xfe\xe4\x35\x31\x75\x7d\xe0\x79\xee\xa6\xec\xa8\xc0\x97\x6b\x37\x3e\xd3\x91\xcd\x27\xb4\x37\x98\x44\xb3\xf8\x7a\x03\x2f\x66\xd3\x65\xbb\xd5\x64\x82\xb3\xd3\x66\xda\xea\x40\x7e\x3c\xa9\x93\xca\xea\x9b\x29\x89\xdb\x4d\x99\x4f\x4f\x5c\x73\xae\xed\x48\x39\xf8\x36\xac\x42\x30\x8f\xa7\x31\xba\x20\x87\x80\x93\xc2\xcf\x92\x65\x09\x26\x4b\x28\xd5\x38\x59\x54\xdd\xea\xeb\x7a\x8b\x73\x11\xd6\x8a\x16\x0a\xd2\xe3\x64\xce\xf9\x5a\x39\xed\xcc\x28\x15\x0e\x3e\xbf\x4c\x7e\x92\x73\x17\x9b\x6c\x63\x6c\xb2\xd4\x6f\x0b\xf1\x4b\x89\x61\x5b\xa3\x3e\x94\x75\x55\x7d\xf7\xba\xdd\x14\x4e\x04\x64\x63\xa1\x14\xea\x11\x9e\x2d\x66\x50\x6f\xb4\x2a\xd1\xb6\x2f\x1c\x28\xe2\x04\x73\x1a\x2c\x9c\xd0\xe3\xf4\xc2\xc9\xcb\xc3\x78\x74\xdc\xb1\x59\x89\xc7\x02\xbb\x70\xfa\xe4\xdb\x67\xe1\xce\x53\x25\xa4\x6e\xdd\x08\x09\x30\x36\x53\xbc\x40\xdb\x5d\xc1\xf0\xf2\x22\x43\x70\x65\x51\x18\xeb\x5d\x38\x35\x48\xe0\x96\xac\xf8\x6a\xe5\x4b\x1d\x5a\xca\xb7\xbe\x13\x86\x5f\xad\x23\xe8\x4d\x64\xb9\xbe\x00\xbf\x02\x13\x0e\xe8\x84\x13\x4c\x8d\xc4\x4d\xef\xdf\x25\xe1\x72\x65\x40\xb1\x48\x01\x15\xc1\xd4\x5c\x1d\x46\xbb\x0f\x37\xbe\xe8\x46\x75\xe3\x98\xe1\x46\x55\x55\xa4\x8e\xf8\x66\xac\xfa\x7c\x71\xf3\x1c\x24\x29\x02\x32\x26\xf5\x81\xc2\xf0\x65\x82\xe8\x23\x34\xcf\x03\x2f\xf5\xec\xde\x35\xb3\x9f\x5f\x6a\x4b\x89\x31\x5e\x70\x04\xaf\x91\x5c\x68\x9b\xd5\x40\x80\xa5\x7e\xe5\xb2\x6c\x02\xab\xa1\xd6\xda\x7d\x98\x13\x9e\x4e\x0b\x7e\x48\xd0\x74\xe7\xdb\xb4\xce\x8d\x2e\xbe\x8e\x0e\x63\xe3\xf4\xe6\x62\x9b\x2a\xb6\x6c\xbf\xb9\xb5\x5f\x7a\x27\x95\xac\x37\x49\x68\xa5\x8b\x6b\x6f\x9a\x33\x8a\xdd\xd9\x13\xac\xf2\x44\xbc\xbc\x6e\x03\xfc\xee\x2a\xcf\xd7\xda\x32\xba\x05\x1a\x71\x6d\x1c\x66\xb6\x9d\xe8\xba\xfe\xfa\x2c\xc2\xd8\x60\x8d\x56\xf8\x88\xcb\xe2\x87\x6b\x60\x64\xf1\x17\xf6\xfd\x92\x86\x21\x22\x2d\x32\x3b\x7f\xfd\x62\x13\x6a\x6d\x50\xbe\x3c\xef\xd3\xae\x72\x29\x17\x71\x0d\xee\xeb\x46\x58\x7c\x16\x7f\xb4\xad\x23\x6e\x4a\xef\xc6\xa1\xf2\xeb\xb0\x5f\x1a\x7c\xd9\x0e\xd3\x4d\xf6\xd0\x7e\x8c\x70\x40\x1b\x5a\xf7\xaa\x72\x4a\x28\x99\x87\xe3\xa5\xed\xf3\xef\xe9\x24\x79\x7d\x6d\x63\xfc\x54\xcf\xb3\xa7\x2d\xb8\x56\x4c\xa1\xb4\xaa\xca\x42\x55\xf7\x63\xf3\xbd\x14\x56\x8e\x38\x1b\x13\x06\xad\xdd\xb8\x18\x3a\xa5\x24\xd0\xf5\xc9\xc2\x00\x6f\x4b\x93\xb4\x2d\xd0\x9e\x79\xc0\x35\xce\x98\xbd\x6d\x33\x04\x91\x64\xb9\x9e\x43\x98\x69\x7d\xc3\x53\x0b\x74\xbc\xb5\x81\x5d\xf9\xf4\x8d\x8a\x2e\x38\x39\x91\x4a\xc4\x4c\xa0\x08\x97\xa8\xa7\xa6\x6d\xab\x48\x02\xd8\xca\xb8\x41\xb8\x66\x2c\x2e\x8b\xd4\xa3\xc5\x5b\x00\x6a\x8a\x86\x40\xaf\x79\x69\x01\xcf\x29\x62\xac\xdb\xa8\x86\x33\xb1\xd5\xea\x5e\xbd\xa2\x89\x1e\x7c\x57\x00\x43\x9d\xc3\xf2\x85\xc8\xc6\xe0\xca\x81\x23\x44\xb5\x5f\xc4\x32\x85\x9f\x71\xda\xfd\x96\x94\xf5\x79\xe2\x49\xa9\xb8\x16\xe4\x70\xf5\x42\xe5\xe8\x5f\x47\x95\xcf\x6d\xb9\xdc\xd8\x57\xd9\x70\xef\x66\xfd\x9e\xf2\x35\x30\xb8\xc1\x66\xf3\xd2\x12\x9a\x2b\x5c\x66\x73\x7e\x23\xff\xb8\xe1\xd4\xde\x87\x21\x5f\x34\x0c\x09\xb5\x35\x3b\x7f\x4b\x66\xfc\x4c\xe6\xbf\xee\xf4\xa6\xa8\xd4\xf6\xb9\x36\x53\xdd\xf8\xe6\xcd\x50\x8e\x4a\xdb\x88\x40\x17\x5d\xfb\x05\xd7\xa2\xe5\x8f\x2e\x13\x89\xe8\x3b\xc8\xea\x2b\x35\x75\x73\xe2\xc7\x42\xf8\x71\xf5\x20\x70\x7a\x60\xf4\xea\xe1\x12\xa7\x61\xfd\x5e\x4c\x88\xa6\x82\x5d\xa9\x3e\x02\x34\x94\xa3\xa5\x91\x51\xb8\x8d\xef\x90\x1b\x5c\x19\x20\xdd\xbd\xac\xac\x90\xed\xaf\x98\x5b\x3e\x4d\xcf\x5d\xef\x93\xfb\x8f\x62\x11\x9a\x53\x4c\xfe\xc6\x1b\xb9\x82\x91\xb3\x4e\xa4\xf3\x23\xe7\x38\xe3\x44\x9a\xcc\x5d\x5d\xd6\x19\xae\xd9\x5a\xca\x53\xc7\xdc\xed\x4f\x38\x5b\x52\x62\xf6\x7f\x8d\xa7\x28\xd8\xab\xef\x42\xdf\x64\x23\xd5\x2f\xd6\xd2\x36\xbf\x30\xb8\x94\xe2\xef\x1d\xda\x23\x3d\x34\xeb\x6a\xfa\xbe\x90\xc1\x3b\x6e\x7f\x35\xef\xde\xd8\x7d\x0d\x5b\xfc\x2b\x6b\xbf\x85\xcf\xc6\x6b\xb8\xf2\x3d\xdf\x3c\xdb\xfa\x34\x60\x80\x71\x25\x37\x86\x5e\x37\xdf\x4c\x2d\xe6\xd9\x68\x65\xb9\xd2\x62\x43\xf6\xf7\x06\x26\x9f\xdd\x05\x03\xdf\x97\x5a\xff\xe7\x8a\x00\x1f\x02\x0e\xdf\x9b\xee\xfd\x2d\x14\x28\xfd\xba\xd6\xbc\xcf\x5f\x7a\xbe\xc8\x96\xf1\xe2\x6a\xae\x0d\x23\x5b\xdb\x36\xeb\x1d\xb2\xeb\x9d\x39\xc3\xbe\xf2\xda\x36\xfe\x1a\x75\x17\x3a\x33\x31\x51\x8d\x1d\xa5\x6b\x58\xfe\x4e\xca\xb0\xb3\xe9\x8e\x5f\x47\x4f\x1f\x32\xef\x2c\x15\xe2\xd3\x02\xb3\x65\xf6\xfb\xf8\xa8\x16\xde\xc6\x77\xb7\x13\x81\x96\x7e\x46\xbb\x7a\xe9\xf1\xd2\xef\xf0\x34\x36\x93\x7f\xf4\xfe\x8a\xdc\x57\xf4\x62\x44\xf5\x0d\x72\x93\xad\xdf\x34\x8b\xb4\x8e\x97\x7f\xaf\x4f\x29\x7e\x31\x1f\xe5\x5e\x41\x7c\x5e\x05\xf1\x3f\x01\x00\x00\xff\xff\x1d\x1f\x03\x41\x19\x80\x00\x00")

func ApiSwaggerV1OauthSwaggerYamlBytes() ([]byte, error) {
	return bindataRead(
		_ApiSwaggerV1OauthSwaggerYaml,
		"../../api/swagger/v1/oauth.swagger.yaml",
	)
}

func ApiSwaggerV1OauthSwaggerYaml() (*asset, error) {
	bytes, err := ApiSwaggerV1OauthSwaggerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../api/swagger/v1/oauth.swagger.yaml", size: 32793, mode: os.FileMode(493), modTime: time.Unix(1608335107, 0)}
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
	"../../api/swagger/v1/oauth.swagger.yaml": ApiSwaggerV1OauthSwaggerYaml,
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
						"oauth.swagger.yaml": &bintree{ApiSwaggerV1OauthSwaggerYaml, map[string]*bintree{}},
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

