package bounded_xdg

import (
	"bytes"
	"crypto/sha512"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed embedded/*
var embedded embed.FS

func newFileManager() *BoundedFileManager {
	return &BoundedFileManager{
		defaults:     map[string]string{},
		environment:  map[string]string{},
		embeddedRoot: "embedded",
	}
}

type BoundedFileManager struct {
	defaults     map[string]string
	environment  map[string]string
	embeddedRoot string
}

func (this *BoundedFileManager) init() error {
	return nil
}

// ///////////////////////////////////////////////////////////////////////////
// Utilities
// ///////////////////////////////////////////////////////////////////////////

func (this *BoundedFileManager) resolveVar(name string) string {
	for k, v := range this.environment {
		if strings.Contains(name, "$"+k) {
			name = strings.ReplaceAll(name, "$"+k, v)
			name = this.resolvePath(name)
		}
	}

	return name
}

func (this *BoundedFileManager) resolvePath(path ...string) string {
	rpath := filepath.Join(path...)
	sep := string(filepath.Separator)
	comps := strings.Split(rpath, sep)

	for i := 0; i < len(comps); i++ {
		comps[i] = this.resolveVar(comps[i])
	}

	return strings.Join(comps, sep)
}

func (this *BoundedFileManager) setDefaults(defaults map[string]string) {
	for k, v := range defaults {
		this.defaults[k] = v
		this.environment[k] = v
	}
}

func (this *BoundedFileManager) ensureVar(name string) error {
	if strings.HasPrefix(name, "$") {
		name = strings.Replace(name, "$", "", 1)
	}

	v := os.Getenv(name)

	if v == "" {
		if found, ok := this.defaults[name]; ok {
			v = found
		}
	}

	if v == "" {
		return fmt.Errorf("environment variable %s not defined", name)
	}

	this.environment[name] = v
	return nil
}

func (this *BoundedFileManager) calcHash(data []byte) string {
	h := sha512.New()
	h.Write(data)
	sum := h.Sum(nil)
	return string(sum)
}

type IsSameResult struct {
	Match bool
	SumA  string
	SumB  string
}

func (this *BoundedFileManager) isSameData(a, b []byte) IsSameResult {
	sumA := this.calcHash(a)
	sumB := this.calcHash(b)
	return IsSameResult{
		Match: sumA == sumB,
		SumA:  sumA,
		SumB:  sumB,
	}
}

func (this *BoundedFileManager) isSameObject(a, b any) (IsSameResult, error) {
	aData, err := json.MarshalIndent(a, _JSON_META.Prefix, _JSON_META.Indent)

	if err != nil {
		return IsSameResult{}, err
	}

	bData, err := json.MarshalIndent(b, _JSON_META.Prefix, _JSON_META.Indent)

	if err != nil {
		return IsSameResult{}, err
	}

	return this.isSameData(aData, bData), nil
}

// ///////////////////////////////////////////////////////////////////////////
// File System
// ///////////////////////////////////////////////////////////////////////////

func (this *BoundedFileManager) fsReadDir(path ...string) ([]fs.DirEntry, error) {
	resolved := this.resolvePath(path...)
	return os.ReadDir(resolved)
}

func (this *BoundedFileManager) fsReadFile(path ...string) ([]byte, error) {
	resolved := this.resolvePath(path...)
	data, err := os.ReadFile(resolved)

	if err != nil {
		return data, err
	}

	data = bytes.TrimSpace(data)
	return data, nil
}

func (this *BoundedFileManager) fsUnmarshalFile(v any, path ...string) error {
	data, err := this.fsReadFile(path...)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	return nil
}

type PathExistsResult struct {
	ResolvedPath string
	Exists       bool
	IsDir        bool
	IsSymLink    bool
}

func (this *BoundedFileManager) fsExists(path ...string) (PathExistsResult, error) {
	result := PathExistsResult{
		ResolvedPath: this.resolvePath(path...),
	}

	info, err := os.Lstat(result.ResolvedPath)

	switch {
	case err != nil && os.IsNotExist(err):
		return result, nil
	case err != nil && !os.IsNotExist(err):
		return result, err
	default:
		mode := info.Mode()
		result.Exists = true
		result.IsDir = info.IsDir()
		result.IsSymLink = mode == fs.ModeSymlink
		return result, nil
	}
}

func (this *BoundedFileManager) fsEnsureDir(path ...string) error {
	resovled := this.resolvePath(path...)
	result, err := this.fsExists(resovled)

	if err != nil {
		return err
	}

	switch {
	case result.Exists && result.IsDir:
		return nil
	case result.Exists && !result.IsDir:
		return fmt.Errorf("%s exists but is not directory", path)
	}

	if err := os.MkdirAll(result.ResolvedPath, os.FileMode(0755)); err != nil {
		return err
	}

	return nil
}

func (this *BoundedFileManager) fsWriteFile(data []byte, path ...string) error {
	rpath := filepath.Join(path...)
	exists, err := this.fsExists(rpath)

	if err != nil {
		return err
	}

	if exists.Exists {
		current, err := this.fsReadFile(rpath)

		if err != nil {
			return err
		}

		result := this.isSameData(data, current)

		if !result.Match {
			return fmt.Errorf("%s already exists and is modified", rpath)
		}

		return nil
	}

	err = os.WriteFile(rpath, data, os.FileMode(0775))

	if err != nil {
		return err
	}

	return nil
}

func (this *BoundedFileManager) fsMarshalFile(jsonAny any, path ...string) error {
	data, err := json.MarshalIndent(jsonAny, _JSON_META.Prefix, _JSON_META.Indent)

	if err != nil {
		return err
	}

	return this.fsWriteFile(data, path...)
}

func (this *BoundedFileManager) fsSymLink(source, dest string) error {
	existSource, err := this.fsExists(source)

	if err != nil {
		return err
	}

	existDest, err := this.fsExists(dest)

	if err != nil {
		return err
	}

	switch {
	case !existSource.Exists:
		return fmt.Errorf("source doesn't exists: %s", source)
	case existDest.Exists && !existDest.IsSymLink:
		return fmt.Errorf("dest exists but isn't a symbolic link: %s", dest)
	default:
		if err := os.Symlink(existSource.ResolvedPath, existDest.ResolvedPath); err != nil {
			return err
		}
	}

	return nil
}

// ///////////////////////////////////////////////////////////////////////////
// Embedded File System
// ///////////////////////////////////////////////////////////////////////////

func (this *BoundedFileManager) embeddedReadDir(path ...string) ([]fs.DirEntry, error) {
	path = append([]string{this.embeddedRoot}, path...)
	joined := filepath.Join(path...)
	return embedded.ReadDir(joined)
}

func (this *BoundedFileManager) embeddedReadFile(path ...string) ([]byte, error) {
	path = append([]string{this.embeddedRoot}, path...)
	joined := filepath.Join(path...)
	data, err := embedded.ReadFile(joined)

	if err != nil {
		return data, err
	}

	data = bytes.TrimSpace(data)
	return data, nil
}

func (this *BoundedFileManager) embeddedUnmarshalFile(v any, path ...string) error {
	data, err := this.embeddedReadFile(path...)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	return nil
}
