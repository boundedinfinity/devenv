package project

import (
    "github.com/boundedinfinity/devenv/config"
    "log"
    "path/filepath"
    "os"
    "fmt"
    "errors"
    "github.com/spf13/afero"
)

type ProjectManager struct {
    GlobalConfig  config.GlobalConfig
    DirConfig     config.DirConfig
    ProjectConfig config.ProjectConfig
    absPath       string
}

func (this *ProjectManager) Validate() error {
    if this.GlobalConfig.Debug() {
        log.Printf("projectPath: %s", this.ProjectConfig.ProjectPath())
    }

    absPath, err1 := filepath.Abs(this.ProjectConfig.ProjectPath())

    if err1 != nil {
        return err1
    }

    this.absPath = absPath

    if this.GlobalConfig.Debug() {
        log.Printf("absProjectPath: %s", absPath)
    }

    pathInfo, err2 := os.Stat(absPath)

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
        log.Printf("Creating %s [mode: %s]", this.absPath, this.DirConfig.FileMode())
    }

    fs := afero.NewOsFs()

    if err := fs.MkdirAll(this.absPath, this.DirConfig.FileMode()); err != nil {
        return err
    }

    return nil
}
