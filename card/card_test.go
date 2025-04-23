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
