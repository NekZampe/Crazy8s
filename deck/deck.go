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
}

func (d *Deck) GetActivePile() []*card.Card {
	return d.activePile
}

func (d *Deck) GetReservePile() []*card.Card {
	return d.reservePile
}

func (d *Deck) GetTopCard() *card.Card {
	pile := d.GetActivePile()
	if len(pile) == 0 {
		return nil
	}
	return pile[len(pile)-1]
}

func (d *Deck) GetReservePileCount() int { return len(d.activePile) }

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

// AddCardToActive adds a card to active
func (d *Deck) AddCardToActive(c *card.Card) {
	d.activePile = append(d.activePile, c)
}

// AddCardToReserve adds card to reserve pile
func (d *Deck) AddCardToReserve(c *card.Card) {
	d.reservePile = append(d.reservePile, c)
}

// RemoveCardFromReserveDeck : Removes top card from reserve pile
func (d *Deck) RemoveCardFromReserveDeck() *card.Card {
	if len(d.reservePile) == 0 {
		fmt.Println("Deck is empty")
		return nil
	}
	lastIndex := len(d.reservePile) - 1
	removedCard := d.reservePile[lastIndex]
	d.reservePile = d.reservePile[:lastIndex]
	return removedCard
}

// InitializeDeck : Creates and shuffles deck
func (d *Deck) initializeDeck() {
	index := 0
	suits := []string{"hearts", "diamonds", "clubs", "spades"}
	values := []string{
		"A", "2", "3", "4", "5", "6", "7",
		"8", "9", "10", "J", "Q", "K",
	}

	for _, value := range values {
		for _, suit := range suits {
			c := card.NewCard(index, suit, value)
			d.reservePile = append(d.reservePile, c)
			index++
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

func (d *Deck) PrintTopCard() {
	fmt.Println("Top Card:", d.GetTopCard().PrintCard())
}

func (d *Deck) ResetReservePile() {
	if len(d.activePile) > 1 {
		// Append all but the top card from activePile to reservePile
		d.reservePile = append(d.reservePile, d.activePile[:len(d.activePile)-1]...)
		// Keep only the top card in activePile
		d.activePile = d.activePile[len(d.activePile)-1:]
		// Shuffle
		d.ShuffleDeck()
	}
}
