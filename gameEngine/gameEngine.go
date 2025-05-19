package gameEngine

import (
	"Crazy8s/card"
	"Crazy8s/deck"
	"Crazy8s/player"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"runtime/debug"
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

func NewGame() *Game {
	return &Game{
		state: initial,
	}
}

func (g *Game) SetCurrentPlayer(idx int) {
	g.currentPlayerIndex = idx
}

func (g *Game) NextPlayer() {
	g.currentPlayerIndex = (g.currentPlayerIndex + 1) % len(g.playerList)
}

// addPlayersLocal add players to game locally
func (g *Game) addPlayersLocal() {
	reader := bufio.NewReader(os.Stdin)
	flag := true

	err := g.Transition(AddPlayers)
	if err != nil {
		fmt.Println("State transition error:", err)
		return
	}

	for flag {
		clearConsole()
		fmt.Println("\nAdd Player:")
		fmt.Println("1: New Human Player")
		fmt.Println("2: New CPU Player (Optimal)")
		fmt.Println("3: New CPU Player (Gambler)")
		fmt.Println("Type 'start' to begin playing.")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "start":
			if len(g.playerList) > 1 {
				flag = false
			} else {
				fmt.Println("Not enough players")
			}
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

	err := g.Transition(Deal)
	if err != nil {
		return
	}

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
func (g *Game) PlayCards(player *player.Player, cards []*card.Card) error {

	//fmt.Println("CARDS PASSED TO playCards:", cards)
	//Pick up cards if twos were played
	//TODO: Make into own function
	if g.countOf2s > 1 {
		for i := 0; i < g.countOf2s; i++ {
			g.PickUpCard(player)
		}
	}

	if len(cards) == 0 || len(cards) > 4 {
		return fmt.Errorf("invalid number of cards")
	}

	topCard := g.gameDeck.GetTopCard()

	// Validate first card against top card if it's not a crazy 8
	if cards[0].GetValue() != "8" {
		if !topCard.ValidatePlay(cards[0]) {
			return fmt.Errorf("first card does not match the top")
		}
	}

	// Validate all cards are the same value
	firstValue := cards[0].GetValue()
	for _, c := range cards {
		if c.GetValue() != firstValue {
			return fmt.Errorf("all played cards must have matching value")
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
		return nil
	case "J":
		g.countOfJacks = len(cards)
		return nil
	case "2":
		if g.countOf2s == 0 {
			g.countOf2s = len(cards)
		} else {
			g.countOf2s += len(cards)
		}
		return nil
	default:
		g.countOf2s = 0
		return nil
	}
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
	//Each player plays until one player is first to have no cards and is declared the winner

	for !g.IsGameOver {

		p := g.playerList[g.currentPlayerIndex]

		err := g.Transition(PlayerTurn)
		if err != nil {
			return
		}
		clearConsole()
		g.gameDeck.PrintTopCard()
		p.PHand.PrintHand()
		var request Request
		for {
			input := g.GetPlayerPlayInput()
			request = g.ParsePlayerRequest(input)

			if request.rType != "" {
				break
			}
			fmt.Println("Please enter a valid command.")
		}

		switch request.rType {
		case "p":
			cards := p.GetCardsByIndexes(request.cards)

			if len(cards) == 0 {
				fmt.Println("No valid cards selected.")
				continue
			}

			err := g.PlayCards(p, cards)
			if err != nil {
				fmt.Println("Invalid play:", err)
				err = g.Transition(CheckWin) //need to transition again
				if err != nil {
					return
				}
				fmt.Println("Please try again...")
				fmt.Println()
				continue // Let player retry
			}

			err = g.Transition(CheckWin)
			if err != nil {
				return
			}

			g.CheckWinner()
			g.NextPlayer()

		case "s":
			g.PickUpCard(p)
			err := g.Transition(CheckWin)
			if err != nil {
				return
			}
			g.NextPlayer()
		case "e":
			fmt.Println(p.GetPlayerName() + " chose to exit the game.")
			return
		default:
			fmt.Println("Invalid request. Please try again.")
		}

	}
}

func (g *Game) Play() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("❗ PANIC RECOVERED:", r)
			debug.PrintStack()
		}
	}()

	g.addPlayersLocal()
	g.initializeGame()
	g.mainLoop()
}

// TODO: Move into UI related folder
func clearConsole() {
	// Try ANSI escape codes first
	time.Sleep(100 * time.Millisecond) // Optional: prevents flicker
	fmt.Print("\033[2J\033[H")

	// Fallback: use OS command if ANSI is not supported
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
