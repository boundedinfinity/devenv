package bounded_xdg

import "github.com/boundedinfinity/rfc3339date"

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
	BinaryName    string `json:"binary-name"`
	Homepage      string `json:"homepage"`
}

type BoundedVariable struct {
	XdgFile      string                      `json:"xdg-file"`
	UserIgnored  bool                        `json:"user-ignored"`
	LastModified rfc3339date.Rfc3339DateTime `json:"last-modified"`
}

type XdgFile struct {
	Name         string        `json:"name"`
	NotCompliant bool          `json:"not-compliant"`
	Variables    []XdgVariable `json:"variables"`
	Ref          []struct {
		Homepage string `json:"homepage"`
		XdgNinja string `json:"xdg-ninja"`
	} `json:"ref"`
}

type XdgVariable struct {
	HomePath string `json:"home-path"`
	XdgPath  string `json:"xdg-path"`
	IsFile   string `json:"is-file"`
}
