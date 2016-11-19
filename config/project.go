package config

import (
    "github.com/spf13/viper"
)

// Flags

const (
    Flag_ProjectName = "project-name"
)

// Default

const (
    Flag_Default_ProjectName = "generatedProject"
)

func init() {
    viper.SetDefault(Flag_ProjectName, Flag_Default_ProjectName)
}

type ProjectConfig struct{}

func (this ProjectConfig) ProjectName() string {
    return viper.GetString(Flag_ProjectName)
}
