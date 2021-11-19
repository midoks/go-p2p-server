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

var _confAppConf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xb1\x8e\xdb\x30\x0c\x86\x77\x3e\x05\xa1\xbd\xae\xe3\xa0\x09\xae\x81\xa7\x0e\x45\x81\x0e\x45\xd1\x2d\x08\x02\x3a\xe2\x39\x6a\x65\x51\xa0\xe8\xcb\xf9\xed\x0b\x25\xf1\xd2\xa9\x23\x3f\xf2\xff\x48\x90\x72\x3e\x27\x9a\x18\x7b\x74\xa3\xe4\x2e\x3b\x18\x94\x92\xff\x17\xea\x9c\xce\x73\x61\xad\x48\x45\xec\x41\x26\xf1\xf7\x21\xcf\x6f\x0e\xe0\x18\x65\x3c\xc1\x01\x7f\x5d\x19\xa3\x8c\xf8\x2a\x3a\x91\x21\x07\xbb\xb2\xa2\xfb\x5d\x24\x39\x14\x45\x67\xfc\x6e\x0e\x9e\xed\x7e\xad\xab\xf6\x9c\xc9\xae\x15\x45\x19\x4b\x55\x2a\xfb\x50\x4e\x08\x07\x2c\x26\xca\x58\xb8\x94\x20\x09\x38\xd1\x10\xeb\xee\x57\x8a\x85\x81\xbc\x57\x2e\xa5\x26\x37\xdd\xbe\x69\x9b\xb6\xd9\x7c\xde\x6d\xf7\x2f\x0e\x32\x95\x72\x13\xf5\xd8\x23\xf8\x01\x7b\x6c\x01\xe0\x78\xe3\x61\x3d\x75\xcd\x9a\xe0\xc0\x18\x43\x31\x4e\xec\x71\x58\xd0\x6a\x37\xe7\x18\x2e\x64\x41\x52\x03\x57\xb3\x7c\xae\xf3\x55\x73\xdf\xd2\xc2\xd3\x92\x45\x0d\xd3\x3c\x0d\xac\xff\x6d\xba\x67\x7a\x74\xdb\x76\xdb\x3a\x80\xc3\x4f\xf6\xf3\x85\x91\xdf\x8d\x35\x51\xc4\xc4\x76\x13\xfd\x83\x9e\x33\x27\xcf\xe9\xb2\x20\x25\x8f\x34\x9b\x4c\x64\xe1\x42\x31\x2e\x58\x58\xdf\x58\xf1\xdb\x8f\x87\xf3\x51\xae\x47\x6e\xf6\x5d\xb3\xdb\x37\x5d\xfb\xd2\xec\x3e\x01\x1c\x47\x96\x13\xac\x3f\xf6\x64\xf4\xf1\x2b\xcb\xf7\x60\xdc\x7d\xf8\x12\x6c\x69\xa6\xc9\x0f\xee\x6f\x00\x00\x00\xff\xff\x73\x77\xd6\x0f\x14\x02\x00\x00"

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
