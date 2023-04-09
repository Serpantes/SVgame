package server

import (
	"net/http"

	game "github.com/Serpantes/SVgame/game"
	log "github.com/sirupsen/logrus"
)

var addr = "127.0.0.1:8081"

var Game *game.Game

func InitServer() {
	log.Info("Init hub")
	Hub := newHub()
	Game = game.NewGame("")
	go Hub.run()
	go refreshStateLoop(Hub)
	log.Info("Init server")
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fileHandler(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(Hub, w, r)
	})
	log.Info("Staring server on ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
