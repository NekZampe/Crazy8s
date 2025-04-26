package deck

import (
	"Crazy8s/card"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Deck struct {
	activePile  []*card.Card
	reservePile []*card.Card
	topCard     *card.Card
}

func (d *Deck) ActivePile() []*card.Card {
	return d.activePile
}

func (d *Deck) ReservePile() []*card.Card {
	return d.reservePile
}

func (d *Deck) TopCard() *card.Card {
	return d.topCard
}

var (
	instance *Deck
	once     sync.Once
)

// GetInstance returns the singleton deck instance
func GetInstance() *Deck {
	once.Do(func() {
		instance = &Deck{
			activePile:  []*card.Card{},
			reservePile: []*card.Card{},
		}
		instance.initializeDeck()
	})
	return instance
}

func (d *Deck) RefreshTopCard() {
	if len(d.activePile) > 0 {
		d.topCard = d.activePile[len(d.activePile)-1]
	}
}

// AddCardToActive adds a card to active
func (d *Deck) AddCardToActive(c *card.Card) {
	d.activePile = append(d.activePile, c)
}

// AddCardToReserve adds card to reserve pile
func (d *Deck) AddCardToReserve(c *card.Card) {
	d.reservePile = append(d.reservePile, c)
}

// RemoveCard : Removes top card from reserve pile
func (d *Deck) RemoveCard() *card.Card {
	if len(d.reservePile) == 0 {
		fmt.Println("Deck is empty")
		return nil
	}
	removedCard := d.reservePile[len(d.reservePile)-1]
	d.reservePile = d.reservePile[1:]
	return removedCard

}

// InitializeDeck : Creates deck, shuffles and refreshes top card
func (d *Deck) initializeDeck() {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{
		"A", "2", "3", "4", "5", "6", "7",
		"8", "9", "10", "J", "Q", "K",
	}

	for _, suit := range suits {
		for _, value := range values {
			c := card.NewCard(suit, value)
			d.reservePile = append(d.reservePile, c)
		}
	}
	d.ShuffleDeck()
}

// ShuffleDeck : shuffles the deck
func (d *Deck) ShuffleDeck() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.reservePile), func(i, j int) {
		d.reservePile[i], d.reservePile[j] = d.reservePile[j], d.reservePile[i]
	})
}
