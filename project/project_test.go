package project

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/spf13/viper"
    "testing"
)

var testDebug = false
var testDir = "./test-output"

func TestMakefile(t *testing.T) {
    viper.Set(config.Flag_Debug, testDebug)
    viper.Set(config.Flag_ProjectPath, testDir)

    manager, err := NewMakefileTemplateFile()

    if err != nil {
        t.Errorf("ensure: %s", err.Error())
    }

    if err := manager.Write(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}

func TestGlide(t *testing.T) {
    viper.Set(config.Flag_Debug, testDebug)
    viper.Set(config.Flag_ProjectPath, testDir)
    viper.Set(config.Flag_GoPackageName, "github.com/boundedinfinity/test")

    manager, err := NewGlideTemplateFile()

    if err != nil {
        t.Errorf("ensure: %s", err.Error())
    }

    if err := manager.Write(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}

func TestEditorConfig(t *testing.T) {
    viper.Set(config.Flag_Debug, testDebug)
    viper.Set(config.Flag_ProjectPath, testDir)

    manager, err := NewEditorConfigTemplateFile()

    if err != nil {
        t.Errorf("ensure: %s", err.Error())
    }

    if err := manager.Write(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}
