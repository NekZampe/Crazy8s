package gameEngine

import (
	"os"
	"testing"
)

func TestParsePlayerRequest(t *testing.T) {
	g := &Game{} // assuming you have a Game type

	tests := []struct {
		input         string
		expectedType  string
		expectedCards []int
	}{
		{"play 5 10 15", "p", []int{5, 10, 15}},
		{"p 1 2 3 4 5", "p", []int{1, 2, 3, 4}}, // Only 4 cards max
		{"skip", "s", nil},
		{"refresh", "r", nil},
		{"exit", "e", nil},
		{"unknown", "", nil}, // invalid command
	}

	for _, tt := range tests {
		req := g.ParsePlayerRequest(tt.input)

		if req.rType != tt.expectedType {
			t.Errorf("Input %q: expected type %q, got %q", tt.input, tt.expectedType, req.rType)
		}

		if len(req.cards) != len(tt.expectedCards) {
			t.Errorf("Input %q: expected %d cards, got %d", tt.input, len(tt.expectedCards), len(req.cards))
			continue
		}

		for i, expectedCard := range tt.expectedCards {
			if req.cards[i] != expectedCard {
				t.Errorf("Input %q: card %d: expected %d, got %d", tt.input, i, expectedCard, req.cards[i])
			}
		}
	}
}

func TestGetPlayerInput(t *testing.T) {
	g := &Game{}

	// Mock stdin
	input := "Play 5 6\n"
	r, w, _ := os.Pipe()
	os.Stdin = r

	go func() {
		defer func(w *os.File) {
			err := w.Close()
			if err != nil {

			}
		}(w)
		_, err := w.Write([]byte(input))
		if err != nil {
			return
		}
	}()

	result := g.GetPlayerInput()

	expected := "play 5 6"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
