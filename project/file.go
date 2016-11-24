package project

import (
    "github.com/boundedinfinity/devenv/file"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/boundedinfinity/devenv/config"
    "path/filepath"
    "fmt"
    "errors"
    "strings"
)

var pfmlogger = logging.ComponentLogger("ProjectFileManager")

func NewProjectFileManager(dataPath string) *ProjectFileManager {
    return NewProjectFileManagerWithData(dataPath, nil)
}

func NewProjectFileManagerWithData(dataPath string, data interface{}) *ProjectFileManager {
    return &ProjectFileManager{
        GlobalConfig: config.GlobalConfig{},
        Project: NewProjectManager(),
        TemplatePath: dataPath,
        TemplateData: data,
    }
}

type ProjectFileManager struct {
    GlobalConfig config.GlobalConfig
    Project      *ProjectManager
    File         *file.FileManager
    TemplatePath string
    TemplateData interface{}
}

func (this *ProjectFileManager) validate() error {
    if this.GlobalConfig.Debug() {
        pfmlogger.Infof("this.TemplatePath: %s", this.TemplatePath)
    }

    if !file.TemplateExists(this.TemplatePath) {
        return errors.New(fmt.Sprintf("template '%s' not found", this.TemplatePath))
    }

    if err := this.Project.validate(); err != nil {
        if strings.Contains(err.Error(), "no such file or directory") {
            if err2 := this.Project.EnsureDirectory(); err2 != nil {
                return err2
            }
        } else {
            return err
        }
    }

    filename := filepath.Base(this.TemplatePath)
    absPath := filepath.Join(this.Project.AbsPath, filename)

    if this.GlobalConfig.Debug() {
        pfmlogger.Infof("filename: %s", filename)
        pfmlogger.Infof("absPath: %s", absPath)
    }

    this.File = file.NewFileManager(absPath)

    if err := this.File.Validate(); err != nil {
        return err
    }

    return nil
}

func (this *ProjectFileManager) Ensure() error {
    if err := this.validate(); err != nil {
        return err
    }

    tm := file.NewTemplateManager(this.TemplatePath, this.TemplateData)
    data, err2 := tm.Render()

    if err2 != nil {
        return err2
    }

    if err3 := this.File.Write(data); err3 != nil {
        return err3
    }

    return nil
}
