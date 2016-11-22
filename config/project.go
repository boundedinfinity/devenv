package config

import (
    "github.com/spf13/viper"
)

// Flags

const (
    Flag_ProjectPath = "project-path"
    Flag_ProjectName = "project-name"
)

// Defaults

const (
    Flag_Default_ProjectName = "generatedProject"
    Flag_Default_ProjectPath = "."
)

func init() {
    viper.SetDefault(Flag_ProjectName, Flag_Default_ProjectName)
    viper.SetDefault(Flag_ProjectPath, Flag_Default_ProjectPath)
}

type ProjectConfig struct{}

func (this ProjectConfig) ProjectName() string {
    return viper.GetString(Flag_ProjectName)
}

func (this ProjectConfig) ProjectPath() string {
    return viper.GetString(Flag_ProjectPath)
}
