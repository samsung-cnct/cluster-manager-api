// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3b\x6d\x6f\xdb\x38\xd2\xdf\xfd\x2b\x06\xfe\xf2\x38\x0f\x1a\xbb\x4d\xdb\xbb\x45\x72\x7d\xf6\x7c\xc9\x5e\x6b\xb4\x49\x83\x3a\xdb\x60\x3f\x19\xb4\x34\x96\x79\x91\x48\x1d\x49\xd9\xf5\x2e\xfa\xdf\x1f\xf0\x4d\x22\xf5\x92\xb6\x69\x16\x7b\x58\x5c\x80\xbb\xb5\xc4\x99\xd1\x70\xde\x87\x9c\xce\x66\x70\xce\xcb\x83\xa0\xd9\x56\xc1\xc9\xd3\x67\x3f\xc0\x92\x14\xb2\x62\x19\x2c\x2f\x96\x70\x9e\xf3\x2a\x85\x2b\xa2\xe8\x0e\xe1\x9c\x17\x65\xa5\x28\xcb\xe0\x06\x49\x01\xa4\x52\x5b\x2e\xe4\x74\x34\x9b\x8d\x66\x33\x78\x47\x13\x64\x12\x53\xa8\x58\x8a\x02\xd4\x16\x61\x5e\x92\x64\x8b\x7e\xe5\x09\x7c\x44\x21\x29\x67\x70\x32\x7d\x0a\x13\x0d\x30\x76\x4b\xe3\xa3\x33\x4d\xe2\xc0\x2b\x28\xc8\x01\x18\x57\x50\x49\x04\xb5\xa5\x12\x36\x34\x47\xc0\x4f\x09\x96\x0a\x28\x83\x84\x17\x65\x4e\x09\x4b\x10\xf6\x54\x6d\xcd\x77\x1c\x15\xcd\x09\xfc\xe2\x68\xf0\xb5\x22\x94\x01\x81\x84\x97\x07\xe0\x9b\x10\x10\x88\x72\x4c\xeb\xbf\xad\x52\xe5\xe9\x6c\xb6\xdf\xef\xa7\xc4\x30\x3c\xe5\x22\x9b\xe5\x16\x54\xce\xde\x2d\xce\x7f\xba\x5a\xfe\x74\x7c\x32\x7d\xea\x90\x7e\x66\x39\x4a\x09\x02\xff\x5d\x51\x81\x29\xac\x0f\x40\xca\x32\xa7\x09\x59\xe7\x08\x39\xd9\x03\x17\x40\x32\x81\x98\x82\xe2\x9a\xe9\xbd\xa0\x5a\x6e\x4f\x40\xf2\x8d\xda\x13\x81\x9a\x4c\x4a\xa5\x12\x74\x5d\xa9\x48\x66\x9e\x45\x2a\x23\x00\xce\x80\x30\x18\xcf\x97\xb0\x58\x8e\xe1\x1f\xf3\xe5\x62\xf9\x44\x13\xb9\x5d\xdc\xbc\x79\xff\xf3\x0d\xdc\xce\x3f\x7c\x98\x5f\xdd\x2c\x7e\x5a\xc2\xfb\x0f\x70\xfe\xfe\xea\x62\x71\xb3\x78\x7f\xb5\x84\xf7\xff\x84\xf9\xd5\x2f\xf0\x76\x71\x75\xf1\x04\x90\xaa\x2d\x0a\xc0\x4f\xa5\xd0\x3b\xe0\x02\xa8\x96\x26\xa6\x46\x74\x4b\xc4\x88\x85\x0d\xb7\x2c\xc9\x12\x13\xba\xa1\x09\xe4\x84\x65\x15\xc9\x10\x32\xbe\x43\xc1\xb4\x25\x94\x28\x0a\x2a\xb5\x56\x25\x10\x96\x6a\x32\x39\x2d\xa8\x22\xca\xbc\xea\xec\x6b\x3a\xd2\x20\x97\x34\xd9\x12\xcc\xe1\x23\x32\xfc\x95\x12\xf8\x5b\xb1\xb3\xbf\xfe\x9e\x15\x84\xe6\xd3\x84\x17\xff\xa7\xe1\xe6\x39\xbd\x23\xf0\x8e\x08\x89\x0c\xfe\x46\xf4\xd3\x34\x37\x4f\x21\xe0\x48\x1e\x98\x22\x9f\xe0\x15\x8c\x4b\xc1\x15\x7f\x3e\x3e\x1b\x8d\x4a\x92\xdc\x69\x56\x93\xbc\x92\x0a\xc5\xaa\x20\x8c\x64\x28\x56\xa4\xa4\x67\xa3\x11\x2d\x4a\x2e\x14\x8c\x33\xce\xb3\x1c\x67\xa4\xa4\x33\xc2\x18\x77\x6c\x4f\x0d\x99\xf1\x59\x0d\x66\x9e\x93\xe3\x0c\xd9\xb1\xdc\x93\x2c\x43\x31\xe3\xa5\x01\xed\x45\x1b\x8d\xec\x2a\x4c\x32\x51\x26\xd3\x8c\x28\xdc\x93\x83\x5d\x4e\x56\x19\xb2\x95\xa3\x32\x75\x54\xa6\xbc\x44\x46\x4a\xba\x3b\xf1\x2b\x47\xf0\x0a\x7e\x1b\x01\x50\xb6\xe1\xa7\xe6\x17\x80\xa2\x2a\xc7\x53\x18\x9f\xdb\x2d\xc1\xa5\xdd\x12\xcc\xaf\x17\xe3\x33\x03\xb1\xb3\x0e\x76\x0a\xe3\xdd\xd3\xe9\xc9\xf4\xa9\x7b\x9d\x70\xa6\x48\xa2\x3c\x1d\xfd\xc7\x48\xa1\x49\x79\x5f\x3f\xbf\x3a\xbf\x71\xc0\xfa\xaf\x12\xf9\x29\x8c\xb5\x63\xc8\xd3\xd9\x2c\xa3\x6a\x5b\xad\xb5\xac\x67\xd2\xc2\x1f\x27\x2c\x51\x33\x27\xda\x63\x27\xda\x63\x52\xd2\x80\x06\x6a\x05\x9d\xc2\x98\xa4\x05\x65\x7f\x0f\x11\xa7\x94\x3b\xb8\xcf\xfa\x3f\xe6\xff\xf0\x93\x42\xc1\x48\xbe\x4a\x79\x22\x3d\xa3\xdf\xcb\x46\x8a\x32\x11\xd4\x88\xf8\x14\xc6\x97\x5c\x20\x90\x35\xaf\x14\x0c\x49\xf0\xf3\x08\x40\x26\x5b\x2c\x50\x9e\xc2\x9b\x9b\x9b\xeb\xe5\x59\xfb\x8d\x7e\x91\x70\x26\x2b\xf3\x66\xec\x1c\x5f\x7f\x62\xf6\x2f\xc9\x99\x21\x53\x0a\x9e\x56\xc9\xd0\xfa\xe7\xb3\xd1\x48\xa2\xd8\xd1\x04\x6b\x46\xec\x7e\xb5\x3f\xd3\x3c\xd7\xf8\x3b\x6a\x22\x25\xf1\xf6\x6b\xd6\x45\x99\xc0\xb9\x40\xa2\xd0\xe3\x4d\xa2\xc7\x4b\x99\x1d\x81\x40\x55\x09\x26\x5b\x4b\x1f\xb0\xcc\x0f\x47\x81\x01\xd4\x16\x6a\x3c\x60\x4a\x4a\x3a\xd5\x82\xf6\x76\xd7\xfc\x95\x95\x82\x53\x18\x1b\x1f\xd9\x3d\xf3\xd2\x1e\x47\x30\x6b\x9e\x1e\x34\xd0\xff\x36\xaf\x3f\x3b\x0d\x47\x1b\x13\xa8\x04\xc5\x9d\x0d\x33\x52\x11\x55\x49\x1d\x9a\xeb\x5d\xea\x10\x02\x54\x49\xb8\xab\xd6\x98\x70\xb6\xa1\x99\x89\x42\x09\x67\x0c\x13\x45\x77\x54\x1d\x6a\x49\xbc\x46\x55\x8b\xa1\xf9\x1d\xcb\xa0\x79\xff\x70\x01\x64\x78\xbf\x00\x7a\x77\x9a\x62\x8e\x0a\x7b\xf4\x77\x61\x16\x6a\xc6\xa3\xc7\x98\xf7\x68\xe9\xe1\xec\x3b\x4e\xbe\x79\x07\x24\xfd\x57\x25\x15\x90\x7b\xad\x71\x6e\x80\x1c\x8f\x57\x3c\x45\x09\x93\xe8\x5d\xbc\xa5\x68\xe9\x3b\x4c\x92\xcb\xdf\xc1\x26\x09\xe4\x54\x2a\x6d\x8f\x8e\x9e\xec\x31\xb5\x77\x1a\x64\x12\x3f\x0f\x99\x9c\x5e\x7b\x6c\xb3\x9b\x69\x1e\xbf\xb0\x23\xca\xa4\x22\x79\x0e\x13\x2e\x40\xa0\x7b\x3a\x02\x45\xf3\x3c\x50\xdd\xb5\x57\xeb\x8d\x79\x0f\x93\xd6\x8b\x78\x57\xad\xc5\xc7\xd3\x9d\xe5\xea\x61\xaa\x1b\xd8\xe8\x16\xf3\x02\x92\x2d\x11\xca\x43\xdf\xe8\x42\x72\xaf\x51\xd6\xa8\xb3\x82\x12\x55\x62\x4a\x5a\x6a\x82\x8f\x06\x85\x2d\x91\x40\x72\x81\x24\x3d\xc0\x1a\x91\x41\x8a\x65\xce\x0f\x98\x36\x75\xa6\x24\x05\x9a\xcc\x59\x0b\x71\x61\xbf\xf9\x06\xf3\xe2\xdc\x50\x99\xb4\xdf\xc4\x62\x6c\xaf\x3e\x5a\x58\xd6\x7b\x7e\x98\x10\x5d\x7c\xa8\x77\xdb\x92\x5e\x13\xb3\x82\x4d\xb6\x5e\xf4\xc5\xad\x47\xd8\x62\x37\x72\xc5\xbb\x1c\x72\xe7\x4a\x30\x5f\x09\x99\xfa\x49\x14\x26\xff\xba\xdc\x4f\x4a\x0a\x3a\xfd\x06\xae\xf0\x1a\x95\x6b\x4d\x16\x01\xf8\xa4\x79\xdd\xf1\x70\xf7\xfe\xd1\xbc\xdb\xb1\xfb\x75\x7b\xab\xca\x4c\x90\x14\xdd\xc7\xa4\xc9\x90\x04\x32\xba\x43\xd6\x09\xd0\xaf\x51\xfd\x6c\xc1\x5d\x50\x6a\xef\x70\x70\xb5\xb3\xe7\x41\xc8\x47\x8f\x71\x6e\x83\x5f\x4a\x50\x4a\x61\x51\x2a\xdd\x62\x79\x89\x74\x13\x54\xcc\x34\x4c\xe2\xe7\x78\x8f\xf1\xda\xa3\xa7\xa7\xee\xb6\xbe\xe4\xa6\x9f\x47\xa3\x02\xa5\xd4\x6d\x4c\xbb\xce\x6b\x2a\xc6\x2b\x1d\x94\x5c\x7f\xeb\xcb\x28\xc5\x75\xa0\xab\x13\x37\xa6\x06\x58\x77\x93\x2c\x33\x01\x0c\x5e\xc1\xb3\xb3\x26\x3a\x3a\x58\xdd\xab\xf9\x66\xcf\x68\xd6\x40\x44\x9f\xbe\x76\x70\xcb\x12\x93\x06\xe9\x15\x9c\x9c\x0d\x72\x6b\x04\x19\x54\xb8\x5b\x34\x4d\x28\x17\xa6\xcf\x0f\xd9\xde\x13\x19\x32\xad\x1b\x6b\x73\x04\xa0\x3b\x6d\x94\x36\x1e\xad\x39\xcf\x81\xdf\x75\x36\x90\xa2\x22\x34\x97\x6d\x49\x38\x54\x10\x28\x4b\xce\xa4\x8d\xdc\xde\x84\x15\x16\x35\x60\x7b\x0b\x51\x45\xf9\x35\xd2\xce\x39\xbf\xd3\x9d\x7c\x79\xaf\xac\xe7\xb7\x4b\x2d\x9d\x14\x99\xa2\x24\xb7\x15\xc6\xfc\x76\x19\xbc\x02\xb2\x97\x96\x1b\x8f\xf2\x6b\x25\xb0\x8b\xa4\xdf\x46\x68\x06\xec\x15\x3c\x3f\xeb\xe3\xd5\xab\x4a\xc2\x84\xec\xe5\x8c\xdc\xc9\xd9\xae\xd8\x13\x81\x33\x54\xc9\x51\xc8\x72\xa0\xd4\x17\x03\x12\x69\x69\x74\x21\x23\x71\x50\x66\x93\xe5\x41\x2a\x2c\xba\x3a\x0b\x35\x70\x61\x94\x76\xaf\x1e\xda\x05\x72\x68\x48\x44\x01\x8d\xbe\xfd\x3f\xd2\x4a\x5c\x71\x9b\xe5\xf9\xe1\xcf\xaa\x8c\x6e\x73\xd0\x08\xe6\x9c\x57\x79\x1a\xa9\xc4\x57\x3d\x3a\xc1\x0f\x7a\xd1\xb2\xee\xc7\x34\x6a\xe8\x73\x8e\x19\xd7\xb0\x0d\x7b\x8a\x2b\x86\x1b\x4e\xbe\x5a\xc0\xcf\x1e\x2a\xe0\x93\xdf\xdf\xda\xeb\x32\xfe\xa1\x16\xef\x90\xde\xf5\x76\x17\x58\xea\x50\x99\xf6\xc5\xa4\xae\xa8\x43\xa0\x86\x99\x8b\x56\x40\x0a\xf7\x47\xd3\x88\x87\x9e\xf0\xd5\xe3\x21\x8d\x50\x43\x1f\x93\x91\x7d\xf4\x60\xd7\xf6\xf1\xbc\x8f\xe9\xc0\xd7\xff\xb3\x59\xef\xc1\x0f\x8e\x23\x14\xf7\xa7\x11\xfa\xe7\x00\xb9\x00\xbe\x6d\x57\xc3\xd9\xb4\x37\xb0\xd5\xf6\x79\x0c\x49\x25\x04\x32\x95\xbb\xa4\x48\xad\xef\x70\x01\x05\x21\xf2\x8b\x09\xde\x17\xc4\x7c\x03\x6f\xab\x35\x0a\x86\x0a\x23\xac\xbb\x1f\xe4\xca\x03\x19\x39\x9a\x45\xce\x90\x6f\x6a\x2e\x56\x61\x39\xdd\x14\x3d\xee\x13\xda\xdb\xbb\xe5\x43\xa7\x84\x98\xdf\x2e\xcd\x7e\xad\xe7\x3f\x3f\x1b\x80\x7a\xeb\xa0\x9c\xa3\xbf\x18\x80\xfb\x78\x79\x4b\x04\x1a\x50\xeb\xe2\xf0\x0a\xfe\xea\x0b\xa8\xaf\xa8\x39\xa8\x84\x37\xf3\xc6\x67\xb7\x34\xdb\xae\xc8\x8e\xd0\x9c\xac\x69\x4e\xd5\x01\x5e\xc1\xcb\x48\x90\x1b\xb2\x16\x34\x71\x49\xbf\x92\xad\xda\x0a\xd5\x9e\x8b\xbb\x95\x03\x7a\x05\x7f\x39\x1b\x0d\x6a\xdf\x0b\xe2\xb7\x51\x4b\x88\x17\x44\x11\x38\x47\xe6\xad\x6a\x7e\xbb\xd4\xaf\xec\x1b\x48\x89\x22\xab\xc4\xfe\x0e\xb5\x1c\x86\x47\xcd\x5d\x45\xe3\x2c\xd0\x17\x7d\x93\xe0\x77\xe8\x3c\xff\xf8\xe5\x3d\x50\x85\x45\x1d\xb2\xaf\x85\x33\xe8\x4a\x60\xaa\x9d\x58\x57\x65\x92\x57\x22\xc1\xd8\x6d\x4c\x87\xcb\x12\x84\x4c\xf0\xaa\x6c\x05\xb9\xf9\xed\xd2\xaf\xbf\xd6\xcb\xb6\x83\x67\x09\xae\x2c\xb4\xd5\x73\xa3\x35\x9a\x6c\x3b\xc2\xf0\xa2\x8c\x85\x12\x19\xa3\x45\x14\x98\x99\x7a\xbd\x92\xc7\x48\xa4\x3a\x7e\xf6\x04\x50\x25\xd3\xa3\x1a\xd2\xe9\xcc\xc1\xd5\xa2\x8c\x88\x44\xa6\xf0\x2b\x67\x28\x03\x82\xeb\x27\xe0\x7f\x9f\x24\xe6\xf7\x1e\xf5\xef\xb4\xfd\xa5\x5a\x00\xee\x93\x21\xd5\x95\xa5\x5a\x4b\xbf\xb1\xda\x7f\x72\x01\xfb\x2d\x32\x90\xbc\x30\x97\x61\x2c\x93\xa0\x2d\xdc\x1f\x53\x24\xc6\x9a\xd2\xb6\x58\x7a\x74\xd5\xf1\xd5\x8f\xd7\xe7\x40\xd3\x27\xb0\xce\x09\xbb\x33\x1d\xa4\xfe\xdf\xd8\x52\xd4\x5e\x6f\x9e\x0f\xbc\x1a\x3f\x81\x0d\xcd\x73\x4c\x81\x6e\xcc\x05\x9d\x66\x40\x9b\xc7\xc7\xeb\xf3\xb6\x24\x77\x65\xb2\x0a\xa3\xb6\x2f\x2c\x30\xa9\x84\x96\x9f\x51\x72\x1b\x49\xba\x55\x6b\x02\x16\xff\xe4\xac\xcd\xef\x62\x7e\x09\x82\xe7\xcd\xa5\x94\xf7\xdf\x09\x11\xec\xc8\x3b\x94\x94\x3c\xa1\x26\xf4\xa4\x69\xfb\x3b\x94\x14\x2b\x4d\x61\x45\x04\x6b\x2c\xb6\x91\x76\xcb\x70\x21\xc5\x0d\x65\xba\xa5\x54\x87\x12\xcd\xc1\x34\xab\x8a\xb5\x0e\x22\x9b\xda\x6c\x65\x5b\xf4\xb1\x75\x47\x52\xaf\xe9\x1b\x7a\x93\xe2\xe5\x34\x27\x22\xc3\x01\xa3\x34\x40\x6d\x41\x5e\x52\x46\x8b\xaa\xe8\x63\x04\x26\x29\x6e\x48\x95\x2b\xe3\xfb\xbf\xa2\xe0\x0d\x49\xca\xd4\xf3\x13\x28\x28\x5b\xfd\xbb\x22\x4c\xd9\xa0\x16\x8b\xf8\x92\x7c\xfa\x0e\xca\xe4\x53\x48\xf9\x79\xd0\xbe\xce\x66\xba\x90\x0a\x93\xcf\xfc\x7a\x01\x4b\x7b\xfc\x12\x94\x5a\xcd\x39\x0b\xfc\xe6\xf0\x6c\xc9\xa5\xd5\xed\xb1\x7d\x49\xda\xc5\x6b\x97\x67\x1b\xe0\x25\x0a\x9b\xa6\x74\x53\xf9\xfe\xed\x40\x1f\xe2\x49\xf5\x1c\xff\x74\x5c\x46\x91\x0c\xb8\xad\xf4\x32\xaa\x3b\xca\x92\x4b\xaa\xb8\x38\xb4\x75\x97\x51\x15\x64\xd2\x67\x1d\x5b\xde\x12\xb9\xf5\xb5\x88\xa6\x94\xf0\xa2\xa0\xaa\x8f\x8a\x5d\xe9\x68\xab\x27\x9d\x29\x81\x68\xb6\x9a\xe4\x48\x98\x0d\x1b\x3a\x03\xf4\x92\xd5\xc0\x2b\x5d\xf2\x60\x9c\x80\x67\x33\x1d\x6d\x4d\x95\x65\xb2\x47\x1b\xd7\xbc\x5c\xa5\x16\xef\x45\x84\xf7\xb1\xd1\x70\xc6\x4d\x5a\xb4\x95\x52\x51\xd2\x1c\x3b\x3c\xf0\x40\x3e\x2f\x23\x3a\xe7\x16\x43\x34\x99\x35\xc0\x4b\xfc\xa2\xc9\xab\x01\xd6\x75\x4e\x94\xd6\x1c\x50\x65\x85\x60\x01\x53\x63\x3e\x33\x10\x15\x33\x97\xd5\x41\x45\xe2\xcb\x7f\x8f\xd8\x53\x34\xf8\x2d\x05\x46\x61\x96\x7a\x6c\xc5\xed\x26\x2a\x8f\x7c\xed\xee\x94\x9e\xc4\xd9\xb9\x72\xf7\xeb\x26\xe0\x1a\x87\x0f\x12\x75\x10\x50\xc2\x64\x1d\xd5\x09\x4b\x4c\x04\xaa\xb7\x78\x58\xd8\x5d\x6a\xbf\x9a\x27\x09\xca\xa8\xa8\x93\x06\x6a\x75\x87\x87\x55\xab\x9c\x6e\x68\x58\xac\xb7\x78\xa8\xe9\x90\x21\x3a\x76\x41\x93\x8b\xaa\x05\x4d\xeb\x83\xcd\xa2\xc3\x24\xea\x34\x6b\xbb\x83\xef\x92\x4b\xbb\x1b\x8c\x2b\xa8\xb2\xbc\x57\x26\xa4\x2c\xfb\x84\x71\x83\x8c\x30\x75\xcf\x06\x94\x05\x68\x6f\xfc\x9a\x48\xb9\xe7\x22\xbd\x07\xb3\xf4\x20\x61\xa1\x64\x14\x50\xad\xeb\x1b\xea\xfb\xa4\x1f\x80\x59\xd6\x5f\x78\xd3\xd2\xe2\x21\x49\xc2\x2b\xa6\xac\xd7\x99\x33\xb3\xe8\x6c\xca\xe5\x74\x5b\x48\x37\x45\xdb\x04\x3f\x9d\x42\xce\x49\x0a\x6b\x92\xeb\x48\x2f\x8e\x5a\x02\xb6\x24\x96\xf6\xaa\x7a\xee\x3e\x12\x89\xfa\x3c\xa7\xc8\xd4\x22\x85\x09\xb9\x23\xa7\x46\xf2\x17\x51\x7b\x9d\x18\x80\x3e\x71\x5b\x54\x6b\x81\x0e\xdd\x8b\xa9\x8f\x82\x35\x40\xef\x54\xfd\xe5\xf4\xdb\xbe\x72\xda\xec\xba\x5d\x43\x3a\xda\x39\x4f\xbc\xaf\x7e\x6b\x19\xdd\xb6\xc0\xa1\x42\xda\x1f\x3c\xfb\x0b\xff\x48\x57\xba\x81\x24\xf9\x9d\xfe\xaf\xd5\x8e\xd1\x49\x6d\xfa\x2d\xe5\x34\xdf\xed\x55\x8c\x63\xcf\x3f\x7e\x7d\x4d\xfe\xf6\x5b\x6a\xf2\x47\xa8\x92\xda\xdf\xeb\x24\x5a\x16\x74\xfa\xbd\x25\x63\xdc\xe0\xb6\x39\xb3\xf5\xd5\x52\x11\x96\x12\x91\xae\x2e\x4e\x56\xbb\x93\xfb\xab\xac\x93\xdf\xad\xca\x7a\xfe\xbb\x55\x59\x2f\xbe\x74\x49\x10\x34\xc6\xb5\x47\x5c\x92\x64\x4b\x75\xb7\xb1\x37\xcd\x8d\x4e\x93\x82\x4a\xec\x98\x77\x6d\x1c\x96\x88\x43\x33\xb4\x12\xce\x94\xe0\xf9\xaa\xcc\x09\xc3\x15\x33\x57\xfe\xa1\xad\x3d\xc2\x27\x74\x07\x8d\xa2\xa6\xfd\x22\x88\x1b\x54\x82\xdc\x9a\xe3\xcd\xb5\x36\xba\x1d\xc9\x2b\x84\x9c\xde\x21\xd0\xf2\xd4\xcc\x6a\xa9\x2d\x51\xf5\x55\x2f\x81\x1d\x15\xaa\x22\x39\x2c\xae\x67\x7a\xd9\x53\xd2\x41\x5b\x07\x4b\x92\x18\xef\xf3\xd3\x47\x90\x54\x52\xf1\x02\x85\x74\xfd\xbd\x99\xe8\xb3\x85\x4c\x51\x31\x9a\xe8\x40\x3a\x7c\xec\x43\x4a\xba\x42\x96\x96\x9c\x1a\x07\x7c\x19\x1d\xfc\x34\xa7\x2e\xef\xc8\x1a\xf3\x38\x4c\x79\x93\x27\x90\xeb\xc5\x2f\x9f\xe6\x98\x8d\xf7\x23\xd8\xb5\x56\xf9\x11\x9d\xcb\xb8\x1b\xc3\x7a\xae\x4f\x4b\xba\xe6\xb3\xab\x8f\x88\xd3\x4a\x6a\x49\x15\x36\x4b\x2f\x97\x6f\x7a\xd2\x55\x0d\xd2\xe6\x7a\xcb\xa5\xba\x07\xcf\x2c\xb7\xd3\xab\x51\x6a\x0f\x8e\x75\x0a\xb3\xda\xce\xaa\x77\x3f\xc8\xba\x67\xa8\xdb\x45\x6b\xb6\x60\xcc\x76\x6a\x2d\x49\x6f\x1b\xa8\x04\xce\xf2\x03\x10\x28\x88\x3d\x10\xda\xf8\xb9\x53\xcc\x53\x33\x86\x69\x42\x5c\x3a\x1d\xc8\xe8\x2f\xa2\x6f\x1b\x6d\xc8\xfa\xb3\x85\xf3\x86\xd8\xe4\xdb\xa6\xe0\x90\xbc\xc5\xd4\xaa\xe8\x4e\x64\x04\x87\xfc\x3e\xc1\xdb\xe9\x8d\xc6\x2b\xdc\x18\x84\x99\x19\x8d\x13\xa9\xbf\x60\x89\xcf\x54\x65\x49\x74\xd0\x1c\x26\x43\x59\xdb\x1c\x2d\x4a\xa8\x29\x53\x13\x33\x6e\xa2\xb6\xa1\x64\xae\xea\xed\xf4\xab\x9d\xd1\x98\xd5\x97\xb4\x8a\x47\xd6\x5a\x37\x02\x41\xbe\xb2\x67\x9f\x8e\xa7\xfa\x52\xf7\x78\x4f\x53\xff\xf6\x47\x0f\xbb\xb4\x1c\x53\x05\x5b\xb2\xc3\x18\xd4\x8c\x1e\x42\x29\xe8\x8e\xe6\x98\xa1\xfc\xb1\xe9\x02\xfd\x58\xa8\x81\x0b\x95\x58\x8b\x44\xda\x50\x42\x55\x18\x71\x5c\x38\xb0\x84\x9d\x7c\x3b\x87\x3b\x7a\x71\xc5\x1a\x3a\xe1\x09\xe2\x77\xdf\x52\xfc\xe5\xdb\x6f\xaf\xfe\xfa\xd0\xcb\x95\x1f\xa2\x00\xd6\x37\x02\x14\x1c\x5a\x13\xab\xb4\xa6\xf1\x96\x95\xf1\xd6\x4d\x95\x77\x9b\x6f\x08\x4e\xba\x1d\x7d\x09\x7b\x14\x68\xc7\x18\xc2\xbd\xfb\xcf\xb7\xef\x42\x7a\x06\x6d\x1e\xd9\x3b\xdc\x60\x94\x44\xa5\xcb\x31\x2b\xab\xd7\xc8\x50\xd0\xc4\x2e\x2d\xed\x8a\xff\x4a\x54\xf7\x99\x41\x99\x65\x0f\x6a\x33\x46\x63\xe7\x8e\x1e\xf5\x92\xf0\x8f\xb3\x8e\xde\xc9\xa6\x3f\xd0\x3c\xba\x23\x4a\xff\x71\xd6\x71\xe5\x27\xc9\xfc\x57\xfe\x4c\xf6\xd0\x37\x05\xf6\x07\x9a\x43\xaf\x66\x7a\x6f\xdc\x9c\xba\xea\xf8\x3d\x94\xfc\x9e\xf5\xde\x14\x06\x27\x9d\x41\x22\xec\x4f\x78\xbd\x1c\x36\xe1\xa1\x97\xbb\xb0\x37\xb2\x23\x7b\x05\x32\x75\x6f\xbd\xd8\x46\x77\x09\x9f\x3b\x7c\x7b\x84\xd0\x8c\xa9\xb7\xf2\x73\x7f\xbe\x8f\x66\x2f\x0c\xb7\xad\xe3\xd0\xfa\xd0\xa7\xe4\x83\xb7\xaa\x16\x91\x0d\x78\xc1\x8b\x18\x87\x08\x74\x16\x63\x87\xde\x26\x0c\xa5\x4e\xba\x07\x52\xe4\x70\x6c\x96\x3e\xea\xc2\x57\x4e\xcd\x1b\x5d\xef\x21\x53\xf2\xa8\x53\x19\xcb\x4e\x65\xfe\x85\x29\xb8\xfe\xe1\x9f\x07\xb9\xe3\x1f\x33\x7a\xd2\x33\x28\xe3\xda\x24\x34\x62\x4d\x38\x93\x9a\x41\xcd\xac\x2e\x5c\x5d\xa1\xd6\x63\x54\x9d\x11\x89\xfb\x67\x02\x83\x88\x4b\x58\x7b\x20\xc5\x7d\x65\x78\x1e\xc5\xb0\xed\xbc\xc5\xde\x7b\x95\x5c\x4a\xaa\x4b\x30\xfb\xaf\xd5\x18\xdf\xf7\x16\x60\x35\x4e\xdb\xc5\x3a\xc3\x7f\xff\x55\x6d\xa8\xda\x1e\xb9\x1b\x22\x7b\xaf\x2c\x93\xcd\xf8\x8f\xfd\xd1\xec\xe5\x3d\xa2\xee\x8d\xfe\x54\xea\x1e\xb4\x0e\xfb\xc3\x53\x7e\x01\xd9\xf6\xbf\x2d\xb8\x77\x1a\xec\x81\x92\x68\x34\xe9\xc7\x33\x74\x4c\x62\x99\xeb\xe2\x5a\x13\x11\x6e\xe9\x5b\xa6\x21\x66\xb3\x68\x1b\x83\xf3\x10\x6d\xb8\xe1\x89\x88\x08\xae\x6f\x22\xe2\x65\x74\xb9\xd1\x2b\xcd\xce\x79\x11\xf4\x1e\xe8\xec\x11\xf6\x84\x29\xdb\x05\xa5\x7d\x87\x21\x70\xff\xe9\x0e\x49\xd3\xfa\x68\xa7\x75\xbd\x39\xf8\x25\x81\x05\xdf\x21\x6c\x04\x2f\xbe\xe6\x73\x1f\x0c\x78\xf8\x51\x4b\xa0\xfe\xee\x49\xaf\x38\x86\x90\x3b\xaa\x1d\x3a\xc7\x80\xce\x59\xc6\xb3\xfa\x43\x43\x46\xfc\xbd\xae\xf1\xff\x01\x00\x00\xff\xff\xb6\x5b\xf4\x24\xc0\x3b\x00\x00")

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

	info := bindataFileInfo{name: "api.proto", size: 15296, mode: os.FileMode(420), modTime: time.Unix(1537166040, 0)}
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
	"api.proto": {apiProto, map[string]*bintree{}},
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
