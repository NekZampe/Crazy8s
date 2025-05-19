package strategy

import (
	"Crazy8s/card"
	"strconv"
	"strings"
)

type OptimalStrategy struct{}

func (s *OptimalStrategy) ChooseCards(hand []*card.Card, topCard *card.Card) string {
	viableMap := make(map[int][]int)

	for i, c := range hand {
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

	// Skip turn if no viable plays
	if len(viableMap) == 0 {
		return "s"
	}

	playSet := GetLargestSet(viableMap)
	return CreatePlayCommand(playSet)
}

func GetLargestSet(viableMap map[int][]int) []int {
	var largest []int
	maxLen := -1

	for _, slice := range viableMap {
		if len(slice) > maxLen {
			maxLen = len(slice)
			largest = slice
		}
	}

	return largest
}

func CreatePlayCommand(playSet []int) string {
	var builder strings.Builder
	builder.WriteString("play ")

	for _, val := range playSet {
		builder.WriteString(strconv.Itoa(val))
		builder.WriteString(" ")
	}

	return strings.TrimSpace(builder.String())
}
