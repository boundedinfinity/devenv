package user

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/spf13/afero"
    "log"
    "strings"
    "github.com/boundedinfinity/devenv/shell"
    "fmt"
)

func NewUserConfigManager() *UserConfigManager {
    return &UserConfigManager{
        GlobalConfig: config.GlobalConfig{},
    }
}

type UserConfigManager struct {
    GlobalConfig config.GlobalConfig
    DirConfig    config.DirConfig
}

func (this *UserConfigManager) EnsureConfigDir() error {
    var realDir string

    if strings.Contains(this.GlobalConfig.UserConfigDir(), "$") {
        output, err := shell.Evaluate(fmt.Sprintf("echo %s", this.GlobalConfig.UserConfigDir()))

        if err != nil {
            return err
        }

        realDir = output
    } else {
        realDir = this.GlobalConfig.UserConfigDir()
    }

    if !this.GlobalConfig.Quiet() {
        log.Printf("Input Configuration Directory: %s", this.GlobalConfig.UserConfigDir())
        log.Printf("Evaluated Configuration Directory: %s", realDir)
    }

    fs := afero.NewOsFs()

    exists, err := afero.DirExists(fs, realDir)

    if err != nil {
        return err
    }

    if !exists {
        log.Printf("Creating: %s", realDir)

        if err := fs.MkdirAll(realDir, this.DirConfig.FileMode()); err != nil {
            return err
        }
    }

    return nil
}
