package cmd

import (
    "github.com/spf13/cobra"
    "github.com/boundedinfinity/devenv/file/glide"
    "log"
)

var glideCmd = &cobra.Command{
    Use:   "glide",
    Short: "Create a glide.yml file",
    Long:  `Create a glide.yml file`,
    Run:   glideCallback,
}

func glideCallback(cmd *cobra.Command, args []string) {
    manager := glide.NewGlideManager()

    if err := manager.Write(); err != nil {
        log.Printf("error: %s", err.Error())
    }
}

func init() {
    assignGoFlags(glideCmd.Flags())
    assignFileFlags(glideCmd.Flags())
    RootCmd.AddCommand(glideCmd)
}
