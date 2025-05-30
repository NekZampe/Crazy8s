package strategy

import (
	"Crazy8s/card"
	"fmt"
	"strconv"
	"strings"
)

func CreatePlayCommand(playSet []int) string {
	var builder strings.Builder
	builder.WriteString("play ")

	for _, val := range playSet {
		builder.WriteString(strconv.Itoa(val))
		builder.WriteString(" ")
	}

	return strings.TrimSpace(builder.String())
}

func GetViablePlays(hand []*card.Card, topCard *card.Card) map[int][]int {
	viableMap := make(map[int][]int)

	for i, c := range hand {
		if c == nil {
			continue // skip nil cards
		}
		if c.GetValue() == "8" || c.EqualValue(topCard) || c.EqualSuit(topCard) {
			currentSet := []int{i} // Start new set with this card

			for j, other := range hand {
				if j != i && other.EqualValue(c) {
					currentSet = append(currentSet, j)
				}
			}

			// Copy currentSet to avoid reference issues
			setCopy := make([]int, len(currentSet))
			copy(setCopy, currentSet)
			viableMap[i] = setCopy
		}
	}
	//Temp printout for troubleshooting
	fmt.Println("----------------------------------------------------")
	fmt.Println("VIABLE PLAYS: ")
	for key, value := range viableMap {
		fmt.Printf("Key: %d\n", key)
		fmt.Printf("Values: ")
		for _, num := range value {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------------------------")

	return viableMap
}
