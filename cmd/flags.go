package cmd

import (
    flag "github.com/spf13/pflag"
    "github.com/boundedinfinity/devenv/config"
    "github.com/spf13/viper"
)

// Flag groups

func assignGlobalFlags(flagSet *flag.FlagSet) {
    stringFlag(flagSet, config.Flag_UserConfigDir, "c", config.Flag_Default_UserConfigDir, "Path to the user configuration directory")
    boolFlag(flagSet, config.Flag_Debug, "d", config.Flag_Default_Debug, "Enable debugging output")
    boolFlag(flagSet, config.Flag_Quiet, "q", config.Flag_Default_Quiet, "Enable quiet mode - silence all output")
}

func assignProjectFlags(flagSet *flag.FlagSet) {
    stringFlag(flagSet, config.Flag_ProjectName, "n", config.Flag_Default_ProjectName, "Project name")
}

func assignFileFlags(flagSet *flag.FlagSet) {
    stringFlag(flagSet, config.Flag_ProjectPath, "p", config.Flag_Default_ProjectPath, "Project path")
    boolFlag(flagSet, config.Flag_Overwrite, "o", config.Flag_Default_Overwrite, "Overwrite file if it exists")
    uint32Flag(flagSet, config.Flag_FileMode, "m", config.Flag_Default_Flag_FileMode, "File mode")
}

func assignGoFlags(flagSet *flag.FlagSet) {
    stringFlag(flagSet, config.Flag_GoPackageName, "", config.Flag_Default_GoPackageName, "Go package name")
}

// Cobra configure and Viper bind

func stringFlag(flagSet *flag.FlagSet, flag string, short string, defaultArg string, description string) {
    if short == "" {
        flagSet.String(flag, defaultArg, description)
    } else {
        flagSet.StringP(flag, short, defaultArg, description)
    }

    viper.BindPFlag(flag, flagSet.Lookup(flag))
}

func boolFlag(flagSet *flag.FlagSet, flag string, short string, defaultArg bool, description string) {
    flagSet.BoolP(flag, short, defaultArg, description)
    viper.BindPFlag(flag, flagSet.Lookup(flag))
}

func uint32Flag(flagSet *flag.FlagSet, flag string, short string, defaultArg uint32, description string) {
    flagSet.Uint32P(flag, short, defaultArg, description)
    viper.BindPFlag(flag, flagSet.Lookup(flag))
}
