package bounded_xdg

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	_XDG_DEFAULTS = map[string]string{
		"XDG_CONFIG_HOME": "$HOME/.config",
		"XDG_DATA_HOME":   "$HOME/.local/share",
		"XDG_STATE_HOME":  "$HOME/.local/state",
		"XDG_CACHE_HOME":  "$HOME/.cache",
		"BOUNDED_CONFIG":  "$XDG_CONFIG_HOME/bounded-xdg",
	}

	_XDG_VARIABLE_NAMES = []string{
		"XDG_CONFIG_HOME",
		"XDG_DATA_HOME",
		"XDG_STATE_HOME",
		"XDG_CACHE_HOME",
	}
)

func NewFileManager() *FileManager {
	return &FileManager{
		environment: map[string]string{},
		data:        map[string]XdgFile{},
	}
}

type FileManager struct {
	environment map[string]string
	data        map[string]XdgFile
}

// ///////////////////////////////////////////////////////////////////////////
// Utilities
// ///////////////////////////////////////////////////////////////////////////

func (this *FileManager) getVar(name string) string {
	path := this.environment[name]
	path = this.resolvePath(path)
	return path
}

func (this *FileManager) ensureVar(name string) error {
	v := os.Getenv(name)

	if found, ok := _XDG_DEFAULTS[name]; v == "" && ok {
		v = found
	}

	if v == "" {
		return fmt.Errorf("environment variable %s not defined", name)
	}

	this.environment[name] = v
	return nil
}

func (this *FileManager) ensureDir(path string) error {
	resolved := this.resolvePath(path)
	info, err := os.Stat(resolved)

	switch {
	case err == nil && !info.IsDir():
		return fmt.Errorf("%s exists but isn't a directory", resolved)
	case os.IsNotExist(err):
		return err
	case err != nil:
		return err
	}

	if err := os.Mkdir(resolved, os.FileMode(0755)); err != nil {
		return err
	}

	return nil
}

func (this *FileManager) resolvePath(path string) string {
	resolve := path

	for _, part := range strings.Split(path, "/") {
		if strings.HasPrefix(part, "$") {
			name := strings.ReplaceAll(part, "$", "")
			if found, ok := this.environment[name]; !ok {
				resolve = strings.ReplaceAll(path, part, found)
			}
		}
	}

	return resolve
}

// ///////////////////////////////////////////////////////////////////////////
// Load Data
// ///////////////////////////////////////////////////////////////////////////

func (this *FileManager) Init() error {
	varNames := []string{"HOME", "SHELL"}
	varNames = append(varNames, _XDG_VARIABLE_NAMES...)

	for _, name := range varNames {
		if err := this.ensureVar(name); err != nil {
			return err
		}
	}

	dirNames := []string{"BOUNDED_XDG_CONFIG"}
	dirNames = append(dirNames, _XDG_VARIABLE_NAMES...)

	for _, name := range dirNames {
		path := this.environment[name]

		if err := this.ensureDir(path); err != nil {
			return err
		}
	}

	if err := this.loadData(); err != nil {
		return err
	}

	if err := this.validateData(); err != nil {
		return err
	}

	return nil
}

func (this *FileManager) loadConfig() error {
	return nil
}

func (this *FileManager) loadData() error {
	files, err := os.ReadDir("")

	if err != nil {
		return err
	}

	for _, file := range files {
		name := file.Name()
		contents, err := os.ReadFile(name)

		if err != nil {
			return err
		}

		var data XdgFile

		if err := json.Unmarshal(contents, &data); err != nil {
			return err
		}

		if _, ok := this.data[name]; ok {
			return fmt.Errorf("path alread loaded: %s", name)
		}

		this.data[name] = data
	}

	return nil
}

func (this *FileManager) validateData() error {
	validate := func(path string, variable string, value string) error {
		for _, part := range strings.Split(value, "/") {
			if strings.HasPrefix(part, "$") {
				name := strings.ReplaceAll(part, "$", "")
				if _, ok := this.environment[name]; !ok {
					return fmt.Errorf(
						"%s.%s=%s: %s not found",
						path, variable, value, part,
					)
				}
			}
		}
		return nil
	}

	for path, data := range this.data {
		for _, variable := range data.Variables {
			if err := validate(path, "home-path", variable.HomePath); err != nil {
				return err
			}

			if err := validate(path, "xdg-path", variable.XdgPath); err != nil {
				return err
			}
		}
	}

	return nil
}
