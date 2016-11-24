package project

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/spf13/afero"
    "path/filepath"
    "strings"
    "fmt"
    "errors"
    "os"
)

var pmroot = logging.ComponentLogger("ProjectManager")

func NewProjectManager() *ProjectManager {
    return &ProjectManager{
        GlobalConfig: config.NewGlobalConfig(),
    }
}

type ProjectManager struct {
    GlobalConfig  config.GlobalConfig
    AbsPath       string
}

func (this *ProjectManager) path2abs() error {
    if this.AbsPath != "" {
        return nil
    }

    absPath, err1 := filepath.Abs(this.GlobalConfig.ProjectConfig.ProjectPath())

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
        pmroot.Infof("Input Project Path: %s", this.GlobalConfig.ProjectConfig.ProjectPath())
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
        pmroot.Infof("Creating %s [mode: %s]", this.AbsPath, this.GlobalConfig.DirConfig.FileMode())
    }

    if err := this.validate(); err != nil {
        if strings.Contains(err.Error(), "no such file or directory") {
            // Ok
        } else {
            return err
        }
    }

    fs := afero.NewOsFs()

    if err := fs.MkdirAll(this.AbsPath, this.GlobalConfig.DirConfig.FileMode()); err != nil {
        return err
    }

    return nil
}

func (this *ProjectManager) DeleteDirectory() error {
    if this.GlobalConfig.Debug() {
        pmroot.Infof("Input Project Path: %s", this.GlobalConfig.ProjectConfig.ProjectPath())
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
