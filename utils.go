package bounded_xdg

import (
	"fmt"
	"os"
)

func pathExists(path string) (bool, bool, error) {
	info, err := os.Stat(path)

	switch {
	case err != nil && os.IsNotExist(err):
		return false, false, nil
	case err != nil && !os.IsNotExist(err):
		return false, false, err
	default:
		return true, info.IsDir(), nil
	}
}

func dirEnsure(path string) error {
	exists, isDir, err := pathExists(path)

	if err != nil {
		return err
	}

	switch {
	case exists && isDir:
		return nil
	case exists && !isDir:
		return fmt.Errorf("%s exists but is not directory", path)
	}

	if err := os.MkdirAll(path, os.FileMode(0755)); err != nil {
		return err
	}

	return nil
}
