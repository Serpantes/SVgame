package main

import (
	"github.com/Serpantes/SVgame/config"
	"github.com/Serpantes/SVgame/engine"
	"github.com/Serpantes/SVgame/server"
	"github.com/Serpantes/SVgame/storage"

	log "github.com/sirupsen/logrus"
)

func main() {
log.Info("Hello world")
config.InitConfig()
config.InitLogger()
server.InitServer()
storage.InitStorage()
engine.InitEngine()
log.Error("DOne!")
}