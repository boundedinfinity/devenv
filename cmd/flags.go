package cmd

import (
    flag "github.com/spf13/pflag"
    "github.com/boundedinfinity/devenv/config"
    "github.com/spf13/viper"
)

func assignGlobalFlags(flagSet *flag.FlagSet) {
    RootCmd.PersistentFlags().StringP(config.Flag_UserConfigDir, "c", config.Flag_Default_UserConfigDir,
        "Path to the user configuration directory")
    viper.BindPFlag(config.Flag_UserConfigDir, RootCmd.PersistentFlags().Lookup(config.Flag_UserConfigDir))

    flagSet.BoolP(config.Flag_Debug, "d", config.Flag_Default_Debug, "Enable debugging output")
    viper.BindPFlag(config.Flag_Debug, flagSet.Lookup(config.Flag_Debug))

    flagSet.BoolP(config.Flag_Quiet, "q", config.Flag_Default_Quiet, "Enable quiet mode - silence all output")
    viper.BindPFlag(config.Flag_Quiet, flagSet.Lookup(config.Flag_Quiet))
}

func assignProjectFlags(flagSet *flag.FlagSet) {
    flagSet.StringP(config.Flag_ProjectName, "n", config.Flag_Default_ProjectName, "Project name")
    viper.BindPFlag(config.Flag_ProjectName, flagSet.Lookup(config.Flag_ProjectName))
}

func assignFileFlags(flagSet *flag.FlagSet) {
    flagSet.BoolP(config.Flag_Overwrite, "o", config.Flag_Default_Overwrite, "Overwrite file if it exists")
    viper.BindPFlag(config.Flag_Overwrite, flagSet.Lookup(config.Flag_Overwrite))

    flagSet.StringP(config.Flag_ProjectPath, "p", config.Flag_Default_ProjectPath, "Project path")
    viper.BindPFlag(config.Flag_ProjectPath, flagSet.Lookup(config.Flag_ProjectPath))

    flagSet.Uint32P(config.Flag_FileMode, "m", config.Flag_Default_Flag_FileMode, "File mode")
    viper.BindPFlag(config.Flag_FileMode, flagSet.Lookup(config.Flag_FileMode))
}

func assignGoFlags(flagSet *flag.FlagSet) {
    flagSet.String(config.Flag_GoPackageName, config.Flag_Default_GoPackageName, "Go package name")
    viper.BindPFlag(config.Flag_GoPackageName, flagSet.Lookup(config.Flag_GoPackageName))
}
