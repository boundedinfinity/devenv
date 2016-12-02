package user

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/file"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/Sirupsen/logrus"
    "path/filepath"
    "strings"
)

func NewUserDirectory() (*UserDirectory, error) {
    logger := logging.ComponentLogger("UserDirectoryManager")
    directory, err := file.NewDirectory(
        logger,
        config.NewGlobalConfig().UserConfigDir(),
    )

    if err != nil {
        return nil, err
    }

    return &UserDirectory{
        directory: directory,
        logger: logger,
    }, nil
}

type UserDirectory struct {
    directory *file.Directory
    logger    *logrus.Entry
}

func (this *UserDirectory) Create() error {
    if err := this.directory.Create(); err != nil {
        return err
    }

    if err := this.ensureBash(); err != nil {
        return err
    }

    return nil
}

func (this *UserDirectory) ensureBash() error {
    if err := this.WriteTemplateFile("user/devenv/bash/load.bash", nil); err != nil {
        return err
    }

    if err := this.WriteTemplateFile("user/devenv/bash/available/go.bash", nil); err != nil {
        return err
    }

    if err := this.CreateDirectory("devenv/bash/enabled"); err != nil {
        return err
    }

    return nil
}

func (this *UserDirectory) NewUserTemplateFile(templatePath string, templateData interface{}) (*file.TemplateFile, error) {
    relativePath := strings.Replace(templatePath, "user/", "", -1)
    fsPath := filepath.Join(this.directory.ExpandedPath, relativePath)
    return file.NewTemplateFile(this.logger, templatePath, templateData, fsPath)
}

func (this *UserDirectory) WriteTemplateFile(templatePath string, templateData interface{}) error {
    f, err := this.NewUserTemplateFile(templatePath, templateData)

    if err != nil {
        return err
    }

    if err := f.Write(); err != nil {
        return err
    }

    return nil
}

func (this *UserDirectory) NewUserDirectory(fsPath string) (*file.Directory, error) {
    newPath := filepath.Join(this.directory.ExpandedPath, fsPath)
    return file.NewDirectory(this.logger, newPath)
}

func (this *UserDirectory) CreateDirectory(fsPath string) error {
    d, err := this.NewUserDirectory(fsPath)

    if err != nil {
        return err
    }

    if err := d.Create(); err != nil {
        return err
    }

    return nil
}
