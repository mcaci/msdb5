package action

import (
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
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

func (ia InputAction) Find(p *player.Player) bool {
	switch ia {
	case "Join":
		return p.IsNameEmpty()
	default:
		return true
	}
}
