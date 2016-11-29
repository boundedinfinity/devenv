package project

import (
    "github.com/Sirupsen/logrus"
    "github.com/boundedinfinity/devenv/logging"
)

func NewEditorConfigManager() *EditorConfigManager {
    return NewEditorConfigManagerWithLogger(logging.ComponentLogger("MakeFileManager"))
}

func NewEditorConfigManagerWithLogger(logger *logrus.Entry) *EditorConfigManager {
    return &EditorConfigManager{
        Path: "project/editorconfig/.editorconfig",
        logger : logger,
    }
}

type EditorConfigManager struct {
    logger *logrus.Entry
    Path   string
    Data   glideTemplateData
}

func (this *EditorConfigManager) Ensure() error {
    manager := NewProjectDirectoryManager(this.logger)

    if err := manager.EnsureFile(this.Path, this.Data); err != nil {
        return err
    }

    return nil
}
