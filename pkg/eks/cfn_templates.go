// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/1.10.0/2018-05-09/amazon-eks-nodegroup.yaml
// assets/1.10.0/2018-05-09/amazon-eks-service-role.yaml
// assets/1.10.0/2018-05-09/amazon-eks-vpc-sample.yaml

package eks

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _amazonEksNodegroupYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x6d\x73\xdb\x36\xf2\x7f\xaf\x4f\xb1\xd5\xdf\x33\xfa\x5f\xc7\x14\x25\xd9\xe9\x03\xe7\xd2\x19\x56\x56\x52\x5e\x6c\x47\x23\x29\xe9\xf4\xda\x8e\x07\x02\x57\x12\x6a\x12\x60\x01\xd0\x8a\x92\xe6\x3e\xfb\x0d\x41\x52\xe2\xa3\x25\xa7\xb9\x69\xf3\x46\x06\xb0\xd8\x87\xdf\x2e\x76\x17\x60\x2c\xcb\xea\xb8\x3f\xce\x17\x18\x46\x01\xd1\xf8\x42\xc8\x90\xe8\xb7\x28\x15\x13\xdc\x81\xde\x68\x30\x1c\x58\x83\x6f\xad\xc1\xb7\xbd\xce\x15\x2a\x2a\x59\xa4\xd3\x15\x37\x24\xef\x05\x87\xc9\xab\x39\x58\x70\x2b\x7c\x84\x97\x52\xc4\x51\xaf\xd3\x99\x12\x49\x42\xd4\x28\x95\xd3\xe9\x00\xbc\xc2\xdd\x2d\x09\xd1\xe9\x00\x00\x94\x78\x2c\x36\x08\x93\xf1\x28\xa1\x80\x29\x61\x12\xb4\x00\x12\x04\x62\x0b\xf3\xf9\x0f\x40\x28\x45\xa5\x92\x39\xbd\x41\x60\x5c\x69\xc2\x29\x2a\xc3\x66\xb1\x8b\xd0\x01\xf7\xc7\xb9\xe3\x4c\xc6\x23\xc7\x79\x85\xbb\x84\x81\xf9\x23\x11\x96\xc8\x4d\x74\xf2\x42\xb2\x46\xcf\x77\x1a\x37\x99\x45\xc7\xf1\xfc\xba\x66\xee\x8d\x07\xcc\x87\x95\x90\x46\x38\x4f\xcc\xdb\x6b\xd0\xdf\x73\xcf\x66\x0c\xe3\x3a\x93\xc4\xb4\x7c\x13\xe8\x5d\x84\xad\xfc\x0e\xda\xcd\xb5\x64\x7c\x9d\xf1\x5a\x91\x38\xd0\x0e\x84\x97\xfd\x80\xc8\x35\x9a\x59\x37\xc1\x07\xfd\xb7\x24\x88\x51\xa5\x42\x2d\xd0\xa3\x7e\x88\x3e\x8b\xc3\xc3\xf8\xb0\xc3\x82\xf0\xa2\xbc\x1c\x5e\x54\x97\xdf\x55\xc6\xa3\xf2\xc4\x65\xbf\x32\xac\x2e\xd7\xe8\x2f\xab\x13\xc3\x41\x79\xe6\x59\xbf\x32\xac\x2e\x8f\xaa\x13\x97\xd5\x89\x61\x8d\x64\x54\xa6\xa1\x65\x3b\x69\xc5\x4e\x5a\xb5\x93\x5e\xf4\x6b\x0c\xbe\x29\x4f\x94\x91\xa0\x15\x24\x68\x15\x09\x5a\x45\x82\x5e\x96\x38\x8e\x05\x57\x5a\x12\xc6\x75\x29\x72\xc2\x58\x69\x58\x22\x10\x78\x20\x01\xf3\xeb\x91\xb4\x0f\x41\x37\xd6\x62\x4e\x49\xc0\xf8\xda\x1c\xbf\x1b\xc6\xe7\xec\x3d\x16\x03\xfe\x36\x0e\x97\x28\xeb\xe1\x79\xc3\x38\x0b\xe3\x10\x14\x7b\x8f\x20\x56\x85\x33\x0c\xee\xfc\x65\xbf\x1c\x83\xc3\x56\x81\xe4\xdd\xc9\x02\xc9\xbb\xd3\x05\x5e\x24\x02\xc7\x41\xac\x34\xca\x47\xd2\x07\x4d\x29\x80\x93\x10\x21\x92\xe2\x81\xf9\xe8\xc3\x76\x83\xdc\x1c\xb4\x7c\x79\x4b\x14\x50\x89\x44\xa3\xdf\x07\xf0\x56\xc0\x34\x30\x05\x8c\x53\x21\x25\x52\x7d\x6e\x8e\xa4\x82\x2d\x0b\x02\xe0\x22\x05\x7f\x19\x60\x92\x7d\x7e\x13\xac\xc4\xac\xe1\xb8\x66\xd8\x18\x5b\x5a\x94\x7d\xc3\xd9\xef\x31\x02\xf3\x91\x6b\xb6\x62\x28\xf7\xc9\xe0\x00\x43\x33\xe7\x0c\x84\xb1\xe0\x5a\x8a\x60\x1a\x10\x8e\x73\xa4\xb1\x64\x7a\x67\x76\xb5\x20\xa3\x32\x1a\x58\x1b\x84\xc5\xaa\x04\x08\x4d\xb9\x41\x94\xb0\xeb\x37\xa6\xc7\xb2\x90\x24\x4d\x76\x00\xde\x46\x34\x4f\xa7\x35\x89\x6f\xa7\xe3\x5c\xcc\x56\xc8\x7b\x94\x47\x92\xf6\xdb\xe9\x38\x67\x3b\x8f\x97\x1c\xb5\x6a\x33\x25\x5d\x4d\xdc\x2a\x73\xde\x0a\x28\xe1\x89\x9b\x72\xb7\x16\x44\x5c\x33\xa5\xff\x59\x30\xc4\x6c\x4f\x44\x7d\xd7\xe9\xdc\xa0\x26\x3e\xd1\x24\x11\x65\x48\xc6\x81\x88\xfd\xb4\xec\x25\x02\x1d\x8f\x6b\x94\x2b\x42\x33\x27\xee\x8b\x99\x81\x21\xd3\x10\xc0\xca\x7e\x01\xae\xc9\x12\x03\x67\x3f\x04\xf0\xf3\x10\xee\x26\xb5\x31\xf3\x5e\x77\x4f\x50\xac\x8e\x87\x4d\x56\x31\xd6\x9b\xe6\x5b\xdd\xff\x14\x85\x7e\x4c\xdd\x62\x22\x6e\x2c\xf8\x8a\xad\x63\x69\xcc\x3e\xae\x5e\x29\xbe\x6b\x2b\x2d\x69\xe8\x38\x5d\x9a\x3d\x6a\x74\xc5\xca\x5a\x5f\x4c\x8b\x7a\x69\x3e\xaf\xfb\x9f\x00\x06\xea\x24\xa2\x9e\x8a\x87\x39\x08\xa5\x99\x2c\x86\x3b\x9d\x19\x2a\x11\x4b\x8a\x69\xfb\x53\x34\x67\x2a\xc5\x8a\x05\x58\x6b\x47\x3c\xf7\x26\x09\xbc\x12\x51\x1a\x7d\x52\x44\x28\x35\xc3\xbd\xf4\x29\xd1\x1b\x07\xba\x76\xae\xe1\x4c\x04\x87\x45\x0b\xbe\x98\xe1\xaa\x24\x32\x59\xaf\xaa\x91\xcc\x35\xeb\x60\xa8\x9b\x05\xbb\x4a\xc5\xa1\xd9\x3a\x15\x01\xa3\xbb\x2b\x41\xe3\x10\xb9\x3e\xc0\x52\xea\x1a\x47\xd6\x70\x60\x0d\xbf\xee\xed\x57\xe7\x9a\x68\x2c\x6f\xb0\x60\xb2\x5a\x21\xd5\x4e\xda\xd5\x14\xd0\x9c\x4a\xc6\x29\x8b\x48\xc9\x75\x00\x73\x94\x0f\x2c\x3f\x99\x07\x2e\x48\x47\x7d\x62\xda\x51\xb2\x55\x7d\x2a\xc2\xc2\xba\x4b\xcd\xb1\x2e\x79\x4a\x69\xe5\x1c\xcc\x69\x01\xd6\x58\xc9\x8a\xd8\xa6\x76\x9b\xfc\x6e\xca\xc5\x21\x46\x5a\x00\x39\x06\x49\x23\x28\x8f\xc1\xd2\x64\x4e\x06\x80\x93\x66\xcd\x25\x7e\x59\x5b\x93\xce\x4b\xd4\x6e\xac\x37\x42\xb2\xf7\x26\xbc\x17\xe2\x1e\x79\x03\xdd\xf7\x44\xd3\xcd\x78\x83\xf4\xfe\x9a\xec\x50\xba\x0f\x84\x05\x64\xc9\x02\xa6\x77\xcd\x5c\xaf\xc4\x96\x07\x82\xf8\x6f\x64\xf0\x42\x48\xb3\xa9\x99\x70\x86\x91\x50\x4c\x0b\xb9\x4b\xc1\x6a\xa0\xca\x0d\xd8\x93\xb2\xac\x6e\x94\xc9\x92\xe4\x6e\x32\x40\xd3\xa2\x31\xe0\x25\xa6\x04\xa5\xf5\xfc\x50\x3a\xd0\xfd\xb2\xdb\xe8\x52\xca\xd9\xdf\xd3\xa3\x59\x86\xda\x17\xa6\xba\xe1\x07\x5a\xaf\x54\x71\xcb\x34\x63\x53\x29\xab\xdc\x1a\x08\x5d\xad\x09\xdd\x9c\x40\x78\x85\x01\x9e\xc4\xf1\x0a\x4f\xe4\x78\x23\x7c\xb6\xda\x55\x09\x5d\xad\x25\x5b\xc6\xba\x51\x59\xa5\xd8\x9a\x4f\x25\x7b\x20\x1a\xbd\xc8\xf5\x7d\x89\x4a\xd5\x00\xd0\x64\xed\x2c\xc8\x7a\x9f\x9b\x9f\x18\x1c\x78\xaf\xfe\x46\xc1\x71\xaf\xf6\x0e\xcf\xda\x83\x47\xcc\xc9\x92\x7f\x43\xcf\xf8\x68\xc7\xd7\x52\x06\xcc\x5a\xa9\x43\x9b\x97\x1b\xcd\xa4\xad\x25\xa6\x8b\x4e\x1a\xea\x72\xeb\x9c\xf1\x28\x34\x91\xc9\x3f\x53\xb3\x8a\xf5\x74\x41\xd6\x85\xb4\xfb\x0a\x77\x0e\x7c\x31\x8f\x97\xd0\xbd\x8f\x97\x28\x39\x6a\x54\x7d\x26\xec\x8c\xa9\x7d\xf6\xa1\xd0\x3b\x7d\x3c\x94\x6e\x73\x47\x76\xa0\x27\xb6\x1c\xfd\x5e\x23\x0e\x1e\x5f\x27\xd1\x72\x02\x1c\x19\x65\xd6\xa0\x46\xc8\x7d\xf5\x9a\x3b\x75\x8e\x2d\xb0\x95\x5f\x17\xcc\x1b\x87\x79\x04\xd0\x02\xa8\x08\xc3\x98\x33\x4a\x34\xc2\x96\xe9\x0d\x20\xa1\x1b\x10\x7a\xb3\xc7\x2b\x55\xc0\x77\x0e\xd5\xbd\xa9\x03\x9c\x1b\x9f\x97\x95\x3e\xb6\xc7\x8b\xa6\x52\x68\x41\x45\xe0\x40\xcf\x1a\xe6\xb1\xfa\x42\x8a\x70\x2a\xa4\x76\x60\x90\xbb\x44\xa4\xe3\xaf\x9e\x3d\xbb\x78\xd6\x08\x65\xb2\xa7\xd8\xa4\xfe\x95\xd0\x66\xd7\x8f\x57\xf1\x32\x49\x50\x0a\x08\xf7\x21\x12\xbe\x79\x4c\x92\x48\x91\x3d\x60\x01\x76\x26\x38\xac\xa4\x08\xdb\x2f\x48\x9f\xcd\x0f\x27\x36\xf3\x45\xb7\x68\x1a\xd5\xbc\x32\x1c\x8c\x9e\xb5\x39\xa6\xc8\x7b\x62\x80\x5d\x88\x4f\x39\xff\x93\xcf\xed\x94\x56\x74\x1b\x0f\x41\xd9\x85\x7b\x0f\x36\x7a\xe2\x44\x54\xaf\x50\x69\xc6\x8d\xbf\x3f\xfd\x94\x3c\xd9\x1d\x47\x74\xfb\x2b\x8f\x49\x7e\x24\x6a\xe0\x17\x5d\xe5\x4e\x3d\xd3\x77\xb7\x64\xa3\x13\xb1\xff\xb3\xc9\xe9\x00\x7b\x8e\xf0\xe5\xe5\x45\xcd\x11\xc9\x5c\xf1\x2d\xa6\x86\x6a\xe1\xfe\x59\x1a\x1c\x01\x8e\x49\xf4\xc7\x24\x22\x94\xe9\x5d\x41\xe9\xc7\x6f\xb3\xd7\x24\xe6\x74\x53\xba\x56\xa6\xbd\xc4\x7e\x7f\x91\x22\xdb\x94\xbf\xdf\x3d\x22\xa4\x74\xb5\xce\x9f\xdf\x4e\x55\xea\xed\x74\xfc\x6f\xc1\xd1\xdb\x3f\x43\x55\xca\x70\x7e\x89\x6d\x2d\xc4\xa5\xeb\x7f\x56\x61\xd3\xe2\x5c\x2e\xc3\xd6\xd9\x87\xd2\x93\xc1\x47\x2b\x19\x16\xee\xd5\x52\x44\x64\x4d\x34\xba\x3a\x85\xc1\x81\x9e\x96\x31\xf6\x1a\xea\x7e\xef\x94\xba\xdf\x6b\xab\xfb\xa7\x49\x7c\x13\xf9\x44\x67\x57\xdb\xfd\x95\xf7\x00\xe5\x4c\x04\xc9\x4f\x4a\x75\xc0\xec\x86\xf1\x7d\xc7\xed\xf1\xfc\x7a\x0a\xbd\xe1\x41\xec\x0d\x79\x67\x6e\x25\xa9\x9b\x92\x95\x2c\x42\x8b\xde\x7f\x3c\x50\x1b\x22\xa9\xfd\x92\x2e\x28\x4b\x0c\x89\x97\x01\xa3\xfb\x76\xb8\x02\xae\x47\xc2\xea\x93\x44\xfd\xe5\xa0\xf8\x0e\x01\x90\x7f\x49\x29\xd0\x95\x9e\x61\x4a\xdf\x42\xea\xcc\x0a\x0f\x39\xf9\xf7\xa0\x94\xa8\xfc\x70\x53\x4a\x03\x4d\xcf\x1a\x4d\x79\xe2\x8d\x42\x79\x95\x3d\xe9\x65\x19\x81\x3b\xce\xf7\x44\xe1\x57\x97\xc5\xd6\x39\x99\xfd\x97\x60\xdc\x81\x9f\x4b\xed\x72\xb7\x7b\x5e\x1a\x97\x57\x01\xba\xff\xf7\x85\xbd\x64\xdc\x5e\x12\xb5\x01\xeb\x1d\xfe\xc2\x2b\x1b\x00\xba\x63\xf7\x6e\x3c\x99\x2d\xbc\x17\xde\xd8\x5d\x4c\xee\xae\xbc\xd9\x64\xbc\x78\x3d\xfb\xe9\xb9\x8d\x9a\xda\x87\xf8\xb5\xa3\x7b\xd6\x3d\x87\xee\x09\x3c\x5e\x78\xd7\x93\xbb\xa9\xbb\xf8\xe1\xf9\x59\x1b\x77\x9b\x92\x3e\x95\xba\x8d\xe1\xcd\xeb\xab\xc9\xf5\x81\x3a\x65\xf6\x1f\xbb\x4f\xb6\xca\xc6\x7b\xf5\xf8\xb6\x82\xf8\x26\x3e\x09\x03\x6b\x34\x18\x7e\x6d\x0d\x87\xd6\x60\xd8\xe7\x42\x86\x24\xe8\xff\xa6\x04\x6f\xe3\x1b\xde\xfb\x4c\x82\x15\x41\xab\x3d\xc7\x77\x36\xa9\xd2\xb6\x8b\xc6\x32\x00\x4b\xe4\x9b\xf6\x06\xc1\x46\xeb\x48\x39\xb6\xad\x2e\xac\x58\x59\x5b\x54\xda\xaa\x3c\x20\xd9\xe9\xc8\xc2\x7b\x65\x8f\x06\xc3\x6f\xac\xc1\xa5\x35\xb8\xfc\x04\x9b\xc9\x56\x25\xcd\x8e\x39\xb8\x08\xc4\xf7\xad\x50\xf8\x18\x80\x65\xa9\x34\x59\x64\x63\x73\x06\x6d\xbb\xa6\xea\x81\xce\x7c\xe1\x78\xc4\x6b\x89\x24\xbc\x57\xe0\x67\x17\x43\x2b\x2f\xe0\x96\x25\x71\xcd\x04\x7f\xde\x3d\x87\x0f\x30\xc3\x95\x03\x5d\x93\x63\x66\x66\xba\x0b\x1f\xcf\xbb\x60\xe5\xe4\x46\x4e\x81\xb4\x90\x62\x33\xc2\xdf\x63\x94\x3b\xe8\xe5\x5f\x43\x3e\xd0\x24\x05\xad\x4c\xf7\x90\x3d\x32\xe9\x9d\x39\x8e\xd0\xb4\xd2\xf7\x89\x26\xe7\x10\x92\x64\xf3\x84\xfb\x91\x60\x5c\x3b\x95\xf1\xc7\x1e\x7c\x07\xb6\x0e\x23\x3b\xb7\xe6\x2e\x13\x77\x27\x51\xc5\x81\x7e\x14\x73\x4a\xf4\xf1\xcd\xf0\x07\xac\x25\x46\x8d\x3a\x26\xda\xc3\x1f\x40\xb6\xf7\xd0\xfb\x10\x49\xc6\x35\x9c\x8d\x3e\xf6\xe0\x0f\x50\xe8\x43\x4f\xd9\x3f\x9f\xff\xd2\xfd\xd5\xb6\xd7\xc9\xd4\xd2\xe4\x19\xb0\x7c\xf8\x0e\x6a\x91\xbd\x77\x64\xeb\x59\x73\xe7\x8b\xc9\xec\x6e\x72\x7b\x35\x7d\xed\xdd\x2e\x9e\x9f\xfd\xff\x93\xb4\x2f\xc3\x76\x9a\xce\xff\x68\xd3\xc5\xbb\x5d\x4c\x66\xb7\xee\xf5\x9d\x37\x4d\xf4\x30\xa7\x47\x99\xc3\xe2\xd8\xf6\xf0\xab\x6f\xfb\xa3\x67\x97\xfd\xec\xd7\x0e\x88\x46\xa5\xed\x10\x35\xb1\x12\x97\xda\x81\xa0\x24\xb0\x58\xf4\x70\xd9\x2a\x20\xd1\xc4\x62\xa0\xce\x2b\x56\x9f\x9f\x55\x27\xd6\x60\x3f\x10\x69\x07\x6c\x69\x72\x67\x80\xda\xfc\xa6\x27\xe9\x28\xfb\xf1\xf5\x1b\xc3\xee\xd6\xbd\x99\x9c\xb7\x85\x32\x74\xff\xa4\x94\x53\x8c\x48\x92\xbf\xda\x29\x8d\xa1\x9f\xfd\xe6\x92\xfa\xd9\xb9\xfe\x5f\x8a\xb1\x22\x29\xde\xed\x4e\x96\x54\xf0\xff\xf9\x59\x71\xf0\x79\x0c\xf9\x04\xf6\x4f\x34\xa0\x78\xf2\xdc\x37\x8b\x1f\x5e\xcf\xbc\xc5\x4f\xe6\x0c\x9e\xb7\x1e\xcc\x23\x41\x00\xc7\x63\xcd\x9b\xdc\x2e\xee\xc6\xee\x71\x39\x27\xa1\xd8\x2a\xd0\x90\x53\x1d\x80\x4f\x30\x14\xdc\x92\x18\x08\xe2\xb7\x62\xb2\x27\x97\xa8\x34\x91\x1a\x32\x41\x70\xc0\xb5\x71\xef\xaf\x9d\xe2\xdf\x9d\xd7\xb1\x8e\xe2\xf4\xc3\x6c\xf3\x17\xa4\xda\xc7\xda\xd2\xff\x6d\x01\x99\x7f\x60\xc9\xaf\x0a\x2f\x51\xbb\x5a\xd7\x98\xf5\x5d\xc9\x3b\x9d\xff\x06\x00\x00\xff\xff\x4f\x1b\x28\xab\x8d\x24\x00\x00")

func amazonEksNodegroupYamlBytes() ([]byte, error) {
	return bindataRead(
		_amazonEksNodegroupYaml,
		"amazon-eks-nodegroup.yaml",
	)
}

func amazonEksNodegroupYaml() (*asset, error) {
	bytes, err := amazonEksNodegroupYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "amazon-eks-nodegroup.yaml", size: 9357, mode: os.FileMode(420), modTime: time.Unix(1528190341, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbf, 0xa4, 0x50, 0x91, 0xa5, 0x49, 0xd1, 0xac, 0x1f, 0x42, 0xf8, 0x7a, 0x3b, 0xb2, 0x96, 0xdf, 0x54, 0x74, 0xaa, 0x89, 0xfa, 0x4f, 0x63, 0x45, 0xaa, 0x96, 0x80, 0x62, 0x42, 0xbb, 0x3b, 0x37}}
	return a, nil
}

var _amazonEksServiceRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x51\x6f\xe2\x46\x10\x7e\xf7\xaf\xd8\x43\x95\x90\x2a\x39\x22\x79\xa9\xce\x6f\xbe\x00\x57\x94\xbb\x34\xc2\x34\x79\x5e\x96\x01\x56\xac\x3d\xd6\xec\x6c\x28\xa9\xfa\xdf\xab\xf5\x02\x05\x63\x3b\x50\x1e\x67\xbe\x6f\xfc\xcd\xc7\x8c\x3d\x71\x1c\x47\xe9\x5b\x36\x83\xbc\x34\x92\x61\x8c\x94\x4b\x7e\x05\xb2\x1a\x8b\x44\xf4\x1f\x06\xf7\x83\x78\xf0\x35\x1e\x7c\xed\x47\x43\xb0\x8a\x74\xc9\x21\x93\xe6\xf2\x03\x0b\x31\x7a\xca\x44\x06\xf4\xae\x15\x88\x29\x1a\xe8\x47\x51\x34\x05\x8b\x8e\x14\xd8\x24\x8a\x84\x48\xdf\xb2\x3d\xc0\xe7\xc7\x48\x81\x39\x7a\xca\x92\x48\x08\x21\x66\xbb\x12\x12\x8f\x4a\x92\x49\xfa\x33\x49\x3c\xaa\x4a\xbc\x10\x96\x40\xac\x7d\x1d\x51\xfd\x52\x6b\x5d\x5e\x95\x79\x41\xa3\xd5\x6e\x88\xca\xe5\x50\xf0\x21\x2f\xc4\x99\xf2\x87\xf8\x7e\x10\xdf\xff\xd6\x3f\x66\x33\x96\x0c\xe7\x84\x58\x8c\x96\x4b\x50\x9c\x88\xd4\x18\xdc\x1e\xe3\xfe\xf1\xba\x50\xba\x94\x26\x39\x09\x8a\x43\xb3\xe7\xc1\x58\xc0\xc6\xde\xc9\xaa\x31\xb9\xb5\x77\x0a\xf3\x93\x7c\xaa\x2a\xcf\xa2\x53\xbc\x65\x9b\xfc\xd7\xce\x3e\x55\x75\x75\xd2\x6f\x1c\x22\xbb\x67\x99\x7b\x8b\x0e\xbe\x9d\xd8\x19\xf2\xc7\xca\x6d\xb6\x7c\x66\x4c\xa3\x35\x5d\xe6\x34\x35\xe5\xf1\xd2\x31\x5a\x25\x8d\x2e\x56\x49\x98\x97\x39\xa4\x8e\x31\x0b\xb1\xef\x84\xae\xb4\x1d\x9c\x3f\xcb\x85\xe4\x0b\x46\xdd\x6b\xf5\x90\xa4\xcc\x52\xad\x5f\xd1\xb8\x1c\x9a\xd2\x8e\xd7\x48\xfa\x03\x32\x50\x8e\x34\xef\xaa\x3a\x93\x62\x45\x60\xeb\xcf\xf7\xf8\x47\x02\xc9\xf0\x0c\xbc\x45\xda\x4c\x0a\x06\x5a\x4a\xd5\x54\xb8\x19\xf8\x02\x94\x6b\xeb\x0d\x6e\xa5\x4c\xd1\x71\x7b\xc1\x33\x99\xad\xa8\x99\x5c\xb5\x8b\x6f\xf5\x62\x08\x06\xae\xea\x2d\x00\xdb\x84\x86\xec\x67\x42\x03\xaa\x43\x4b\x98\x89\x49\x61\x59\x16\x0a\x9a\xda\x39\x60\xea\x8a\xbb\xb0\x95\xe8\x99\x9c\x9b\x4e\xd4\x99\xf8\x4e\xa0\x9b\x17\xc0\x5d\x88\xd0\x61\x27\xa2\x54\xcd\xe9\xce\xc1\xfd\x89\x0b\xbd\xdc\x1d\xec\x49\x99\x49\xcf\x9b\xff\x8f\x80\xac\x9b\xd4\xc5\x98\xc2\x3b\x6e\xae\xda\x08\x23\x2d\x6b\x65\x50\x2e\xe6\xd2\xc8\x42\xf9\xd5\x4c\xcb\xd2\xec\xce\x1d\x9c\xe1\x0f\x94\x8b\x6f\x15\x04\xe8\xaa\x22\xd5\xda\x9e\xb2\x66\xd8\x62\x76\x13\xfd\x11\x8b\xa5\x5e\x39\x82\xdf\x41\x1a\x5e\x3f\xae\x41\x6d\xae\xe2\x55\x1b\x72\xab\xd8\x4b\xd6\x0f\x6d\x19\x0a\xa0\xeb\xc4\x5e\xd0\x6b\x6f\xeb\x0e\x6e\x58\xa3\x5b\x05\x5f\xb2\x6e\x12\x3c\x04\x82\x95\x27\xd0\x71\x3d\xc7\x84\xf9\xed\x2a\xc2\x02\x9c\xf2\x8e\x63\x79\xa5\x90\xcb\x0a\x57\x12\xeb\xe3\xe5\x1b\xb8\x61\xc0\xc2\x52\xfd\x7f\xe5\xd3\xba\x81\x6f\x9a\xd7\xb7\x1a\x98\x01\x5f\x4c\x8d\x06\x3b\x46\xfa\x26\xd5\x06\x8a\x85\x3f\x00\x6a\xa5\x0e\x37\x57\x22\x7a\xbf\xf6\xa2\x28\xfa\xc3\x71\xe9\x38\x5c\x60\xfe\x50\x48\x69\xff\xad\x3e\xbb\xe2\x66\x6b\x10\x84\x06\x04\xaf\x25\x57\xb7\xdc\x56\x1b\x23\x9c\x05\xc1\x28\x54\x35\xc0\xfe\x34\x13\x74\x38\xe9\xc4\x12\x49\x3c\xb9\x39\x50\x01\x0c\x56\x28\xe3\x7c\xbf\xc1\x9c\x57\x69\x1c\x24\xe2\xcb\x77\xe0\x94\xb9\xe3\xf2\xbb\x4b\x29\x7c\x29\x47\x7f\x95\x48\xc7\x9b\x23\xdc\x39\x5f\x32\x37\x17\xbd\x5f\xfe\xae\x4e\xc2\x8c\xa5\xda\xf8\xf8\x3f\xf1\xbe\x8d\x5e\x14\xfd\x1b\x00\x00\xff\xff\x6e\xcc\xb0\x9e\xb5\x0a\x00\x00")

func amazonEksServiceRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_amazonEksServiceRoleYaml,
		"amazon-eks-service-role.yaml",
	)
}

func amazonEksServiceRoleYaml() (*asset, error) {
	bytes, err := amazonEksServiceRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "amazon-eks-service-role.yaml", size: 2741, mode: os.FileMode(420), modTime: time.Unix(1528190341, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfa, 0x59, 0xf0, 0x91, 0x72, 0xbd, 0xb5, 0xc0, 0x57, 0x9a, 0x68, 0x95, 0x91, 0x50, 0x62, 0x31, 0xf8, 0xc3, 0xa5, 0x76, 0x90, 0x86, 0x55, 0xcf, 0x37, 0x22, 0x9e, 0x5f, 0x38, 0xef, 0x12, 0x9e}}
	return a, nil
}

var _amazonEksVpcSampleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\xdf\x6f\xdb\x46\x0c\x7e\xd7\x5f\xc1\x1a\x05\xdc\x02\xb1\x23\x19\x43\x90\xe8\xcd\x53\x7e\x2c\x6b\xd7\x1a\x96\xe1\x00\x2d\xfa\x70\x96\x68\xfb\x90\xd3\x9d\x70\x47\x25\xc8\x8a\xfe\xef\x83\x4e\x3f\x2c\xd9\xd2\x66\x63\xcb\xda\x3e\x04\x12\x79\xe4\x77\xfc\x3e\xd2\xd4\x68\x34\x72\xa6\x0f\xe1\x02\x93\x54\x30\xc2\x5b\xa5\x13\x46\x4b\xd4\x86\x2b\xe9\xc3\x70\xe2\x7a\xee\xc8\xbd\x1a\xb9\x57\x43\xe7\x1a\x4d\xa4\x79\x4a\x85\x65\x9a\xb0\x3f\x95\x84\x9b\x0f\x21\x84\x2c\x49\x05\xc2\x72\x16\x0c\x1d\x67\xc6\x34\x4b\x90\x50\x1b\xdf\x01\x08\x44\x66\x08\xf5\x27\x96\x60\xfe\x08\xd0\x0a\xb2\xd8\x22\x44\x85\x07\x48\x96\x20\xd0\x96\x11\x3c\x73\x21\x60\x85\x90\x19\x8c\x81\x14\x44\x1a\x19\xe5\x36\xb4\xd9\xca\x03\x63\x1b\x6e\xf1\x92\xa2\x0f\x21\x69\x2e\x37\x8e\x03\xb0\x4c\xa3\x5f\x85\x8a\x1e\xfd\x43\x6b\x91\x7d\xcd\x32\x41\x3e\x78\x57\x93\xb1\x77\x71\x39\x76\xc7\xee\xb9\x77\xd1\x8d\x2c\xb8\xbf\x9e\x83\x66\x72\x83\xb0\x56\xda\xe6\x5f\xce\x82\x31\x2c\xb6\xdc\x80\xd9\xaa\x4c\xc4\x39\x4c\x06\x4f\x4c\xf0\x18\x52\xcd\x9f\x72\x9c\xef\xe6\xb7\x01\x78\x57\xde\xe5\xfb\x46\x84\x71\x0e\x2e\xcc\x56\x12\xc9\xf5\x4e\x42\x78\xf1\x4b\x0e\xf1\xf2\x10\x62\xc0\x63\x6d\x23\x59\x74\xc6\xc6\x06\xd7\x83\x67\x4e\x5b\x2e\x2b\xb8\x8d\xbc\x93\x93\xf2\x7a\x93\xcb\x93\x12\x4f\x0e\x12\xff\x81\xc4\x62\x46\x2c\xcf\x38\x7d\x08\x7d\x3f\x10\x2a\x8b\x0b\x85\xe5\x81\xfc\x7b\x49\xa8\xd7\x2c\x2a\xa5\x51\x2b\xe7\x4e\xab\x2c\x35\xc5\x4b\x80\x51\xf9\x17\xe0\x23\x5b\xa1\xf0\xeb\x47\x80\xb8\x02\x3d\x28\x75\x06\x81\x92\x6b\xbe\xc9\xb4\xcd\x30\xa8\x5d\xdb\xa2\xac\xfe\x8d\x9a\xf2\x3c\x25\xdb\x83\xd2\x8f\xa8\xe1\x13\xd2\xb3\xd2\x8f\xa7\x26\xad\x34\xda\x7a\xd9\xd2\x46\x97\xa5\x60\xcf\x71\xe6\x68\x54\xa6\x23\xb4\x41\x97\xb3\xa0\xc9\xa7\xad\xf2\x4d\x30\xf1\xfd\x9c\x00\x8b\x41\xab\x14\x35\x71\xac\x31\xd4\xec\xf9\x00\x6f\xe6\xb8\xde\x87\x73\x23\xd9\x4a\xe0\xb5\x34\x61\x96\xa6\x4a\x93\x0f\xa4\x33\xdc\x37\xfe\xa6\x0c\xe5\x1d\x6b\x5a\xe6\x05\xdb\xec\x68\x83\x0f\xf8\xe2\x43\xa3\xb4\x00\x4b\x26\x32\xf4\xe1\x4d\x98\xad\x60\xf8\xf6\xbb\x85\x1b\x12\x8b\x1e\x73\xaf\x1f\x23\x3b\x40\x5a\xa7\x0b\xcf\xc7\x6c\x85\x5a\x22\xa1\x19\x73\x75\x5e\xf6\xff\xf9\xdb\xef\x0d\xf6\x7e\x0c\xf7\x93\x0c\xcd\x96\x69\x8c\x87\x79\x03\x58\xa1\x49\xa4\x3b\x46\xf8\xcc\x5e\x9a\x25\x1b\xec\x6a\xb6\xe7\x35\x70\x8a\x02\x97\x8f\x53\x22\x16\x6d\x13\x94\xd4\x73\xbc\xcb\x75\xd0\x43\xc2\x5e\xaa\xfb\xd8\x2f\xb8\xd8\x7b\x5f\x7a\x2f\xd3\xa8\xf6\x28\x5b\x7a\xae\x32\xc2\x45\x4e\x46\x37\xff\x3b\x7b\x0f\x82\xfd\x98\x27\x11\x38\xcb\x56\x82\x47\xa5\x32\xcd\x9e\x7f\xd1\x14\xdd\x47\x6a\xe8\xd5\x8f\x41\x8a\x32\x36\x9f\xa5\xdf\x59\xe8\xfe\x9b\xf5\x5c\x6a\x77\xeb\xfa\x6e\x7b\x85\xb0\xa3\x8c\xb8\xb4\xad\xda\x68\x05\x77\x6c\xff\x9f\xbb\xa5\xd7\x3f\x11\xd3\x98\xe6\xdd\x04\x14\x56\x6b\x6a\xce\x41\xdb\x81\x2a\xb1\x32\x2a\x23\x80\xeb\xf5\x5c\x67\xfa\xc4\xb8\x60\x2b\x2e\x38\xbd\x7c\x51\x12\x77\x63\xe4\x56\xfa\x7e\x88\x02\x23\xda\xbd\x1b\xc1\xd0\x1d\x36\x9e\x72\x9f\x3b\xa4\xe9\x97\xd6\xf8\x01\x98\xe3\xba\x04\x3a\xc7\x0d\x57\xf2\x60\x2e\x38\x2d\xcf\xae\xc1\x54\x88\xa7\xed\x77\xb2\x88\x6c\x6f\x0f\x0e\xa7\x40\x95\x70\xf0\x2a\xa3\xa0\x9a\xa6\xff\x9e\xb5\xc9\x7f\xc5\x9a\xf7\x7a\xac\x4d\xfe\x6f\xd6\x26\xaf\xca\x9a\xb7\xeb\xe6\xa9\x31\x2a\xe2\xc5\x22\xf1\x37\x54\x76\x1e\xe8\x21\xae\x38\x51\xb7\x7c\x95\xf3\xa8\xd9\xd2\x50\xd6\x4f\xc0\x38\x39\x16\x63\xa0\x24\x69\x25\x66\x82\x49\x0c\x31\xca\x34\xa7\x17\xbb\x6d\xf5\xe0\x6b\xba\xf4\x20\xb2\xb6\xf6\x7e\x58\xee\x62\x91\x4a\x92\x4c\xf2\xc8\xde\xc7\x6e\x87\xf0\x5c\x2c\x4e\x52\xc5\x68\xfa\x7e\xdc\x3e\x67\x94\x66\x74\xe4\xd7\x83\xfd\x6a\x50\xeb\xfd\x8f\x03\xa7\xa9\xd8\x3c\x72\x73\xd1\xab\xb9\x32\x1d\x41\xa7\x42\x94\x2b\xad\x81\xc6\x32\xdb\x8c\xf7\xbb\xe2\x12\xbe\xc2\x00\x06\x67\xf0\xb5\xad\x95\xb3\x36\x2d\xf0\x0d\xbe\x35\xd2\x7d\xe4\x86\x8e\x4e\x09\xef\xf2\xfa\xb1\x91\xc1\x94\x69\x46\x18\xbf\xef\x04\x71\x76\x34\x88\xb2\x04\xf5\xca\x78\x50\xcb\xfb\xeb\xaa\x92\x79\x7a\x4e\x06\xc5\xfa\xa0\x90\xd5\x27\xc5\xa1\x78\x5a\xf1\x2a\x3b\x6c\x72\x87\xfa\xfb\x29\xaa\xa5\x61\x75\x08\x69\x2e\xc4\xa3\x84\xd2\x22\xb3\x4f\xc5\x8e\xf3\x57\x00\x00\x00\xff\xff\xd4\x28\xd7\x20\xcf\x0e\x00\x00")

func amazonEksVpcSampleYamlBytes() ([]byte, error) {
	return bindataRead(
		_amazonEksVpcSampleYaml,
		"amazon-eks-vpc-sample.yaml",
	)
}

func amazonEksVpcSampleYaml() (*asset, error) {
	bytes, err := amazonEksVpcSampleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "amazon-eks-vpc-sample.yaml", size: 3791, mode: os.FileMode(420), modTime: time.Unix(1528190448, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0xbc, 0x6, 0xc0, 0x37, 0x67, 0x9d, 0x35, 0xfe, 0x98, 0x7c, 0xba, 0x47, 0x77, 0xfd, 0x1f, 0xd, 0xe9, 0x60, 0x61, 0x46, 0x34, 0x7f, 0x48, 0xa5, 0x29, 0x45, 0x39, 0xf2, 0x8f, 0x4f, 0xb1}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"amazon-eks-nodegroup.yaml": amazonEksNodegroupYaml,

	"amazon-eks-service-role.yaml": amazonEksServiceRoleYaml,

	"amazon-eks-vpc-sample.yaml": amazonEksVpcSampleYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"amazon-eks-nodegroup.yaml":    &bintree{amazonEksNodegroupYaml, map[string]*bintree{}},
	"amazon-eks-service-role.yaml": &bintree{amazonEksServiceRoleYaml, map[string]*bintree{}},
	"amazon-eks-vpc-sample.yaml":   &bintree{amazonEksVpcSampleYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}