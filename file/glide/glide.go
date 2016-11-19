package glide

import (
    "github.com/boundedinfinity/devenv/config"
    "log"
    "github.com/boundedinfinity/devenv/file"
)

func NewGlideManager() *GlideManager {
    return &GlideManager{
        GlobalConfig: config.GlobalConfig{},
        FileConfig: config.FileConfig{},
        GoConfig: config.GoConfig{},
    }
}

type GlideManager struct {
    GlobalConfig config.GlobalConfig
    FileConfig   config.FileConfig
    GoConfig     config.GoConfig
}

type templateData struct{
    PackageName string
}

func (this *GlideManager) Write() error {
    fm := file.FileManager{
        GlobalConfig: this.GlobalConfig,
        FileConfig: this.FileConfig,
        FileName: "glide.yml",
    }

    if err1 := fm.Validate(); err1 != nil {
        return err1
    }

    tm := file.TemplateManager{
        GlobalConfig: this.GlobalConfig,
        TemplatePath: "glide/glide.yml",
        TemplateData: templateData{},
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
