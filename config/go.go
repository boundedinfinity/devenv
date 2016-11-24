package config

import (
    "github.com/spf13/viper"
)

// Flags

const (
    Flag_GoPackageName = "go-package-name"
    Flag_GoPath = "go-path"
)

// Defaults

const (
    Flag_Default_GoPackageName = "main"
    Flag_Default_GoPath = "$HOME/go"
)

func init() {
    viper.SetDefault(Flag_GoPackageName, Flag_Default_GoPackageName)
    viper.SetDefault(Flag_GoPath, Flag_Default_GoPath)
}

type GoConfig struct{}

func (this GoConfig) GoPackageName() string {
    return viper.GetString(Flag_GoPackageName)
}

func (this GoConfig) GoPath() string {
    return viper.GetString(Flag_GoPath)
}
