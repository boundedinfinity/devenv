package makefile

import (
    "testing"
    "github.com/spf13/viper"
    "github.com/boundedinfinity/devenv/config"
)

func TestMakefile(t *testing.T) {
    manager := NewMakefileManager()

    viper.Set(config.Flag_Debug, true)
    viper.Set(config.Flag_ProjectPath, "./test-project")

    if err := manager.Pfm.Project.DeleteDirectory(); err != nil {
        t.Errorf("delete: %s", err.Error())
    }

    if err := manager.Ensure(); err != nil {
        t.Errorf("ensure: %s", err.Error())
    }
}
