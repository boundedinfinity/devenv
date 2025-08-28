package bounded_xdg

// https://specifications.freedesktop.org/basedir-spec/latest/

type BoundedXdgDefaults struct {
	EnvironmentDefaults map[string]string `json:"environment-defaults"`
}

type BoundedXdgConfig struct {
	Shells    []BoundedShellConfig `json:"shells"`
	Variables []BoundedVariable    `json:"variables"`
}

type BoundedShellConfig struct {
	Source       string            `json:"-"`
	Name         string            `json:"name"`
	BinaryName   string            `json:"binary-name"`
	ScriptExt    string            `json:"script-ext"`
	TemplatePath string            `json:"template-path"`
	Homepage     string            `json:"homepage"`
	NotCompliant bool              `json:"not-compliant"`
	Variables    []BoundedVariable `json:"variables"`
	Ref          struct {
		Homepage string `json:"homepage"`
	} `json:"ref"`
}

type BoundedShellState struct {
	Config   BoundedShellConfig
	IsInPath bool
	Current  bool
}

type BoundedVariable struct {
	Name          string `json:"name"`
	XdgPath       string `json:"xdg-path"`
	DefaultPath   string `json:"default-path"`
	StockDisabled bool   `json:"stock-disabled"`
}

type BoundedProgramConfig struct {
	Source        string            `json:"-"`
	Name          string            `json:"name"`
	NotCompliant  bool              `json:"not-compliant"`
	StockDisabled bool              `json:"stock-disabled"`
	Variables     []BoundedVariable `json:"variables"`
	Ref           struct {
		Homepage string `json:"homepage"`
		XdgNinja string `json:"xdg-ninja"`
	} `json:"ref"`
}

type BoundedProgramState struct {
	Config BoundedProgramConfig
	Shells []BoundedProgramShell
}

type BoundedProgramShell struct {
	State     BoundedShellState
	Available bool
	Enabled   bool
}
