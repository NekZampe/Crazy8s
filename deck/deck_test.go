package deck_test

import (
	"Crazy8s/card"
	"Crazy8s/deck"
	"testing"
)

// Helper to reset singleton state for testing
func resetDeckInstanceForTest() {
	// Only possible by exposing a non-exported `resetInstance()` in the `deck` package
	// or avoiding singleton during testing
}

func TestInitializeDeck(t *testing.T) {
	d := deck.GetInstance()

	totalCards := len(d.GetReservePile())
	if totalCards != 52 {
		t.Errorf("Expected 52 cards in reserve pile, got %d", totalCards)
	}
}

func TestAddCardToActive(t *testing.T) {
	d := deck.GetInstance()

	c := d.GetReservePile()[51]

	d.AddCardToActive(c)

	if len(d.GetActivePile()) == 0 {
		t.Errorf("Expected card in active pile")
	}
	if d.GetTopCard() != c {
		t.Errorf("Expected top card to be the newly added one")
	}
}

func TestAddCardToReserve(t *testing.T) {
	d := deck.GetInstance()
	initial := len(d.GetReservePile())
	c := card.NewCard(333, "Spades", "K")

	d.AddCardToReserve(c)

	if len(d.GetReservePile()) != initial+1 {
		t.Errorf("Card was not added to reserve pile")
	}
}

func TestRemoveCard(t *testing.T) {
	d := deck.GetInstance()

	initial := len(d.GetReservePile())
	if initial == 0 {
		t.Skip("Reserve pile is empty, skipping test")
	}

	d.RemoveCardFromReserveDeck()

	if len(d.GetReservePile()) != initial-1 {
		t.Errorf("Card was not removed from reserve pile")
	}
}

func TestShuffleDeck(t *testing.T) {
	d := deck.GetInstance()

	before := make([]*card.Card, len(d.GetReservePile()))
	copy(before, d.GetReservePile())

	d.ShuffleDeck()

	sameOrder := true
	for i := range before {
		if before[i] != d.GetReservePile()[i] {
			sameOrder = false
			break
		}
	}

	if sameOrder {
		t.Log("Deck appears to have same order after shuffle (unlikely)")
	}
}

func TestPrintTopCard(t *testing.T) {
	d := deck.GetInstance()
	d.AddCardToActive(d.RemoveCardFromReserveDeck())
	d.PrintTopCard()
}

func TestResetReservePile(t *testing.T) {
	d := deck.GetInstance()

	d.AddCardToActive(d.RemoveCardFromReserveDeck())
	d.AddCardToActive(d.RemoveCardFromReserveDeck())
	d.AddCardToActive(d.RemoveCardFromReserveDeck())
	d.RefreshTopCard()

	top := d.GetTopCard()

	d.ResetReservePile()

	if len(d.GetReservePile()) != 51 {
		t.Errorf("Not all cards returned to reserve pile")
	}
	if top != d.GetTopCard() {
		t.Errorf("TopCard changed unexpectedly")
	}
}
