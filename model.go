package bounded_xdg

// https://specifications.freedesktop.org/basedir-spec/latest/

type BoundedXdgConfig struct {
	Shells    []BoundedShellConfig `json:"shells"`
	Variables []BoundedVariable    `json:"variables"`
}

type BoundedShellConfig struct {
	Source       string            `json:"-"`
	Name         string            `json:"name"`
	BinaryName   string            `json:"binary-name"`
	ScriptExt    string            `json:"script-ext"`
	ConfigRoot   string            `json:"config-root"`
	TemplatePath string            `json:"template-path"`
	Homepage     string            `json:"homepage"`
	NotCompliant bool              `json:"not-compliant"`
	Variables    []BoundedVariable `json:"variables"`
	Ref          map[string]string `json:"ref"`
}

type BoundedVariable struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	XdgPath     string `json:"xdg-path"`
	LegacyPath  string `json:"legacy-path"`
	Disabled    bool   `json:"disabled"`
}

type BoundedProgramConfig struct {
	Source        string            `json:"-"`
	Name          string            `json:"name"`
	NotCompliant  bool              `json:"not-compliant"`
	StockDisabled bool              `json:"stock-disabled"`
	ConfigRoot    string            `json:"config-root"`
	Variables     []BoundedVariable `json:"variables"`
	Ref           struct {
		Homepage string `json:"homepage"`
		XdgNinja string `json:"xdg-ninja"`
	} `json:"ref"`
}

type BoundedGlobalConfig struct {
	EnvironmentDefaults map[string]string `json:"environment-defaults"`
}

type BoundededUserState struct {
	Shells []*BoundedShellState `json:"programs"`
}

type BoundedShellState struct {
	Name     string   `json:"name"`
	IsInPath bool     `json:"-"`
	Managed  bool     `json:"managed"`
	Programs []string `json:"programs"`
}

type BoundedProgramState struct {
	Name    string `json:"name"`
	Enabled bool   `json:"disabled"`
	Managed bool   `json:"managed"`
}
