package bounded_xdg

// https://specifications.freedesktop.org/basedir-spec/latest/

type File struct {
}

type FileXdgConfig struct {
	From string
	To   string
}

var (
	XDG_MAP = map[string]string{
		"XDG_CONFIG_HOME": "$HOME/.config",
		"XDG_DATA_HOME":   "$HOME/.local/share",
		"XDG_STATE_HOME":  "$HOME/.local/state",
		"XDG_CACHE_HOME":  "$HOME/.cache",
	}
)
