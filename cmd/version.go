package cmd

import (
    "github.com/spf13/cobra"
    "github.com/boundedinfinity/devenv/config"
    "fmt"
)

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Version information",
    Long:  `Version information`,
    Run:   func(cmd *cobra.Command, args []string) {
        globalConfig := config.GlobalConfig{}
        if globalConfig.Quiet() {
            fmt.Printf(config.Version)
        } else {
            fmt.Printf("Version: %s\n", config.Version)
        }
    },
}

func init() {
    RootCommand.AddCommand(versionCmd)
}
