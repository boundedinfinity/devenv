package cmd

import (
    "github.com/spf13/cobra"
    "github.com/boundedinfinity/devenv/file/makefile"
    "log"
)

var makefileCommand = &cobra.Command{
    Use:   "makefile",
    Short: "Create a Makefile",
    Long:  `Create a Makefile`,
    Run:   func(cmd *cobra.Command, args []string) {
        manager := makefile.NewMakefileManager()
        if err := manager.Write(); err != nil {
            log.Printf("error: %s", err.Error())
        }
    },
}

func init() {
    RootCommand.AddCommand(makefileCommand)

    assignProjectFlags(makefileCommand.Flags())
    assignFileFlags(makefileCommand.Flags())
}
