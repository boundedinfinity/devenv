package makefile

import (
    "github.com/boundedinfinity/devenv/project"
    "github.com/boundedinfinity/devenv/config"
)

func NewMakefileManager() *MakefileManager {
    projectConfig := config.ProjectConfig{}

    return &MakefileManager{
        Pfm: project.NewProjectFileManagerWithData("project/makefile/Makefile", templateData{
            ProjectName: projectConfig.ProjectName(),
        }),
    }
}

type MakefileManager struct {
    Pfm *project.ProjectFileManager
}

type templateData struct {
    ProjectName string
}

func (this *MakefileManager) Ensure() error {
    if err := this.Pfm.Validate(); err != nil {
        return err
    }

    if err := this.Pfm.Ensure(); err != nil {
        return err
    }

    return nil
}
