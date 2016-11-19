package editorconfig

import (
    "github.com/boundedinfinity/devenv/config"
    "log"
    "github.com/boundedinfinity/devenv/file"
)

func NewEditorConfigManager() *EditorConfigManager {
    return &EditorConfigManager{
        GlobalConfig: config.GlobalConfig{},
        FileConfig: config.FileConfig{},
    }
}

type EditorConfigManager struct {
    GlobalConfig config.GlobalConfig
    FileConfig   config.FileConfig
}

type templateData struct{}

func (this *EditorConfigManager) Write() error {
    fm := file.FileManager{
        GlobalConfig: this.GlobalConfig,
        FileConfig: this.FileConfig,
        FileName: ".editorconfig",
    }

    if err1 := fm.Validate(); err1 != nil {
        return err1
    }

    tm := file.TemplateManager{
        GlobalConfig: this.GlobalConfig,
        TemplatePath: "editorconfig/.editorconfig",
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
