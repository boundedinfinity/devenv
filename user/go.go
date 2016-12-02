package user

import (
    "github.com/boundedinfinity/devenv/logging"
    "github.com/Sirupsen/logrus"
)

func NewGoEnvManager() *GoEnvManager {
    return &GoEnvManager{
        logger : logging.ComponentLogger("GoEnvManager"),
    }
}

type GoEnvManager struct {
    logger *logrus.Entry
}

func (this *GoEnvManager) Enable() error {
    userDir, err := NewUserDirectory()

    if err != nil {
        return err
    }

    if err := userDir.Enable("go"); err != nil {
        return err
    }

    return nil
}

func (this *GoEnvManager) Disable() error {
    //manager := NewUserDirectoryManagerWithLogger(this.logger)

    return nil
}
