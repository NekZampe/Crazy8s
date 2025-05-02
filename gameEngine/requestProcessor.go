package gameEngine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Request struct {
	rType string
	cards []int
}

// rTypes: play[p] , Skip[s] , refresh[r] , exit[e]
// Cards: 1 4 7... [ max 4 card IDs ]

func (g *Game) GetPlayerPlayInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Your turn: ")
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
				fmt.Println("Invalid number:", words[i])
				continue
			}
			if len(r.cards) < 4 {
				r.cards = append(r.cards, num)
			} else {
				break
			}
		}
		return r
	case "skip", "s":
		r.rType = "s"
		return r
	case "refresh", "r":
		r.rType = "r"
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
