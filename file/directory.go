package file

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/Sirupsen/logrus"
    "github.com/spf13/afero"
    "path/filepath"
    "os"
)

func NewDirectory(logger *logrus.Entry, fsPath string) (*Directory, error) {
    expaned := os.ExpandEnv(fsPath)
    abs, err := filepath.Abs(expaned)

    if err != nil {
        return nil, err
    }

    return &Directory{
        FsPath: fsPath,
        ExpandedPath: abs,
        logger: logger,
        GlobalConfig: config.NewGlobalConfig(),
        fileSystem: afero.NewOsFs(),
    }, nil
}

type Directory struct {
    FsPath         string
    ExpandedPath   string
    logger         *logrus.Entry
    GlobalConfig   config.GlobalConfig
    fileSystem     afero.Fs
}

func (this *Directory) Create() error {
    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("FsPath: %s", this.FsPath)
        this.logger.Infof("ExpandedPath: %s", this.ExpandedPath)
    }

    exists, err := afero.DirExists(this.fileSystem, this.ExpandedPath)

    if err != nil {
        return err
    }

    if exists {
        return nil
    }

    if err := this.fileSystem.MkdirAll(this.ExpandedPath, this.GlobalConfig.FileConfig.FileMode()); err != nil {
        return err
    }

    return nil
}
