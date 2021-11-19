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

var _confAppConf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xb1\xce\xdb\x30\x0c\x84\x77\x3e\x05\xa1\xbd\xae\xe3\xa0\x09\xfe\x06\x9e\x3a\x14\x05\x3a\x14\x45\x37\xc3\x30\x64\x8b\xb6\xd5\x4a\xa2\x20\xd1\x7f\x92\xb7\x2f\x94\xc4\x4b\xa7\x8e\xbc\x23\xbf\x3b\x41\x3a\xc6\x21\x68\x4f\xd8\xa2\x5a\x38\x36\x51\xc1\x98\x74\x30\xff\x8a\x69\x0b\xc3\x96\x29\x15\x29\x31\xcb\x53\xf1\x6c\x1e\x4b\x86\xde\x15\x40\xe7\x78\xe9\xe1\x82\xbf\x56\x42\xc7\x0b\xce\x9c\xbc\x16\x24\x2b\x2b\x25\x54\xbf\x33\x07\x85\x9c\x50\x09\xdd\x44\xc1\xcb\x6e\xf7\xb9\x60\x87\xa8\x65\x2d\x92\xe3\x25\x17\x64\x22\x63\x73\x8f\x70\xc1\x2c\x9c\x08\x33\xe5\x6c\x39\x00\x05\x3d\xba\x92\x3d\x6b\x97\x09\xb4\x31\x89\x72\x2e\x97\x87\xe6\x5c\xd5\x55\x5d\x1d\x3e\x9f\x8e\xe7\x37\x05\x51\xe7\x7c\xe5\x64\xb0\x45\x30\x23\xb6\x58\x43\x4c\x34\xdb\x1b\xb6\xb8\xc4\x01\xa0\xf3\xe4\x7b\x80\xcb\xb8\xcd\x33\x25\x14\xeb\x09\xbb\x4c\x13\x07\xd3\xc3\xa4\xa7\x95\x86\x69\xd5\x21\x90\x1b\x1e\x5e\x8b\xa7\x1a\xa0\xbb\xd2\xb8\x3f\x76\x4f\x17\xc6\x91\xd0\xd9\x2c\x14\xc8\xe0\x78\x47\x29\x6e\x8c\xce\x4e\x5a\x2c\x87\x0a\x56\x91\x38\x94\xfd\x52\xe4\xd1\xb3\x86\x17\x25\x72\x12\x0c\x9b\x1f\x4b\x89\xff\x24\x3d\x6e\x5a\x54\xc7\xfa\x58\x2b\x80\xcb\x4f\x32\xdb\x44\x48\x37\xa1\x14\xb4\xc3\x40\x72\xe5\xf4\x07\x0d\x45\x0a\x86\xc2\x74\x47\x1d\x0c\xea\x4d\xd8\x6b\xb1\x93\x76\xee\x8e\x99\xd2\x3b\x25\xfc\xf6\xe3\xc9\x7c\x8e\x7b\xc9\xc3\xb9\xa9\x4e\xe7\xaa\xa9\xdf\xaa\xd3\x27\x80\x6e\x21\xee\x61\xff\x25\xa3\x45\x7f\xfc\x4a\xfc\xdd\x0a\x35\x1f\xbe\x58\xb9\x57\xde\x9b\x51\x01\xfc\x0d\x00\x00\xff\xff\xb7\x90\x58\xec\x58\x02\x00\x00"

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
