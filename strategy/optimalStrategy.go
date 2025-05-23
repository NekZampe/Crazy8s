package strategy

import (
	"Crazy8s/card"
)

type OptimalStrategy struct{}

// ChooseCards finds all viable choices of playable cards of the cpu and selects the one with the most cards played, returns the play string ex: 'play 2 3 4'
func (s *OptimalStrategy) ChooseCards(hand []*card.Card, topCard *card.Card) string {
	viableMap := GetViablePlays(hand, topCard)

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

// HandleCrazy8 selects suit based on the highest freq suit in cpu's hand
func (s *OptimalStrategy) HandleCrazy8(hand []*card.Card) string {

	// Order of suits: Clubs, Diamonds, Hearts, and Spades
	countOfSuits := []int{0, 0, 0, 0}

	// total up number of each suit in hand
	for _, c := range hand {
		switch c.GetSuit() {
		case "clubs":
			countOfSuits[0]++
		case "diamonds":
			countOfSuits[1]++
		case "hearts":
			countOfSuits[2]++
		case "spades":
			countOfSuits[3]++
		}
	}

	//Get max index and return the suit string
	maximum := GetMaxIndex(countOfSuits)

	switch maximum {
	case 0:
		return "clubs"
	case 1:
		return "diamonds"
	case 2:
		return "hearts"
	case 3:
		return "spades"
	default:
		return "hearts"
	}

}

func GetMaxIndex(array []int) int {

	maxIdx := 0
	maxVal := -1

	for i, v := range array {
		if v > maxVal {
			maxVal = v
			maxIdx = i
		}
	}

	return maxIdx
}
