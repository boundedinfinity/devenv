package project

import (
    "github.com/Sirupsen/logrus"
)

func NewEditorConfigManager() *EditorConfigManager {
    return &EditorConfigManager{
        Path: "project/editorconfig/.editorconfig",
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
