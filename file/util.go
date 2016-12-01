package file

import (
    "github.com/boundedinfinity/devenv/shell"
    "fmt"
    "path/filepath"
    "strings"
)

func CalcRealPath(input string, expandPath bool) (string, error) {
    realPath := input

    if expandPath {
        if strings.Contains(realPath, "$") {
            expandedPath, err := shell.Evaluate(fmt.Sprintf("echo -n %s", realPath))

            if err != nil {
                return realPath, err
            }

            realPath = expandedPath
        }
    }

    realPath, err := filepath.Abs(realPath)

    if err != nil {
        return realPath, err
    }

    return realPath, nil
}
