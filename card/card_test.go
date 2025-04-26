package card

import "testing"

func TestVerifySuit(t *testing.T) {
	tests := []struct {
		input   string
		isValid bool
	}{
		{"hearts", true},
		{"diamonds", true},
		{"clubs", true},
		{"spades", true},
		{"banana", false},
		{"", false},
	}

	for _, test := range tests {
		card := &Card{}
		_, err := card.VerifySuit(test.input)
		if (err == nil) != test.isValid {
			t.Errorf("VerifySuit(%s): expected valid=%v, got err=%v", test.input, test.isValid, err)
		}
	}
}

func TestVerifyValue(t *testing.T) {
	tests := []struct {
		input   string
		isValid bool
	}{
		{"9", true},
		{"K7", false},
		{"5", true},
		{"Q", true},
		{"0", false},
		{"", false},
	}

	for _, test := range tests {
		card := &Card{}
		_, err := card.VerifyValue(test.input)
		if (err == nil) != test.isValid {
			t.Errorf("VerifyValue(%s): expected valid=%v, got err=%v", test.input, test.isValid, err)
		}
	}
}

func TestComparisons(t *testing.T) {
	card1 := NewCard(1, "spade", "3")
	card2 := NewCard(2, "spade", "Q")
	card3 := NewCard(3, "diamonds", "3")

	tests := []struct {
		name     string
		got      bool
		expected bool
	}{
		{"Same suit (spade vs spade)", card1.EqualSuit(card2), true},
		{"Different suits (spade vs diamonds)", card1.EqualSuit(card3), false},
		{"Different values (3 vs Q)", card1.EqualValue(card2), false},
		{"Same value (3 vs 3)", card1.EqualValue(card3), true},
	}

	for _, tt := range tests {
		if tt.got != tt.expected {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, tt.got)
		}
	}
}
