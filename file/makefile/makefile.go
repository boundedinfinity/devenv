package makefile

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/file"
    "log"
)

func NewMakefileManager() *MakefileManager {
    return &MakefileManager{
        GlobalConfig: config.GlobalConfig{},
        ProjectConfig: config.ProjectConfig{},
        FileConfig: config.FileConfig{},
    }
}

type MakefileManager struct {
    GlobalConfig  config.GlobalConfig
    ProjectConfig config.ProjectConfig
    FileConfig    config.FileConfig
}

type templateData struct {
    ProjectName string
}

func (this *MakefileManager) Write() error {
    fm := file.FileManager{
        GlobalConfig: this.GlobalConfig,
        FileConfig: this.FileConfig,
        FileName: "Makefile",
    }

    if err1 := fm.Validate(); err1 != nil {
        return err1
    }

    tm := file.TemplateManager{
        GlobalConfig: this.GlobalConfig,
        TemplatePath: "makefile/Makefile",
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
