package file

import "os"

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
