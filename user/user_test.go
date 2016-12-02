package user

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/spf13/viper"
    "testing"
)

var testDebug = false
var testDir = "./test-output"

func TestConfigDir(t *testing.T) {
    viper.Set(config.Flag_Debug, testDebug)
    viper.Set(config.Flag_UserConfigDir, testDir)

    manager, err := NewUserDirectory()

    if err != nil {
        t.Errorf("ensure: %s", err.Error())
    }

    if err := manager.Create(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}

func TestGoEnv(t *testing.T) {
    viper.Set(config.Flag_Debug, testDebug)
    viper.Set(config.Flag_UserConfigDir, testDir)

    manager := NewGoEnvManager()

    if err := manager.Enable(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}
