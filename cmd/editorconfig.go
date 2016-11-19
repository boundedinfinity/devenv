package cmd

import (
    "github.com/spf13/cobra"
    "github.com/boundedinfinity/devenv/file/editorconfig"
    "log"
)

var editorconfigCommand = &cobra.Command{
    Use:   "editorconfig",
    Short: "Create an .editorconfig file",
    Long:  `Create an .editorconfig file`,
    Run:   func(cmd *cobra.Command, args []string) {
        manager := editorconfig.NewEditorConfigManager()
        if err := manager.Write(); err != nil {
            log.Printf("error: %s", err.Error())
        }
    },
}

func init() {
    RootCommand.AddCommand(editorconfigCommand)

    assignFileFlags(editorconfigCommand.Flags())
}

