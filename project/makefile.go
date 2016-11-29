package project

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/Sirupsen/logrus"
)

func NewMakefileManager() *MakefileManager {
    return &MakefileManager{
        logger : logging.ComponentLogger("MakeFileManager"),
        Path: "project/makefile/Makefile",
        Data: makefileTemplateData{
            ProjectName: config.NewGlobalConfig().ProjectConfig.ProjectName(),
        },
    }
}

type MakefileManager struct {
    logger *logrus.Entry
    Path   string
    Data   makefileTemplateData
}

type makefileTemplateData struct {
    ProjectName string
}

func (this *MakefileManager) Ensure() error {
    manager := NewProjectDirectoryManager(this.logger)

    if err := manager.EnsureFile(this.Path, this.Data); err != nil {
        return err
    }

    return nil
}
