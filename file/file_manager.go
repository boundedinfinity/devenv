package file

import (
    "path/filepath"
    "os"
    "fmt"
    "github.com/boundedinfinity/devenv/config"
    "errors"
    "github.com/spf13/afero"
    "github.com/boundedinfinity/devenv/logging"
)

var logger = logging.ComponentLogger("FileManager")

func NewFileManager(path string) *FileManager {
    return &FileManager{
        GlobalConfig: config.NewGlobalConfig(),
        Path: path,
    }
}

type FileManager struct {
    GlobalConfig config.GlobalConfig
    Path         string
    absFilePath  string
    absDirPath   string
}

func (this *FileManager) Validate() error {
    if this.GlobalConfig.Debug() {
        logger.Infof("Path: %s", this.Path)
    }

    absFilePath, err1 := filepath.Abs(this.Path)

    if err1 != nil {
        return err1
    }

    this.absFilePath = absFilePath

    if this.GlobalConfig.Debug() {
        logger.Infof("this.absFilePath: %s", this.absFilePath)
    }

    this.absDirPath = filepath.Dir(this.absFilePath)

    if this.GlobalConfig.Debug() {
        logger.Infof("this.absDirPath: %s", this.absDirPath)
    }

    _, absFileErr := os.Stat(this.absFilePath)

    if os.IsNotExist(absFileErr) {
        if this.GlobalConfig.Debug() {
            logger.Infof("%s doesn't exists", this.absFilePath)
        }
    } else {
        if !this.GlobalConfig.FileConfig.Overwrite() {
            return errors.New(fmt.Sprintf("%s already exists", this.absFilePath))
        }
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
        if err := os.MkdirAll(this.absDirPath, this.GlobalConfig.FileConfig.FileMode()); err != nil {
            return err
        }
    }

    if !this.GlobalConfig.Quiet() {
        logger.Infof("Writing %s [mode: %s]", this.absFilePath, this.GlobalConfig.FileConfig.FileMode())
    }

    if err := afero.WriteFile(fs, this.absFilePath, data, this.GlobalConfig.FileConfig.FileMode()); err != nil {
        return err
    }

    return nil
}
