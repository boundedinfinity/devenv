package file

import (
    "path/filepath"
    "os"
    "fmt"
    "github.com/boundedinfinity/devenv/config"
    "io/ioutil"
    "errors"
    "log"
)

type FileManager struct {
    GlobalConfig   config.GlobalConfig
    FileConfig     config.FileConfig
    FileName       string
    absProjectPath string
    absFilePath    string
}

func (this *FileManager) Validate() error {
    if this.GlobalConfig.Debug() {
        log.Printf("projectPath: %s", this.FileConfig.ProjectPath())
    }

    absProjectPath, err1 := filepath.Abs(this.FileConfig.ProjectPath())

    if err1 != nil {
        return err1
    }

    this.absProjectPath = absProjectPath

    if this.GlobalConfig.Debug() {
        log.Printf("absProjectPath: %s", absProjectPath)
    }

    projectPathInfo, err2 := os.Stat(absProjectPath)

    if err2 != nil {
        return err2
    }

    if !projectPathInfo.IsDir() {
        return errors.New(fmt.Sprintf("%s must be directory", config.Flag_ProjectPath))
    }

    absFilePath := filepath.Join(absProjectPath, this.FileName)
    this.absFilePath = absFilePath

    if this.GlobalConfig.Debug() {
        log.Printf("absFilePath: %s", absFilePath)
    }

    _, absFileErr := os.Stat(absFilePath)

    if this.GlobalConfig.Debug() {
        log.Printf("absFileErr: %v", absFileErr)
    }

    if os.IsNotExist(absFileErr) {
        if this.GlobalConfig.Debug() {
            log.Printf("%s doesn't exists", this.FileName)
        }
    } else {
        if !this.FileConfig.Overwrite() {
            return errors.New(fmt.Sprintf("%s already exists", absFilePath))
        }
    }

    if this.GlobalConfig.Debug() {
        log.Printf("validated")
    }

    return nil
}

func (this *FileManager) Write(data []byte) error {
    log.Printf("Writing %s to %s", this.FileName, this.absFilePath)

    if err := ioutil.WriteFile(this.absFilePath, data, this.FileConfig.FileMode()); err != nil {
        return err
    }

    return nil
}
