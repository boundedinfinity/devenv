package bounded_xdg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/boundedinfinity/bounded_xdg/shells"
)

func NewShellManager() *BoundedShellManager {
	return &BoundedShellManager{
		internalConfigs: map[string]BoundedShellConfig{},
	}
}

type BoundedShellManager struct {
	internalConfigs map[string]BoundedShellConfig
}

func (this *BoundedShellManager) init() error {
	if err := this.loadInternal(); err != nil {
		return err
	}

	return nil
}

func (this *BoundedShellManager) loadInternal() error {
	files, err := shells.FS.ReadDir("data")

	if err != nil {
		return err
	}

	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && filepath.Ext(name) == ".json" {
			path := filepath.Join("data", name)
			data, err := shells.FS.ReadFile(path)

			if err != nil {
				return err
			}

			var config BoundedShellConfig

			if err := json.Unmarshal(data, &config); err != nil {
				return err
			}

			this.internalConfigs[config.Name] = config
		}
	}

	return nil
}

func (this *BoundedShellManager) EnsureConfigDir(config BoundedShellConfig) error {
	if err := dirEnsure(config.XdgConfigHome); err != nil {
		return err
	}

	return nil
}

func (this *BoundedShellManager) IsInPath(config BoundedShellConfig) (bool, error) {
	var found bool
	name := config.Name

	if config.BinaryName != "" {
		name = config.BinaryName
	}

	abs, err := exec.LookPath(name)

	if err != nil {
		return found, err
	}

	if abs != "" {
		found = true
	}

	return found, nil
}

func (this *BoundedShellManager) CurrentConfig() (BoundedShellConfig, error) {
	var config BoundedShellConfig

	shell := os.Getenv("SHELL")

	if shell == "" {
		return config, errors.New("can't determine shell")
	}

	shell = filepath.Base(shell)

	return this.GetConfig(shell)
}

func (this *BoundedShellManager) GetConfig(shell string) (BoundedShellConfig, error) {
	var config BoundedShellConfig
	var ok bool

	if shell == "" {
		return config, errors.New("can't determine shell")
	}

	shell = filepath.Base(shell)
	config, ok = this.internalConfigs[shell]

	if !ok {
		return config, fmt.Errorf("no config for shell %s", shell)
	}

	return config, nil
}
