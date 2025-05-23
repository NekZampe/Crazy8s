package strategy

import (
	"Crazy8s/card"
	"testing"
)

func TestChooseCards_NoPlayableCards(t *testing.T) {
	hand := []*card.Card{
		card.NewCard(1, "Hearts", "2"),
		card.NewCard(2, "Clubs", "4"),
		card.NewCard(3, "Diamonds", "7"),
	}
	top := card.NewCard(0, "Spades", "K")

	strategy := &OptimalStrategy{}
	result := strategy.ChooseCards(hand, top)

	if result != "s" {
		t.Errorf("Expected 's' (skip), got: %s", result)
	}
}

func TestChooseCards_Play8(t *testing.T) {
	hand := []*card.Card{
		card.NewCard(1, "Hearts", "8"),
		card.NewCard(2, "Clubs", "4"),
	}
	top := card.NewCard(0, "Spades", "K")

	strategy := &OptimalStrategy{}
	result := strategy.ChooseCards(hand, top)

	if result != "play 0" {
		t.Errorf("Expected 'play 0', got: %s", result)
	}
}

func TestChooseCards_PlayMatchingSuit(t *testing.T) {
	hand := []*card.Card{
		card.NewCard(1, "Spades", "3"), // matches suit
		card.NewCard(2, "Hearts", "3"),
		card.NewCard(3, "Diamonds", "3"),
	}
	top := card.NewCard(0, "Spades", "K")

	strategy := &OptimalStrategy{}
	result := strategy.ChooseCards(hand, top)

	expected := "play 0 1 2" // or any group starting with a match and all 3s
	if result != expected {
		t.Errorf("Expected '%s', got: %s", expected, result)
	}
}

func TestChooseCards_PlayMatchingValue(t *testing.T) {
	hand := []*card.Card{
		card.NewCard(1, "Hearts", "K"),
		card.NewCard(2, "Diamonds", "K"),
		card.NewCard(3, "Clubs", "K"),
	}
	top := card.NewCard(0, "Spades", "K")

	strategy := &OptimalStrategy{}
	result := strategy.ChooseCards(hand, top)

	if result != "play 0 1 2" {
		t.Errorf("Expected 'play 0 1 2', got: %s", result)
	}
}

func TestHandleCrazy8(t *testing.T) {
	hand := []*card.Card{
		card.NewCard(1, "hearts", "A"),
		card.NewCard(2, "diamonds", "Q"),
		card.NewCard(3, "clubs", "K"),
		card.NewCard(4, "clubs", "K"),
		card.NewCard(5, "clubs", "K"),
		card.NewCard(6, "clubs", "K"),
		card.NewCard(7, "clubs", "K"),
	}

	strategy := &OptimalStrategy{}
	result := strategy.HandleCrazy8(hand)

	if result != "clubs" {
		t.Errorf("Expected clubs, got: %s", result)
	}
}
