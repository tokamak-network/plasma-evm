// Code generated by go-bindata. DO NOT EDIT.
// sources:
// faucet.html

package main


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataFaucethtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x7a\x6d\x93\xdb\x36\x92\xff\xeb\xf1\xa7\xe8\xf0\x6f\xaf\xa4\xbf\x87" +
	"\xa4\x66\xc6\xf6\xfa\x24\x52\x29\xaf\x37\xd9\xf3\xd5\x5d\x92\x4a\x9c\xba\xdb\xca\xa6\xae\x40\xb2\x25\xc2\x03\x02" +
	"\x0c\x00\x4a\xa3\x4c\xe9\xbb\x5f\x35\x40\x52\xd4\xc3\x4c\xec\xb5\xaf\xea\xfc\x62\x4c\x02\x8d\x46\xa3\xfb\xd7\xe8" +
	"\x07\x2a\xf9\xea\xaf\xdf\xbf\x7d\xff\xf7\x1f\xbe\x81\xd2\x56\x62\xf1\x24\xa1\xff\x40\x30\xb9\x4a\x03\x94\xc1\xe2" +
	"\xc9\x45\x52\x22\x2b\x16\x4f\x2e\x2e\x92\x0a\x2d\x83\xbc\x64\xda\xa0\x4d\x83\xc6\x2e\xc3\xd7\xc1\x7e\xa2\xb4\xb6" +
	"\x0e\xf1\xb7\x86\xaf\xd3\xe0\xbf\xc2\x9f\xdf\x84\x6f\x55\x55\x33\xcb\x33\x81\x01\xe4\x4a\x5a\x94\x36\x0d\xde\x7d" +
	"\x93\x62\xb1\xc2\xc1\x3a\xc9\x2a\x4c\x83\x35\xc7\x4d\xad\xb4\x1d\x90\x6e\x78\x61\xcb\xb4\xc0\x35\xcf\x31\x74\x2f" +
	"\x97\xc0\x25\xb7\x9c\x89\xd0\xe4\x4c\x60\x7a\x15\x2c\x9e\x10\x1f\xcb\xad\xc0\xc5\xfd\x7d\xf4\x1d\xda\x8d\xd2\xb7" +
	"\xbb\xdd\x0c\xde\x34\xb6\x44\x69\x79\xce\x2c\x16\xf0\x2d\x6b\x72\xb4\x49\xec\x29\xdd\x22\xc1\xe5\x2d\x94\x1a\x97" +
	"\x69\x40\xa2\x9b\x59\x1c\xe7\x85\xfc\x60\xa2\x5c\xa8\xa6\x58\x0a\xa6\x31\xca\x55\x15\xb3\x0f\xec\x2e\x16\x3c\x33" +
	"\xb1\xdd\x70\x6b\x51\x87\x99\x52\xd6\x58\xcd\xea\xf8\x26\xba\x89\xfe\x1c\xe7\xc6\xc4\xfd\x58\x54\x71\x19\xe5\xc6" +
	"\x04\xa0\x51\xa4\x81\xb1\x5b\x81\xa6\x44\xb4\x01\xc4\x8b\x7f\x6e\xdf\xa5\x92\x36\x64\x1b\x34\xaa\xc2\xf8\x45\xf4" +
	"\xe7\x68\xea\xb6\x1c\x0e\x3f\xbe\x2b\x6d\x6b\x72\xcd\x6b\x0b\x46\xe7\x1f\xbd\xef\x87\xdf\x1a\xd4\xdb\xf8\x26\xba" +
	"\x8a\xae\xda\x17\xb7\xcf\x07\x13\x2c\x92\xd8\x33\x5c\x7c\x16\xef\x50\x2a\xbb\x8d\xaf\xa3\x17\xd1\x55\x5c\xb3\xfc" +
	"\x96\xad\xb0\xe8\x76\xa2\xa9\xa8\x1b\xfc\x62\xfb\x3e\x64\xc3\x0f\xc7\x26\xfc\x12\x9b\x55\xaa\x42\x69\xa3\x0f\x26" +
	"\xbe\x8e\xae\x5e\x47\xd3\x6e\xe0\x94\xbf\xdb\x80\x8c\x46\x5b\x5d\x44\x6b\xd4\x84\x5c\x11\xe6\x28\x2d\x6a\xb8\xa7" +
	"\xd1\x8b\x8a\xcb\xb0\x44\xbe\x2a\xed\x0c\xae\xa6\xd3\x67\xf3\x73\xa3\xeb\xd2\x0f\x17\xdc\xd4\x82\x6d\x67\xb0\x14" +
	"\x78\xe7\x87\x98\xe0\x2b\x19\x72\x8b\x95\x99\x81\xe7\xec\x26\x76\x6e\xcf\x5a\xab\x95\x46\x63\xda\xcd\x6a\x65\xb8" +
	"\xe5\x4a\xce\x08\x51\xcc\xf2\x35\x9e\xa3\x35\x35\x93\x27\x0b\x58\x66\x94\x68\x2c\x1e\x09\x92\x09\x95\xdf\xfa\x31" +
	"\xe7\xcd\xc3\x43\xe4\x4a\x28\x3d\x83\x4d\xc9\xdb\x65\xe0\x36\x82\x5a\x63\xcb\x1e\x6a\x56\x14\x5c\xae\x66\xf0\xaa" +
	"\x6e\xcf\x03\x15\xd3\x2b\x2e\x67\x30\xdd\x2f\x49\xe2\x4e\x8d\x49\xec\x2f\xae\x27\x17\x49\xa6\x8a\xad\xb3\x61\xc1" +
	"\xd7\x90\x0b\x66\x4c\x1a\x1c\xa9\xd8\x5d\x48\x07\x04\x74\x0f\x31\x2e\xbb\xa9\x83\x39\xad\x36\x01\xb8\x8d\xd2\xc0" +
	"\x0b\x11\x66\xca\x5a\x55\xcd\xe0\x8a\xc4\x6b\x97\x1c\xf1\x13\xa1\x58\x85\x57\xd7\xdd\xe4\x45\x52\x5e\x75\x4c\x2c" +
	"\xde\xd9\xd0\xd9\xa7\xb7\x4c\xb0\x48\x78\xb7\x76\xc9\x60\xc9\xc2\x8c\xd9\x32\x00\xa6\x39\x0b\x4b\x5e\x14\x28\xd3" +
	"\xc0\xea\x06\x09\x47\x7c\x01\xc3\xeb\xef\x81\xdb\xaf\xbc\xea\xe4\x8a\x0b\xbe\x6e\x8f\x35\x78\x3c\x3a\xe1\xc3\x87" +
	"\x78\x0d\xed\x83\x5a\x2e\x0d\xda\x70\x70\xa6\x01\x31\x97\x75\x63\xc3\x95\x56\x4d\xdd\xcf\x5f\x24\x6e\x14\x78\x91" +
	"\x06\x8d\x16\x41\x7b\xfd\xbb\x47\xbb\xad\x5b\x55\x04\xfd\xc1\x95\xae\x42\xb2\x84\x56\x22\x80\x5a\xb0\x1c\x4b\x25" +
	"\x0a\xd4\x69\xf0\x93\xca\x39\x13\x20\xfd\x99\xe1\xe7\x1f\xff\x1d\x5a\x93\x71\xb9\x82\xad\x6a\x34\x7c\x63\x4b\xd4" +
	"\xd8\x54\xc0\x8a\x82\xe0\x1a\x45\xd1\x40\x10\x87\xdd\x53\x51\xc3\xcc\xca\x3d\xd5\x45\x92\x35\xd6\xaa\x9e\x30\xb3" +
	"\x12\x32\x2b\xc3\x02\x97\xac\x11\x16\x0a\xad\xea\x42\x6d\x64\x68\xd5\x6a\x45\x91\xce\x1f\xc2\x2f\x0a\xa0\x60\x96" +
	"\xb5\x53\x69\xd0\xd1\x76\x36\x64\xa6\x56\x75\x53\xb7\x56\xf4\x83\x78\x57\x33\x59\x60\x41\x36\x17\x06\x83\xc5\xdf" +
	"\xf8\x1a\xa1\x42\x7f\x96\x8b\x63\x48\xe4\x4c\xa3\x0d\x87\x4c\x4f\x80\x91\xc4\x5e\x18\x7f\x24\x68\xff\x25\x8d\xe8" +
	"\x38\xf5\x47\xa8\x50\x36\x70\xf0\x16\x6a\xba\x57\x82\xc5\xfd\xbd\x66\x72\x85\xf0\x94\x17\x77\x97\xf0\x94\x55\xaa" +
	"\x91\x16\x66\x29\x44\x6f\xdc\xa3\xd9\xed\x0e\xb8\x03\x24\x82\x2f\x12\xf6\x18\xbc\x41\xc9\x5c\xf0\xfc\x36\x0d\x2c" +
	"\x47\x9d\xde\xdf\x13\xf3\xdd\x6e\x0e\xf7\xf7\x7c\x09\x4f\xa3\x1f\x31\x67\xb5\xcd\x4b\xb6\xdb\xad\x74\xf7\x1c\xe1" +
	"\x1d\xe6\x8d\xc5\xf1\xe4\xfe\x1e\x85\xc1\xdd\xce\x34\x59\xc5\xed\xb8\x5b\x4e\xe3\xb2\xd8\xed\x48\xe6\x56\xce\xdd" +
	"\x0e\x62\x62\x2a\x0b\xbc\x83\xa7\xd1\x0f\xa8\xb9\x2a\x0c\x78\xfa\x24\x66\x8b\x24\x16\x7c\xd1\xae\x3b\x54\x52\xdc" +
	"\x88\x3d\x5e\x62\x02\x4c\x8f\x73\xe7\x36\x4e\xd4\xa1\xa4\x67\xbc\x60\x15\xf6\xd2\xb7\x78\x30\xdc\xe2\x2d\x6e\xd3" +
	"\xe0\xfe\x7e\xb8\xb6\x9d\xcd\x99\x10\x19\x23\xbd\xf8\xa3\xf5\x8b\x7e\x47\xc2\xe9\x9a\x1b\x97\x52\x2d\x3a\x09\xf6" +
	"\x62\x7f\xa4\x5b\x1f\x5d\x5c\x56\xd5\x33\xb8\xb9\x1e\xdc\x5a\xe7\x3c\xfe\xd5\x91\xc7\xdf\x9c\x25\xae\x99\x44\x01" +
	"\xee\x6f\x68\x2a\x26\xba\xe7\xd6\x5b\x06\xce\x77\xbc\x28\xa4\x3b\xba\x17\xad\xbf\xeb\xa7\x73\x50\x6b\xd4\x4b\xa1" +
	"\x36\x33\x60\x8d\x55\x73\xa8\xd8\x5d\x1f\xef\x6e\xa6\xd3\xa1\xdc\x94\x0a\xb2\x4c\xa0\xbb\x5d\x34\xfe\xd6\xa0\xb1" +
	"\xa6\xbf\x4b\xfc\x94\xfb\x4b\x57\x4a\x81\xd2\x60\x71\xa4\x0d\xda\x91\x54\xeb\xa8\x06\xa6\xef\x95\x79\x56\xf6\xa5" +
	"\x52\x7d\x08\x19\x8a\xd1\xb2\x1e\x44\xbb\x60\x91\x58\xbd\xa7\xbb\x48\x6c\xf1\x49\x21\x40\x53\x8a\xf7\x50\x04\xf0" +
	"\x37\x1a\x9d\xbd\x46\xd4\x3e\xbf\x20\xc8\x82\x7b\x4d\x62\x5b\x7c\xc6\xce\x04\xc2\x8c\x19\xfc\x98\xed\x5d\xa4\xdf" +
	"\x6f\xef\x5e\x3f\x77\xff\x12\x99\xb6\x19\x32\xfb\x31\x02\x2c\x1b\x59\x0c\xce\xef\xee\xce\xcf\x15\xa0\x91\x7c\x8d" +
	"\xda\x70\xbb\xfd\x58\x09\xb0\xd8\x8b\xe0\xdf\x0f\x45\x48\x62\xab\x1f\xc7\xda\xf0\xe5\x0b\x39\xf7\x1f\xa5\x24\x37" +
	"\x8b\x7f\x55\x1b\x28\x14\x1a\xb0\x25\x37\x40\xc1\xf5\xeb\x24\x2e\x6f\x7a\x92\x7a\xf1\x9e\x26\xbe\x65\x9a\x15\x6c" +
	"\x0b\xef\xd1\x58\x89\xd6\x2b\x19\x96\x2e\xd5\x00\x6e\x40\x37\xd2\x45\x62\x25\xc1\x96\x78\x98\x9e\xb4\x41\x3b\x82" +
	"\xf7\x8a\x52\xbc\x35\x4a\x0b\x15\x13\x3c\xe7\xaa\x31\xc0\x72\xab\xb4\x81\xa5\x56\x15\xe0\x5d\xc9\x1a\x63\x89\x11" +
	"\x5d\x27\x6c\xcd\xb8\x70\xbe\xe5\x4c\x0c\x4a\x03\xcb\xf3\xa6\x6a\x28\x45\x95\x2b\x40\xa9\x9a\x55\xd9\xca\x62\x15" +
	"\xf8\x40\x25\x94\x5c\xf5\xf2\x98\x9a\x55\xc0\xac\x65\xf9\xad\xb9\x84\xee\x96\x00\xa6\x11\x2c\xc7\x82\x56\xe5\xaa" +
	"\xaa\x94\x84\x1b\x5d\x40\xcd\xb4\xdd\x82\x39\xcc\x35\x58\x9e\xbb\xa8\x17\xc1\x1b\xb9\x55\x12\xa1\x64\x6b\x27\x21" +
	"\xbc\xf7\xe5\x05\xc9\xf5\x2d\xcb\x31\x53\xaa\xa7\x86\x8a\x6d\xbb\xed\x5a\xe9\x37\xdc\x96\xdc\xab\xa7\x46\x5d\xd1" +
	"\xd2\x02\x04\xaf\xb8\x35\x51\x12\xd7\xfb\x1b\x76\x1f\xab\x45\x58\x2a\xcd\x7f\xa7\x44\x47\x0c\xaf\x53\x7b\x74\xd9" +
	"\x74\x77\xa5\x43\x81\xc0\xa5\x9d\xc1\x0b\x7f\x57\x1e\xe3\xba\xad\x88\xce\x81\xba\xe3\xe9\x2a\x4d\x0a\x40\x33\xb8" +
	"\xf1\xe9\xad\x4f\x2c\x0a\x3b\x90\xa0\x38\x82\x9e\xdf\xf4\xf5\xeb\xfa\xae\x97\xa3\xcf\x91\xa7\x3d\x13\x42\xc0\xa1" +
	"\x52\xd6\xbc\x57\xe3\x25\x54\xec\x16\x81\x41\xc2\x8e\x2a\xe6\x56\x68\x57\x6f\x71\xd7\x2f\x88\xed\x06\xd1\x7e\x4d" +
	"\xae\x9c\xfe\xe8\x19\x72\xb9\x7a\x76\x3d\xf5\x88\xa4\x07\x62\xff\xec\x7a\xca\xa5\x55\xcf\xae\xa7\xd3\xbb\xe9\x47" +
	"\xfe\x7b\x76\x3d\x55\xf2\xd9\xf5\xd4\x96\xf8\xec\x7a\xfa\xec\xfa\x66\x88\x65\x3f\xd2\x65\x9a\x44\x85\x86\x76\xeb" +
	"\x20\x1e\x80\x65\x7a\x85\x36\x0d\xfe\x9b\x65\xaa\xb1\xb3\x4c\x30\x79\x1b\x2c\x9c\xb8\x94\x7d\x38\x14\x9c\xcf\x57" +
	"\xa1\x66\x86\x20\x41\x12\x3b\x94\xb4\xbd\x11\x03\x63\xd3\x68\xad\x1a\x49\x51\x12\xe8\xcc\xce\x63\xe5\x88\x50\x46" +
	"\x8a\x99\x44\x49\xa6\xe3\xc5\x5b\x55\x6f\x43\xc7\xc4\x2d\x3f\x51\xa3\x69\xea\x5a\x69\x1b\x0d\xd5\xc9\xa8\x2e\x12" +
	"\x68\xe2\xd7\xd3\x97\xaf\x5f\x3d\x2a\xbe\xa1\xac\xdb\x9d\xa1\x97\x90\x65\x6a\x8d\xe0\x73\xfc\x4c\xdd\x01\x93\x05" +
	"\x2c\xb9\x46\x60\x1b\xb6\xfd\x2a\x89\x0b\x57\x91\x7d\x3e\x6a\x97\xad\x77\xfd\x9f\x82\x6d\xe7\xf2\x97\x50\x37\x99" +
	"\xe0\xa6\x04\x06\x12\x37\x90\x18\xab\x95\x5c\x2d\xdc\x68\x4e\x25\xaa\x7b\x85\x5a\x19\xfb\x98\xf9\xb1\xca\xb0\x28" +
	"\xce\x00\xe0\x4b\xd9\x7f\xb3\xd9\x44\x9d\x26\x9d\xf1\x4b\x14\x75\x4c\xd7\x5f\x23\xb9\xdd\xc6\xde\x8d\x94\x8c\xbf" +
	"\xe6\x45\x7a\xfd\xfa\xfa\xd5\xab\xeb\x17\xff\xf2\xfa\xe5\xcb\xeb\xd7\x2f\x5e\x3e\x84\x0c\x3a\xd4\x67\x02\xc3\xa7" +
	"\xd5\xdf\x29\xaa\x62\xfb\x9c\xda\xe3\xa5\xcb\xe5\x28\x62\x17\x54\x93\xe8\xe0\x9f\xc6\x50\x23\x29\x31\x09\x99\x38" +
	"\x9b\x53\x7c\x02\x8a\x1c\x8c\x1e\x91\xec\x33\xa1\xd5\xc1\x87\x90\xa2\x1a\x4b\x27\xec\x8a\x7b\xae\x64\x0f\xa7\x4b" +
	"\x30\xbc\xaa\xc5\x16\xf2\xbd\xd5\xcf\xe3\xea\x41\xa3\xfc\x21\xac\x0e\xcd\xe6\x41\xe6\xb2\x81\x4a\x15\x48\x51\xdf" +
	"\x34\x26\xc7\xda\x75\x7d\x29\x92\xfe\x65\xfb\x3b\x93\x96\x4b\xec\x22\x6e\x04\xdf\x4b\xb1\x85\xc6\x20\x2c\x95\x86" +
	"\x02\xb3\x66\xb5\x72\x69\x82\x86\x5a\xf3\x35\xb3\xd8\x85\x59\xd3\xa2\xa2\x07\xc5\xa0\xd2\xa1\x14\x48\x0c\x32\x92" +
	"\xbf\xab\x06\x72\x26\xc1\x6a\x96\xdf\x7a\x4f\x69\xb4\x26\x4f\xa9\xd1\x9f\xa6\x0f\xf4\x19\x0a\xb5\x71\x24\xfe\xdc" +
	"\x4b\x8e\xc2\x45\x7d\x83\x08\xa5\xda\x40\xd5\xe4\xce\x21\x29\xaa\xbb\x43\x6c\x18\xb7\xd0\x48\xcb\x85\xd7\xa7\x6d" +
	"\xb4\xa4\x1c\x01\x0f\xa2\xf4\x49\x2d\x98\x60\xb5\x78\x5f\xe2\x99\x94\xa8\xaf\xe2\x40\xe3\x5b\x4f\x0e\xb5\x56\x16" +
	"\x73\x32\x28\xb0\x15\xe3\xd2\x90\x45\x5c\x1e\x80\xd5\x47\x54\x79\xfd\x53\xfb\xb0\xef\x58\xba\xe9\x38\x86\xbf\x09" +
	"\x95\x31\x01\x6b\x42\x7a\x26\x28\xbd\x53\x50\x2a\x3a\xfa\x40\x5b\xc6\x32\xdb\x18\x50\x4b\x37\xea\x25\xa7\xf5\x6b" +
	"\xa6\xc9\x82\x58\xd5\x16\xd2\xb6\xdf\x46\x63\x06\xf5\xba\xed\x22\xd2\x2b\x55\xf2\x07\xf3\xbd\xd6\x53\xf8\xe5\xd7" +
	"\xf9\x93\x56\x94\xbf\xe2\xd2\x41\x82\xf0\xed\x8f\x6c\x4b\x66\x21\xd7\xc8\x2c\x1a\xc8\x85\x32\x8d\xf6\x12\x16\x5a" +
	"\xd5\x40\x52\x76\x9c\x3a\xce\x34\x51\xbb\xdd\x3a\x26\xe3\x92\x99\x72\xd2\xb6\x0b\x35\x3a\x2b\xf5\x73\xdd\xf8\x05" +
	"\xa1\x6e\x4c\x0c\x78\x3a\x9d\x03\x4f\x3a\xbe\x91\x40\xb9\xb2\xe5\x1c\xf8\xf3\xe7\x3d\xf1\x05\x5f\xc2\xb8\xa3\xf8" +
	"\x85\xff\x1a\xd9\xbb\x88\x76\x81\x34\x85\xe1\x6e\x6e\xc3\x96\x8f\xa9\x05\xcf\x71\xcc\x2f\xe1\x6a\x32\xef\x66\x33" +
	"\x8d\xec\xb6\x7b\x6b\xed\xe8\xff\x73\x7f\x77\xf3\x43\xcd\x38\xe5\x1f\xe8\xc6\xf7\x02\x0c\x30\x58\x71\x63\xa1\xd1" +
	"\x02\x5a\x1f\xf6\x26\xe8\x0d\xe2\xe8\x86\x5a\x39\xc1\x65\xfb\xd0\x62\xaa\x3b\x82\x67\x13\x19\x94\xc5\xf8\xdf\x7e" +
	"\xfa\xfe\xbb\xc8\x58\xcd\xe5\x8a\x2f\xb7\xe3\xfb\x46\x8b\x19\x3c\x1d\x07\xff\xaf\xd1\x22\x98\xfc\x32\xfd\x35\x5a" +
	"\x33\xd1\xe0\xa5\xb3\xf7\xcc\xfd\x3d\xd9\xe5\x12\xda\xc7\x19\x1c\x6e\xb8\x9b\x4c\xe6\xe7\xfb\x26\x83\x36\x8f\x46" +
	"\x83\x76\x4c\x84\x3d\xf0\x8f\x75\xc4\xa0\x42\x5b\x2a\xe7\xba\x1a\x73\x25\x25\xe6\x16\x9a\x5a\xc9\x56\x25\x20\x94" +
	"\x31\x7b\x20\x76\x14\xe9\x29\x28\x5a\xfa\xd4\x05\xeb\xff\xc4\xec\x27\x95\xdf\xa2\x1d\x8f\xc7\x1b\x2e\x0b\xb5\x89" +
	"\x84\xf2\x57\x6d\x44\x4e\xaa\x72\x25\x20\x4d\x53\x68\xa3\x68\x30\x81\xaf\x21\xd8\x18\x8a\xa7\x01\xcc\xe8\x91\x9e" +
	"\x26\xf0\x1c\x8e\x97\x97\x14\xef\x9f\x43\x10\xb3\x9a\x07\x13\xef\x0e\x9d\xe2\x95\xac\xd0\x18\xb6\xc2\xa1\x80\xae" +
	"\x32\xea\x41\x46\xe7\xa8\xcc\x0a\x52\x70\x06\xaa\x99\x36\xe8\x49\x22\xaa\xce\x3b\xb4\x11\x66\x1d\x59\x9a\x82\x6c" +
	"\x84\xd8\x83\xd4\x3b\xc5\xbc\x83\xdf\x01\x79\xe4\x63\xcd\x57\x69\x0a\x54\xaa\x92\x8a\x8b\xfd\x4a\x32\xbe\x2f\xaa" +
	"\x27\x11\xc5\x85\xfd\x8a\xc9\x7c\x88\xe6\x03\x6e\x58\xfc\x11\x3b\x2c\x8e\xf9\x61\xf1\x00\x43\xd7\xc3\x78\x8c\x9f" +
	"\xef\x79\x0c\xd8\xb9\x81\x07\xb8\xc9\xa6\xca\x50\x3f\xc6\xce\xf7\x30\x5a\x76\x4e\xd5\xef\xa4\x1d\xac\xbd\x84\xab" +
	"\x57\x93\x07\xb8\xa3\xd6\xea\x41\xe6\x52\xd9\xed\xf8\x5e\xb0\x2d\xe5\x4c\x30\xb2\xaa\x7e\xeb\x5a\x0e\xa3\x4b\x17" +
	"\x71\x67\xd0\x73\xb8\x74\xcd\xe4\x19\x8c\xdc\x1b\xcd\xf3\x0a\xdd\xaa\x97\xd3\xe9\xf4\x12\xba\xaf\x30\x7f\x61\xe4" +
	"\x84\xba\xc1\xdd\x03\xf2\x98\x26\xcf\x29\xee\x7f\x8e\x44\x2d\x8f\x5e\xa6\xf6\xfd\x33\xa4\xea\x63\xc3\x81\x58\xf0" +
	"\xa7\x3f\xc1\xc9\xec\x21\x8c\xe3\x18\xfe\x83\x51\x19\x2e\x84\xeb\x1e\xb8\xa6\x41\x4f\x5f\x71\x63\x5c\x31\x6e\xa0" +
	"\x50\x12\xdb\x35\x9f\x76\xed\x9f\xc8\xd8\x92\xc1\x02\xa6\xc7\x02\xd2\x75\x38\x08\x0b\x67\xa2\xc5\x80\xef\x61\x20" +
	"\xb8\xd8\x0d\xf7\x3b\x58\xc9\x2b\x84\xaf\x52\x08\x82\xe1\xe2\x13\x0a\x22\xe8\x99\x5d\x18\xb4\xef\xbd\x2d\xc6\x6d" +
	"\x74\x3c\x17\xbb\x26\x97\x70\x33\x9d\x4e\x27\x27\x42\xec\xf6\xea\x7d\x53\x53\xda\x04\x4c\x6e\xdd\x95\xd8\xeb\xd6" +
	"\x25\x8e\x94\x02\xd1\x95\x26\x20\x57\x42\xf8\x9c\xa5\x5d\x4a\x0a\x6e\x9b\x27\x29\x84\x57\xf3\x33\x51\x74\xa0\xc9" +
	"\xc1\xd1\x8e\xcd\x73\x46\xf7\xc7\x26\x3a\xd4\xd9\x11\x71\x78\x75\x60\x94\x03\x7b\x9d\x37\xcc\x45\x2f\x37\xdf\x6b" +
	"\xf4\xc8\x5c\x7b\x7b\x1d\xeb\x6c\x20\xbf\xe7\xf3\xfc\xea\x23\x8f\xd1\x4f\xd7\x8d\x29\xc7\x47\x82\x4e\xe6\xa7\xb6" +
	"\x79\x67\x51\x53\x96\xac\x28\x64\x91\x2d\xa8\x14\xd0\x78\x62\x12\x97\xaa\x6b\x0c\x35\xca\x02\x75\x97\x52\xf8\xcc" +
	"\x9e\x12\xc0\x03\x93\xf9\xaa\x72\x08\xa7\x4f\x74\x18\x97\x92\x29\x89\x00\x00\x47\x4e\xe0\x80\x7a\x80\x54\x22\x46" +
	"\xc1\x6a\x83\x05\xa4\xe0\x3f\x8a\x8f\x27\x51\x23\xf9\xdd\x78\x12\xb6\xef\xc7\x3c\xba\xf9\x79\x5f\x26\x76\x62\x3f" +
	"\x4f\x21\x48\xac\x06\x5e\xa4\xa3\x00\x9e\x9f\x73\x41\x8a\xba\xa3\xc5\x5e\x82\xe1\x52\x80\xc4\x16\x0b\xd7\x17\xf5" +
	"\xf5\xda\x3f\x82\x8c\xe5\xb7\x2b\x57\x08\xcd\x28\xd5\x1a\x9f\xb0\x65\x6b\x66\x99\x76\x5c\x27\x73\xd8\x93\xb7\x85" +
	"\x62\x4e\xc6\x99\x83\xaf\x48\x5d\xfb\x15\xfa\x4f\x16\xee\x2d\x53\xba\x40\x1d\x6a\x56\xf0\xc6\xcc\xe0\x45\x7d\x37" +
	"\xff\x47\xf7\x49\xc7\x35\x89\x1f\x15\xb5\xd6\xb8\x38\x91\xa8\xed\x32\x3e\x87\x20\x89\x89\xe0\x8f\xd8\xf4\x87\x1d" +
	"\x7e\x8c\x87\x33\xad\x70\xe8\x3f\x95\xb7\xe3\x15\x2f\x0a\x81\x24\xf0\x9e\x3d\x39\x23\xd9\x7f\xe8\x52\x87\x5b\x42" +
	"\xdb\x03\xdf\xaf\xd9\x01\x0a\x83\x8f\x2c\xe8\xdb\xe9\x23\x02\x40\x48\x47\xe6\x4e\xe7\x6d\xb1\xed\x86\xf5\xc8\xe9" +
	"\xa2\xfd\x69\x45\xd1\x68\x97\x6b\x8d\xc3\x16\x60\x97\x30\x32\x94\xfb\x15\x66\x34\x89\xca\xa6\x62\x92\xff\x8e\x63" +
	"\x8a\x4b\x13\xaf\x2b\xd7\x9f\x0f\x4e\xaf\xe4\x13\x61\xf6\x8d\xf3\x51\x17\xe3\x46\xad\x12\x47\x9d\x75\x5f\xec\x6b" +
	"\xfb\x19\x4c\xe7\xa3\x4f\xd4\xd0\xf9\x5d\xc2\x8c\x69\x18\xbe\x84\x5d\xf0\x05\xad\x68\xf7\x6e\x2e\x63\x7a\xe4\x3b" +
	"\x19\x2e\x3f\x97\x6a\x93\x8e\x6e\xa6\xbd\x90\xde\xd0\xce\xce\xa3\x16\x6b\x27\xc6\x20\x29\x3b\xd7\x5c\xc0\xcd\xf4" +
	"\x4b\x48\xeb\xbb\x21\x47\x27\xb0\x9a\xd7\x58\x00\xcb\x2d\x5f\xe3\xff\xc2\x41\xbe\x80\x92\x3f\x59\x44\xc2\x61\xa7" +
	"\x3c\x07\xd3\x03\x79\x69\xb6\xd7\xed\xff\x27\x7f\x83\xd8\x69\xf8\x39\x04\x67\x0f\xf2\x20\x12\x8f\x08\x8f\x5c\xfb" +
	"\x61\xbf\x77\x1f\x9c\x82\xe3\x98\x42\xd9\x6e\xff\xb1\x74\x12\x95\xb6\x12\xe3\x20\xb1\xee\x47\x33\x24\x73\xcf\xc1" +
	"\x31\xf0\xc3\x87\x29\xdd\xee\xb0\x90\xa1\xfa\x1d\x8f\xea\x2c\x18\x24\x27\x7d\x2d\xd6\x65\x22\xb0\xdb\xff\xb6\x28" +
	"\x8e\xe1\x27\xcb\xb4\x05\x06\x3f\xbf\x83\xa6\x2e\x98\xf5\x9f\x72\x28\x3e\xfa\x4f\x25\xdd\x8f\x8f\x32\xa6\x0d\x2c" +
	"\x95\xde\x30\x5d\xb4\xfd\x19\x5b\xe2\xd6\x7d\xca\xe9\x52\x3f\x83\xf6\x1d\xdd\x62\x6b\x26\xc6\x27\x75\xdf\xd3\xf1" +
	"\x28\x1a\x9a\x7c\x34\x89\x90\xe5\xe5\x29\xa1\x8b\x58\xfd\xbe\x29\x7c\xe7\x4a\x80\xf1\xd3\xb1\x2d\xb9\x99\x44\xcc" +
	"\x5a\x3d\x1e\x1d\x80\x61\x34\x21\xbb\x5e\x0d\x4a\xb2\x7e\x79\x72\xe0\x56\x8f\xf1\xd8\x27\xd3\x7d\x22\xd0\x91\xe7" +
	"\xc6\x8c\x3d\xae\x46\x97\x03\xde\x87\xb0\x1a\x3d\x1b\xf5\x86\xda\xbb\xf7\xfe\x1c\xe9\x59\x49\x0e\x58\x8f\xc8\xcb" +
	"\x46\x27\xdb\xb3\xa2\x78\x4b\xfe\x33\x0e\xce\x78\xfa\x31\x3a\x26\xbd\xb2\xfd\x7d\xfd\xa8\x96\xfd\xcf\x34\x1e\x50" +
	"\x31\x2f\x46\x93\xc8\x34\x99\xef\x4d\x8c\x5f\xf6\x05\x58\x47\xe6\xc0\x7b\x1c\x0a\x4e\x12\x0a\xda\xe2\x30\xa9\x08" +
	"\x8f\x92\x90\x47\xa2\x46\xbb\xa5\x3f\xd5\xee\x92\x14\x3e\x9d\xf4\xad\xad\x6f\x0c\x25\x57\xbe\xf5\xbf\xc1\xcc\xb8" +
	"\x4e\x02\xb4\x78\x77\xdd\x1c\xdf\xb5\x79\xf3\xc3\xbb\x41\xe7\xa6\xf7\x88\xb1\xe3\xde\xff\x2e\xf0\x5c\x9f\xe4\xec" +
	"\x0f\x11\x37\x9b\x4d\xb4\x52\x6a\x25\xfc\x4f\x10\xfb\x46\x4a\xcc\x6a\x1e\x7d\x30\x01\x30\xb3\x95\x39\x14\xb8\x44" +
	"\xbd\x18\xb0\x6f\xbb\x2b\x49\xec\x7f\x22\x97\xc4\xfe\x57\xc0\xff\x13\x00\x00\xff\xff\x6f\xe9\x5d\x53\x16\x2c\x00" +
	"\x00")

func bindataFaucethtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataFaucethtml,
		"faucet.html",
	)
}



func bindataFaucethtml() (*asset, error) {
	bytes, err := bindataFaucethtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "faucet.html",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"faucet.html": bindataFaucethtml,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"faucet.html": {Func: bindataFaucethtml, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
