package file

import (
    "github.com/boundedinfinity/devenv/shell"
    "github.com/boundedinfinity/devenv/data"
    "github.com/boundedinfinity/devenv/config"
    "github.com/boundedinfinity/devenv/logging"
    "github.com/spf13/afero"
    "github.com/Sirupsen/logrus"
    "path/filepath"
    "fmt"
    "os"
    "errors"
    "strings"
    "bytes"
    "text/template"
)

var engine = template.New("template")

type FileExistMode int

const (
    FailOnExists FileExistMode = iota
    IgnoreIfExists
    OverwriteIfExists
)

type DirectoryDescriptor struct {
    FsPath     string
    FileMode   os.FileMode
    ExistMode  FileExistMode
    ExpandPath bool
}

type TemplateCopyFileDescriptor struct {
    TemplatePath string
    TemplateData interface{}
    FsPath       string
    FileMode     os.FileMode
    ExpandPath   bool
    ExistMode    FileExistMode
}

type SymbolicLinkFileDescriptor struct {
    TemplatePath string
    Target       string
}

func NewFileManager2(logger *logrus.Entry) *FileManager2 {
    return &FileManager2{
        logger: logging.SubComponentLogger(logger, "FileManager"),
        GlobalConfig: config.NewGlobalConfig(),
    }
}

type FileManager2 struct {
    logger       *logrus.Entry
    GlobalConfig config.GlobalConfig
}

func (this *FileManager2) CalcRealPath(input string, expandPath bool) (string, error) {
    realPath := input

    if expandPath && strings.Contains(realPath, "$") {
        expandedPath, err := shell.Evaluate(fmt.Sprintf("echo -n %s", realPath))

        if err != nil {
            return realPath, err
        }

        realPath = expandedPath
    }

    realPath, err := filepath.Abs(realPath)

    if err != nil {
        return realPath, err
    }

    return realPath, nil
}

func (this *FileManager2) CreateDir(desc DirectoryDescriptor) error {
    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Input Dir Path: %s", desc.FsPath)
    }

    realFsPath, err := this.CalcRealPath(desc.FsPath, desc.ExpandPath)

    if err != nil {
        return err
    }

    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Absolute Dir Path: %s", realFsPath)
    }

    fs := afero.NewOsFs()

    exists, err := afero.DirExists(fs, realFsPath)

    if err != nil {
        return err
    }

    if exists {
        switch desc.ExistMode {
        case IgnoreIfExists:
            return nil
        case FailOnExists:
            return errors.New("already exists")
        case OverwriteIfExists:
            return nil
        }
    }

    if err := fs.Mkdir(realFsPath, desc.FileMode); err != nil {
        return err
    }

    return nil
}

func (this *FileManager2) DeleteDir(desc DirectoryDescriptor) error {
    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Input Dir Path: %s", desc.FsPath)
    }

    realFsPath, err := this.CalcRealPath(desc.FsPath, desc.ExpandPath)

    if err != nil {
        return err
    }

    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Absolute Dir Path: %s", realFsPath)
    }

    fs := afero.NewOsFs()

    exists, err := afero.DirExists(fs, realFsPath)

    if err != nil {
        return err
    }

    if exists {
        if err := fs.RemoveAll(realFsPath); err != nil {
            return err
        }
    }

    return nil
}

func (this *FileManager2) CopyFile(desc TemplateCopyFileDescriptor) error {
    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Input File Path: %s", desc.FsPath)
    }

    realFsPath, err := this.CalcRealPath(desc.FsPath, desc.ExpandPath)

    if err != nil {
        return err
    }

    if !this.GlobalConfig.Quiet() {
        this.logger.Infof("Absolute File Path: %s", realFsPath)
    }

    fs := afero.NewOsFs()

    exists, err := afero.Exists(fs, realFsPath)

    if err != nil {
        return err
    }

    if exists {
        switch desc.ExistMode {
        case IgnoreIfExists:
            return nil
        case FailOnExists:
            return errors.New("already exists")
        case OverwriteIfExists:
            if err := fs.Remove(realFsPath); err != nil {
                return nil
            }
        }
    }

    if this.GlobalConfig.Debug() {
        this.logger.Infof("Template File Path: %s", desc.FsPath)
    }

    content, err := data.Asset(desc.TemplatePath)

    if err != nil {
        return err
    }

    buffer := new(bytes.Buffer)
    template, err := engine.Parse(string(content))

    if err := template.Execute(buffer, desc.TemplateData); err != nil {
        return err
    }

    if err := afero.WriteFile(fs, realFsPath, buffer.Bytes(), desc.FileMode); err != nil {
        return err
    }

    return nil
}
