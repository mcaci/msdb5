package game

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func nextPhase(g *Game, request string) phase.ID {
	current, nextPhase := g.phase, g.phase+1
	predicateToNextPhase := func() bool { return true }
	switch current {
	case phase.Joining:
		predicateToNextPhase = func() bool {
			return team.Count(g.players, func(p *player.Player) bool { return p.IsNameEmpty() }) == 0
		}
	case phase.InsideAuction:
		predicateToNextPhase = func() bool {
			return team.Count(g.players, func(p *player.Player) bool { return p.Folded() }) == 4
		}
		if !g.IsSideUsed() {
			nextPhase = current + 2
		}
	case phase.ExchangingCards:
		predicateToNextPhase = func() bool {
			data := strings.Split(request, "#")
			return len(data) > 1 && data[1] == "0"
		}
	case phase.ChoosingCompanion:
		nextPhase = phase.PlayingCards
	case phase.PlayingCards:
		predicateToNextPhase = func() bool {
			return team.Count(g.players, func(p *player.Player) bool { return p.IsHandEmpty() }) == 5
		}
	}
	if predicateToNextPhase() {
		return nextPhase
	}
	return current
}