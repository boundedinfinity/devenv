package bounded_xdg

import (
	"path/filepath"
	"strings"
)

func newProgramManager(fm *BoundedFileManager, sm *BoundedShellManager) *BoundedProgramManager {
	return &BoundedProgramManager{
		availableRoot: "$BOUNDED_CONFIG/programs/available",
		enabledRoot:   "$BOUNDED_CONFIG/programs/enabled",
		embeddedRoot:  "programs",
		fm:            fm,
		sm:            sm,
		states:        map[string]*BoundedProgramState{},
	}
}

type BoundedProgramManager struct {
	availableRoot string
	enabledRoot   string
	embeddedRoot  string
	fm            *BoundedFileManager
	sm            *BoundedShellManager
	states        map[string]*BoundedProgramState
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

		var state BoundedProgramState
		path := this.fm.resolvePath(this.embeddedRoot, name)

		if err := this.fm.embeddedUnmarshalFile(&state.Config, path); err != nil {
			return err
		}

		state.Config.Source = path
		this.states[state.Config.Name] = &state
	}

	for _, programState := range this.states {
		for _, shellState := range this.sm.States() {
			availableExists, err := this.IsAvailable(shellState.Config, programState.Config)

			if err != nil {
				return err
			}

			enabledExists, err := this.IsEnabled(shellState.Config, programState.Config)

			if err != nil {
				return err
			}

			programShellState := BoundedProgramShell{
				State:     *shellState,
				Available: availableExists.Exists,
				Enabled:   enabledExists.Exists,
			}

			programState.Shells = append(programState.Shells, programShellState)
		}
	}

	return nil
}

func (this *BoundedProgramManager) States() []*BoundedProgramState {
	var states []*BoundedProgramState

	for _, state := range this.states {
		states = append(states, state)
	}

	return states
}

func (this *BoundedProgramManager) Available(shell BoundedShellConfig, config BoundedProgramConfig, on bool) error {
	tmpl, err := this.sm.GetTemplate(shell)

	if err != nil {
		return err
	}

	writer := new(strings.Builder)

	if err := tmpl.Funcs(funcMap).Execute(writer, config); err != nil {
		return err
	}

	content := writer.String()
	filename := config.Name + "." + shell.ScriptExt

	if err := this.fm.fsWriteFile([]byte(content), this.availableRoot, filename); err != nil {
		return err
	}

	return nil
}

func (this *BoundedProgramManager) IsAvailable(shell BoundedShellConfig, config BoundedProgramConfig) (PathExistsResult, error) {
	filename := config.Name + "." + shell.ScriptExt
	return this.fm.fsExists(this.availableRoot, filename)
}

func (this *BoundedProgramManager) Enable(shell BoundedShellConfig, config BoundedProgramConfig, on bool) error {
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
		this.fm.resolvePath(this.availableRoot, filename),
		this.fm.resolvePath(this.enabledRoot, filename),
	); err != nil {
		return err
	}

	return nil
}

func (this *BoundedProgramManager) IsEnabled(shell BoundedShellConfig, config BoundedProgramConfig) (PathExistsResult, error) {
	filename := config.Name + "." + shell.ScriptExt
	return this.fm.fsExists(this.enabledRoot, filename)
}
