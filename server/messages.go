package server

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

const refreshRate = 3 * time.Second

type message struct {
	Command string `json:"command"`
	Data map[string]interface{} `json:"data"`
}
func handleMessage(bytes []byte, c *Client) {
	log.Debug("Handling new message")
	var msg message
	json.Unmarshal(bytes, &msg)

	switch msg.Command{
	case "hello":
		log.Info("New player: ", c.conn.RemoteAddr())
	case "auth":
		authHandler(msg.Data["name"], c)
	}
}

func makeMessage(command string) message{
	var msg message
	msg.Data = make(map[string]interface{})
	msg.Command = command
	return msg
}

func authHandler(name interface{},c *Client) {
	log.Info("Player ", c.conn.RemoteAddr() , " is ", name, "!")

	result := "success"

	msg := makeMessage("result")

	for client := range c.hub.clients {
		if client.username == name {
			result = "error"
			msg.Data["reason"] = "Name taken"
		}
	}

	if result == "success" {
		c.username = fmt.Sprintf("%v", name)
	}
	
	msg.Data["command"] = "auth"
	msg.Data["result"] = result
	msg.Data["username"] = name
	var bytes []byte

	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Error("Can't marshall message ", err)
	}
	c.send <- bytes
}

func refreshStateLoop(hub *Hub){
	ticker := time.NewTicker(refreshRate)
	defer func() {
		ticker.Stop()
	}()

	for range ticker.C{
		msg := makeMessage("stateFull")
		var clients []string //TODO users for client are usernames only now. They can be map[string]string. but what [key] should i use?
		for client := range hub.clients {
			if client.username != "" {
				clients = append(clients, client.username)
			}
		}
		if len(clients) < 1 {
			log.Trace("No data for stateFull")
			continue
		}
		msg.Data["players"] = clients

		if Game.Active {
			msg.Data["game"] = Game
		} else {
			msg.Data["game"] = "inactive" 
		}

		bytes, err := json.Marshal(msg)
		if err != nil {
			log.Error("Can't marshal stateFull ", err)
		}

		hub.broadcast <- bytes
	}
}