package project

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/Sirupsen/logrus"
)

func NewGlideManager() *GlideManager {
    return &GlideManager{
        Path: "project/glide/glide.yaml",
        Data: glideTemplateData{
            PackageName: config.NewGlobalConfig().GoConfig.GoPackageName(),
        },
    }
}

type GlideManager struct {
    logger *logrus.Entry
    Path   string
    Data   glideTemplateData
}

type glideTemplateData struct {
    PackageName string
}

func (this *GlideManager) Ensure() error {
    manager := NewProjectDirectoryManager(this.logger)

    if err := manager.EnsureFile(this.Path, this.Data); err != nil {
        return err
    }

    return nil
}
