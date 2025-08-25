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
	}
)

func NewFileManager() (*FileManager, error) {
	fm := &FileManager{
		environment: map[string]string{},
		data:        map[string]XdgFile{},
	}

	if err := fm.getVarOrErr("HOME"); err != nil {
		return nil, err
	}

	if err := fm.getVarOrErr("SHELL"); err != nil {
		return nil, err
	}

	fm.getVarOrDefault("XDG_CONFIG_HOME")
	fm.getVarOrDefault("XDG_DATA_HOME")
	fm.getVarOrDefault("XDG_STATE_HOME")
	fm.getVarOrDefault("XDG_CACHE_HOME")

	return fm, nil
}

type FileManager struct {
	environment map[string]string
	data        map[string]XdgFile
}

// ///////////////////////////////////////////////////////////////////////////
// Ensure Data
// ///////////////////////////////////////////////////////////////////////////

func (this *FileManager) ensureDir(path string) error {

	return nil
}

// ///////////////////////////////////////////////////////////////////////////
// Load Data
// ///////////////////////////////////////////////////////////////////////////

func (this *FileManager) getVarOrErr(name string) error {
	v := os.Getenv(name)

	if v == "" {
		return fmt.Errorf("environment variable %s not defined", name)
	}

	this.environment[name] = v

	return nil
}

func (this *FileManager) getVarOrDefault(name string) {
	v := os.Getenv(name)

	if v == "" {
		v = _XDG_DEFAULTS[name]
	}

	this.environment[name] = v
}

func (this *FileManager) LoadData() error {
	files, err := os.ReadDir("")

	if err != nil {
		return err
	}

	for _, file := range files {
		if err := this.loadData(file.Name()); err != nil {
			return err
		}
	}

	if err := this.validateData(); err != nil {
		return err
	}

	return nil
}

func (this *FileManager) loadData(path string) error {
	contents, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	var data XdgFile

	if err := json.Unmarshal(contents, &data); err != nil {
		return err
	}

	if _, ok := this.data[path]; ok {
		return fmt.Errorf("path alread loaded: %s", path)
	}

	this.data[path] = data

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
