package phasesupplier

import (
	"github.com/nikiforosFreespirit/msdb5/app/phase"
)

type InputAction string

func (ia InputAction) ID() phase.ID {
	switch ia {
	case "Join":
		return phase.Joining
	case "Auction":
		return phase.InsideAuction
	case "Exchange":
		return phase.ExchangingCards
	case "Companion":
		return phase.ChosingCompanion
	case "Card":
		return phase.PlayingCards
	default:
		return phase.End
	}
}
