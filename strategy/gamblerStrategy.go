package strategy

import (
	"Crazy8s/card"
	"math/rand"
	"time"
)

type GamblerStrategy struct{}

func (s *GamblerStrategy) ChooseCards(hand []*card.Card, topCard *card.Card) string {
	viableMap := GetViablePlays(hand, topCard)

	// Skip turn if no viable plays
	if len(viableMap) == 0 {
		return "s"
	}

	playSet := GetRandomSet(viableMap)
	return CreatePlayCommand(playSet)
}

func GetRandomSet(viableMap map[int][]int) []int {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Convert map values to a slice of slices
	matrix := make([][]int, 0, len(viableMap))
	for _, val := range viableMap {
		matrix = append(matrix, val)
	}

	// Guard against empty input
	if len(matrix) == 0 {
		return nil
	}

	// Pick a random index and return the slice
	randomIndex := r.Intn(len(matrix))
	return matrix[randomIndex]

}

func (s *GamblerStrategy) HandleCrazy8(hand []*card.Card) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	random := r.Intn(3)

	switch random {
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
