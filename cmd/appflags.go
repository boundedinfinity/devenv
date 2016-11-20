package cmd

import (
    flag "github.com/spf13/pflag"
    "github.com/boundedinfinity/devenv/config"
    "github.com/spf13/viper"
    "log"
)

// Flag initialization

var appFlags *flag.FlagSet
var debug = false

func init() {
    if appFlags == nil {
        if debug {
            log.Printf("=======================> init flags")
        }

        appFlags = flag.NewFlagSet("application-flags", flag.PanicOnError)

        initGlobalFlags()
        initProjectFlags()
        initFileFlags()
        initGoFlags()
    }
}

func initGlobalFlags() {
    appFlags.StringP(config.Flag_UserConfigDir, "c", config.Flag_Default_UserConfigDir, "Path to the user configuration directory")
    appFlags.BoolP(config.Flag_Debug, "d", config.Flag_Default_Debug, "Enable debugging output")
    appFlags.BoolP(config.Flag_Quiet, "q", config.Flag_Default_Quiet, "Enable quiet mode - silence all output")
}

func initProjectFlags() {
    appFlags.StringP(config.Flag_ProjectName, "n", config.Flag_Default_ProjectName, "Project name")
}

func initFileFlags() {
    appFlags.StringP(config.Flag_ProjectPath, "p", config.Flag_Default_ProjectPath, "Project path")
    appFlags.BoolP(config.Flag_Overwrite, "o", config.Flag_Default_Overwrite, "Overwrite file if it exists")
    //appFlags.Uint32P(config.Flag_FileMode, "m", config.Flag_Default_Flag_FileMode, "File mode")
    appFlags.StringP(config.Flag_FileMode, "m", config.Flag_Default_Flag_FileMode, "File mode")
}

func initGoFlags() {
    appFlags.String(config.Flag_GoPackageName, config.Flag_Default_GoPackageName, "Go package name")
}

// Flag group assignment

func assignGlobalFlags(flagSet *flag.FlagSet) {
    bindFlag(flagSet, config.Flag_UserConfigDir)
    bindFlag(flagSet, config.Flag_Debug)
    bindFlag(flagSet, config.Flag_Quiet)
}

func assignProjectFlags(flagSet *flag.FlagSet) {
    bindFlag(flagSet, config.Flag_ProjectName)
}

func assignFileFlags(flagSet *flag.FlagSet) {
    bindFlag(flagSet, config.Flag_ProjectPath)
    bindFlag(flagSet, config.Flag_Overwrite)
    bindFlag(flagSet, config.Flag_FileMode)
}

func assignDirFlags(flagSet *flag.FlagSet) {
    bindFlag(flagSet, config.Flag_FileMode)
}

func assignGoFlags(flagSet *flag.FlagSet) {
    bindFlag(flagSet, config.Flag_GoPackageName)
}

// Utility functions

func bindFlag(flagSet *flag.FlagSet, flagName string) {
    flag := appFlags.Lookup(flagName)
    flagSet.AddFlag(flag)
    viper.BindPFlag(flagName, flag)
}


