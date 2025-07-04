package gameEngine

import "fmt"

type GameState string

const (
	StartMenu  GameState = "StartMenu"
	AddPlayers           = "AddPlayers"
	Start                = "Start"
	Deal                 = "DealCards"
	PlayerTurn           = "PlayerTurn"
	CheckWin             = "CheckWin"
	End                  = "EndGame"
)

func (g *Game) Transition(newState GameState) error {
	validTransitions := map[GameState][]GameState{
		StartMenu:  {AddPlayers},
		AddPlayers: {Start},
		Start:      {Deal},
		Deal:       {PlayerTurn},
		PlayerTurn: {CheckWin},
		CheckWin:   {End, PlayerTurn},
		End:        {Start},
	}

	valid := false
	for _, s := range validTransitions[g.state] {
		if s == newState {
			valid = true
			break
		}
	}

	if !valid {
		return fmt.Errorf("invalid transition from %s to %s", g.state, newState)
	}

	g.state = newState
	return nil
}
