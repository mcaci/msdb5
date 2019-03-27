package nextphase

import (
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type NextPhaseStruct struct {
	current  game.Phase
	players  playerset.Players
	sideDeck bool
	request  string
}

func NewChanger(current game.Phase, players playerset.Players, sideDeck bool,
	request string) action.NextPhaseChanger {
	return &NextPhaseStruct{current, players, sideDeck, request}
}

func (nps NextPhaseStruct) NextPhase() game.Phase {
	switch nps.current {
	case game.Joining:
		return nps.NextPhaseJoin()
	case game.InsideAuction:
		if nps.sideDeck {
			return nps.NextPhaseAuctionWithSide()
		}
		return nps.NextPhaseAuctionNoSide()
	case game.ExchangingCards:
		return nps.NextPhaseExchange()
	case game.ChosingCompanion:
		return game.PlayingCards
	case game.PlayingCards:
		return nps.NextPhasePlay()
	default:
		return game.End
	}
}

func (nps NextPhaseStruct) NextPhaseJoin() game.Phase {
	var isPlayerEmpty = func(p *player.Player) bool { return p.IsNameEmpty() }
	if nps.players.Count(isPlayerEmpty) == 0 {
		return game.InsideAuction
	}
	return game.Joining
}

func (nps NextPhaseStruct) NextPhasePlay() game.Phase {
	var isHandEmpty = func(p *player.Player) bool { return p.IsHandEmpty() }
	if nps.players.All(isHandEmpty) {
		return game.End
	}
	return game.PlayingCards
}

func (nps NextPhaseStruct) nextPhase(phase game.Phase) game.Phase {
	var isFolded = func(p *player.Player) bool { return p.Folded() }
	if nps.players.Count(isFolded) == 4 {
		return phase
	}
	return game.InsideAuction
}

func (nps NextPhaseStruct) NextPhaseAuctionNoSide() game.Phase {
	return nps.nextPhase(game.ChosingCompanion)
}

func (nps NextPhaseStruct) NextPhaseAuctionWithSide() game.Phase {
	return nps.nextPhase(game.ExchangingCards)
}

func (nps NextPhaseStruct) NextPhaseExchange() game.Phase {
	data := strings.Split(nps.request, "#")
	if len(data) > 1 {
		number, err := strconv.Atoi(data[1])
		if number == 0 || err != nil {
			return game.ChosingCompanion
		}
	}
	return game.ExchangingCards
}
