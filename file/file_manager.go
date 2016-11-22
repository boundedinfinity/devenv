package file

import (
    "path/filepath"
    "os"
    "fmt"
    "github.com/boundedinfinity/devenv/config"
    "errors"
    "log"
    "github.com/spf13/afero"
)

type FileManager struct {
    GlobalConfig config.GlobalConfig
    FileConfig   config.FileConfig
    Path         string
    absFilePath  string
    absDirPath   string
}

func (this *FileManager) Validate() error {
    if this.GlobalConfig.Debug() {
        log.Printf("Path: %s", this.Path)
    }

    aaaa, err1 := filepath.Abs(this.Path)

    if err1 != nil {
        return err1
    }

    this.absFilePath = filepath.Dir(aaaa)

    if this.GlobalConfig.Debug() {
        log.Printf("absFilePath: %s", this.absFilePath)
    }

    this.absDirPath = filepath.Dir(this.absFilePath)

    if this.GlobalConfig.Debug() {
        log.Printf("absDirPath: %s", this.absDirPath)
    }

    _, absFileErr := os.Stat(this.absFilePath)

    if this.GlobalConfig.Debug() {
        log.Printf("absFileErr: %v", absFileErr)
    }

    if os.IsNotExist(absFileErr) {
        if this.GlobalConfig.Debug() {
            log.Printf("%s doesn't exists", this.absFilePath)
        }
    } else {
        if !this.FileConfig.Overwrite() {
            return errors.New(fmt.Sprintf("%s already exists", this.absFilePath))
        }
    }

    if this.GlobalConfig.Debug() {
        log.Printf("validated")
    }

    return nil
}

func (this *FileManager) Write(data []byte) error {
    fs := afero.NewOsFs()

    dirExists, err := afero.DirExists(fs, this.absDirPath)

    if err != nil {
        return err
    }

    if !dirExists {
        if err := os.MkdirAll(this.absDirPath, this.FileConfig.FileMode()); err != nil {
            return err
        }
    }

    if !this.GlobalConfig.Quiet() {
        log.Printf("Writing %s [mode: %s]", this.absFilePath, this.FileConfig.FileMode())
    }

    if err := afero.WriteFile(fs, this.absFilePath, data, this.FileConfig.FileMode()); err != nil {
        return err
    }

    return nil
}
