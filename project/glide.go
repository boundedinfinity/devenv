package project

import (
    "github.com/boundedinfinity/devenv/config"
)

func NewGlideManager() *GlideManager {
    return &GlideManager{
        Pfm: NewProjectFileManagerWithData("project/glide/glide.yaml", glideTemplateData{
            PackageName: config.NewGlobalConfig().GoConfig.GoPackageName(),
        }),
    }
}

type GlideManager struct {
    Pfm      *ProjectFileManager
}

type glideTemplateData struct {
    PackageName string
}

func (this *GlideManager) Ensure() error {
    if err := this.Pfm.Ensure(); err != nil {
        return err
    }

    return nil
}
