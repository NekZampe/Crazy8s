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
}

func (g *Game) SetCurrentPlayer(idx int) {
	g.currentPlayerIndex = idx
}

func (g *Game) NextPlayer() {
	g.currentPlayerIndex = (g.currentPlayerIndex + 1) % len(g.playerList)
}

// AddPlayersLocal add players to game
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

// distributeCards builds deck, distributes cards, only done at start of game
func (g *Game) distributeCards() {

	g.gameDeck = deck.GetInstance()

	for _, currentP := range g.playerList {
		for i := 0; i < 8; i++ {
			currentP.PHand.AddCard(g.gameDeck.RemoveCardFromDeck())
		}
	}
}

func (g *Game) ShufflePlayers() {
	rng.Shuffle(len(g.playerList), func(i, j int) {
		g.playerList[i], g.playerList[j] = g.playerList[j], g.playerList[i]
	})
}

func (g *Game) setTopCard() {
	g.gameDeck.AddCardToActive(g.gameDeck.RemoveCardFromDeck())
	g.gameDeck.RefreshTopCard()
}

func (g *Game) initializeGame() {
	g.distributeCards()
	g.ShufflePlayers()
	g.setTopCard()
	g.currentPlayerIndex = 0

}

// PickUpCard : Player takes top card from reserve deck (played on start of turn )
func (g *Game) PickUpCard(player *player.Player) {
	player.PHand.AddCard(g.gameDeck.RemoveCardFromDeck())
}

// PlayCards play up to 4 cards at once
func (g *Game) PlayCards(player *player.Player, cards []*card.Card) {
	if len(cards) == 0 || len(cards) > 4 {
		fmt.Println("Invalid number of cards")
		return
	}

	topCard := g.gameDeck.GetTopCard()

	// Validate first card against top card
	if !topCard.ValidatePlay(cards[0]) {
		fmt.Println("First card does not match the top")
		return
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
	g.CheckWinner()
}

func (g *Game) PlayCrazy8Card() {

}

func (g *Game) PlayJackCard() {

}

func (g *Game) PlayTwosCard() {

}

func (g *Game) PlayCardSelector(player *player.Player, cards []*card.Card) {
	switch cards[0].GetValue() {
	case "8":
		g.PlayCrazy8Card()
	case "J":
		g.PlayJackCard()
	case "2":
		g.PlayTwosCard()
	default:
		g.PlayCards(player, cards)
	}
}

func (g *Game) CheckWinner() {
	if g.playerList[g.currentPlayerIndex].PHand.GetCount() == 0 {
		fmt.Printf("END OF GAME, Player %s won!", g.playerList[g.currentPlayerIndex].GetPlayerName())
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
