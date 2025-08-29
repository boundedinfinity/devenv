package bounded_xdg

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	_XDG_VARIABLE_NAMES = []string{
		"XDG_CONFIG_HOME",
		"XDG_DATA_HOME",
		"XDG_STATE_HOME",
		"XDG_CACHE_HOME",
	}

	_JSON_META = struct {
		Prefix string
		Indent string
	}{
		Prefix: "",
		Indent: "    ",
	}
)

func NewBoundeManager() (*BoundeManager, error) {
	fm := newFileManager()

	var defaults BoundedGlobalConfig

	if err := fm.embeddedUnmarshalFile(&defaults, "config/config.json"); err != nil {
		return nil, err
	}

	fm.setDefaults(defaults.EnvironmentDefaults)

	if err := fm.init(); err != nil {
		return nil, err
	}

	varNames := append([]string{"HOME", "SHELL", "BOUNDED_CONFIG"}, _XDG_VARIABLE_NAMES...)

	for _, name := range varNames {
		if err := fm.ensureVar(name); err != nil {
			return nil, err
		}
	}

	sm := newShellManager(fm)

	if err := sm.init(); err != nil {
		return nil, err
	}

	pm := newProgramManager(fm, sm)

	if err := pm.init(); err != nil {
		return nil, err
	}

	bm := &BoundeManager{config: defaults, fm: fm, sm: sm, pm: pm}

	if err := bm.init(); err != nil {
		return nil, err
	}

	return bm, nil
}

type BoundeManager struct {
	config BoundedGlobalConfig
	state  BoundededUserState
	sm     *BoundedShellManager
	fm     *BoundedFileManager
	pm     *BoundedProgramManager
}

// ///////////////////////////////////////////////////////////////////////////
// Load Data
// ///////////////////////////////////////////////////////////////////////////

func (this *BoundeManager) init() error {
	return nil
}

func (this *BoundeManager) Save() error {

	return nil
}

func (this *BoundeManager) find(shell string, program string) (*BoundedShellConfig, *BoundedProgramConfig, error) {
	shellConfig, err := this.sm.GetState(shell)

	if err != nil {
		return nil, nil, err
	}

	programConfig, err := this.pm.GetConfig(program)

	if err != nil {
		return nil, nil, err
	}

	return shellConfig, programConfig, nil
}

func (this *BoundeManager) Shells() []*BoundedShellState {
	shells := map[string]*BoundedShellState{}

	for _, shell := range this.state.Shells {
		shells[strings.ToLower(shell.Name)] = shell
	}

	for _, shell := range this.sm.All() {
		if _, ok := shells[strings.ToLower(shell.Name)]; !ok {
			shells[strings.ToLower(shell.Name)] = &BoundedShellState{
				Name: shell.Name,
			}
		}
	}

	var list []*BoundedShellState

	for _, shell := range shells {
		list = append(list, shell)
	}

	return list
}

func (this *BoundeManager) Programs(shell string) []*BoundedShellState {
	shells := map[string]*BoundedShellState{}

	for _, shell := range this.state.Shells {
		shells[strings.ToLower(shell.Name)] = shell
	}

	for _, shell := range this.sm.All() {
		if _, ok := shells[strings.ToLower(shell.Name)]; !ok {
			shells[strings.ToLower(shell.Name)] = &BoundedShellState{
				Name: shell.Name,
			}
		}
	}

	var list []*BoundedShellState

	for _, shell := range shells {
		list = append(list, shell)
	}

	return list
}

// func (this *BoundeManager) Available(shell string, program string, status bool) error {
// 	sstate, pstate, err := this.find(shell, program)

// 	if err != nil {
// 		return err
// 	}

// 	if err := this.pm.Available(sstate.Config, pstate.Config, status); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (this *BoundeManager) Enabled(shell string, program string, status bool) error {
// 	sstate, pstate, err := this.find(shell, program)

// 	if err != nil {
// 		return err
// 	}

// 	if err := this.pm.Enable(sstate.Config, pstate.Config, status); err != nil {
// 		return err
// 	}

// 	return nil
// }

func PP(v any) {
	if data, err := json.MarshalIndent(v, "", "    "); err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}
}

// func (this *BoundedShellManager) AddProgram(shell *BoundedShellState, program *BoundedProgramState) {
// 	var found *BoundedProgramState

// 	for _, working := range shell.Programs {
// 		if working.Config.Name == program.Config.Name {
// 			found = working
// 			break
// 		}
// 	}

// 	if found == nil {
// 		shell.Programs = append(shell.Programs, program)
// 	}
// }

// func (this *BoundedShellManager) RemoveProgram(shell *BoundedShellState, program *BoundedProgramState) {
// 	var programs []*BoundedProgramState

// 	for _, working := range shell.Programs {
// 		if working.Config.Name != program.Config.Name {
// 			programs = append(programs, working)
// 		}
// 	}

// 	shell.Programs = programs
// }
