package player

import (
	"Crazy8s/card"
	"Crazy8s/hand"
	"fmt"
	"math/rand"
	"time"
)

var (
	usedIDs      = make(map[int]bool)
	rng          = rand.New(rand.NewSource(time.Now().UnixNano()))
	totalPlayers = 0
	totalCPUs    = 0
)

type Player struct {
	name       string
	id         int
	PHand      *hand.Hand
	playerType string // "human" or "cpu"
	difficulty string // empty for humans, or "optimal"/"gambler" for CPUs
}

// Getter methods
func (p *Player) GetPlayerName() string {
	return p.name
}

func (p *Player) GetPlayerId() int {
	return p.id
}

func (p *Player) GetType() string {
	return p.playerType
}

func (p *Player) GetDifficulty() string {
	return p.difficulty
}

// Setter methods
func (p *Player) SetPlayerName(name string) {
	p.name = name
}

func (p *Player) SetPlayerId(id int) {
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
		difficulty: "",
	}
}

// CreateCPUPlayer creates a CPU player with specified difficulty ("optimal" or "gambler")
func CreateCPUPlayer(difficulty string) *Player {
	totalCPUs++
	name := fmt.Sprintf("CPU %d", totalCPUs)

	return &Player{
		name:       name,
		id:         generateUniqueID(),
		PHand:      &hand.Hand{},
		playerType: "cpu",
		difficulty: difficulty,
	}
}

// generateUniqueID returns a unique 5-digit ID
func generateUniqueID() int {
	for {
		id := rng.Intn(90000) + 10000 // range: 10000–99999
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
