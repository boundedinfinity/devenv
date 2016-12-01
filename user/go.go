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
    GoPath      string
    PackageName string
}

func (this *GoEnvManager) Ensure() error {
    //manager := NewUserDirectoryManagerWithLogger(this.logger)


    return nil
}

func (this *GoEnvManager) Enable() error {
    //manager := NewUserDirectoryManagerWithLogger(this.logger)

    return nil
}
