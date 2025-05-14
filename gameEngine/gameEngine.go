package gameEngine

import (
	"Crazy8s/card"
	"Crazy8s/deck"
	"Crazy8s/player"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	playerCount = 0
	rng         = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Game struct {
	currentPlayerIndex int
	playerList         []*player.Player
	gameDeck           *deck.Deck
	state              GameState
	countOf2s          int
	countOfJacks       int
	IsGameOver         bool
}

func (g *Game) SetCurrentPlayer(idx int) {
	g.currentPlayerIndex = idx
}

func (g *Game) NextPlayer() {
	g.currentPlayerIndex = (g.currentPlayerIndex + 1) % len(g.playerList)
}

// AddPlayersLocal add players to game locally
func (g *Game) AddPlayersLocal() {
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

// distributeCards builds deck, distributes cards, ( only done at start of game )
func (g *Game) distributeCards() {

	g.gameDeck = deck.GetInstance()

	for _, currentP := range g.playerList {
		for i := 0; i < 8; i++ {
			currentP.PHand.AddCard(g.gameDeck.RemoveCardFromReserveDeck())
		}
	}
}

func (g *Game) ShufflePlayers() {
	rng.Shuffle(len(g.playerList), func(i, j int) {
		g.playerList[i], g.playerList[j] = g.playerList[j], g.playerList[i]
	})
}

func (g *Game) setTopCard() {
	g.gameDeck.AddCardToActive(g.gameDeck.RemoveCardFromReserveDeck())
	g.gameDeck.RefreshTopCard()
}

func (g *Game) initializeGame() {

	err := g.Transition(Start)
	if err != nil {
		fmt.Println("State transition error:", err)
		return
	}

	g.IsGameOver = false
	g.distributeCards()
	g.ShufflePlayers()
	g.setTopCard()
	g.currentPlayerIndex = 0

}

// PickUpCard : Player takes top card from reserve deck (played on start of turn )
func (g *Game) PickUpCard(player *player.Player) {
	player.PHand.AddCard(g.gameDeck.RemoveCardFromReserveDeck())
}

// PlayCards play up to 4 cards at once
func (g *Game) PlayCards(player *player.Player, cards []*card.Card) {

	//Pick up cards if twos were played
	//TODO: Make into own function
	if g.countOf2s > 1 {
		for i := 0; i < g.countOf2s; i++ {
			g.PickUpCard(player)
		}
	}

	if len(cards) == 0 || len(cards) > 4 {
		fmt.Println("Invalid number of cards")
		return
	}

	topCard := g.gameDeck.GetTopCard()

	// Validate first card against top card if it's not a crazy 8
	if cards[0].GetValue() != "8" {
		if !topCard.ValidatePlay(cards[0]) {
			fmt.Println("First card does not match the top")
			return
		}
	}

	// Validate all cards are the same value
	firstValue := cards[0].GetValue()
	for _, c := range cards {
		if c.GetValue() != firstValue {
			fmt.Println("All played cards must have the same value")
			return
		}
	}

	// All checks passed, play the cards
	for _, c := range cards {
		g.gameDeck.AddCardToActive(player.PHand.RemoveCardFromHand(c))
	}
	g.gameDeck.RefreshTopCard()

	switch firstValue {
	case "8":
		suit := g.GetPlayerC8Input()
		g.gameDeck.GetTopCard().SetSuit(suit) //If we set suit of top card to player requested value we ensure only requested suit can be played on next turn
	case "J":
		g.countOfJacks = len(cards)
	case "2":
		if g.countOf2s == 0 {
			g.countOf2s = len(cards)
		} else {
			g.countOf2s += len(cards)
		}
	default:
		g.countOf2s = 0
	}
	g.CheckWinner() //Check to see if player has zero cards left
}

// CheckWinner : Placeholder win function
func (g *Game) CheckWinner() {
	if g.playerList[g.currentPlayerIndex].PHand.GetCount() == 0 {
		fmt.Printf("END OF GAME, Player %s won!", g.playerList[g.currentPlayerIndex].GetPlayerName())
		g.IsGameOver = true
		os.Exit(1)
	}
}

func (g *Game) IsValidPlay(player *player.Player, cards []int) bool {
	handMap := make(map[int]bool)
	for _, c := range player.PHand.GetCards() {
		handMap[c.GetID()] = true
	}
	for _, c := range cards {
		if !handMap[c] {
			return false // Card is not in the hand
		}
	}

	return true // All cards are valid
}

func (g *Game) mainLoop() {

	//Players join
	//Cards shuffled and distributed
	//each player plays until one has no cards and is declared the winner

	for !g.IsGameOver {

	}
}
