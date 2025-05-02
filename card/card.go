package card

import (
	"fmt"
	"strings"
)

// Card structure
type Card struct {
	id    int
	suit  string
	value string
}

func NewCard(id int, suit string, value string) *Card {
	card := &Card{id, suit, value}
	return card
}

func (c *Card) SetSuit(s string) { c.suit = s }

// Getters
func (c *Card) GetSuit() string {
	return c.suit
}

func (c *Card) GetValue() string {
	return c.value
}

func (c *Card) GetID() int {
	return c.id
}

// ♠ ♥ ♦ ♣
func (c *Card) VerifySuit(suit string) (bool, error) {
	suit = strings.ToLower(suit)

	validSuits := []string{"hearts", "diamonds", "clubs", "spades"}

	for i := 0; i < len(validSuits); i++ {
		validSuit := validSuits[i]
		if suit == validSuit {
			c.suit = suit
			return true, nil
		}
	}
	return false, fmt.Errorf("invalid suit: %s", suit)
}

func (c *Card) VerifyValue(value string) (bool, error) {

	validValues := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	for i := 0; i < len(validValues); i++ {
		validValue := validValues[i]
		if value == validValue {
			c.value = value
			return true, nil
		}
	}
	return false, fmt.Errorf("invalid suit: %s", value)
}

func (c *Card) EqualSuit(other *Card) bool {
	return c.GetSuit() == other.GetSuit()
}

func (c *Card) EqualValue(other *Card) bool {
	return c.GetValue() == other.GetValue()
}

func (c *Card) ValidatePlay(other *Card) bool {
	return c.EqualValue(other) || c.EqualSuit(other)
}
