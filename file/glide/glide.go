package glide

import (
    "github.com/boundedinfinity/devenv/config"
    "log"
    "github.com/boundedinfinity/devenv/file"
    "github.com/boundedinfinity/devenv/project"
    "path"
)

func NewGlideManager() *GlideManager {
    return &GlideManager{
        GlobalConfig: config.GlobalConfig{},
        FileConfig: config.FileConfig{},
        ProjectConfig: config.ProjectConfig{},
        DirConfig: config.DirConfig{},
        GoConfig: config.GoConfig{},
    }
}

type GlideManager struct {
    GlobalConfig  config.GlobalConfig
    ProjectConfig config.ProjectConfig
    DirConfig     config.DirConfig
    FileConfig    config.FileConfig
    GoConfig      config.GoConfig
}

type templateData struct {
    PackageName string
}

func (this *GlideManager) Write() error {
    pm := project.ProjectManager{
        GlobalConfig: this.GlobalConfig,
        ProjectConfig: this.ProjectConfig,
        DirConfig: this.DirConfig,
    }

    if err := pm.validate(); err != nil {
        return err
    }

    fm := file.FileManager{
        GlobalConfig: this.GlobalConfig,
        FileConfig: this.FileConfig,
        Path: path.Join(this.ProjectConfig.ProjectPath(), "glide.yaml"),
    }

    if err1 := fm.Validate(); err1 != nil {
        return err1
    }

    tm := file.TemplateManager{
        GlobalConfig: this.GlobalConfig,
        TemplatePath: "project/glide/glide.yaml",
        TemplateData: templateData{
            PackageName: this.GoConfig.GoPackageName(),
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
