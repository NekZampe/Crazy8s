package gameEngine

import (
	"Crazy8s/deck"
	"Crazy8s/player"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	playerCount = 0
)

type Game struct {
	playerList []*player.Player
	gameDeck   *deck.Deck
	state      GameState
}

// AddPlayers add players to game
func (g *Game) AddPlayers() {
	reader := bufio.NewReader(os.Stdin)
	flag := true

	err := g.Transition(AddPlayers)
	if err != nil {
		fmt.Println("State transition error:", err)
		return
	}

	for flag {
		fmt.Println("\nAdd Player:")
		fmt.Println("1: New Human Player")
		fmt.Println("2: New CPU Player (Optimal)")
		fmt.Println("3: New CPU Player (Gambler)")
		fmt.Println("Type 'start' to begin playing.")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "start":
			flag = false
		case "1":
			newPlayer := player.CreatePlayer()
			g.playerList = append(g.playerList, newPlayer)
			fmt.Println("Created new human newPlayer:", newPlayer.GetPlayerName())
		case "2":
			cpu := player.CreateCPUPlayer("optimal")
			g.playerList = append(g.playerList)
			fmt.Println("Created new CPU newPlayer (optimal):", cpu.GetPlayerName())
		case "3":
			cpu := player.CreateCPUPlayer("gambler")
			g.playerList = append(g.playerList, cpu)
			fmt.Println("Created new CPU newPlayer (gambler):", cpu.GetPlayerName())
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

// initialize game { builds deck, distribute cards }

func (g *Game) distributeCards() {

	g.gameDeck = deck.GetInstance()

}

// Pick up card {pick up card}
// organise hand {organizes hand}
// play card(s)
// Pick one up and skip turn
