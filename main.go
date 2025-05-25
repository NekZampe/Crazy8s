package main

import (
	"Crazy8s/gameEngine"
	"Crazy8s/ilogger"
)

func main() {
	log, _ := ilogger.NewFileLogger("game.log")
	game := gameEngine.NewGame(log)
	game.Play()
}
