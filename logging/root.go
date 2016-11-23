package logging

import (
    "github.com/Sirupsen/logrus"
    "os"
)

var root *logrus.Entry

func rootLogger() *logrus.Entry {
    if root == nil {
        root = logrus.WithFields(logrus.Fields{
            "application": "devenv",
        })
    }

    return root
}

func ComponentLogger(name string) *logrus.Entry {
    return rootLogger().WithFields(logrus.Fields{
        "component": name,
    })
}

func init() {
    logrus.SetFormatter(&logrus.TextFormatter{})
    logrus.SetOutput(os.Stdout)
}
