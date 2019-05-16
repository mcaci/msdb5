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
	var predicateToNextPhase func() bool
	switch nps.current {
	case game.Joining:
		predicateToNextPhase = func() bool {
			var isPlayerEmpty = func(p *player.Player) bool { return p.IsNameEmpty() }
			return nps.players.Count(isPlayerEmpty) == 0
		}
	case game.InsideAuction:
		predicateToNextPhase = func() bool {
			var isFolded = func(p *player.Player) bool { return p.Folded() }
			return nps.players.Count(isFolded) == 4
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
		predicateToNextPhase = func() bool { return true }
	case game.PlayingCards:
		predicateToNextPhase = func() bool {
			var isHandEmpty = func(p *player.Player) bool { return p.IsHandEmpty() }
			return nps.players.All(isHandEmpty)
		}
	default:
		return game.End
	}
	return nextPhaseNext(current, nextPhase, predicateToNextPhase)
}

func nextPhaseNext(self, next game.Phase, canStepToNext func() bool) game.Phase {
	if canStepToNext() {
		return next
	}
	return self
}

func (nps NextPhaseStruct) NextPhasePlay() game.Phase {
	var isHandEmpty = func(p *player.Player) bool { return p.IsHandEmpty() }
	if nps.players.All(isHandEmpty) {
		return game.End
	}
	return game.PlayingCards
}
