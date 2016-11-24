package project

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/spf13/viper"
    "testing"
)

var testDebug = true
var testDir = "./test-project"

func TestMakefile(t *testing.T) {
    viper.Set(config.Flag_Debug, testDebug)
    viper.Set(config.Flag_ProjectPath, testDir)

    manager := NewMakefileManager()

    //if err := manager.Pfm.Project.DeleteDirectory(); err != nil {
    //    t.Errorf("delete: %s", err.Error())
    //}

    if err := manager.Ensure(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}

func TestGlide(t *testing.T) {
    viper.Set(config.Flag_Debug, testDebug)
    viper.Set(config.Flag_ProjectPath, testDir)
    viper.Set(config.Flag_GoPackageName, "github.com/boundedinfinity/test")

    manager := NewGlideManager()

    //if err := manager.Pfm.Project.DeleteDirectory(); err != nil {
    //    t.Errorf("delete: %s", err.Error())
    //}

    if err := manager.Ensure(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}
