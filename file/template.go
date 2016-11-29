package file

import "github.com/boundedinfinity/devenv/data"

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
