package project

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/boundedinfinity/devenv/file"
)

func NewMakefileTemplateFile() (*file.TemplateFile, error) {
    return NewProjectTemplateFile(
        logging.ComponentLogger("MakeFileManager"),
        "project/makefile/Makefile",
        makefileTemplateData{
            ProjectName: config.NewGlobalConfig().ProjectConfig.ProjectName(),
        },
    )
}

type makefileTemplateData struct {
    ProjectName string
}
