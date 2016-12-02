package project

import (
    "github.com/boundedinfinity/devenv/logging"
    "github.com/boundedinfinity/devenv/file"
)

func NewEditorConfigTemplateFile() (*file.TemplateFile, error) {
    return NewProjectTemplateFile(
        logging.ComponentLogger("EditorConfigTManager"),
        "project/editorconfig/.editorconfig",
        nil,
    )
}
