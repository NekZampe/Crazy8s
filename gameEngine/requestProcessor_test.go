package gameEngine

import (
	"testing"
)

func TestParsePlayerRequest(t *testing.T) {
	g := &Game{}

	tests := []struct {
		input    string
		expType  string
		expCards []int
	}{
		{"p 1 2 3 4", "p", []int{1, 2, 3, 4}},
		{"play 10 20 30", "p", []int{10, 20, 30}},
		{"skip", "s", nil},
		{"s", "s", nil},
		{"exit", "e", nil},
		{"e", "e", nil},
		{"invalid command", "", nil},
		{"p 1 2 3 4 5 6", "p", []int{1, 2, 3, 4}}, // max 4 cards only
		{"p abc 1 2", "p", []int{1, 2}},           // ignores invalid numbers
		{"", "", nil},                             // empty input returns empty Request
	}

	for _, tt := range tests {
		req := g.ParsePlayerRequest(tt.input)

		if req.rType != tt.expType {
			t.Errorf("Input %q: expected rType %q, got %q", tt.input, tt.expType, req.rType)
		}

		if len(req.cards) != len(tt.expCards) {
			t.Errorf("Input %q: expected %d cards, got %d", tt.input, len(tt.expCards), len(req.cards))
			continue
		}

		for i := range req.cards {
			if req.cards[i] != tt.expCards[i] {
				t.Errorf("Input %q: expected card %d at pos %d, got %d", tt.input, tt.expCards[i], i, req.cards[i])
			}
		}
	}
}
