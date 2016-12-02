package cmd

import (
    //"github.com/boundedinfinity/devenv/user"
    "github.com/spf13/cobra"
)

func init() {
    RootCommand.AddCommand(userCommand)

    //userCommand.AddCommand(userConfigDirCommand)
    //assignDirFlags(userConfigDirCommand.Flags())
    //
    //userCommand.AddCommand(userGoCommand)
    //assignGoFlags(userGoCommand.Flags())
}

var userCommand = &cobra.Command{
    Use:   "user",
    Short: "User configuration management",
    Long:  `User configuration management`,
    Run:   func(cmd *cobra.Command, args []string) {
    },
}

//var userConfigDirCommand = &cobra.Command{
//    Use:   "configDir",
//    Short: "User configuration directory management",
//    Long:  `User configuration directory management`,
//    Run:   func(cmd *cobra.Command, args []string) {
//        manager := user.NewUserDirectoryManager()
//
//        if err := manager.Ensure(); err != nil {
//            logger.Printf("error: %s", err.Error())
//        }
//    },
//}
//
//var userGoCommand = &cobra.Command{
//    Use:   "go",
//    Short: "Configure go environment",
//    Long:  `Configure go environment`,
//    Run:   func(cmd *cobra.Command, args []string) {
//        manager := user.NewGoEnvManager()
//
//        if err := manager.Ensure(); err != nil {
//            logger.Printf("error: %s", err.Error())
//        }
//    },
//}
