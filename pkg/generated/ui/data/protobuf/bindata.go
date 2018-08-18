// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x73\xdb\x36\x12\x7f\xd7\xa7\xd8\xd1\xcb\x39\x37\xb6\x98\x38\xed\x5d\xc7\xaa\x6f\xaa\xca\x69\xaa\x49\x22\x7b\x42\x37\x9e\x3c\x69\x20\x70\x45\xe1\x0c\x02\x28\x00\x4a\x61\x32\xfe\xee\x37\xf8\x43\x8a\xa4\xa8\xe4\xae\xed\xc3\x69\x26\x31\x09\xec\xfe\xb8\xfb\xdb\x3f\x58\x32\x49\x60\x2e\x55\xa5\x59\xbe\xb5\x70\xf9\xfc\xc5\x0f\x90\x92\xc2\x94\x22\x87\xf4\x26\x85\x39\x97\x65\x06\x4b\x62\xd9\x0e\x61\x2e\x0b\x55\x5a\x26\x72\xb8\x47\x52\x00\x29\xed\x56\x6a\x33\x19\x25\xc9\x28\x49\xe0\x2d\xa3\x28\x0c\x66\x50\x8a\x0c\x35\xd8\x2d\xc2\x4c\x11\xba\xc5\x7a\xe7\x1c\x3e\xa0\x36\x4c\x0a\xb8\x9c\x3c\x87\x33\x27\x30\x8e\x5b\xe3\x67\x53\x07\x51\xc9\x12\x0a\x52\x81\x90\x16\x4a\x83\x60\xb7\xcc\xc0\x86\x71\x04\xfc\x44\x51\x59\x60\x02\xa8\x2c\x14\x67\x44\x50\x84\x3d\xb3\x5b\xff\x9c\x88\xe2\x2c\x81\x8f\x11\x43\xae\x2d\x61\x02\x08\x50\xa9\x2a\x90\x9b\xb6\x20\x10\x1b\x8d\x76\xbf\xad\xb5\xea\x2a\x49\xf6\xfb\xfd\x84\x78\x83\x27\x52\xe7\x09\x0f\xa2\x26\x79\xbb\x98\xbf\x5a\xa6\xaf\x2e\x2e\x27\xcf\xa3\xd2\x6f\x82\xa3\x31\xa0\xf1\xf7\x92\x69\xcc\x60\x5d\x01\x51\x8a\x33\x4a\xd6\x1c\x81\x93\x3d\x48\x0d\x24\xd7\x88\x19\x58\xe9\x8c\xde\x6b\xe6\x78\x3b\x07\x23\x37\x76\x4f\x34\x3a\x98\x8c\x19\xab\xd9\xba\xb4\x1d\xce\x6a\x13\x99\xe9\x08\x48\x01\x44\xc0\x78\x96\xc2\x22\x1d\xc3\xcf\xb3\x74\x91\x9e\x3b\x90\x87\xc5\xfd\xaf\xb7\xbf\xdd\xc3\xc3\xec\xfd\xfb\xd9\xf2\x7e\xf1\x2a\x85\xdb\xf7\x30\xbf\x5d\xde\x2c\xee\x17\xb7\xcb\x14\x6e\x7f\x81\xd9\xf2\x23\xbc\x59\x2c\x6f\xce\x01\x99\xdd\xa2\x06\xfc\xa4\xb4\xf3\x40\x6a\x60\x8e\x4d\xcc\x3c\x75\x29\x62\xc7\x84\x8d\x0c\x26\x19\x85\x94\x6d\x18\x05\x4e\x44\x5e\x92\x1c\x21\x97\x3b\xd4\xc2\x65\x82\x42\x5d\x30\xe3\xa2\x6a\x80\x88\xcc\xc1\x70\x56\x30\x4b\xac\x5f\x3a\xf2\x6b\x32\x72\x22\xef\x18\xdd\x12\xe4\xf0\x01\x05\x7e\x66\x04\x7e\x2c\x76\xe1\xea\xa7\xbc\x20\x8c\x4f\xa8\x2c\xfe\xe5\xe4\x66\x9c\x3d\x12\x78\x4b\xb4\x41\x01\x3f\x12\x77\x37\xe1\xfe\xae\x2d\x38\x32\x95\xb0\xe4\x13\x5c\xc3\x58\x69\x69\xe5\xcb\xf1\x74\x34\x52\x84\x3e\x3a\x53\x29\x2f\x8d\x45\xbd\x2a\x88\x20\x39\xea\x15\x51\x6c\x3a\x1a\xb1\x42\x49\x6d\x61\x9c\x4b\x99\x73\x4c\x88\x62\x09\x11\x42\x46\xb3\x27\x1e\x66\x3c\x6d\xc4\xfc\x3d\xbd\xc8\x51\x5c\x98\x3d\xc9\x73\xd4\x89\x54\x5e\x74\x50\x6d\x34\x0a\xbb\x70\x96\x6b\x45\x27\x39\xb1\xb8\x27\x55\xd8\xa6\xab\x1c\xc5\x2a\xa2\x4c\x22\xca\x44\x2a\x14\x44\xb1\xdd\x65\xbd\xf3\x0c\xae\xe1\xcb\x08\x80\x89\x8d\xbc\xf2\x57\x00\x96\x59\x8e\x57\x30\x9e\x07\x97\xe0\x5d\x70\x09\x66\x77\x8b\xf1\xd4\x4b\xec\x42\x81\x5d\xc1\x78\xf7\x7c\x72\x39\x79\x1e\x97\xa9\x14\x96\x50\x5b\xe3\xb8\x9f\x20\x85\x83\xaa\x6b\x7d\xbe\x9c\xdf\x47\x61\xf7\x2b\x35\xbf\x82\xb1\x2b\x0c\x73\x95\x24\x39\xb3\xdb\x72\xed\xb8\x4e\x4c\x90\xbf\xa0\x82\xda\x24\x52\x7b\x11\xa9\xbd\x20\x8a\xb5\x30\xd0\x05\xe8\x0a\xc6\x24\x2b\x98\xf8\xa9\xad\x38\x61\x32\xca\x3d\xb9\x3f\xfe\x3f\xfc\x64\x51\x0b\xc2\x57\x99\xa4\xa6\x36\xf4\xcf\x9a\x91\xa1\xa1\x9a\x79\x8a\xaf\x60\xfc\x4e\x6a\x04\xb2\x96\xa5\x85\x53\x0c\x3e\x8d\x00\x0c\xdd\x62\x81\xe6\x0a\x7e\xbd\xbf\xbf\x4b\xa7\xfd\x15\xb7\x40\xa5\x30\xa5\x5f\x19\xc7\xc2\x77\x8f\x48\xfe\x6d\xa4\xf0\x30\x4a\xcb\xac\xa4\xa7\xf6\x9f\xa6\xa3\x91\x41\xbd\x63\x14\x1b\x43\x82\xbf\xae\x9e\x19\xe7\x4e\x7f\xc7\x7c\xa7\x24\x75\xfe\xfa\x7d\xad\x28\xcc\x35\x12\x8b\xb5\xde\x59\xe7\xf6\x9d\xc9\x9f\x81\x46\x5b\x6a\x61\x7a\x5b\xef\x51\xf1\xea\x59\x2b\x01\x9a\x0c\xf5\x15\x30\x21\x8a\x4d\x1c\xd1\x75\xde\x1d\x7e\xaa\xb4\x70\x05\x63\x5f\x23\xbb\x17\x35\xdb\xe3\x8e\xcc\x5a\x66\x95\x13\xfa\xfb\x61\xf9\x29\x46\xb8\xe3\x98\x46\xab\x19\xee\x42\x9b\x31\x96\xd8\xd2\xb8\xd6\xdc\x78\xe9\x5a\x08\x30\x6b\xe0\xb1\x5c\x23\x95\x62\xc3\x72\xdf\x85\xa8\x14\x02\xa9\x65\x3b\x66\xab\x86\x89\xd7\x68\x1b\x1a\x0e\xd7\x5d\x0e\x0e\xeb\x7f\x9c\x80\x1c\xbf\x4e\xc0\xa0\xa7\x19\x72\xb4\x38\x10\xbf\x1b\xbf\xd1\x18\xde\xb9\xed\xda\xde\xd9\xfa\xe3\xe6\x47\x4b\xfe\x67\x0f\x9a\x58\x11\xe0\xcc\x58\x17\xa7\xa8\x68\x06\x42\xf0\xd6\x89\x9c\x75\xef\x4f\x85\xc2\xed\xfd\xd5\xe1\x48\x9c\x8d\xdf\xf6\xa8\xd4\xa2\x6e\x92\xbe\xb5\xea\xc2\x97\x66\x6c\x0b\x44\x31\x70\x95\xd9\x0a\xd7\x6b\xb4\x71\x6a\x59\xb4\xc4\xcf\x0e\xcb\x47\x4e\xc6\xf5\xbf\xcc\xc1\x68\xee\x80\x6f\x4f\xa3\x51\x81\xc6\xb8\x53\xae\xdf\x06\x0e\x0d\x65\x49\x0a\xac\xc7\x9f\xba\xca\xac\x84\x35\x1e\xba\x0c\x66\x5e\xd8\x0d\x1b\x22\xf7\x27\x03\x5c\xc3\x8b\x69\x8d\x70\xbf\x8d\xb2\xee\x28\xaf\x67\x01\xcf\x83\x97\xe8\x3c\xfa\x2e\xca\xa5\x0a\xe9\x41\xe9\x1a\x2e\xa7\x27\xad\xf5\x44\xb5\x1a\xe0\x16\xfd\x8c\x22\xb5\x1f\x03\xdb\x66\xef\x89\x69\x1b\xed\xe6\x2e\x3f\x21\xba\x41\x0c\x8d\x1d\x85\x4e\x24\x39\xc8\xc7\x23\x07\x32\xb4\x84\x71\xd3\x67\x22\xaa\x82\x46\xa3\xa4\x30\x18\x3c\x0a\x9b\x0b\x8b\x45\x23\xd8\x77\xa1\xd3\x70\xfe\x1b\xb6\xb9\x94\x8f\x6e\xd0\x53\xc3\x5c\x0f\x42\xf7\xa8\x59\x98\x0e\x2e\x13\xa1\x8d\x56\xc6\x62\x71\xec\x7c\xdb\x95\x1b\xef\xfd\x57\x1d\xea\x37\xa2\x76\x44\x88\x75\xe3\x68\xeb\xd9\x7f\x33\xc1\x74\x2b\xdd\x19\x6b\xb5\xac\xbe\xe9\xd5\x71\x37\x3b\x3c\x61\x2e\x4b\x9e\x75\x7c\x5b\x63\x0d\x1c\x93\x73\x28\xae\x69\x73\x80\x38\xd5\x76\x16\x44\x43\xe2\x09\x73\x3a\x76\xb1\x4b\xc1\x97\xd3\xdb\x7f\x2a\x06\x51\xe9\xed\x60\xff\x44\xe5\xaa\x20\x1b\x4a\xb7\x63\x9b\xdb\x42\x07\x63\x6e\x7a\xb9\xd6\x76\x9e\x65\x1d\x1b\x06\x32\x73\x20\x66\x97\xd3\xa1\xa8\x9b\x0e\xd1\x03\xda\x0d\xd1\x2f\x87\x8c\x6e\x65\xdf\xff\xb7\xe9\x03\xfa\xad\x41\xc4\xca\x7a\x0e\x71\x97\x27\xe0\x5a\xf2\xd7\xf0\xdd\xe9\xae\xd7\x69\x94\x83\xa5\xd6\x74\xcf\x0b\xa0\xa5\xd6\x28\x2c\x8f\xfd\x8e\x19\x20\x7b\xff\xf6\x56\x10\x62\x4e\xf7\xee\xda\xa4\x1f\xcc\xaa\x3e\xf3\x3c\x4d\x7e\x53\x0a\x94\x9b\xe6\x21\xab\xf6\x61\xf8\xa5\xdd\x38\x67\x0f\xe9\xb7\x9a\xfe\xec\x21\xf5\x6e\x38\xa3\x1a\x1e\x4f\x1f\x4f\xb5\xf8\xd1\x63\x6e\x88\x25\x30\x47\x51\x53\x3a\x7b\x48\xdd\x52\x58\x81\x8c\x58\xb2\xa2\xe1\xba\x9d\x1d\x73\x8d\x19\x0a\xcb\x08\x37\xbe\xd5\x96\xac\xdb\x4b\x6a\xa8\xb6\x1c\x6d\x5d\xb7\x33\xe7\xe7\x8f\xb7\xc0\x2c\x16\xa6\x56\xba\xd3\x31\x9a\xa5\xc6\xcc\x65\xb0\x3b\x6d\x8c\x2c\x35\xc5\x6e\xce\x2c\x84\xb1\xfe\xa3\x44\xae\x65\xa9\x7a\x15\x3e\x7b\x48\xeb\xfd\xd7\x6e\x1b\x58\xbc\x5b\x05\xe9\x90\x2a\x87\x24\x60\x74\x7b\x44\x46\x4d\x65\x97\x94\xc3\xe4\xd0\x28\x6a\xcc\xfd\x9c\x51\x9a\x0b\x24\xc6\x5e\xbc\x38\x07\xb4\x74\xf2\xac\x91\x8c\x69\x11\xe5\x1a\x2a\x3b\x20\x64\x47\x18\x27\x6b\xc6\x99\xad\xe0\xb3\x14\x68\x5a\x80\xeb\x73\xa8\xaf\x2f\xa9\xbf\xde\xa3\xbb\xce\xfa\x4f\x6a\x08\x88\x8f\x6c\xa3\xae\x02\x6a\xc3\xfe\xd3\xa8\x9d\x0e\xb4\x1b\xd4\x32\x7e\x8d\xa0\x2e\x8f\x1c\x52\x3f\xbe\x2d\x76\xda\x71\xee\xd0\xe3\x70\x53\xa4\x1a\xed\x1b\xac\x16\x99\x07\x9c\xdd\x2d\x60\x46\x29\x1a\xd3\xa7\xc7\x78\xc9\xd5\x23\x56\xab\x76\x3b\x3a\xc2\x0a\xda\x6f\xb0\x6a\xf0\xc8\xd7\xf0\xc2\xa6\x83\x1d\x72\xfd\x17\xa9\x61\xbf\x45\x01\x46\x16\xfe\xf3\x97\xc8\x0d\x10\xf7\xda\xca\x35\x92\xac\x0a\x04\xc4\x23\xb1\xe5\xf3\x40\x9a\x1e\xb9\xfe\xe1\x6e\x0e\x2c\x3b\x87\x35\x27\xe2\xd1\x1b\xeb\xfe\x8d\x03\xa2\x6b\x07\xfe\xbe\x92\xe5\xf8\x1c\x36\x8c\x73\xcc\x80\x6d\xfc\x27\x39\x67\x80\xab\x8c\x0f\x77\xf3\xbe\x57\x3b\x45\x87\xe8\x49\x91\x96\xda\xa5\x8e\xcf\xef\x01\x2a\xfc\x6e\xc8\xfe\xa0\x7f\x44\x45\xaf\xa0\x20\xc3\x0d\x13\xee\x65\xc4\x56\x0a\xfd\x7b\xa2\x28\x8b\xb5\x1b\x14\x37\x4d\x39\x99\x3e\x2f\xdd\xaa\xeb\x50\xd2\xe0\x7b\xbc\xb3\xe2\xfb\x09\x27\x3a\xc7\x13\xc5\xe2\x85\xfa\x5e\x2e\x8f\x0d\x80\xb3\x0c\x37\xa4\xe4\xd6\xa7\xed\x67\xd4\xf2\x00\xc5\x84\x7d\x79\x09\xbf\x97\x44\x58\x47\x4d\xcb\xe7\x27\xff\x51\xec\x35\xda\xe6\xdd\x44\x6e\x7c\x2a\xa5\xe1\x65\xa4\x35\x99\x1c\xde\x3a\xc2\xd0\x92\x24\x10\x26\x14\x17\xbc\x5a\xbb\x1e\x85\x8e\xf5\xfa\xd3\xcc\x06\xa4\x42\x1d\xda\xbe\x1b\xaf\x6f\xdf\x9c\x18\x24\x6b\xa8\x81\x97\xa1\xa3\x4c\xb3\x24\x07\x19\x06\xa3\x9c\xb9\xd9\x5a\x49\xc3\xac\xd4\x55\x9f\xd5\x9c\xd9\xd6\xc9\x74\x5c\x61\x5b\x62\xb6\xf5\xd1\xed\x90\xa8\x2c\x0a\x66\x87\x50\xc2\xce\x81\x53\x38\xf9\x32\x61\x35\xa2\x77\x95\x72\x24\x22\x54\x9b\x3b\x33\x06\x61\x9d\xf0\xca\x4d\x08\x78\x68\xf6\x11\xfa\xc6\x57\xcd\x26\x9c\x37\x7d\x5d\xbf\xb8\xca\x82\xde\x77\x1d\xbd\x0f\x87\x08\xe7\xbe\xb1\x65\x61\xb0\x28\x14\xe3\x78\x64\x83\x6c\xf1\xf3\x7d\x07\x67\x1e\x34\xb4\x87\xe8\xeb\xd1\x7a\xf3\x1a\xfe\xd1\xd1\xba\xe3\xc4\xba\xc8\x01\xb3\x81\x84\x20\x18\x9a\x61\x02\xba\x14\xfe\xab\x6e\x3c\xe5\x5b\x88\xaa\x56\xbc\x86\x7f\xf6\x4b\xb5\x76\xa9\x95\x14\x7e\x6b\x20\x57\xa2\x37\x9d\x71\xa3\x1e\x75\x47\xff\x09\x00\x00\xff\xff\x39\x51\x30\x25\x8b\x18\x00\x00")

func apiProtoBytes() ([]byte, error) {
	return bindataRead(
		_apiProto,
		"api.proto",
	)
}

func apiProto() (*asset, error) {
	bytes, err := apiProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.proto", size: 6283, mode: os.FileMode(420), modTime: time.Unix(1534551504, 0)}
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
	"api.proto": apiProto,
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
	"api.proto": &bintree{apiProto, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
