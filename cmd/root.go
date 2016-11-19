package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
    "os"
    "github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
    Use:   "devenv",
    Short: "Tame your development environment",
    Long:  `Tame your development environment`,
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
    assignGlobalFlags(RootCmd.PersistentFlags())
    RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.devenv.yaml)")
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
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
