package glide

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/project"
)

func NewGlideManager() *GlideManager {
    goConfig := config.GoConfig{}

    return &GlideManager{
        Pfm: project.NewProjectFileManagerWithData("project/glide/glide.yml", templateData{
            PackageName: goConfig.GoPackageName(),
        }),
        GoConfig: goConfig,
    }
}

type GlideManager struct {
    Pfm      *project.ProjectFileManager
    GoConfig config.GoConfig
}

type templateData struct {
    PackageName string
}

func (this *GlideManager) Ensure() error {
    if err := this.Pfm.Ensure(); err != nil {
        return err
    }

    return nil
}
