package strategy

import "Crazy8s/card"

type PlayStrategy interface {
	ChooseCard(hand []*card.Card, topCard *card.Card) *card.Card
}
