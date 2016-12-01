package user

import (
    "testing"
    "github.com/spf13/viper"
    "github.com/boundedinfinity/devenv/config"
)

var testDebug = false
var testDir = "./test-output"

func TestConfigDir(t *testing.T) {
    viper.Set(config.Flag_Debug, testDebug)
    viper.Set(config.Flag_UserConfigDir, testDir)

    manager := NewUserDirectoryManager()

    if err := manager.Ensure(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}

func TestGoEnv(t *testing.T) {
    viper.Set(config.Flag_Debug, testDebug)
    viper.Set(config.Flag_UserConfigDir, testDir)

    manager := NewGoEnvManager()

    if err := manager.Ensure(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}
