package file

import (
    "bytes"
    "text/template"
    "github.com/boundedinfinity/devenv/data"
    "github.com/boundedinfinity/devenv/config"
    "log"
)

type TemplateManager struct{
    GlobalConfig config.GlobalConfig
    TemplatePath string
    TemplateData interface{}
}

func (this *TemplateManager) Render() ([]byte, error) {
    if this.GlobalConfig.Debug() {
        log.Printf("TemplatePath: %s", this.TemplatePath)
        log.Printf("TemplateData: %v", this.TemplateData)
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
