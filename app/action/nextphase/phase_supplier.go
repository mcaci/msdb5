package nextphase

import (
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

type NextPhaseStruct struct {
	current  phase.ID
	players  team.Players
	sideDeck bool
	request  string
}

func NewChanger(current phase.ID, players team.Players, sideDeck bool,
	request string) action.NextPhaseChanger {
	return &NextPhaseStruct{current, players, sideDeck, request}
}

func (nps NextPhaseStruct) NextPhase() phase.ID {
	current, nextPhase := nps.current, nps.current+1
	predicateToNextPhase := func() bool { return true }
	switch nps.current {
	case phase.Joining:
		predicateToNextPhase = func() bool {
			return team.Count(nps.players, func(p *player.Player) bool { return p.IsNameEmpty() }) == 0
		}
	case phase.InsideAuction:
		predicateToNextPhase = func() bool {
			return team.Count(nps.players, func(p *player.Player) bool { return p.Folded() }) == 4
		}
		if !nps.sideDeck {
			nextPhase = nps.current + 2
		}
	case phase.ExchangingCards:
		predicateToNextPhase = func() bool {
			data := strings.Split(nps.request, "#")
			if len(data) > 1 {
				number, err := strconv.Atoi(data[1])
				return number == 0 || err != nil
			}
			return false
		}
	case phase.ChosingCompanion:
		nextPhase = phase.PlayingCards
	case phase.PlayingCards:
		predicateToNextPhase = func() bool {
			return team.Count(nps.players, func(p *player.Player) bool { return p.IsHandEmpty() }) == 5
		}
	default:
		nextPhase = phase.End
	}
	if predicateToNextPhase() {
		return nextPhase
	}
	return current
}
