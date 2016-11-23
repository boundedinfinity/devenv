package project

import (
    "github.com/boundedinfinity/devenv/config"
    "path/filepath"
    "os"
    "fmt"
    "errors"
    "github.com/spf13/afero"
    "github.com/boundedinfinity/devenv/logging"
    "strings"
)

var pmroot = logging.ComponentLogger("ProjectManager")

func NewProjectManager() *ProjectManager {
    return &ProjectManager{
        GlobalConfig: config.GlobalConfig{},
        DirConfig: config.DirConfig{},
        ProjectConfig: config.ProjectConfig{},
    }
}

type ProjectManager struct {
    GlobalConfig  config.GlobalConfig
    DirConfig     config.DirConfig
    ProjectConfig config.ProjectConfig
    AbsPath       string
}

func (this *ProjectManager) path2abs() error {
    if this.AbsPath != "" {
        return nil
    }

    absPath, err1 := filepath.Abs(this.ProjectConfig.ProjectPath())

    if err1 != nil {
        return err1
    }

    this.AbsPath = absPath

    if this.GlobalConfig.Debug() {
        pmroot.Infof("Absolute Project Path: %s", absPath)
    }

    return nil
}

func (this *ProjectManager) validate() error {
    if this.GlobalConfig.Debug() {
        pmroot.Infof("Input Project Path: %s", this.ProjectConfig.ProjectPath())
    }

    if err := this.path2abs(); err != nil {
        return err
    }

    pathInfo, err2 := os.Stat(this.AbsPath)

    if err2 != nil {
        return err2
    }

    if !pathInfo.IsDir() {
        return errors.New(fmt.Sprintf("%s must be directory", config.Flag_ProjectPath))
    }

    return nil
}

func (this *ProjectManager) EnsureDirectory() error {
    if !this.GlobalConfig.Quiet() {
        pmroot.Infof("Creating %s [mode: %s]", this.AbsPath, this.DirConfig.FileMode())
    }

    if err := this.validate(); err != nil {
        if strings.Contains(err.Error(), "no such file or directory") {
            // Ok
        } else {
            return err
        }
    }

    fs := afero.NewOsFs()

    if err := fs.MkdirAll(this.AbsPath, this.DirConfig.FileMode()); err != nil {
        return err
    }

    return nil
}

func (this *ProjectManager) DeleteDirectory() error {
    if this.GlobalConfig.Debug() {
        pmroot.Infof("Input Project Path: %s", this.ProjectConfig.ProjectPath())
    }

    if err := this.path2abs(); err != nil {
        return err
    }

    if !this.GlobalConfig.Quiet() {
        pmroot.Infof("Deleting %s", this.AbsPath)
    }

    if err := this.validate(); err != nil {
        if strings.Contains(err.Error(), "no such file or directory") {
            return nil
        } else {
            return err
        }
    }

    fs := afero.NewOsFs()

    if err := fs.RemoveAll(this.AbsPath); err != nil {
        return err
    }

    return nil
}
