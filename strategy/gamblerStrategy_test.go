package strategy

import (
	"Crazy8s/card"
	"strings"
	"testing"
)

func TestGamblerStrategy_ChooseCards_NoPlayableCards(t *testing.T) {
	hand := []*card.Card{
		card.NewCard(1, "Hearts", "2"),
		card.NewCard(2, "Clubs", "4"),
		card.NewCard(3, "Diamonds", "7"),
	}
	top := card.NewCard(0, "Spades", "K")

	strategy := &GamblerStrategy{}
	result := strategy.ChooseCards(hand, top)

	if result != "s" {
		t.Errorf("Expected 's' (skip), got: %s", result)
	}
}

func TestGamblerStrategy_ChooseCards_PlayableExists(t *testing.T) {
	hand := []*card.Card{
		card.NewCard(1, "Spades", "2"),   // match suit
		card.NewCard(2, "Hearts", "K"),   // match rank
		card.NewCard(3, "Diamonds", "8"), // 8 (wild)
	}
	top := card.NewCard(0, "Spades", "K")

	strategy := &GamblerStrategy{}
	result := strategy.ChooseCards(hand, top)

	if !strings.HasPrefix(result, "play") {
		t.Errorf("Expected result to start with 'play', got: %s", result)
	}
}

func TestGamblerStrategy_HandleCrazy8(t *testing.T) {
	strategy := &GamblerStrategy{}
	validSuits := map[string]bool{
		"clubs":    true,
		"diamonds": true,
		"hearts":   true,
		"spades":   true,
	}

	for i := 0; i < 10; i++ { // Run multiple times to account for randomness
		suit := strategy.HandleCrazy8(nil)
		if !validSuits[suit] {
			t.Errorf("Invalid suit returned: %s", suit)
		}
	}
}

func TestGamblerStrategy_Name(t *testing.T) {
	strategy := &GamblerStrategy{}
	if strategy.Name() != "gambler" {
		t.Errorf("Expected 'gambler', got: %s", strategy.Name())
	}
}
