package user

import (
    "github.com/spf13/viper"
    "github.com/boundedinfinity/devenv/config"
    "testing"
)

var testDebug = false
var testDir = "./test-output"

func TestConfig(t *testing.T) {
    viper.Set(config.Flag_Debug, testDebug)
    viper.Set(config.Flag_UserConfigDir, testDir)

    manager := NewUserConfigManager()

    if err := manager.EnsureConfigDir(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}
