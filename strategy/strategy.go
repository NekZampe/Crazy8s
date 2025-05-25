package strategy

import "Crazy8s/card"

type PlayStrategy interface {
	ChooseCards(hand []*card.Card, topCard *card.Card) string
	HandleCrazy8(hand []*card.Card) string
	Name() string
}
