package project

import (
    "github.com/boundedinfinity/devenv/config"
)

func NewMakefileManager() *MakefileManager {
    projectConfig := config.ProjectConfig{}

    return &MakefileManager{
        Pfm: NewProjectFileManagerWithData("project/makefile/Makefile", makefileTemplateData{
            ProjectName: projectConfig.ProjectName(),
        }),
    }
}

type MakefileManager struct {
    Pfm *ProjectFileManager
}

type makefileTemplateData struct {
    ProjectName string
}

func (this *MakefileManager) Ensure() error {
    if err := this.Pfm.Ensure(); err != nil {
        return err
    }

    return nil
}
