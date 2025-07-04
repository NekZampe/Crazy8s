package player

import (
	"Crazy8s/card"
	"Crazy8s/hand"
	"Crazy8s/strategy"
	"fmt"
	"math/rand"
	"time"
)

var (
	usedIDs      = make(map[byte]bool)
	rng          = rand.New(rand.NewSource(time.Now().UnixNano()))
	totalPlayers = 0
	totalCPUs    = 0
)

type Player struct {
	name       string
	id         byte
	PHand      *hand.Hand
	playerType string                // "human" or "cpu"
	Strategy   strategy.PlayStrategy // empty for humans, or "optimal"/"gambler" for CPUs
}

// GetPlayerName Getter methods
func (p *Player) GetPlayerName() string {
	return p.name
}

func (p *Player) GetPlayerId() byte {
	return p.id
}

func (p *Player) GetType() string {
	return p.playerType
}

func (p *Player) GetStrategy() string {
	if p.Strategy == nil {
		return "human"
	}
	return p.Strategy.Name()
}

// SetPlayerName Setter methods
func (p *Player) SetPlayerName(name string) {
	p.name = name
}

func (p *Player) SetPlayerId(id byte) {
	p.id = id
}

// CreatePlayer creates a human player
func CreatePlayer() *Player {
	totalPlayers++
	name := fmt.Sprintf("Player %d", totalPlayers)

	return &Player{
		name:       name,
		id:         generateUniqueID(),
		PHand:      &hand.Hand{},
		playerType: "human",
		Strategy:   nil,
	}
}

// CreateCPUPlayer creates a CPU player with the specified strategy ("optimal" or "gambler")
func CreateCPUPlayer(strategyName string) *Player {
	totalCPUs++
	name := fmt.Sprintf("CPU %d", totalCPUs)

	var strat strategy.PlayStrategy
	switch strategyName {
	case "optimal":
		strat = &strategy.OptimalStrategy{}
	case "gambler":
		strat = &strategy.GamblerStrategy{}
	default:
		strat = nil
	}

	return &Player{
		name:       name,
		id:         generateUniqueID(),
		PHand:      &hand.Hand{},
		playerType: "cpu",
		Strategy:   strat,
	}
}

// generateUniqueID returns a unique ID between 1 and 255
func generateUniqueID() byte {
	for {
		id := byte(rng.Intn(255) + 1) // range: 1–255 (or 0–255 if you prefer)
		if !usedIDs[id] {
			usedIDs[id] = true
			return id
		}
	}
}

// GetCardsByIndexes Returns the players cards by index ( used with requestProcessor to retrieve played cards )
func (p *Player) GetCardsByIndexes(indexes []int) []*card.Card {
	var selected []*card.Card
	cards := p.PHand.GetCards()

	for _, idx := range indexes {
		if idx >= 0 && idx < len(cards) {
			selected = append(selected, cards[idx])
		} else {
			fmt.Printf("Warning: index %d is out of range\n", idx)
		}
	}
	return selected
}
