package config

import (
    "github.com/spf13/viper"
    "os"
    "strconv"
)

// Flags

const (
    Flag_Overwrite = "overwrite"
    Flag_ProjectPath = "project-path"
    Flag_FileMode = "file-mode"
)

// Defaults

const (
    Flag_Default_Overwrite = false
    Flag_Default_ProjectPath = "."
    Flag_Default_Flag_FileMode = "0755"
)

func init() {
    viper.SetDefault(Flag_Overwrite, Flag_Default_Overwrite)
    viper.SetDefault(Flag_ProjectPath, Flag_Default_ProjectPath)
    viper.SetDefault(Flag_FileMode, Flag_Default_Flag_FileMode)
}

type FileConfig struct{}

func (this FileConfig) Overwrite() bool {
    return viper.GetBool(Flag_Overwrite)
}

func (this FileConfig) ProjectPath() string {
    return viper.GetString(Flag_ProjectPath)
}

func (this FileConfig) FileMode() os.FileMode {
    val1 := viper.GetString(Flag_FileMode)
    val2, _ := strconv.ParseUint(val1, 0, 32)
    val3 := os.FileMode(val2)
    return os.FileMode(val3)
}

type DirConfig struct{}

func (this DirConfig) FileMode() os.FileMode {
    val1 := viper.GetString(Flag_FileMode)
    val2, _ := strconv.ParseUint(val1, 0, 32)
    val3 := os.FileMode(val2)
    return os.FileMode(val3)
}
