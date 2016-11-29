package file

import (
    "github.com/boundedinfinity/devenv/shell"
    "github.com/boundedinfinity/devenv/data"
    "fmt"
    "path/filepath"
    "errors"
    "strings"
)

func TemplateExists(path string) bool {
    found := false

    for _, currentPath := range data.AssetNames() {
        if currentPath == path {
            found = true
            break
        }
    }

    return found
}

func Template2FsPath(fsPath string, templatePath string) (string, error) {
    var tempPath string

    if !TemplateExists(templatePath) {
        return fsPath, errors.New(fmt.Sprintf("template '%s' not found", templatePath))
    }

    filename := filepath.Base(templatePath)
    tempPath = filepath.Join(fsPath, filename)

    return tempPath, nil
}

func CalcRealPath(input string, expandPath bool) (string, error) {
    realPath := input

    if expandPath && strings.Contains(realPath, "$") {
        expandedPath, err := shell.Evaluate(fmt.Sprintf("echo -n %s", realPath))

        if err != nil {
            return realPath, err
        }

        realPath = expandedPath
    }

    realPath, err := filepath.Abs(realPath)

    if err != nil {
        return realPath, err
    }

    return realPath, nil
}
