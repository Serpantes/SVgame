package game

import (
	log "github.com/sirupsen/logrus"
)

type Game struct {
	Players []*Player
	Title   string
	Active bool
}

type Player struct {
	Username string
	Score int
}

func NewGame(name string) *Game {
	log.Debug("Creating new game '", name, "'")
	var game Game
	game.Title = name
	return &game
}

func NewPlayer(username string) *Player {
	var player Player
	player.Username = username
	return &player
}

func (game *Game) addPlayer(username string) {
	log.Debug("Adding new player ", username, " to game ", game.Title)
	player := NewPlayer(username)
	game.Players = append(game.Players, player)
}

func (game *Game) removePlayer(p Player) {
	log.Debug("Removing player ", p.Username, " from game ", game.Title)
	if len(game.Players) > 1 {
		for i, player := range game.Players {
			if player.Username == p.Username {
				game.Players[i] = game.Players[len(game.Players)-1]
				game.Players = game.Players[:len(game.Players)-1]
			}
		}
	}
	if len(game.Players) <= 1 {
		log.Debug("One or less players in game ", game.Title, " remaining")
		game.deleteGame()
	}
}

func (game Game) deleteGame() {
	log.Error("Game ", game.Title, "should be deleted!") //TODO delete game
}
