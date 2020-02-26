package main

import (
	"github.com/bresilla/tent/source/cmd"
    log "github.com/sirupsen/logrus")

func main() {
 	log.SetFormatter(&log.TextFormatter{
        DisableTimestamp: true,
        DisableLevelTruncation: true,
	})


	cmd.Execute()
}
