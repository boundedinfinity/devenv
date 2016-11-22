package project

import "github.com/boundedinfinity/devenv/config"

func NewProjectFileManager(dataPath string) *ProjectFileManager {
    return &ProjectFileManager{
        GlobalConfig: config.GlobalConfig{},
        ProjectConfig: config.ProjectConfig{},
        DirConfig: config.DirConfig{},
        FileConfig: config.FileConfig{},
    }
}

type ProjectFileManager struct {
    GlobalConfig config.GlobalConfig
    ProjectConfig config.ProjectConfig
    DirConfig     config.DirConfig
    FileConfig   config.FileConfig
}
