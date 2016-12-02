package project

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/file"
    "github.com/Sirupsen/logrus"
    "path/filepath"
)

func NewProjectTemplateFile(logger *logrus.Entry, templatePath string, templateData interface{}) (*file.TemplateFile, error) {
    filename := filepath.Base(templatePath)
    fsPath := filepath.Join(config.NewGlobalConfig().ProjectConfig.ProjectPath(), filename)
    return file.NewTemplateFile(logger, templatePath, templateData, fsPath)
}
