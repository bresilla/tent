package cmd

import (
    "os"
    "github.com/kyoh86/xdg"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    log "github.com/sirupsen/logrus"
)

var name string = "tent"
var cfgDir string = xdg.ConfigHome() + "/" + name;
var cfgFile string
var rootFlag bool = false
var subArg string
var selectSys system

var rootCmd = &cobra.Command{
    Use:   "tent",
    Short: "a better chroot/jail manager",
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", cfgDir + "/config.yml", "specify a config file")
}

func initConfig() {
    if cfgFile != "" {
        viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))
        viper.SetConfigFile(cfgFile)
    } else {
        viper.AddConfigPath(cfgDir)
        viper.SetConfigName("config.yml")
        log.Info(cfgDir)
    }
    viper.SetConfigType("yaml")
    viper.AutomaticEnv()
    if os.Args[1] != "init" {
        err := viper.ReadInConfig() // Find and read the config file
        if err != nil { // Handle errors reading the config file
            log.Error(err)
            log.Error("use ", name, " init first\n")
        }
    }
}
