package gameEngine

import "fmt"

func (g *Game) mainLoopMultiplayer() {

	for !g.IsGameOver {

		// ADD BREAKPOINT HERE <-------------------------------------

		// On start of turn ensure enough cards in reserve pile
		if g.gameDeck.GetReservePileCount() < 8 {
			g.gameDeck.ResetReservePile()
			g.logger.Info("RESHUFFLE RESERVE PILE")
		}

		p := g.playerList[g.currentPlayerIndex]

		g.logger.Info("Current player " + p.GetPlayerName())

		err := g.Transition(PlayerTurn)
		if err != nil {
			return
		}

		var request Request
		g.logger.Info("TOP CARD: " + g.gameDeck.GetTopCard().PrintCard())

		if p.GetType() == "human" {
			for {
				g.gameDeck.PrintTopCardUI()
				fmt.Println("     ------------------------ YOUR HAND ------------------------")
				p.PHand.PrintHandUI()

				input := g.GetPlayerPlayInput()
				request = g.ParsePlayerRequest(input)

				g.logger.Debug("INPUT: " + input + " ,received from: " + p.GetPlayerName())

				if request.rType != "" {
					break
				}
				fmt.Println("Please enter a valid command.")
			}
		} else {
			input := p.Strategy.ChooseCards(p.PHand.GetCards(), g.gameDeck.GetTopCard())
			request = g.ParsePlayerRequest(input)
			g.logger.Debug("INPUT: " + input + " ,received from: " + p.GetPlayerName())
		}

		switch request.rType {
		case "p":
			cards := p.GetCardsByIndexes(request.cards)

			if len(cards) == 0 {
				fmt.Println("No valid cards selected.")
				err = g.Transition(CheckWin)
				continue
			}

			err := g.PlayCards(p, cards)
			if err != nil {
				fmt.Println("Invalid play:", err)
				err = g.Transition(CheckWin) //need to transition again
				if err != nil {
					return
				}
				fmt.Println("Please try again...")
				fmt.Println()
				continue // Let player retry
			}

			err = g.Transition(CheckWin)
			if err != nil {
				return
			}

			g.CheckWinner()
			g.NextPlayer()

		case "s":
			g.PickUpCard(p)
			err := g.Transition(CheckWin)
			if err != nil {
				return
			}
			g.NextPlayer()
		case "e":
			fmt.Println(p.GetPlayerName() + " chose to exit the game.")
			return
		default:
			fmt.Println("Invalid request. Please try again.")
		}

	}
}
