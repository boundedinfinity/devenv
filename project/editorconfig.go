package project

func NewEditorConfigManager() *EditorConfigManager {
    return &EditorConfigManager{
        Pfm: NewProjectFileManager("project/editorconfig/.editorconfig"),
    }
}

type EditorConfigManager struct {
    Pfm *ProjectFileManager
}

func (this *EditorConfigManager) Ensure() error {
    if err := this.Pfm.Ensure(); err != nil {
        return err
    }

    return nil
}
