package file

import (
    "github.com/boundedinfinity/devenv/data"
    "github.com/boundedinfinity/devenv/config"
    "github.com/spf13/afero"
    "github.com/Sirupsen/logrus"
    "bytes"
    "text/template"
    "errors"
    "fmt"
    "os"
    "path/filepath"
)

func NewTemplateFile(logger *logrus.Entry, templatePath string, templateData interface{}, fsPath string) (*TemplateFile, error) {
    expaned := os.ExpandEnv(fsPath)
    abs, err := filepath.Abs(expaned)

    if err != nil {
        return nil, err
    }

    return &TemplateFile{
        TemplatePath: templatePath,
        TemplateData: templateData,
        FsPath: fsPath,
        ExpandedPath: abs,
        logger: logger,
        GlobalConfig: config.NewGlobalConfig(),
        templateEngine: template.New("template"),
        fileSystem: afero.NewOsFs(),
    }, nil
}

type TemplateFile struct {
    TemplatePath   string
    TemplateData   interface{}
    FsPath         string
    ExpandedPath   string
    logger         *logrus.Entry
    GlobalConfig   config.GlobalConfig
    templateEngine *template.Template
    fileSystem     afero.Fs
}

func (this *TemplateFile) Write() error {
    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("FsPath: %s", this.FsPath)
        this.logger.Infof("ExpandedPath: %s", this.ExpandedPath)
    }

    exists, err := afero.Exists(this.fileSystem, this.ExpandedPath)

    if err != nil {
        return err
    }

    if exists {
        if !this.GlobalConfig.FileConfig.Overwrite() {
            return errors.New(fmt.Sprintf("%s already exists: %s", this.ExpandedPath))
        }
    }

    dir := filepath.Dir(this.ExpandedPath)

    exists, err = afero.DirExists(this.fileSystem, dir)

    if err != nil {
        return err
    }

    if !exists {
        if err := this.fileSystem.MkdirAll(dir, this.GlobalConfig.FileConfig.FileMode()); err != nil {
            return err
        }
    }

    content, err := data.Asset(this.TemplatePath)

    if err != nil {
        return err
    }

    buffer := new(bytes.Buffer)
    template, err := this.templateEngine.Parse(string(content))

    if err := template.Execute(buffer, this.TemplateData); err != nil {
        return err
    }

    if err := afero.WriteFile(this.fileSystem, this.ExpandedPath, buffer.Bytes(), this.GlobalConfig.FileConfig.FileMode()); err != nil {
        return err
    }

    return nil
}
