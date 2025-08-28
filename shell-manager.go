package bounded_xdg

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
)

func newShellManager(fm *BoundedFileManager) *BoundedShellManager {
	return &BoundedShellManager{
		fsRoot:       "$BOUNDED_CONFIG/shells",
		embeddedRoot: "shells",
		states:       map[string]*BoundedShellState{},
		fm:           fm,
	}
}

type BoundedShellManager struct {
	fsRoot       string
	embeddedRoot string
	states       map[string]*BoundedShellState
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

		var state BoundedShellState

		if err := this.fm.embeddedUnmarshalFile(&state.Config, this.embeddedRoot, name); err != nil {
			return err
		}

		inPath, err := this.IsShellInPath(state.Config)

		if err != nil {
			return err
		}

		state.IsInPath = inPath
		state.Config.Source = this.fm.resolvePath(this.embeddedRoot, name)

		this.states[state.Config.Name] = &state
	}

	current, err := this.CurrentConfig()

	if err != nil {
		return err
	}

	for _, state := range this.states {
		if state.Config.Name == current.Config.Name {
			state.Current = true
		}
	}

	return nil
}

func (this *BoundedShellManager) States() []*BoundedShellState {
	var states []*BoundedShellState

	for _, state := range this.states {
		states = append(states, state)
	}

	return states
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

func (this *BoundedShellManager) CurrentConfig() (BoundedShellState, error) {
	var state BoundedShellState

	shell := os.Getenv("SHELL")

	if shell == "" {
		return state, errors.New("can't determine shell")
	}

	shell = filepath.Base(shell)

	return this.GetConfig(shell)
}

func (this *BoundedShellManager) GetConfig(shell string) (BoundedShellState, error) {
	var state *BoundedShellState
	var ok bool

	if shell == "" {
		return *state, errors.New("can't determine shell")
	}

	shell = filepath.Base(shell)
	state, ok = this.states[shell]

	if !ok {
		return *state, fmt.Errorf("no config for shell %s", shell)
	}

	return *state, nil
}

func (this *BoundedShellManager) GetTemplate(config BoundedShellConfig) (*template.Template, error) {
	var path string

	if config.TemplatePath != "" {
		path = config.TemplatePath
	} else {
		path = config.Name + ".tmpl"
	}

	data, err := this.fm.embeddedReadFile(this.embeddedRoot, path)

	if err != nil {
		return nil, err
	}

	content := string(data)
	tmpl, err := template.New(config.Name).Parse(content)

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
