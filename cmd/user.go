package cmd

import (
    "github.com/spf13/cobra"
    "log"
    "github.com/boundedinfinity/devenv/user"
)

func init() {
    RootCommand.AddCommand(userCommand)

    userCommand.AddCommand(userConfigDirCommand)
    assignDirFlags(userConfigDirCommand.Flags())
}

var userCommand = &cobra.Command{
    Use:   "user",
    Short: "User configuration management",
    Long:  `User configuration management`,
    Run:   func(cmd *cobra.Command, args []string) {
    },
}

var userConfigDirCommand = &cobra.Command{
    Use:   "configDir",
    Short: "User configuration directory management",
    Long:  `User configuration directory management`,
    Run:   func(cmd *cobra.Command, args []string) {
        manager := user.NewUserConfigManager()
        if err := manager.EnsureConfigDir(); err != nil {
            log.Printf("error: %s", err.Error())
        }
    },
}


