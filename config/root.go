package config

import "github.com/spf13/viper"

// Flags

const (
    Flag_Debug = "debug"
    Flag_UserConfigDir = "user-config-dir"
    Flag_Quiet = "quiet"
)

// Default

const (
    Flag_Default_Debug = false
    Flag_Default_UserConfigDir = "$HOME/.config"
    Flag_Default_Quiet = false
)

func init() {
    viper.SetDefault(Flag_Debug, Flag_Default_Debug)
    viper.SetDefault(Flag_UserConfigDir, Flag_Default_UserConfigDir)
    viper.SetDefault(Flag_Quiet, Flag_Default_Quiet)
}

type GlobalConfig struct{}

func (this GlobalConfig) Debug() bool {
    return viper.GetBool(Flag_Debug)
}

func (this GlobalConfig) Quiet() bool {
    return viper.GetBool(Flag_Quiet)
}

func (this GlobalConfig) UserConfigDir() string {
    return viper.GetString(Flag_UserConfigDir)
}
