package nextphase

import (
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/game"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

type NextPhaseStruct struct {
	current  game.Phase
	players  team.Players
	sideDeck bool
	request  string
}

func NewChanger(current game.Phase, players team.Players, sideDeck bool,
	request string) action.NextPhaseChanger {
	return &NextPhaseStruct{current, players, sideDeck, request}
}

func (nps NextPhaseStruct) NextPhase() game.Phase {
	current, nextPhase := nps.current, nps.current+1
	predicateToNextPhase := func() bool { return true }
	switch nps.current {
	case game.Joining:
		predicateToNextPhase = func() bool {
			return team.Count(nps.players, func(p *player.Player) bool { return p.IsNameEmpty() }) == 0
		}
	case game.InsideAuction:
		predicateToNextPhase = func() bool {
			return team.Count(nps.players, func(p *player.Player) bool { return p.Folded() }) == 4
		}
		if !nps.sideDeck {
			nextPhase = nps.current + 2
		}
	case game.ExchangingCards:
		predicateToNextPhase = func() bool {
			data := strings.Split(nps.request, "#")
			if len(data) > 1 {
				number, err := strconv.Atoi(data[1])
				return number == 0 || err != nil
			}
			return false
		}
	case game.ChosingCompanion:
		nextPhase = game.PlayingCards
	case game.PlayingCards:
		predicateToNextPhase = func() bool {
			return team.Count(nps.players, func(p *player.Player) bool { return p.IsHandEmpty() }) == 5
		}
	default:
		nextPhase = game.End
	}
	if predicateToNextPhase() {
		return nextPhase
	}
	return current
}
