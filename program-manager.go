package bounded_xdg

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

func newProgramManager(fm *BoundedFileManager, sm *BoundedShellManager) *BoundedProgramManager {
	return &BoundedProgramManager{
		availablePath: "available",
		enabledPath:   "enabled",
		embeddedRoot:  "programs",
		configRoot:    "$BOUNDED_CONFIG",
		fm:            fm,
		sm:            sm,
		states:        map[string]*BoundedProgramConfig{},
	}
}

type BoundedProgramManager struct {
	availablePath string
	enabledPath   string
	embeddedRoot  string
	configRoot    string
	fm            *BoundedFileManager
	sm            *BoundedShellManager
	states        map[string]*BoundedProgramConfig
}

func (this *BoundedProgramManager) init() error {
	embeddedFiles, err := this.fm.embeddedReadDir(this.embeddedRoot)

	if err != nil {
		return err
	}

	for _, file := range embeddedFiles {
		name := file.Name()

		if filepath.Ext(name) != ".json" {
			continue
		}

		var config BoundedProgramConfig
		path := this.fm.resolvePath(this.embeddedRoot, name)

		if err := this.fm.embeddedUnmarshalFile(&config, path); err != nil {
			return err
		}

		config.Source = path
		this.states[config.Name] = &config
		this.states[strings.ToLower(config.Name)] = &config
	}

	return nil
}

func (this *BoundedProgramManager) States() []*BoundedProgramConfig {
	var states []*BoundedProgramConfig

	for _, state := range this.states {
		states = append(states, state)
	}

	return states
}

func (this *BoundedProgramManager) availableDir(shell BoundedShellConfig) string {
	return this.fm.resolvePath(this.configRoot, shell.Name, this.availablePath)
}

func (this *BoundedProgramManager) availableJsonFile(shell BoundedShellConfig, config BoundedProgramConfig) string {
	return this.fm.resolvePath(this.configRoot, shell.Name, this.availablePath, strings.ToLower(config.Name)+".json")
}

func (this *BoundedProgramManager) availableShellFile(shell BoundedShellConfig, config BoundedProgramConfig) string {
	return this.fm.resolvePath(this.configRoot, shell.Name, this.availablePath, strings.ToLower(config.Name)+"."+shell.ScriptExt)
}

func (this *BoundedProgramManager) GetConfig(program string) (*BoundedProgramConfig, error) {
	var state *BoundedProgramConfig
	var ok bool

	if program == "" {
		return state, errors.New("can't determine shell")
	}

	state, ok = this.states[program]

	if !ok {
		return state, fmt.Errorf("no config for shell %s", program)
	}

	return state, nil
}

func (this *BoundedProgramManager) Available(shell BoundedShellConfig, config BoundedProgramConfig, on bool) error {
	availableDir := this.availableDir(shell)

	if err := this.fm.fsEnsureDir(availableDir); err != nil {
		return err
	}

	availableJsonFile := this.availableJsonFile(shell, config)
	configExists, err := this.fm.fsExists(availableJsonFile)

	if err != nil {
		return err
	}

	if !configExists.Exists {
		if err := this.fm.fsMarshalFile(config, availableJsonFile); err != nil {
			return err
		}
	}

	tmpl, err := this.sm.getTemplate(shell)

	if err != nil {
		return err
	}

	writer := new(strings.Builder)

	if err := tmpl.Execute(writer, config); err != nil {
		return err
	}

	content := writer.String()
	availableShellFile := this.availableShellFile(shell, config)

	if err := this.fm.fsWriteFile([]byte(content), availableShellFile); err != nil {
		return err
	}

	return nil
}

func (this *BoundedProgramManager) IsAvailable(shell BoundedShellConfig, config BoundedProgramConfig) (PathExistsResult, error) {
	filename := this.availableShellFile(shell, config)
	return this.fm.fsExists(filename)
}

func (this *BoundedProgramManager) Enable(shell BoundedShellConfig, config BoundedProgramConfig, on bool) error {
	if on {
		if err := this.fm.fsEnsureDir(shell.ConfigRoot, this.enabledPath); err != nil {
			return err
		}

		availableExists, err := this.IsAvailable(shell, config)

		if err != nil {
			return err
		}

		if !availableExists.Exists {
			if err := this.Available(shell, config, true); err != nil {
				return err
			}
		}

		filename := config.Name + "." + shell.ScriptExt

		if err := this.fm.fsSymLink(
			this.fm.resolvePath(shell.ConfigRoot, this.availablePath, filename),
			this.fm.resolvePath(shell.ConfigRoot, this.enabledPath, filename),
		); err != nil {
			return err
		}

	} else {

	}

	return nil
}

func (this *BoundedProgramManager) IsEnabled(shell BoundedShellConfig, config BoundedProgramConfig) (PathExistsResult, error) {
	filename := config.Name + "." + shell.ScriptExt
	return this.fm.fsExists(shell.ConfigRoot, this.enabledPath, filename)
}
