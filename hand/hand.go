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

// Add card
func (h *Hand) AddCard(c *card.Card) {
	h.cards = append(h.cards, c)
}

// Remove card
func (h *Hand) RemoveCard(target *card.Card) {
	for i, c := range h.cards {
		if c.GetSuit() == target.GetSuit() && c.GetValue() == target.GetValue() {
			// Remove the card by slicing
			h.cards = append(h.cards[:i], h.cards[i+1:]...)
			break
		}
	}
}

// Get card count
func (h *Hand) GetCount() int {
	return len(h.cards)
}

// Get cards
func (h *Hand) GetCards() []*card.Card {
	return h.cards
}

// OrganizeHand : Sort cards by value, then by suit
func (h *Hand) OrganizeHand() {
	// Sort the cards in h.cards slice
	sort.Slice(h.cards, func(i, j int) bool {
		// First, sort by card value
		if cardValue(h.cards[i].GetValue()) != cardValue(h.cards[j].GetValue()) {
			return cardValue(h.cards[i].GetValue()) < cardValue(h.cards[j].GetValue())
		}
		// If the values are the same, sort by suit
		return h.cards[i].GetSuit() < h.cards[j].GetSuit()
	})
}

// Helper function to map card values to an order
func cardValue(value string) int {
	switch value {
	case "A":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "10":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	}
	return 0 // Default value in case of an invalid card value
}

// PrintHand prints the cards in the hand with their index
func (h *Hand) PrintHand() {
	for i, c := range h.cards {
		fmt.Printf("Index: %d, Card: %s of %s\n", i, c.GetValue(), c.GetSuit())
	}
}
