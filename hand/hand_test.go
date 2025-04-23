package hand_test

import (
	"Crazy8s/card"
	"Crazy8s/hand"
	"testing"
)

func TestAddCard(t *testing.T) {
	h := hand.Hand{}
	c := card.NewCard("Spades", "K")

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
	c1 := card.NewCard("Spades", "K")
	c2 := card.NewCard("Hearts", "A")

	h.AddCard(c1)
	h.AddCard(c2)

	h.RemoveCard(c1)

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

	h.AddCard(card.NewCard("Diamonds", "5"))
	h.AddCard(card.NewCard("Clubs", "J"))

	if h.GetCount() != 2 {
		t.Errorf("expected hand count to be 2, got %d", h.GetCount())
	}
}

func TestGetCards(t *testing.T) {
	h := hand.Hand{}
	c1 := card.NewCard("Spades", "7")
	c2 := card.NewCard("Hearts", "3")

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
	c1 := card.NewCard("Spades", "K")
	c2 := card.NewCard("Clubs", "2")
	c3 := card.NewCard("Hearts", "K")
	c4 := card.NewCard("Diamonds", "A")

	h.AddCard(c1)
	h.AddCard(c2)
	h.AddCard(c3)
	h.AddCard(c4)

	h.OrganizeHand()
	cards := h.GetCards()

	expectedOrder := []*card.Card{c4, c2, c3, c1} // A, 2, K of Hearts, K of Spades

	for i, expected := range expectedOrder {
		if cards[i] != expected {
			t.Errorf("at index %d, expected %v of %v, got %v of %v",
				i, expected.GetValue(), expected.GetSuit(),
				cards[i].GetValue(), cards[i].GetSuit())
		}
	}
}
