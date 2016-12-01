package user

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/file"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/Sirupsen/logrus"
    "path/filepath"
    "strings"
)

func NewUserDirectoryManager() *UserDirectoryManager {
    return NewUserDirectoryManagerWithLogger(logging.ComponentLogger("UserDirectoryManager"))
}

func NewUserDirectoryManagerWithLogger(logger *logrus.Entry) *UserDirectoryManager {
    globalConfig := config.NewGlobalConfig()

    manager := &UserDirectoryManager{
        logger: logger,
        GlobalConfig: globalConfig,
        Descriptor: file.DirectoryDescriptor{
            FsPath: globalConfig.UserConfigDir(),
            FileMode: globalConfig.DirConfig.FileMode(),
            ExistMode: file.IgnoreIfExists,
            ExpandPath: true,
        },
    }

    return manager
}

type UserDirectoryManager struct {
    logger       *logrus.Entry
    GlobalConfig config.GlobalConfig
    Descriptor   file.DirectoryDescriptor
}

func (this *UserDirectoryManager) newDirDesc(path string) file.DirectoryDescriptor {
    fsPath := strings.Replace(path, "user", this.Descriptor.FsPath, -1)
    fsDir := filepath.Dir(fsPath)

    return file.DirectoryDescriptor{
        FsPath: fsDir,
        FileMode: this.Descriptor.FileMode,
        ExpandPath: this.Descriptor.ExpandPath,
        ExistMode: this.Descriptor.ExistMode,
    }
}

func (this *UserDirectoryManager) Ensure() error {
    manager := file.NewDirectoryManager(this.logger, this.Descriptor)

    if err := manager.CreateDir(); err != nil {
        return err
    }

    if err := this.ensureBash(); err != nil {
        return err
    }

    if err := this.ensureFish(); err != nil {
        return err
    }

    return nil
}

func (this *UserDirectoryManager) ensureBash() error {
    if err := this.ensureFile("user/devenv/bash/load.bash", nil); err != nil {
        return err
    }

    if err := this.ensureFile("user/devenv/bash/available/go.bash", nil); err != nil {
        return err
    }

    enabledDesc := file.DirectoryDescriptor{
        FsPath: filepath.Join(this.Descriptor.FsPath, "devenv/bash/enabled"),
        FileMode: this.Descriptor.FileMode,
        ExpandPath: this.Descriptor.ExpandPath,
        ExistMode: this.Descriptor.ExistMode,
    }

    if err := file.NewDirectoryManager(this.logger, enabledDesc).CreateDir(); err != nil {
        return err
    }

    return nil
}

func (this *UserDirectoryManager) ensureFish() error {
    if err := this.ensureFile("user/devenv/fish/load.fish", nil); err != nil {
        return err
    }

    if err := this.ensureFile("user/devenv/fish/available/go.fish", nil); err != nil {
        return err
    }

    enabledDesc := file.DirectoryDescriptor{
        FsPath: filepath.Join(this.Descriptor.FsPath, "devenv/fish/enabled"),
        FileMode: this.Descriptor.FileMode,
        ExpandPath: this.Descriptor.ExpandPath,
        ExistMode: this.Descriptor.ExistMode,
    }

    if err := file.NewDirectoryManager(this.logger, enabledDesc).CreateDir(); err != nil {
        return err
    }

    return nil
}

func (this *UserDirectoryManager) Delete() error {
    manager := file.NewDirectoryManager(this.logger, this.Descriptor)

    if err := manager.DeleteDir(this.Descriptor); err != nil {
        return err
    }

    return nil
}

func (this *UserDirectoryManager) ensureFile(path string, data interface{}) error {
    manager := file.NewDirectoryManager(this.logger, this.newDirDesc(path))

    if err := manager.EnsureFile(path, data); err != nil {
        return err
    }

    return nil
}

func (this *UserDirectoryManager) EnsureFile(path string, data interface{}) error {
    if err := this.Ensure(); err != nil {
        return err
    }

    return this.ensureFile(path, data)
}
