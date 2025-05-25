package player

import (
	"Crazy8s/card"
	"Crazy8s/hand"
	"testing"
)

func TestCreatePlayer(t *testing.T) {
	p := CreatePlayer()

	if p.GetType() != "human" {
		t.Errorf("Expected type 'human', got '%s'", p.GetType())
	}
	if p.GetDifficulty() != "" {
		t.Errorf("Expected strategy to be empty, got '%s'", p.GetDifficulty())
	}
	if p.PHand == nil {
		t.Error("Expected hand to be initialized")
	}
}

func TestCreateCPUPlayer(t *testing.T) {
	cpu := CreateCPUPlayer("optimal")

	if cpu.GetType() != "cpu" {
		t.Errorf("Expected type 'cpu', got '%s'", cpu.GetType())
	}
	if cpu.GetDifficulty() != "optimal" {
		t.Errorf("Expected strategy 'optimal', got '%s'", cpu.GetDifficulty())
	}
	if cpu.PHand == nil {
		t.Error("Expected hand to be initialized")
	}
}

func TestGetCardsByIndexes(t *testing.T) {
	c1 := card.NewCard(101, "hearts", "7")
	c2 := card.NewCard(102, "spades", "K")
	c3 := card.NewCard(103, "diamonds", "5")

	h := &hand.Hand{}
	h.AddCard(c1)
	h.AddCard(c2)
	h.AddCard(c3)

	p := &Player{
		name:  "Test Player",
		id:    999,
		PHand: h,
	}

	selected := p.GetCardsByIndexes([]int{101, 103})
	if len(selected) != 2 {
		t.Fatalf("Expected 2 cards, got %d", len(selected))
	}

	if selected[0].GetID() != 101 || selected[1].GetID() != 103 {
		t.Errorf("Expected IDs 101 and 103, got %d and %d", selected[0].GetID(), selected[1].GetID())
	}
}
