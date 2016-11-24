package user

import (
    "github.com/boundedinfinity/devenv/config"
    "github.com/spf13/afero"
    "strings"
    "github.com/boundedinfinity/devenv/shell"
    "fmt"
    "github.com/boundedinfinity/devenv/data"
    "path/filepath"
    "github.com/boundedinfinity/devenv/file"
    "github.com/boundedinfinity/devenv/logging"
)

var logger = logging.ComponentLogger("UserConfigManager")

func NewUserConfigManager() *UserConfigManager {
    return &UserConfigManager{
        GlobalConfig: config.NewGlobalConfig(),
        data: newDataDescriptor(),
    }
}

type UserConfigManager struct {
    GlobalConfig config.GlobalConfig
    realDir      string
    data         *dataDescriptor
}

func newDataDescriptor() *dataDescriptor {
    return &dataDescriptor{
        dirs: make(map[string]string),
        files: make(map[string]string),
    }
}

type dataDescriptor struct {
    dirs  map[string]string
    files map[string]string
}

func (this *UserConfigManager) EnsureConfigDir() error {
    if err := this.ensureConfigDir(); err != nil {
        return err
    }

    if err := this.ensureScriptDDirectories(); err != nil {
        return err
    }

    return nil
}

func (this *UserConfigManager) ensureScriptDDirectories() error {
    if err := this.getList("user/config"); err != nil {
        return nil
    }

    if this.GlobalConfig.Debug() {
        for dataPath, fsPath := range this.data.dirs {
            logger.Printf("dataPath: [ %s ], fsPath: [ %s ]", dataPath, fsPath)
        }
    }

    fs := afero.NewOsFs()

    for _, fsPath := range this.data.dirs {
        exists, err := afero.DirExists(fs, fsPath)

        if err != nil {
            return err
        }

        if !exists {
            logger.Printf("Creating: %s", fsPath)

            if err := fs.MkdirAll(fsPath, this.GlobalConfig.DirConfig.FileMode()); err != nil {
                return err
            }
        }
    }

    for dataPath, fsPath := range this.data.files {
        if err := this.copyFile(dataPath, fsPath); err != nil {
            return err
        }
    }

    return nil
}

type templateData struct {}

func (this *UserConfigManager) copyFile(dataPath string, fsPath string) error {
    fm := file.FileManager{
        GlobalConfig: this.GlobalConfig,
        Path: fsPath,
    }

    if err1 := fm.Validate(); err1 != nil {
        return err1
    }

    tm := file.TemplateManager{
        GlobalConfig: this.GlobalConfig,
        TemplatePath: dataPath,
        TemplateData: templateData{},
    }

    data, err2 := tm.Render()

    if err2 != nil {
        return err2
    }

    if err3 := fm.Write(data); err3 != nil {
        return err3
    }

    return nil
}

func (this *UserConfigManager) getList(path string) error {
    list, err := data.AssetDir(path)
    fsPath := filepath.Join(this.realDir, strings.Replace(path, "user/config", "", -1))

    if err != nil {
        if strings.Contains(err.Error(), fmt.Sprintf("%s not found", path)) {
            this.data.files[path] = fsPath
            return nil
        } else {
            return err
        }
    }

    this.data.dirs[path] = fsPath

    for _, name := range list {
        subPath := filepath.Join(path, name)

        if err := this.getList(subPath); err != nil {
            return err
        }
    }

    return nil
}

func (this *UserConfigManager) ensureConfigDir() error {
    if strings.Contains(this.GlobalConfig.UserConfigDir(), "$") {
        output, err := shell.Evaluate(fmt.Sprintf("echo -n %s", this.GlobalConfig.UserConfigDir()))

        if err != nil {
            return err
        }

        this.realDir = output
    } else {
        this.realDir = this.GlobalConfig.UserConfigDir()
    }

    if !this.GlobalConfig.Quiet() {
        logger.Printf("Input Configuration Directory: %s", this.GlobalConfig.UserConfigDir())
        logger.Printf("Evaluated Configuration Directory: %s", this.realDir)
    }

    fs := afero.NewOsFs()

    exists, err := afero.DirExists(fs, this.realDir)

    if err != nil {
        return err
    }

    if !exists {
        logger.Printf("Creating: %s", this.realDir)

        if err := fs.MkdirAll(this.realDir, this.GlobalConfig.DirConfig.FileMode()); err != nil {
            return err
        }
    }

    return nil
}
