package server

import (
	"Crazy8s/network/protocol"
	"fmt"
	"net"
	"sync"
)

type Lobby struct {
	Players map[byte]net.Conn
	Mu      sync.Mutex // for safe concurrent access
}

// NewLobby : Creates a new empty lobby
func NewLobby() *Lobby {
	return &Lobby{
		Players: make(map[byte]net.Conn),
	}
}

// AddPlayer : Adds a new player to the lobby
func (l *Lobby) AddPlayer(playerID byte, conn net.Conn) {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	if len(l.Players) >= 3 {
		fmt.Println("Lobby full, cannot add new player")
		return
	}
	l.Players[playerID] = conn
	l.BroadcastLobbyState()
}

// RemovePlayer : Removes a player
func (l *Lobby) RemovePlayer(playerID byte) {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	delete(l.Players, playerID)
	fmt.Printf("[Lobby] Player %d left\n", playerID)
	l.BroadcastLobbyState()
}

// BroadcastLobbyState : Broadcasts the list of players
func (l *Lobby) BroadcastLobbyState() {
	l.Mu.Lock()
	// Build list of player IDs
	ids := ""
	for id := range l.Players {
		if ids != "" {
			ids += ","
		}
		ids += fmt.Sprintf("%d", id)
	}
	l.Mu.Unlock()

	// Build packet
	packet := protocol.BuildBroadcastLobbyState(ids)

	// Send
	l.Broadcast(packet)
}

// Broadcast sends a packet to all connected players
func (l *Lobby) Broadcast(packet []byte) {
	l.Mu.Lock()
	defer l.Mu.Unlock()

	for id, conn := range l.Players {
		err := protocol.SendMessage(conn, packet)
		if err != nil {
			fmt.Printf("[Lobby] Error sending to player %d: %v\n", id, err)
		}
	}
}
