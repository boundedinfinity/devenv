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

    if err := fs.Mkdir(realFsPath, this.Descriptor.FileMode); err != nil {
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
    manager := NewFileManager(this.logger)

    if this.GlobalConfig.Debug() {
        this.logger.Infof("Template Path: %s", path)
    }

    fsPath, err := Template2FsPath(this.Descriptor.FsPath, path)

    if err != nil {
        return err
    }

    dirDesc := DirectoryDescriptor{
        FsPath: filepath.Dir(fsPath),
        FileMode: this.Descriptor.FileMode,
        ExistMode: this.Descriptor.ExistMode,
        ExpandPath: this.Descriptor.ExpandPath,
    }

    if err := NewDirectoryManager(this.logger, dirDesc).CreateDir(); err != nil {
        return err
    }

    existMode := IgnoreIfExists

    if this.GlobalConfig.FileConfig.Overwrite() {
        existMode = OverwriteIfExists
    }

    if err != nil {
        return err
    }

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
