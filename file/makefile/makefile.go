package makefile

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/file"
    "log"
    "github.com/boundedinfinity/devenv/project"
    "path"
)

func NewMakefileManager() *MakefileManager {
    return &MakefileManager{
        GlobalConfig: config.GlobalConfig{},
        ProjectConfig: config.ProjectConfig{},
        DirConfig: config.DirConfig{},
        FileConfig: config.FileConfig{},
    }
}

type MakefileManager struct {
    GlobalConfig  config.GlobalConfig
    ProjectConfig config.ProjectConfig
    DirConfig     config.DirConfig
    FileConfig    config.FileConfig
}

type templateData struct {
    ProjectName string
}

func (this *MakefileManager) Write() error {
    pm := project.ProjectManager{
        GlobalConfig: this.GlobalConfig,
        ProjectConfig: this.ProjectConfig,
        DirConfig: this.DirConfig,
    }

    if err := pm.Validate(); err != nil {
        return err
    }

    fm := file.FileManager{
        GlobalConfig: this.GlobalConfig,
        FileConfig: this.FileConfig,
        Path: path.Join(this.ProjectConfig.ProjectPath(), "Makefile"),
    }

    if err1 := fm.Validate(); err1 != nil {
        return err1
    }

    tm := file.TemplateManager{
        GlobalConfig: this.GlobalConfig,
        TemplatePath: "project/makefile/Makefile",
        TemplateData: templateData{
            ProjectName: this.ProjectConfig.ProjectName(),
        },
    }

    data, err2 := tm.Render()

    if err2 != nil {
        return err2
    }

    if this.GlobalConfig.Debug() {
        log.Printf("rendered")
    }

    if err3 := fm.Write(data); err3 != nil {
        return err3
    }

    if this.GlobalConfig.Debug() {
        log.Printf("written")
    }

    return nil
}
