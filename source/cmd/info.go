package cmd

import (
    "fmt"
    // "gopkg.in/yaml.v2"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    log "github.com/sirupsen/logrus")

// infoCmd represents the list command
var infoCmd = &cobra.Command{
    Use:   "info",
    Short: "Show detailed information about chroot",
    Run: func(cmd *cobra.Command, args []string) {
        getInfo()
    },
    Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      log.Fatal("requires an argument")
    }
    if !isSystemValid(args[0]) {
      log.Fatal(args[0], " not a system in ", cfgFile, " configuration file")
    }
    // return fmt.Errorf("invalid color specified: %s", args[0])
    subArg = args[0]
    return nil
    },
}

func init() {
    rootCmd.AddCommand(infoCmd)
}

func getInfo(){
    getSystem(&selectSys)
    fmt.Println("name: ", selectSys.name)
    fmt.Println("path: ", selectSys.path)
    fmt.Println("url: ", selectSys.url)
    fmt.Println("dist: ", selectSys.dist)
    fmt.Println("user: ", selectSys.user)
    fmt.Println("groups: ", selectSys.groups)
    fmt.Println("xorg: ", selectSys.xorg)
}

func isSystemValid(name string) bool {
    toprint := viper.GetStringMapString("systems")
    for i := range toprint{
        if i == name { return true }
    }
    return false
}
