package project

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/boundedinfinity/devenv/file"
)

func NewGlideTemplateFile() (*file.TemplateFile, error) {
    return NewProjectTemplateFile(
        logging.ComponentLogger("GlideManager"),
        "project/glide/glide.yaml",
        glideTemplateData{
            PackageName: config.NewGlobalConfig().GoConfig.GoPackageName(),
        },
    )
}

type glideTemplateData struct {
    PackageName string
}
