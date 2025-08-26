package bounded_xdg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func NewFileManager() *BoundedFileManager {
	return &BoundedFileManager{
		defaults:    map[string]string{},
		environment: map[string]string{},
	}
}

type BoundedFileManager struct {
	defaults    map[string]string
	environment map[string]string
}

func (this *BoundedFileManager) init() error {
	return nil
}

func (this *BoundedFileManager) setDefaults(defaults map[string]string) {
	for k, v := range defaults {
		this.defaults[k] = v
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

func (this *BoundedFileManager) resolveVar(name string) string {
	for k, v := range this.environment {
		if strings.Contains(name, "$"+k) {
			name = strings.ReplaceAll(name, "$"+k, v)
			name = this.resolvePath(name)
		}
	}

	return name
}

func (this *BoundedFileManager) resolvePath(path string) string {
	sep := string(filepath.Separator)
	comps := strings.Split(path, sep)

	for i := 0; i < len(comps); i++ {
		comps[i] = this.resolveVar(comps[i])
	}

	return strings.Join(comps, sep)
}

func (this *BoundedFileManager) pathExists(path string) (string, bool, bool, error) {
	path = this.resolvePath(path)
	info, err := os.Stat(path)

	switch {
	case err != nil && os.IsNotExist(err):
		return path, false, false, nil
	case err != nil && !os.IsNotExist(err):
		return path, false, false, err
	default:
		return path, true, info.IsDir(), nil
	}
}

func (this *BoundedFileManager) dirEnsure(path string) error {
	resolved, exists, isDir, err := this.pathExists(path)

	if err != nil {
		return err
	}

	switch {
	case exists && isDir:
		return nil
	case exists && !isDir:
		return fmt.Errorf("%s exists but is not directory", path)
	}

	if err := os.MkdirAll(resolved, os.FileMode(0755)); err != nil {
		return err
	}

	return nil
}
