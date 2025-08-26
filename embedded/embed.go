package embedded

import (
	"embed"
	"encoding/json"
	"io/fs"
	"path/filepath"
)

//go:embed data/*
var _fs embed.FS

func ReadDir(elem ...string) ([]fs.DirEntry, error) {
	join := append([]string{"data"}, elem...)
	path := filepath.Join(join...)
	return _fs.ReadDir(path)
}

func ReadFile(elem ...string) ([]byte, error) {
	join := append([]string{"data"}, elem...)
	path := filepath.Join(join...)
	return _fs.ReadFile(path)
}

func UnmarshalFile(v any, elem ...string) error {
	data, err := ReadFile(elem...)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	return nil
}
