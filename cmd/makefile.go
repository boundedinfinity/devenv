package cmd

import (
    "github.com/spf13/cobra"
    "github.com/boundedinfinity/devenv/file/makefile"
    "log"
)

var makefileCmd = &cobra.Command{
    Use:   "makefile",
    Short: "Create a Makefile",
    Long:  `Create a Makefile`,
    Run:   makefileCallback,
}

func makefileCallback(cmd *cobra.Command, args []string) {
    manager := makefile.NewMakefileManager()

    if err := manager.Write(); err != nil {
        log.Printf("error: %s", err.Error())
    }
}

func init() {
    assignProjectFlags(makefileCmd.Flags())
    assignFileFlags(makefileCmd.Flags())

    RootCmd.AddCommand(makefileCmd)
}
