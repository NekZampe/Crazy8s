package card

import (
	"fmt"
	"strings"
)

// Card structure
type Card struct {
	suit  string
	value string
}

func NewCard(suit string, value string) *Card {
	card := &Card{}
	card.SetValue(value)
	card.SetSuit(suit)
	return card
}

// Getters
func (c *Card) GetSuit() string {
	return c.suit
}

func (c *Card) GetValue() string {
	return c.value
}

// Setters
func (c *Card) SetSuit(suit string) {
	valid, err := c.VerifySuit(suit)

	if valid && err == nil {
		c.suit = suit
	} else {
		return
	}
}

func (c *Card) SetValue(value string) {
	valid, err := c.VerifyValue(value)
	if valid && err == nil {
		c.value = value
	} else {
		return
	}
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
