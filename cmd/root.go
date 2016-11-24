package cmd

import (
    "github.com/boundedinfinity/devenv/logging"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "os"
)

var cfgFile string
var logger = logging.ComponentLogger("RootCommand")

var RootCommand = &cobra.Command{
    Use:   "devenv",
    Short: "Tame your development environment",
    Long:  `Tame your development environment`,
}

func Execute() {
    if err := RootCommand.Execute(); err != nil {
        logger.Errorf(err.Error())
        os.Exit(-1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
    assignGlobalFlags(RootCommand.PersistentFlags())
    RootCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.devenv.yaml)")
}

func initConfig() {
    if cfgFile != "" {
        // enable ability to specify config file via flag
        viper.SetConfigFile(cfgFile)
    }

    viper.SetConfigName(".devenv") // name of config file (without extension)
    viper.AddConfigPath("$HOME")   // adding home directory as first search path
    viper.AutomaticEnv()           // read in environment variables that match

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        logger.Printf("Using config file:", viper.ConfigFileUsed())
    }
}
