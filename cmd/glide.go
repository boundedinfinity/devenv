package cmd

import (
    "github.com/spf13/cobra"
    "github.com/boundedinfinity/devenv/file/glide"
    "log"
)

var glideCommand = &cobra.Command{
    Use:   "glide",
    Short: "Create a glide.yaml file",
    Long:  `Create a glide.yaml file`,
    Run:   func(cmd *cobra.Command, args []string) {
        manager := glide.NewGlideManager()
        if err := manager.Write(); err != nil {
            log.Printf("error: %s", err.Error())
        }
    },
}

func init() {
    RootCommand.AddCommand(glideCommand)

    assignGoFlags(glideCommand.Flags())
    assignFileFlags(glideCommand.Flags())
}
