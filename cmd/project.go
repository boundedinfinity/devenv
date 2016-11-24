package cmd

import (
    "github.com/spf13/cobra"
    "github.com/boundedinfinity/devenv/file/makefile"
    "log"
    "github.com/boundedinfinity/devenv/file/editorconfig"
    "github.com/boundedinfinity/devenv/file/glide"
)

func init() {
    RootCommand.AddCommand(projectCommand)

    projectCommand.AddCommand(projectMakefileCommand)
    assignProjectFlags(projectMakefileCommand.Flags())
    assignFileFlags(projectMakefileCommand.Flags())

    projectCommand.AddCommand(projectEditorconfigCommand)
    assignFileFlags(projectEditorconfigCommand.Flags())

    projectCommand.AddCommand(projectGlideCommand)
    assignGoFlags(projectGlideCommand.Flags())
    assignFileFlags(projectGlideCommand.Flags())
}

var projectCommand = &cobra.Command{
    Use:   "project",
    Short: "Project file management",
    Long:  `Project file management`,
    Run:   func(cmd *cobra.Command, args []string) {
    },
}

var projectMakefileCommand = &cobra.Command{
    Use:   "makefile",
    Short: "Create a Makefile",
    Long:  `Create a Makefile`,
    Run:   func(cmd *cobra.Command, args []string) {
        manager := makefile.NewMakefileManager()
        if err := manager.Ensure(); err != nil {
            log.Printf("error: %s", err.Error())
        }
    },
}

var projectEditorconfigCommand = &cobra.Command{
    Use:   "editorconfig",
    Short: "Create an .editorconfig file",
    Long:  `Create an .editorconfig file`,
    Run:   func(cmd *cobra.Command, args []string) {
        manager := editorconfig.NewEditorConfigManager()
        if err := manager.Ensure(); err != nil {
            log.Printf("error: %s", err.Error())
        }
    },
}

var projectGlideCommand = &cobra.Command{
    Use:   "glide",
    Short: "Create a glide.yml file",
    Long:  `Create a glide.yml file`,
    Run:   func(cmd *cobra.Command, args []string) {
        manager := glide.NewGlideManager()
        if err := manager.Ensure(); err != nil {
            log.Printf("error: %s", err.Error())
        }
    },
}
