package bounded_xdg

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func NewShellManager(fm *BoundedFileManager) *BoundedShellManager {
	return &BoundedShellManager{
		fsRoot:          "$BOUNDED_CONFIG/shells",
		embeddedRoot:    "shells",
		internalConfigs: map[string]BoundedShellConfig{},
		fm:              fm,
	}
}

type BoundedShellManager struct {
	fsRoot          string
	embeddedRoot    string
	internalConfigs map[string]BoundedShellConfig
	fm              *BoundedFileManager
}

func (this *BoundedShellManager) init() error {
	if err := this.loadInternal(); err != nil {
		return err
	}

	return nil
}

func (this *BoundedShellManager) loadInternal() error {
	files, err := embedded.ReadDir(this.embeddedRoot)

	if err != nil {
		return err
	}

	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && filepath.Ext(name) == ".json" {
			var config BoundedShellConfig
			if err := this.fm.embeddedUnmarshalFile(&config, this.embeddedRoot, name); err != nil {
				return err
			}

			this.internalConfigs[config.Name] = config
		}
	}

	return nil
}

func (this *BoundedShellManager) WriteConfig(config BoundedShellConfig) error {
	if err := this.fm.fsEnsureDir(this.fsRoot); err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.json", config.Name)
	path := filepath.Join(this.fsRoot, filename)

	if err := this.fm.fsMarshalFile(config, path); err != nil {
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
