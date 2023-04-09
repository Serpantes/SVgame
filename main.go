package main

//April 3 2023

import (
	"github.com/Serpantes/SVgame/config"
	"github.com/Serpantes/SVgame/server"

	log "github.com/sirupsen/logrus"
)

func main() {
log.Info("Hello world")
config.InitConfig()
config.InitLogger()
server.InitServer()
log.Error("Done!")
}