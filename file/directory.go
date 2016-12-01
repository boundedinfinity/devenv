package file

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/Sirupsen/logrus"
    "github.com/spf13/afero"
    "errors"
    "path/filepath"
)

func NewDirectoryManager(logger *logrus.Entry, descriptor DirectoryDescriptor) *DirectoryManager {
    return &DirectoryManager{
        logger: logger,
        GlobalConfig: config.NewGlobalConfig(),
        Descriptor: descriptor,
    }
}

type DirectoryManager struct {
    logger       *logrus.Entry
    GlobalConfig config.GlobalConfig
    Descriptor   DirectoryDescriptor
}

func (this *DirectoryManager) CreateDir() error {
    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Input Dir Path: %s", this.Descriptor.FsPath)
    }

    realFsPath, err := CalcRealPath(this.Descriptor.FsPath, this.Descriptor.ExpandPath)

    if err != nil {
        return err
    }

    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Absolute Dir Path: %s", realFsPath)
    }

    fs := afero.NewOsFs()

    exists, err := afero.DirExists(fs, realFsPath)

    if err != nil {
        return err
    }

    if exists {
        switch this.Descriptor.ExistMode {
        case IgnoreIfExists:
            return nil
        case FailOnExists:
            return errors.New("already exists")
        case OverwriteIfExists:
            return nil
        }
    }

    if err := fs.MkdirAll(realFsPath, this.Descriptor.FileMode); err != nil {
        return err
    }

    return nil
}

func (this *DirectoryManager) DeleteDir(desc DirectoryDescriptor) error {
    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Input Dir Path: %s", desc.FsPath)
    }

    realFsPath, err := CalcRealPath(desc.FsPath, desc.ExpandPath)

    if err != nil {
        return err
    }

    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Absolute Dir Path: %s", realFsPath)
    }

    fs := afero.NewOsFs()

    exists, err := afero.DirExists(fs, realFsPath)

    if err != nil {
        return err
    }

    if exists {
        if err := fs.RemoveAll(realFsPath); err != nil {
            return err
        }
    }

    return nil
}

func (this *DirectoryManager) EnsureFile(path string, data interface{}) error {
    if err := this.CreateDir(); err != nil {
        return err
    }

    manager := NewFileManager(this.logger)

    if this.GlobalConfig.Debug() {
        this.logger.Infof("Template Path: %s", path)
    }

    existMode := IgnoreIfExists

    if this.GlobalConfig.FileConfig.Overwrite() {
        existMode = OverwriteIfExists
    }

    filename := filepath.Base(path)
    fsPath := filepath.Join(this.Descriptor.FsPath, filename)

    desc := TemplateCopyFileDescriptor{
        TemplatePath: path,
        TemplateData: data,
        FsPath: fsPath,
        ExpandPath: true,
        FileMode: this.GlobalConfig.FileConfig.FileMode(),
        ExistMode: existMode,
    }

    if err := manager.CopyFile(desc); err != nil {
        return err
    }

    return nil
}
