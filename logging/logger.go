package logging

import (
    "github.com/Sirupsen/logrus"
    "os"
    "fmt"
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

func getFieldName(entry *logrus.Entry, name string) string {
    var found string

    for k, v := range entry.Data {
        if k == name {
            found = v.(string)
            break
        }
    }

    return found
}

func ComponentLogger(name string) *logrus.Entry {
    return rootLogger().WithFields(logrus.Fields{
        "component": name,
    })
}

func SubComponentLogger(entry *logrus.Entry, name string) *logrus.Entry {
    var parentComponent = getFieldName(entry, "component")
    var subComponent string

    if parentComponent == "" {
        subComponent = name
    } else {
        subComponent = fmt.Sprintf("%s/%s", parentComponent, name)
    }

    return ComponentLogger(subComponent)
}

func init() {
    logrus.SetFormatter(&logrus.TextFormatter{})
    logrus.SetOutput(os.Stdout)
}
