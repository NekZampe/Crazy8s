package hand

import (
	"Crazy8s/card"
	"fmt"
	"sort"
)

type Hand struct {
	count int
	cards []*card.Card
}

// AddCard adds card to hand
func (h *Hand) AddCard(c *card.Card) {
	h.cards = append(h.cards, c)
}

// RemoveCardFromHand Removes a card from the players hand
func (h *Hand) RemoveCardFromHand(target *card.Card) *card.Card {
	for i, c := range h.cards {
		if c.GetID() == target.GetID() {
			h.cards = append(h.cards[:i], h.cards[i+1:]...)
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
// TODO : Organizes hand for each print might lead to unnecessary calculations, check for more optimized method later on
func (h *Hand) PrintHand() {
	h.OrganizeHand()
	for i, c := range h.cards {
		fmt.Printf("Index: %d, Card: %s of %s\n", i, c.GetValue(), c.GetSuit())
	}
}
