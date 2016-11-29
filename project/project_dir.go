package project

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/file"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/Sirupsen/logrus"
)

func NewProjectDirectoryManager(logger *logrus.Entry) *ProjectDirectoryManager {
    globalConfig := config.NewGlobalConfig()

    return &ProjectDirectoryManager{
        logger: logging.SubComponentLogger(logger, "ProjectDirectoryManager"),
        GlobalConfig: globalConfig,
        Descriptor: file.DirectoryDescriptor{
            FsPath: globalConfig.ProjectConfig.ProjectPath(),
            FileMode: globalConfig.DirConfig.FileMode(),
            ExistMode: file.IgnoreIfExists,
            ExpandPath: true,
        },
    }
}

type ProjectDirectoryManager struct {
    logger       *logrus.Entry
    GlobalConfig config.GlobalConfig
    Descriptor   file.DirectoryDescriptor
}

func (this *ProjectDirectoryManager) Ensure() error {
    manager := file.NewDirectoryManager(this.logger, this.Descriptor)

    if err := manager.CreateDir(); err != nil {
        return err
    }

    return nil
}

func (this *ProjectDirectoryManager) Delete() error {
    manager := file.NewDirectoryManager(this.logger, this.Descriptor)

    if err := manager.DeleteDir(this.Descriptor); err != nil {
        return err
    }

    return nil
}

func (this *ProjectDirectoryManager) EnsureFile(path string, data interface{}) error {
    if err := this.Ensure(); err != nil {
        return err
    }

    manager := file.NewDirectoryManager(this.logger, this.Descriptor)

    if err := manager.EnsureFile(path, data); err != nil {
        return err
    }

    return nil
}