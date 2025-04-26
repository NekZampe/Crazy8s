package hand_test

import (
	"Crazy8s/card"
	"Crazy8s/hand"
	"testing"
)

func TestAddCard(t *testing.T) {
	h := hand.Hand{}
	c := card.NewCard(1, "Spades", "K")

	h.AddCard(c)

	if h.GetCount() != 1 {
		t.Errorf("expected hand to have 1 card, got %d", h.GetCount())
	}

	if h.GetCards()[0] != c {
		t.Errorf("expected first card to be %v, got %v", c, h.GetCards()[0])
	}
}

func TestRemoveCard(t *testing.T) {
	h := hand.Hand{}
	c1 := card.NewCard(1, "Spades", "K")
	c2 := card.NewCard(2, "Hearts", "A")

	h.AddCard(c1)
	h.AddCard(c2)

	h.RemoveCardFromHand(c1)

	if h.GetCount() != 1 {
		t.Errorf("expected 1 card after removal, got %d", h.GetCount())
	}

	if h.GetCards()[0] != c2 {
		t.Errorf("expected remaining card to be %v, got %v", c2, h.GetCards()[0])
	}
}

func TestGetCount(t *testing.T) {
	h := hand.Hand{}

	if h.GetCount() != 0 {
		t.Errorf("expected empty hand, got count %d", h.GetCount())
	}

	h.AddCard(card.NewCard(1, "Diamonds", "5"))
	h.AddCard(card.NewCard(2, "Clubs", "J"))

	if h.GetCount() != 2 {
		t.Errorf("expected hand count to be 2, got %d", h.GetCount())
	}
}

func TestGetCards(t *testing.T) {
	h := hand.Hand{}
	c1 := card.NewCard(1, "Spades", "7")
	c2 := card.NewCard(2, "Hearts", "3")

	h.AddCard(c1)
	h.AddCard(c2)

	cards := h.GetCards()
	if len(cards) != 2 {
		t.Errorf("expected 2 cards, got %d", len(cards))
	}
	if cards[0] != c1 || cards[1] != c2 {
		t.Errorf("cards do not match what was added")
	}
}

func TestOrganizeHand(t *testing.T) {
	h := hand.Hand{}
	c1 := card.NewCard(6, "Spades", "K")
	c2 := card.NewCard(2, "Clubs", "2")
	c3 := card.NewCard(8, "Hearts", "K")
	c4 := card.NewCard(4, "Diamonds", "A")

	h.AddCard(c1)
	h.AddCard(c2)
	h.AddCard(c3)
	h.AddCard(c4)

	h.OrganizeHand()
	cards := h.GetCards()

	expectedOrder := []*card.Card{c2, c4, c1, c3}

	for i, expected := range expectedOrder {
		if cards[i] != expected {
			t.Errorf("at index %d, expected %v of %v, got %v of %v",
				i, expected.GetValue(), expected.GetSuit(),
				cards[i].GetValue(), cards[i].GetSuit())
		}
	}
}
