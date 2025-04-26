package gameEngine

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type request struct {
	rType string
	cards []int
}

// rTypes: Play [p] , Skip[s] , refresh[r]  exit[e]
// Cards: 1 2 3... [ max 4 card IDs ]

func (g *Game) ParsePlayerRequest(request string) {
	reader := bufio.NewReader(os.Stdin)
	flag := true

	for flag {
		fmt.Printf("Your turn:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		words := strings.Fields(input)
		re := regexp.MustCompile(`^[a-zA-Z]+$`)

		var validWords []string
		for _, word := range words {
			if re.MatchString(word) {
				validWords = append(validWords, word)
			}
		}

		flag = false
	}

}
