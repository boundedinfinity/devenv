package bounded_xdg

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/boundedinfinity/bounded_xdg/embedded"
)

var (
	_XDG_VARIABLE_NAMES = []string{
		"XDG_CONFIG_HOME",
		"XDG_DATA_HOME",
		"XDG_STATE_HOME",
		"XDG_CACHE_HOME",
	}
)

func NewBoundeManager() (*BoundeManager, error) {
	fm := NewFileManager()
	bm := &BoundeManager{
		data: map[string]XdgFile{},
		sm:   NewShellManager(fm),
		fm:   fm,
	}

	if err := bm.init(); err != nil {
		return nil, err
	}

	return bm, nil
}

type BoundeManager struct {
	defaults BoundedXdgDefaults
	data     map[string]XdgFile
	sm       *BoundedShellManager
	fm       *BoundedFileManager
}

// ///////////////////////////////////////////////////////////////////////////
// Load Data
// ///////////////////////////////////////////////////////////////////////////

func (this *BoundeManager) init() error {
	if err := this.sm.init(); err != nil {
		return err
	}

	if err := this.fm.init(); err != nil {
		return err
	}

	if err := embedded.UnmarshalFile(&this.defaults, "defaults/config.json"); err != nil {
		return err
	}

	this.fm.setDefaults(this.defaults.EnvironmentDefaults)

	varNames := append([]string{"HOME", "SHELL", "BOUNDED_CONFIG"}, _XDG_VARIABLE_NAMES...)

	for _, name := range varNames {
		if err := this.fm.ensureVar(name); err != nil {
			return err
		}
	}

	dirNames := append([]string{"BOUNDED_CONFIG"}, _XDG_VARIABLE_NAMES...)

	for _, name := range dirNames {
		if err := this.fm.dirEnsure("$" + name); err != nil {
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

func (this *BoundeManager) loadConfig() error {
	return nil
}

func (this *BoundeManager) loadData() error {
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

func (this *BoundeManager) validateData() error {
	validate := func(path string, variable string, value string) error {
		for _, part := range strings.Split(value, "/") {
			if strings.HasPrefix(part, "$") {
				name := strings.ReplaceAll(part, "$", "")
				if _, ok := this.fm.environment[name]; !ok {
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
