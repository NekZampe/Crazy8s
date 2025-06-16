package card

import (
	"fmt"
	"strings"
)

func (c *Card) BuildCardAscii() string {
	val := c.GetValue()
	suit := c.GetSuitSymbol()

	valTop := fmt.Sprintf("%-2s", val)   // left-align
	valBottom := fmt.Sprintf("%2s", val) // right-align

	ascii := fmt.Sprintf(
		` _______ 
|%s     |
|   %s   |
|       |
|   %s   |
|_____%s|`, valTop, suit, suit, valBottom)

	return ascii
}

func (c *Card) GetSuitSymbol() string {
	switch strings.ToLower(c.GetSuit()) {
	case "hearts":
		return "♥"
	case "diamonds":
		return "♦"
	case "clubs":
		return "♣"
	case "spades":
		return "♠"
	default:
		return "?"
	}
}

func BuildCardBackAscii() string {
	return ` _______
| \ ~ / |
| }}:{{ |
| }}:{{ |
| }}:{{ |
| /_~_\ |`
}
