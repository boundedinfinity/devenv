package config

import (
    "github.com/spf13/viper"
    "os"
    "github.com/boundedinfinity/devenv/util"
)

// Flags

const (
    Flag_Overwrite = "overwrite"
    Flag_FileMode = "file-mode"
)

// Defaults

const (
    Flag_Default_Overwrite = false

    Flag_Default_Flag_FileMode = "0755"
)

func init() {
    viper.SetDefault(Flag_Overwrite, Flag_Default_Overwrite)
    viper.SetDefault(Flag_FileMode, Flag_Default_Flag_FileMode)
}

type FileConfig struct{}

func (this FileConfig) Overwrite() bool {
    return viper.GetBool(Flag_Overwrite)
}

func (this FileConfig) FileMode() os.FileMode {
    val, _ := util.String2FileMode(viper.GetString(Flag_FileMode))
    return os.FileMode(val)
}

type DirConfig struct{}

func (this DirConfig) FileMode() os.FileMode {
    val, _ := util.String2FileMode(viper.GetString(Flag_FileMode))
    return os.FileMode(val)
}
