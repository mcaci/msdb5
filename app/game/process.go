package game

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/game/action"
	"github.com/mcaci/msdb5/app/game/end"
	"github.com/mcaci/msdb5/app/game/next"
	"github.com/mcaci/msdb5/app/game/track"
	"github.com/mcaci/msdb5/dom/phase"
)

// Process func
func (g *Game) Process(inputRequest, origin string) Round {
	// verify phase
	phInfo := phase.NewInfo(g.Phase(), parse(inputRequest, com))
	err := phase.Check(phInfo)
	if err != nil {
		return Round{Game: g, req: inputRequest, rErr: err}
	}

	// verify player
	es := action.NewExpectedSender(g.Players(), origin, g.CurrentPlayer())
	err = action.CheckOrigin(es)
	if err != nil {
		return Round{Game: g, req: inputRequest, rErr: err}
	}

	// play
	gInfo := Round{Game: g, req: inputRequest}
	err = action.Play(gInfo)
	if err != nil {
		return Round{Game: g, req: inputRequest, rErr: err}
	}

	// next phase
	startPhase := g.Phase()
	nextPhInfo := next.NewPhInfo(startPhase, g.Players(), g.Briscola(), g.IsSideUsed(),
		g.Caller(), g.Companion(), len(*g.PlayedCards()) == 5, parse(inputRequest, val))
	g.setPhase(next.Phase(nextPhInfo))

	// next player
	plInfo := next.NewPlInfo(startPhase, g.Players(), g.Briscola(), g.PlayedCards(), origin)
	nextPl := next.Player(plInfo)
	track.Player(g.LastPlaying(), nextPl)

	// collect cards
	cardToCollect := end.Collector(g.Phase(), g.Players(), g.SideDeck(), g.PlayedCards())
	set.Move(cardToCollect(), g.CurrentPlayer().Pile())
	return Round{Game: g, req: inputRequest}
}
