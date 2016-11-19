package config

import (
    "github.com/spf13/viper"
    "os"
)

// Flags

const (
    Flag_Overwrite = "overwrite"
    Flag_ProjectPath = "project-path"
    Flag_FileMode = "file-mode"
)

// Default

const (
    Flag_Default_Overwrite = false
    Flag_Default_ProjectPath = "."
    Flag_Default_Flag_FileMode uint32 = 755
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
    return os.FileMode(viper.GetInt(Flag_FileMode))
}
