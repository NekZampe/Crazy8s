package gameEngine

import (
	"Crazy8s/player"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Request struct {
	player *player.Player
	rType  string
	cards  []int
}

// rTypes: play[p] , skip[s] , exit[e]
// Cards: 1 4 7... [ max 4 card IDs ]

// GetPlayerPlayInput shows the player it's their turn and their current hand, requests an input that is parsed in ParsePlayerRequest
func (g *Game) GetPlayerPlayInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Your turn: ")
	printPlayOptions()
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input))
}

func (g *Game) ParsePlayerRequest(input string) Request {
	words := strings.Fields(input)

	if len(words) == 0 {
		fmt.Println("No input detected.")
		return Request{}
	}

	r := Request{}

	switch words[0] {
	case "play", "p":
		r.rType = "p"
		for i := 1; i < len(words); i++ {
			num, err := strconv.Atoi(words[i])
			if err != nil {
				fmt.Println("Invalid index:", words[i])
				continue
			}
			if len(r.cards) < 4 {
				r.cards = append(r.cards, num)
			} else {
				break
			}
		}
		//fmt.Printf("%v\n", r.cards)
		return r
	case "skip", "s":
		r.rType = "s"
		return r
	case "exit", "e":
		r.rType = "e"
		return r
	default:
		fmt.Println("Error reading user request.")
		return Request{}
	}
}

func (g *Game) GetPlayerC8Input() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Entire desired suit: hearts[h] , diamonds[d], spades[s] , clovers[c]")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "h", "heart", "hearts":
			return "hearts"
		case "d", "diamond", "diamonds":
			return "diamonds"
		case "s", "spade", "spades":
			return "spades"
		case "c", "clover", "clovers":
			return "clovers"
		default:
			fmt.Printf("Error Invalid Input: %s\n", input)
		}
	}
}

func printPlayOptions() {
	fmt.Println("Options:")
	fmt.Println("  play [p]  - Play cards by specifying their indexes. Example: 'play 2 4 5'")
	fmt.Println("  skip [s]  - Skip your turn and pick up a card")
	fmt.Println("  exit [e]  - Leave the game")
}
