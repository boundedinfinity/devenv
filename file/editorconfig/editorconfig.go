package editorconfig

import (
    "github.com/boundedinfinity/devenv/project"
)

func NewEditorConfigManager() *EditorConfigManager {
    return &EditorConfigManager{
        Pfm: project.NewProjectFileManager("project/editorconfig/.editorconfig"),
    }
}

type EditorConfigManager struct {
    Pfm *project.ProjectFileManager
}

func (this *EditorConfigManager) Ensure() error {
    if err := this.Pfm.Ensure(); err != nil {
        return err
    }

    return nil
}
