package file

import (
    "github.com/boundedinfinity/devenv/data"
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/spf13/afero"
    "github.com/Sirupsen/logrus"
    "bytes"
    "text/template"
    "errors"
)

func NewFileManager(logger *logrus.Entry) *FileManager {
    return &FileManager{
        logger: logging.SubComponentLogger(logger, "FileManager"),
        templateEngine: template.New("template"),
        GlobalConfig: config.NewGlobalConfig(),
    }
}

type FileManager struct {
    logger         *logrus.Entry
    templateEngine *template.Template
    GlobalConfig   config.GlobalConfig
}

func (this *FileManager) CopyFile(desc TemplateCopyFileDescriptor) error {
    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Input File Path: %s", desc.FsPath)
    }

    realFsPath, err := CalcRealPath(desc.FsPath, desc.ExpandPath)

    if err != nil {
        return err
    }

    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Absolute File Path: %s", realFsPath)
    }

    fs := afero.NewOsFs()

    exists, err := afero.Exists(fs, realFsPath)

    if err != nil {
        return err
    }

    if exists {
        switch desc.ExistMode {
        case IgnoreIfExists:
            return nil
        case FailOnExists:
            return errors.New("already exists")
        case OverwriteIfExists:
            if err := fs.Remove(realFsPath); err != nil {
                return nil
            }
        }
    }

    if this.GlobalConfig.Debug() {
        this.logger.Infof("Template File Path: %s", desc.FsPath)
    }

    content, err := data.Asset(desc.TemplatePath)

    if err != nil {
        return err
    }

    buffer := new(bytes.Buffer)
    template, err := this.templateEngine.Parse(string(content))

    if err := template.Execute(buffer, desc.TemplateData); err != nil {
        return err
    }

    if err := afero.WriteFile(fs, realFsPath, buffer.Bytes(), desc.FileMode); err != nil {
        return err
    }

    return nil
}
