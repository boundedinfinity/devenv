package project

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/file"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/Sirupsen/logrus"
    "fmt"
    "errors"
    "path/filepath"
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
    manager := file.NewFileManager2(this.logger)

    if err := manager.CreateDir(this.Descriptor); err != nil {
        return err
    }

    return nil
}

func (this *ProjectDirectoryManager) Delete() error {
    manager := file.NewFileManager2(this.logger)

    if err := manager.DeleteDir(this.Descriptor); err != nil {
        return err
    }

    return nil
}

func (this *ProjectDirectoryManager) Template2FsPath(templatePath string) (string, error) {
    var fsPath string

    if this.GlobalConfig.Debug() {
        this.logger.Infof("Template Path: %s", templatePath)
    }

    if !file.TemplateExists(templatePath) {
        return fsPath, errors.New(fmt.Sprintf("template '%s' not found", templatePath))
    }

    filename := filepath.Base(templatePath)
    fsPath = filepath.Join(this.Descriptor.FsPath, filename)

    return fsPath, nil
}

func (this *ProjectDirectoryManager) EnsureFile(path string, data interface{}) error {
    if err := this.Ensure(); err != nil {
        return err
    }

    manager := file.NewFileManager2(this.logger)
    fsPath, err := this.Template2FsPath(path)
    existMode := file.IgnoreIfExists

    if this.GlobalConfig.FileConfig.Overwrite() {
        existMode = file.OverwriteIfExists
    }

    if err != nil {
        return err
    }

    desc := file.TemplateCopyFileDescriptor{
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
