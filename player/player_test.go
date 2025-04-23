package player_test

import (
	"Crazy8s/player"
	"testing"
)

func TestCreatePlayer(t *testing.T) {
	p := player.CreatePlayer("Alice", 42)

	if p.GetPlayerName() != "Alice" {
		t.Errorf("expected name 'Alice', got '%s'", p.GetPlayerName())
	}

	if p.GetPlayerId() != 42 {
		t.Errorf("expected ID 42, got %d", p.GetPlayerId())
	}
}

func TestSetPlayerName(t *testing.T) {
	p := player.CreatePlayer("Temp", 1)
	p.SetPlayerName("Bob")

	if p.GetPlayerName() != "Bob" {
		t.Errorf("expected name 'Bob', got '%s'", p.GetPlayerName())
	}
}

func TestSetPlayerId(t *testing.T) {
	p := player.CreatePlayer("Temp", 0)
	p.SetPlayerId(99)

	if p.GetPlayerId() != 99 {
		t.Errorf("expected ID 99, got %d", p.GetPlayerId())
	}
}
