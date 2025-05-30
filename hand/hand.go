package hand

import (
	"Crazy8s/card"
	"fmt"
	"sort"
	"strings"
)

type Hand struct {
	count int // For debugging
	cards []*card.Card
}

// AddCard adds card to hand
func (h *Hand) AddCard(c *card.Card) {
	h.cards = append(h.cards, c)
	h.count = len(h.cards)
}

// RemoveCardFromHand Removes a card from the players hand
func (h *Hand) RemoveCardFromHand(target *card.Card) *card.Card {
	for i, c := range h.cards {
		if c.GetID() == target.GetID() {
			h.cards = append(h.cards[:i], h.cards[i+1:]...)
			h.count = len(h.cards)
			return c
		}
	}
	return nil
}

// GetCount Gets card count
func (h *Hand) GetCount() int {
	return len(h.cards)
}

// GetCards Gets cards
func (h *Hand) GetCards() []*card.Card {
	return h.cards
}

// OrganizeHand Sorts cards in players Hand
func (h *Hand) OrganizeHand() {
	sort.Slice(h.cards, func(i, j int) bool {
		return h.cards[i].GetID() < h.cards[j].GetID()
	})
}

// PrintHand prints the cards in the hand with their index
// TODO : Organizes hand for each print might lead to unnecessary overhead, check for more optimized method later on
func (h *Hand) PrintHand() {
	h.OrganizeHand()
	for i, c := range h.cards {
		fmt.Printf("Index: %d, Card: %s of %s\n", i, c.GetValue(), c.GetSuit())
	}
}

func (h *Hand) PrintHandUI() {
	h.OrganizeHand()

	// Collect ASCII lines for all cards
	cardLines := [][]string{}
	for _, c := range h.cards {
		ascii := c.BuildCardAscii()
		lines := strings.Split(ascii, "\n")
		cardLines = append(cardLines, lines)
	}

	// Determine the number of lines per card (usually 6)
	numLines := len(cardLines[0])

	// Print all cards side-by-side
	for i := 0; i < numLines; i++ {
		for _, card := range cardLines {
			// Ensure fixed width for each line
			fmt.Printf("%-9s", card[i]) // adjust width as needed
		}
		fmt.Println()
	}

	// Print indexes underneath
	for i := range h.cards {
		fmt.Printf("   [%d]   ", i)
	}
	fmt.Println()
}
