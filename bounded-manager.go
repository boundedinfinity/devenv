package bounded_xdg

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

	var defaults BoundedXdgDefaults

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

	bm := &BoundeManager{defaults: defaults, fm: fm, sm: sm, pm: pm}

	if err := bm.init(); err != nil {
		return nil, err
	}

	return bm, nil
}

type BoundeManager struct {
	defaults BoundedXdgDefaults
	sm       *BoundedShellManager
	fm       *BoundedFileManager
	pm       *BoundedProgramManager
}

// ///////////////////////////////////////////////////////////////////////////
// Load Data
// ///////////////////////////////////////////////////////////////////////////

func (this *BoundeManager) init() error {
	dirNames := append([]string{"BOUNDED_CONFIG"}, _XDG_VARIABLE_NAMES...)

	for _, name := range dirNames {
		if err := this.fm.fsEnsureDir("$" + name); err != nil {
			return err
		}
	}

	return nil
}

func (this *BoundeManager) Shells() []*BoundedShellState {
	return this.sm.States()
}

func (this *BoundeManager) Programs() []*BoundedProgramState {
	return this.pm.States()
}
