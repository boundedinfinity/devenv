package bounded_xdg

// https://specifications.freedesktop.org/basedir-spec/latest/

type File struct {
}

type BoundedXdgConfig struct {
	Shells    []BoundedShellConfig `json:"shells"`
	Variables []BoundedVariable    `json:"variables"`
}

type BoundedShellConfig struct {
	Name          string `json:"name"`
	XdgConfigHome string `json:"xdg-config-home"`
}

type BoundedVariable struct {
	Variable XdgVariable `json:"home-path"`
	Ignore   string      `json:"ignore"`
}

type XdgFile struct {
	Name         string        `json:"name"`
	Filename     string        `json:"filename"`
	NotCompliant bool          `json:"not-compliant"`
	Variables    []XdgVariable `json:"variables"`
	Ref          []struct {
		Homepage string `json:"homepage"`
		XdgNinja string `json:"xdg-ninja"`
	} `json:"ref"`
}

type XdgVariable struct {
	HomePath     string `json:"home-path"`
	XdgPath      string `json:"xdg-path"`
	NotSupported string `json:"not-supported"`
}
