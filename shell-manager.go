package bounded_xdg

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func newShellManager(fm *BoundedFileManager) *BoundedShellManager {
	return &BoundedShellManager{
		fsRoot:       "$BOUNDED_CONFIG/shells",
		embeddedRoot: "shells",
		configs:      map[string]*BoundedShellConfig{},
		fm:           fm,
	}
}

type BoundedShellManager struct {
	fsRoot       string
	embeddedRoot string
	configs      map[string]*BoundedShellConfig
	fm           *BoundedFileManager
}

func (this *BoundedShellManager) init() error {
	files, err := this.fm.embeddedReadDir(this.embeddedRoot)

	if err != nil {
		return err
	}

	for _, file := range files {
		name := file.Name()

		if filepath.Ext(name) != ".json" {
			continue
		}

		var config BoundedShellConfig

		if err := this.fm.embeddedUnmarshalFile(&config, this.embeddedRoot, name); err != nil {
			return err
		}

		config.Source = this.fm.resolvePath(this.embeddedRoot, name)

		this.configs[config.Name] = &config
		this.configs[strings.ToLower(config.Name)] = &config
	}

	return nil
}

func (this *BoundedShellManager) All() []*BoundedShellConfig {
	var configs []*BoundedShellConfig

	for _, config := range this.configs {
		configs = append(configs, config)
	}

	return configs
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

func (this *BoundedShellManager) IsShellInPath(config BoundedShellConfig) (bool, error) {
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

func (this *BoundedShellManager) CurrentConfig() (*BoundedShellConfig, error) {
	var config *BoundedShellConfig

	shell := os.Getenv("SHELL")

	if shell == "" {
		return config, errors.New("can't determine shell")
	}

	shell = filepath.Base(shell)
	return this.GetState(shell)
}

func (this *BoundedShellManager) GetState(shell string) (*BoundedShellConfig, error) {
	var config *BoundedShellConfig
	var ok bool

	if shell == "" {
		return config, errors.New("can't determine shell")
	}

	config, ok = this.configs[shell]

	if !ok {
		return config, fmt.Errorf("no config for shell %s", shell)
	}

	return config, nil
}

func (this *BoundedShellManager) getTemplate(config BoundedShellConfig) (*template.Template, error) {
	var path string

	if config.TemplatePath != "" {
		path = config.TemplatePath
	} else {
		path = strings.ToLower(config.Name) + ".tmpl"
	}

	data, err := this.fm.embeddedReadFile(this.embeddedRoot, path)

	if err != nil {
		return nil, err
	}

	content := string(data)
	tmpl, err := template.New(strings.ToLower(config.Name)).Funcs(funcMap).Parse(content)

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
