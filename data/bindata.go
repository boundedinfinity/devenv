// Code generated by go-bindata.
// sources:
// bindata.go
// project/editorconfig/.editorconfig
// project/glide/glide.yaml
// project/makefile/Makefile
// user/config/bash/load-env.bash
// user/config/bash/scripts.d/.gitkeep
// user/config/devenv/available/bash/scripts.d/anyenv.bash
// user/config/devenv/available/bash/scripts.d/go.bash
// user/config/devenv/available/fish/scripts.d/anyenv.fish
// user/config/devenv/available/fish/scripts.d/go.fish
// user/config/devenv/enabled/.gitkeep
// user/config/fish/load-env.fish
// user/config/fish/scripts.d/.gitkeep
// DO NOT EDIT!

package data

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

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1480404288, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _projectEditorconfigEditorconfig = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\xcb\x41\x0a\xc2\x40\x0c\x85\xe1\x7d\x4e\x91\xb5\x6b\xb7\x3d\x82\x27\x28\x25\x44\xfb\x02\xc1\x90\x91\x99\x88\xe8\xe9\xa5\x22\x08\x76\xfb\xbd\xf7\xf7\xd6\x8a\x27\xae\x7e\x07\xd1\x7c\x58\x08\xb9\x4a\x33\x09\x4f\xf0\xc4\x61\xe4\x39\xd0\x4b\xcc\x53\x43\x12\x8f\xef\xf2\x29\x3c\x57\x64\xc9\xa8\x67\x6c\x36\x6e\x7a\xf9\xa1\xbf\x36\x3b\x12\xcd\x27\xbd\xc2\x3c\xb0\xfc\x07\xa5\xe7\xdd\xfd\x1d\x00\x00\xff\xff\xc8\xd9\x9e\x08\x92\x00\x00\x00")

func projectEditorconfigEditorconfigBytes() ([]byte, error) {
	return bindataRead(
		_projectEditorconfigEditorconfig,
		"project/editorconfig/.editorconfig",
	)
}

func projectEditorconfigEditorconfig() (*asset, error) {
	bytes, err := projectEditorconfigEditorconfigBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "project/editorconfig/.editorconfig", size: 146, mode: os.FileMode(420), modTime: time.Unix(1479504906, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _projectGlideGlideYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\xb5\x52\xa8\xae\xd6\x0b\x80\xb0\xfd\x12\x73\x53\x6b\x6b\xb9\x32\x73\x0b\xf2\x8b\x4a\xac\x14\xa2\x63\xb9\x00\x01\x00\x00\xff\xff\x5d\x8a\xe0\xe3\x25\x00\x00\x00")

func projectGlideGlideYamlBytes() ([]byte, error) {
	return bindataRead(
		_projectGlideGlideYaml,
		"project/glide/glide.yaml",
	)
}

func projectGlideGlideYaml() (*asset, error) {
	bytes, err := projectGlideGlideYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "project/glide/glide.yaml", size: 37, mode: os.FileMode(420), modTime: time.Unix(1479515893, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _projectMakefileMakefile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcd\xc1\xaa\x82\x40\x14\xc6\xf1\xb5\xf3\x14\x07\xae\x30\x2a\xdc\x71\x7f\x40\xe8\x05\x8a\xf6\xa6\x72\xd2\x93\x5a\x33\x39\xcc\x0c\xb5\x30\xdf\x3d\x8c\x02\x97\xbf\x8f\x3f\x7c\x86\x6e\x7c\x19\x35\x37\xdd\xe8\x22\x2c\x20\x4e\xe8\xec\x2d\x85\x01\xe2\xc4\x0f\xac\x35\xd8\x67\x97\xa6\x42\x90\xb5\xcd\x9d\x0c\x47\x6b\x35\xcf\xea\xe8\xa6\x2b\xb7\xe1\x40\x86\x97\x45\x08\x3d\xfa\x80\x22\xda\xf5\x8e\x2d\xc8\xba\xac\xff\x4a\xf4\x96\x5a\xc6\xaa\x52\x19\x4a\xd8\x7f\x8f\xe0\x05\x9f\xe6\xff\x01\x12\x0b\xb9\x65\x7d\x52\xab\x3d\x77\x20\x7d\x8e\x2a\xcb\xf3\x7e\x3b\xfc\x38\xb9\x20\xde\x01\x00\x00\xff\xff\xaa\x44\xaf\x76\xb8\x00\x00\x00")

func projectMakefileMakefileBytes() ([]byte, error) {
	return bindataRead(
		_projectMakefileMakefile,
		"project/makefile/Makefile",
	)
}

func projectMakefileMakefile() (*asset, error) {
	bytes, err := projectMakefileMakefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "project/makefile/Makefile", size: 184, mode: os.FileMode(420), modTime: time.Unix(1479504751, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userConfigBashLoadEnvBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\x4a\x2c\xce\xe0\xe2\x4a\xcb\x2f\x52\x28\x4e\x2e\xca\x2c\x28\x51\xc8\xcc\x53\x50\x71\xf6\xf7\x73\xf3\x74\x8f\x0f\xf2\xf7\x0f\xd1\x07\x29\xd0\x87\xc8\x15\xeb\xa5\xe8\x6b\xe9\x81\x04\xac\x15\x52\xf2\xb9\x14\x14\x14\x14\x8a\xf3\x4b\x8b\x92\x53\x15\x54\x20\x0a\xb8\x52\xf2\xf3\x52\xb9\x00\x01\x00\x00\xff\xff\xd7\x56\x2e\x93\x62\x00\x00\x00")

func userConfigBashLoadEnvBashBytes() ([]byte, error) {
	return bindataRead(
		_userConfigBashLoadEnvBash,
		"user/config/bash/load-env.bash",
	)
}

func userConfigBashLoadEnvBash() (*asset, error) {
	bytes, err := userConfigBashLoadEnvBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "user/config/bash/load-env.bash", size: 98, mode: os.FileMode(420), modTime: time.Unix(1479702201, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userConfigBashScriptsDGitkeep = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func userConfigBashScriptsDGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_userConfigBashScriptsDGitkeep,
		"user/config/bash/scripts.d/.gitkeep",
	)
}

func userConfigBashScriptsDGitkeep() (*asset, error) {
	bytes, err := userConfigBashScriptsDGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "user/config/bash/scripts.d/.gitkeep", size: 0, mode: os.FileMode(420), modTime: time.Unix(1479701462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userConfigDevenvAvailableBashScriptsDAnyenvBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\x4a\x2c\xce\xe0\xe2\x4a\xad\x28\xc8\x2f\x2a\x51\x70\xf4\x8b\x74\xf5\x0b\x8b\x0f\xf2\xf7\x0f\xb1\x55\x71\xf6\xf7\x73\xf3\x74\x07\x73\xf4\x13\xf3\x2a\x53\xf3\xca\x60\xca\x02\x1c\x43\x3c\x6c\x95\x54\x90\x54\x83\x8c\xb3\x52\x01\x89\x2b\x71\x71\xa5\x96\x25\xe6\x28\x28\xa9\x68\x40\x34\x29\x64\xe6\x65\x96\x28\xe8\x6a\x2a\x71\x01\x02\x00\x00\xff\xff\x91\xbb\xbf\x90\x7a\x00\x00\x00")

func userConfigDevenvAvailableBashScriptsDAnyenvBashBytes() ([]byte, error) {
	return bindataRead(
		_userConfigDevenvAvailableBashScriptsDAnyenvBash,
		"user/config/devenv/available/bash/scripts.d/anyenv.bash",
	)
}

func userConfigDevenvAvailableBashScriptsDAnyenvBash() (*asset, error) {
	bytes, err := userConfigDevenvAvailableBashScriptsDAnyenvBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "user/config/devenv/available/bash/scripts.d/anyenv.bash", size: 122, mode: os.FileMode(420), modTime: time.Unix(1479701722, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userConfigDevenvAvailableBashScriptsDGoBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\x4a\x2c\xce\xe0\xe2\x4a\xad\x28\xc8\x2f\x2a\x51\x70\xf7\x0f\x70\x0c\xf1\xb0\x55\xf1\xf0\xf7\x75\xd5\x4f\x49\x2d\xd3\x2f\x28\xca\xcf\x4a\x4d\x2e\x29\xd6\x4f\xcf\x87\xa9\x81\xa8\x00\x91\x56\x2a\x10\xf5\x20\xb3\xb8\x00\x01\x00\x00\xff\xff\xc3\x61\xc9\x35\x57\x00\x00\x00")

func userConfigDevenvAvailableBashScriptsDGoBashBytes() ([]byte, error) {
	return bindataRead(
		_userConfigDevenvAvailableBashScriptsDGoBash,
		"user/config/devenv/available/bash/scripts.d/go.bash",
	)
}

func userConfigDevenvAvailableBashScriptsDGoBash() (*asset, error) {
	bytes, err := userConfigDevenvAvailableBashScriptsDGoBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "user/config/devenv/available/bash/scripts.d/go.bash", size: 87, mode: os.FileMode(420), modTime: time.Unix(1479970520, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userConfigDevenvAvailableFishScriptsDAnyenvFish = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\xcb\x2c\xce\xe0\xe2\x2a\x4e\x2d\x51\xd0\x4d\xaf\x50\x70\xf4\x8b\x74\xf5\x0b\x8b\x0f\xf2\xf7\x0f\x51\x50\x71\xf6\xf7\x73\xf3\x74\x07\x73\xf4\x13\xf3\x2a\x53\xf3\xca\xe0\xea\x02\x1c\x43\x3c\x14\x54\x90\x54\x83\xcc\x53\x50\x01\x09\x73\x71\xa5\x96\x25\xe6\x28\x28\x69\x40\xb4\x28\x64\xe6\x65\x96\x28\xe8\x6a\x2a\x71\x01\x02\x00\x00\xff\xff\x7f\x42\xdf\x3b\x79\x00\x00\x00")

func userConfigDevenvAvailableFishScriptsDAnyenvFishBytes() ([]byte, error) {
	return bindataRead(
		_userConfigDevenvAvailableFishScriptsDAnyenvFish,
		"user/config/devenv/available/fish/scripts.d/anyenv.fish",
	)
}

func userConfigDevenvAvailableFishScriptsDAnyenvFish() (*asset, error) {
	bytes, err := userConfigDevenvAvailableFishScriptsDAnyenvFishBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "user/config/devenv/available/fish/scripts.d/anyenv.fish", size: 121, mode: os.FileMode(420), modTime: time.Unix(1479701739, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userConfigDevenvAvailableFishScriptsDGoFish = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\xcb\x2c\xce\xe0\xe2\x2a\x4e\x2d\x51\xd0\x4d\xaf\x50\x70\xf7\x0f\x70\x0c\xf1\x50\xa8\xae\x56\xd0\x73\xcf\x0f\x48\x2c\xc9\x50\xa8\xad\x85\x4b\x82\xa5\x54\x20\x24\x44\x21\xc8\x14\x2e\x40\x00\x00\x00\xff\xff\xf0\xe5\x4f\x3e\x51\x00\x00\x00")

func userConfigDevenvAvailableFishScriptsDGoFishBytes() ([]byte, error) {
	return bindataRead(
		_userConfigDevenvAvailableFishScriptsDGoFish,
		"user/config/devenv/available/fish/scripts.d/go.fish",
	)
}

func userConfigDevenvAvailableFishScriptsDGoFish() (*asset, error) {
	bytes, err := userConfigDevenvAvailableFishScriptsDGoFishBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "user/config/devenv/available/fish/scripts.d/go.fish", size: 81, mode: os.FileMode(420), modTime: time.Unix(1480372080, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userConfigDevenvEnabledGitkeep = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func userConfigDevenvEnabledGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_userConfigDevenvEnabledGitkeep,
		"user/config/devenv/enabled/.gitkeep",
	)
}

func userConfigDevenvEnabledGitkeep() (*asset, error) {
	bytes, err := userConfigDevenvEnabledGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "user/config/devenv/enabled/.gitkeep", size: 0, mode: os.FileMode(420), modTime: time.Unix(1480372095, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userConfigFishLoadEnvFish = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\xcb\x2c\xce\xe0\xe2\x4a\xcb\x2f\x52\x28\x4e\x2e\xca\x2c\x28\x51\xc8\xcc\x53\x50\x71\xf6\xf7\x73\xf3\x74\x8f\x0f\xf2\xf7\x0f\xd1\x07\x29\xd0\x87\xc8\x15\xeb\xa5\xe8\x6b\xe9\x81\x75\x28\x28\x28\x28\x14\xe7\x97\x16\x25\xa7\x2a\xa8\x40\x24\xb9\x52\xf3\x52\xb8\xb8\x00\x01\x00\x00\xff\xff\x65\x82\x06\x6c\x5e\x00\x00\x00")

func userConfigFishLoadEnvFishBytes() ([]byte, error) {
	return bindataRead(
		_userConfigFishLoadEnvFish,
		"user/config/fish/load-env.fish",
	)
}

func userConfigFishLoadEnvFish() (*asset, error) {
	bytes, err := userConfigFishLoadEnvFishBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "user/config/fish/load-env.fish", size: 94, mode: os.FileMode(420), modTime: time.Unix(1479702141, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userConfigFishScriptsDGitkeep = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func userConfigFishScriptsDGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_userConfigFishScriptsDGitkeep,
		"user/config/fish/scripts.d/.gitkeep",
	)
}

func userConfigFishScriptsDGitkeep() (*asset, error) {
	bytes, err := userConfigFishScriptsDGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "user/config/fish/scripts.d/.gitkeep", size: 0, mode: os.FileMode(420), modTime: time.Unix(1479701462, 0)}
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
	"bindata.go": bindataGo,
	"project/editorconfig/.editorconfig": projectEditorconfigEditorconfig,
	"project/glide/glide.yaml": projectGlideGlideYaml,
	"project/makefile/Makefile": projectMakefileMakefile,
	"user/config/bash/load-env.bash": userConfigBashLoadEnvBash,
	"user/config/bash/scripts.d/.gitkeep": userConfigBashScriptsDGitkeep,
	"user/config/devenv/available/bash/scripts.d/anyenv.bash": userConfigDevenvAvailableBashScriptsDAnyenvBash,
	"user/config/devenv/available/bash/scripts.d/go.bash": userConfigDevenvAvailableBashScriptsDGoBash,
	"user/config/devenv/available/fish/scripts.d/anyenv.fish": userConfigDevenvAvailableFishScriptsDAnyenvFish,
	"user/config/devenv/available/fish/scripts.d/go.fish": userConfigDevenvAvailableFishScriptsDGoFish,
	"user/config/devenv/enabled/.gitkeep": userConfigDevenvEnabledGitkeep,
	"user/config/fish/load-env.fish": userConfigFishLoadEnvFish,
	"user/config/fish/scripts.d/.gitkeep": userConfigFishScriptsDGitkeep,
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
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"project": &bintree{nil, map[string]*bintree{
		"editorconfig": &bintree{nil, map[string]*bintree{
			".editorconfig": &bintree{projectEditorconfigEditorconfig, map[string]*bintree{}},
		}},
		"glide": &bintree{nil, map[string]*bintree{
			"glide.yaml": &bintree{projectGlideGlideYaml, map[string]*bintree{}},
		}},
		"makefile": &bintree{nil, map[string]*bintree{
			"Makefile": &bintree{projectMakefileMakefile, map[string]*bintree{}},
		}},
	}},
	"user": &bintree{nil, map[string]*bintree{
		"config": &bintree{nil, map[string]*bintree{
			"bash": &bintree{nil, map[string]*bintree{
				"load-env.bash": &bintree{userConfigBashLoadEnvBash, map[string]*bintree{}},
				"scripts.d": &bintree{nil, map[string]*bintree{
					".gitkeep": &bintree{userConfigBashScriptsDGitkeep, map[string]*bintree{}},
				}},
			}},
			"devenv": &bintree{nil, map[string]*bintree{
				"available": &bintree{nil, map[string]*bintree{
					"bash": &bintree{nil, map[string]*bintree{
						"scripts.d": &bintree{nil, map[string]*bintree{
							"anyenv.bash": &bintree{userConfigDevenvAvailableBashScriptsDAnyenvBash, map[string]*bintree{}},
							"go.bash": &bintree{userConfigDevenvAvailableBashScriptsDGoBash, map[string]*bintree{}},
						}},
					}},
					"fish": &bintree{nil, map[string]*bintree{
						"scripts.d": &bintree{nil, map[string]*bintree{
							"anyenv.fish": &bintree{userConfigDevenvAvailableFishScriptsDAnyenvFish, map[string]*bintree{}},
							"go.fish": &bintree{userConfigDevenvAvailableFishScriptsDGoFish, map[string]*bintree{}},
						}},
					}},
				}},
				"enabled": &bintree{nil, map[string]*bintree{
					".gitkeep": &bintree{userConfigDevenvEnabledGitkeep, map[string]*bintree{}},
				}},
			}},
			"fish": &bintree{nil, map[string]*bintree{
				"load-env.fish": &bintree{userConfigFishLoadEnvFish, map[string]*bintree{}},
				"scripts.d": &bintree{nil, map[string]*bintree{
					".gitkeep": &bintree{userConfigFishScriptsDGitkeep, map[string]*bintree{}},
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

