package file

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/spf13/afero"
    "github.com/Sirupsen/logrus"
    "os"
    "path/filepath"
    "fmt"
    "errors"
)

func NewSymLinkFile(logger *logrus.Entry, linkPath string, targetPath string) (*SymLinkFile, error) {
    expanedLink := os.ExpandEnv(linkPath)
    absLink, err := filepath.Abs(expanedLink)

    if err != nil {
        return nil, err
    }

    expanedTarget := os.ExpandEnv(targetPath)
    absTarget, err := filepath.Abs(expanedTarget)

    if err != nil {
        return nil, err
    }

    return &SymLinkFile{
        LinkPath: linkPath,
        TargetPath: targetPath,
        ExpandedLinkPath: absLink,
        ExpandedTargetPath: absTarget,
        logger: logger,
        GlobalConfig: config.NewGlobalConfig(),
        fileSystem: afero.NewOsFs(),
    }, nil
}

type SymLinkFile struct {
    LinkPath           string
    TargetPath         string
    ExpandedLinkPath   string
    ExpandedTargetPath string
    logger             *logrus.Entry
    GlobalConfig       config.GlobalConfig
    fileSystem         afero.Fs
}

func (this *SymLinkFile) Create() error {
    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("LinkPath: %s", this.LinkPath)
        this.logger.Infof("ExpandedLinkPath: %s", this.ExpandedLinkPath)
        this.logger.Infof("TargetPath: %s", this.TargetPath)
        this.logger.Infof("ExpandedTargetPath: %s", this.ExpandedTargetPath)
    }

    exists, err := afero.Exists(this.fileSystem, this.ExpandedTargetPath)

    if err != nil {
        return err
    }

    if !exists {
        return errors.New(fmt.Sprintf("target %s doesn't exists: %s", this.ExpandedTargetPath))
    }

    dir := filepath.Dir(this.ExpandedLinkPath)

    exists, err = afero.DirExists(this.fileSystem, dir)

    if err != nil {
        return err
    }

    if !exists {
        if err := this.fileSystem.MkdirAll(dir, this.GlobalConfig.FileConfig.FileMode()); err != nil {
            return err
        }
    }

    if err := os.Symlink(this.ExpandedTargetPath, this.ExpandedLinkPath); err != nil {
        return err
    }

    return nil
}
