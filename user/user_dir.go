package user

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/file"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/Sirupsen/logrus"
    "path/filepath"
    "strings"
    "fmt"
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
    files := []string{
        "user/devenv/bash/load.bash",
        "user/devenv/bash/available/go.bash",
        "user/devenv/bash/available/anyenv.bash",
    }

    for _, file := range files {
        if err := this.WriteTemplateFile(file, nil); err != nil {
            return err
        }
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

func (this *UserDirectory) NewSymLinkFile(link string, target string) (*file.SymLinkFile, error) {
    userLink := filepath.Join(this.directory.ExpandedPath, link)
    userTarget := filepath.Join(this.directory.ExpandedPath, target)
    return file.NewSymLinkFile(this.logger, userLink, userTarget)
}

func (this *UserDirectory) WriteSymLinkFile(link string, target string) error {
    f, err := this.NewSymLinkFile(link, target)

    if err != nil {
        return err
    }

    if err := f.Create(); err != nil {
        return err
    }

    return nil
}

func (this *UserDirectory) Enable(thing string) error {
    if err := this.enableShell("bash", thing); err != nil {
        return err
    }

    return nil
}

func (this *UserDirectory) enableShell(shell string, thing string) error {
    enabled := fmt.Sprintf("devenv/%s/enabled/%s.bash", shell, thing)
    available := fmt.Sprintf("devenv/%s/available/%s.bash", shell, thing)

    if err := this.WriteSymLinkFile(enabled, available); err != nil {
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
