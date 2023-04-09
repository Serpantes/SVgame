package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)
func fileHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("file recieved")
}