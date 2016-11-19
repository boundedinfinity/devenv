package config

import (
    "github.com/spf13/viper"
)

// Flags

const (
    Flag_GoPackageName = "go-package-name"
)

// Defaults

const (
    Flag_Default_GoPackageName = "main"
)

func init() {
    viper.SetDefault(Flag_GoPackageName, Flag_Default_GoPackageName)
}

type GoConfig struct{}

func (this GoConfig) GoPackageName() string {
    return viper.GetString(Flag_GoPackageName)
}
