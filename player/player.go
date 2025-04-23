package player

import (
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
	username   string
	id         int
	hand       *hand.Hand
	playerType string // "human" or "cpu"
	difficulty string // empty for humans, or "optimal"/"gambler" for CPUs
}

// Getter methods
func (p *Player) GetPlayerName() string {
	return p.username
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
	p.username = name
}

func (p *Player) SetPlayerId(id int) {
	p.id = id
}

// CreatePlayer creates a human player
func CreatePlayer() *Player {
	totalPlayers++
	name := fmt.Sprintf("Player %d", totalPlayers)

	return &Player{
		username:   name,
		id:         generateUniqueID(),
		hand:       &hand.Hand{},
		playerType: "human",
		difficulty: "",
	}
}

// CreateCPUPlayer creates a CPU player with specified difficulty ("optimal" or "gambler")
func CreateCPUPlayer(difficulty string) *Player {
	totalCPUs++
	name := fmt.Sprintf("CPU %d", totalCPUs)

	return &Player{
		username:   name,
		id:         generateUniqueID(),
		hand:       &hand.Hand{},
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
