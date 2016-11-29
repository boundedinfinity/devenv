package user

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/file"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/Sirupsen/logrus"
)

func NewUserDirectoryManager() *UserDirectoryManager {
    return NewUserDirectoryManagerWithLogger(logging.ComponentLogger("UserDirectoryManager"))
}

func NewUserDirectoryManagerWithLogger(logger *logrus.Entry) *UserDirectoryManager {
    globalConfig := config.NewGlobalConfig()

    return &UserDirectoryManager{
        logger: logger,
        GlobalConfig: globalConfig,
        Descriptor: file.DirectoryDescriptor{
            FsPath: globalConfig.UserConfigDir(),
            FileMode: globalConfig.DirConfig.FileMode(),
            ExistMode: file.IgnoreIfExists,
            ExpandPath: true,
        },
    }
}

type UserDirectoryManager struct {
    logger       *logrus.Entry
    GlobalConfig config.GlobalConfig
    Descriptor   file.DirectoryDescriptor
}

func (this *UserDirectoryManager) Ensure() error {
    manager := file.NewDirectoryManager(this.logger, this.Descriptor)

    if err := manager.CreateDir(); err != nil {
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

func (this *UserDirectoryManager) EnsureFile(path string, data interface{}) error {
    if err := this.Ensure(); err != nil {
        return err
    }

    manager := file.NewDirectoryManager(this.logger, this.Descriptor)

    if err := manager.EnsureFile(path, data); err != nil {
        return err
    }

    return nil
}
