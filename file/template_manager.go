package file

import (
    "bytes"
    "text/template"
    "github.com/boundedinfinity/devenv/data"
    "github.com/boundedinfinity/devenv/config"
    "github.com/pkg/errors"
    "fmt"
    "github.com/boundedinfinity/devenv/logging"
)

var tlogger = logging.ComponentLogger("TemplateManager")

func NewTemplateManager(templatePath string, templateData interface{}) *TemplateManager {
    return &TemplateManager {
        GlobalConfig: config.GlobalConfig{},
        TemplatePath: templatePath,
        TemplateData: templateData,
    }
}

type TemplateManager struct{
    GlobalConfig config.GlobalConfig
    TemplatePath string
    TemplateData interface{}
}

func (this *TemplateManager) Render() ([]byte, error) {
    if this.GlobalConfig.Debug() {
        tlogger.Infof("TemplatePath: %s", this.TemplatePath)
        tlogger.Infof("TemplateData: %v", this.TemplateData)
    }

    if !TemplateExists(this.TemplatePath) {
        return []byte{}, errors.New(fmt.Sprintf("template '%s' not found", this.TemplatePath))
    }

    content, err1 := data.Asset(this.TemplatePath)

    if err1 != nil {
        return []byte{}, err1
    }

    buffer := new(bytes.Buffer)
    tmpl, err2 := template.New("template").Parse(string(content))

    if err2 != nil {
        return []byte{}, err2
    }

    if err := tmpl.Execute(buffer, this.TemplateData); err != nil {
        return []byte{}, err
    }

    return buffer.Bytes(), nil
}
