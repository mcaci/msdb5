package phasesupplier

import (
	"github.com/nikiforosFreespirit/msdb5/api/game"
)

type InputAction string

func (ia InputAction) Phase() game.Phase {
	switch ia {
	case "Join":
		return game.Joining
	case "Auction":
		return game.InsideAuction
	case "Exchange":
		return game.ExchangingCards
	case "Companion":
		return game.ChosingCompanion
	case "Card":
		return game.PlayingCards
	default:
		return game.End
	}
}
