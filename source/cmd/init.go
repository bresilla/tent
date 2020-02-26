package cmd

import (
    "io/ioutil"
    "os"
	"github.com/spf13/cobra"
    "github.com/spf13/viper"
    log "github.com/sirupsen/logrus"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
        if (initUpdateFlag){
            updateInitTent()
            return
        }
        initTent()
	},
}

var initUpdateFlag = false
var initForceFlag = false

func init() {
	rootCmd.AddCommand(initCmd)
    initCmd.Flags().BoolVar(&initUpdateFlag, "update", false, "update initial settings with newer version ones")
    initCmd.Flags().BoolVar(&initForceFlag, "force", false, "force initialisation of new settings")
}

func updateInitTent(){
    log.Info("UPDATING...")
}
func initTent(){
    if (exists(cfgDir)){
        if (!exists(cfgFile)){
            err := ioutil.WriteFile(cfgFile, yamlExample, 0777)
            if err != nil { log.Fatal(err) }
            log.Info(cfgFile, " created")
        } else {
            if initForceFlag {
                err := ioutil.WriteFile(cfgFile, yamlExample, 0777)
                if err != nil { log.Fatal(err) }
                log.Info(cfgFile, " overrided")
            } else {
                log.Fatal(cfgFile, " already exists, please use --force to override")
            }
        }
    } else {
        err := os.MkdirAll(cfgDir, 0777)
        if err != nil {
            log.Fatal(err)
        } else {
            log.Info(cfgDir, " created")
            err := ioutil.WriteFile(cfgFile, yamlExample, 0777)
            if err != nil { log.Fatal(err) }
        }
        log.Info(cfgFile, " created")
        // err = os.Chown(cfgDir, os.Getuid(), os.Getgid())
        // err = os.Chown(cfgDir, 1000, 1000)
        if err != nil { log.Fatal(err) }
    }
}

var yamlExample = []byte(`
systems:
  ubuntu:
    path: "/opt/chroot/ubuntu"
    url: "https://"
    dist: ubuntu
    user: bresilla
    groups:
      - "storage"
      - "video"
    xorg: true
  arch:
    path: "/opt/chroot/arch"
    dist: archlinux
    url: "https://"
    user: bresilla
    groups:
      - "storage"
      - "video"
    xorg: false
`)

type system struct {
    name string
    path string
    url string
    dist string
    user string
    groups []string
    xorg bool
}

func getSystem(system *system){
    system.name = subArg
    system.path = viper.GetString("systems." + subArg + ".path")
    system.url = viper.GetString("systems." + subArg + ".url")
    system.dist = viper.GetString("systems." + subArg + ".dist")
    system.user = viper.GetString("systems." + subArg + ".user")
    system.groups = viper.GetStringSlice("systems." + subArg + ".groups")
    system.xorg = viper.GetBool("systems." + subArg + ".xorg")
}
