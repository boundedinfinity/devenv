package user

import (
    "github.com/boundedinfinity/devenv/logging"
    "github.com/boundedinfinity/devenv/config"
    "github.com/Sirupsen/logrus"
)

func NewGoEnvManager() *GoEnvManager {
    return &GoEnvManager{
        logger : logging.ComponentLogger("GoEnvManager"),
        Data: goTemplateData{
            PackageName: config.NewGlobalConfig().GoConfig.GoPackageName(),
            GoPath: config.NewGlobalConfig().GoConfig.GoPath(),
        },
    }
}

type GoEnvManager struct {
    logger *logrus.Entry
    Data   goTemplateData
}

type goTemplateData struct {
    GoPath string
    PackageName string
}

func (this *GoEnvManager) Ensure() error {
    manager := NewUserDirectoryManagerWithLogger(this.logger)

    if err := manager.EnsureFile("user/config/bash/load-env.bash", this.Data); err != nil {
        return err
    }

    if err := manager.EnsureFile("user/config/fish/load-env.fish", this.Data); err != nil {
        return err
    }

    return nil
}
