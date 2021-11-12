// Code generated for package conf by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../conf/app.conf
package conf

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _confAppConf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xcb\x6e\xdc\x30\x0c\xbc\xeb\x2b\x06\x3a\xb7\x86\x93\x02\x4d\xdb\x85\x4f\xfd\x85\xde\x82\x60\x21\x5b\xb4\xad\x42\x16\x05\x92\xca\xe3\xef\x0b\x6d\x76\x7b\x68\x2e\x3d\xd1\x1e\x0e\x67\x38\x84\x42\xad\xe7\x12\x0e\xc2\x04\xbf\x71\xbd\xaf\xde\xcd\x12\x4a\xfc\x17\x94\x56\xce\x4d\x49\x3a\x24\xcc\xf6\x8e\x1c\x1c\x2f\xa4\x48\xcf\xde\xb9\xc7\xcc\xdb\x93\x3b\xe1\xd7\x4e\xc8\xbc\x61\x65\x39\x82\x81\x92\xed\x24\xf0\xbf\x95\x8b\x07\x0b\xbc\xd1\xab\x79\x77\x6d\x4f\xb7\xff\x2e\x7b\xae\xc1\xf6\x0e\x65\xde\xb4\x4b\x0a\xc5\xa4\x4f\x70\x27\xa8\xb1\x10\x94\x54\x13\x17\x47\x25\xcc\xb9\x7b\xaf\x21\x2b\xb9\x10\xa3\x90\x6a\x9f\xbc\xbb\x7f\x18\xc6\x61\x1c\xee\x7e\x7c\xfd\xf2\xf0\xdd\xbb\x1a\x54\x5f\x58\x62\xef\x79\x17\xe7\x5e\x47\xef\x9c\x7b\x7c\xa1\xf9\xb6\x6e\x15\x36\x5e\x38\xc3\xf6\x60\x48\x8a\xa6\x14\x61\x0c\x25\x79\x26\xc4\x24\xb4\x18\x4c\xc2\xba\xa6\xa5\xe3\xb6\x13\x42\xad\x39\x2d\xc1\x12\x97\xc1\x9d\xf0\xb3\x89\x50\xb1\xfc\x06\x6d\xb5\xb2\x98\xc2\xef\x66\xd5\x7f\x7a\xaf\xda\x3f\xd6\x65\x4b\x1e\xa1\x44\xf8\x56\xd2\xab\x1f\xdc\x5f\xeb\x09\x9d\x75\x5d\xe8\x16\xc8\x18\x33\x21\x27\x35\x2a\x14\x31\xbf\x7d\x74\xee\x53\xe7\xce\xc7\x84\xf1\x12\x7d\x74\xb7\x58\x2c\x86\xd2\x8e\x99\xe4\xbf\x95\x2e\x33\xfd\x8e\xe3\xb7\x7e\xa6\xab\x50\x9b\x73\x5a\x3e\xaf\x61\x49\x65\x43\xe4\x23\xa4\x82\xcb\x13\x59\x59\x3e\x0a\x5d\x09\x13\x32\x2f\x21\xef\xac\xe6\xfe\x04\x00\x00\xff\xff\x99\x57\xc3\xf9\x69\x02\x00\x00"

func confAppConfBytes() ([]byte, error) {
	return bindataRead(
		_confAppConf,
		"conf/app.conf",
	)
}

func confAppConf() (*asset, error) {
	bytes, err := confAppConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/app.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"conf/app.conf": confAppConf,
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
	"conf": &bintree{nil, map[string]*bintree{
		"app.conf": &bintree{confAppConf, map[string]*bintree{}},
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
