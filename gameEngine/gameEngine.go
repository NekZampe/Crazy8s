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
}

func (g *Game) SetCurrentPlayer(idx int) {
	g.currentPlayerIndex = idx
}

func (g *Game) NextPlayer() {
	g.currentPlayerIndex = (g.currentPlayerIndex + 1) % len(g.playerList)
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

// distributeCards builds deck, distributes cards
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

// play card(s)
func (g *Game) PlaySingleCard(player *player.Player, card1 *card.Card) {
	topCard := g.gameDeck.GetTopCard()

	if topCard.EqualSuit(card1) || topCard.EqualValue(card1) {
		removedCard := player.PHand.RemoveCardFromHand(card1)
		g.gameDeck.AddCardToActive(removedCard)
		g.gameDeck.RefreshTopCard()
	}
}

func (g *Game) PlayDoubleCard(card1 *card.Card, card2 *card.Card) {
	topCard := g.gameDeck.GetTopCard()

}

func (g *Game) PlayTripleCard(card1 *card.Card, card2 *card.Card, card3 *card.Card) {
	topCard := g.gameDeck.GetTopCard()

}

func (g *Game) PlayQuadroCard(card1 *card.Card, card2 *card.Card, card3 *card.Card, card4 *card.Card) {
	topCard := g.gameDeck.GetTopCard()

}

func (g *Game) CheckWinner() {
	if g.playerList[g.currentPlayerIndex].PHand.GetCount() == 0 {
		fmt.Printf("END OF GAME, Player %s won!", g.playerList[g.currentPlayerIndex].GetPlayerName())
		os.Exit(1)
	}
}
