package game

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/game/action"
	"github.com/mcaci/msdb5/app/game/input"
	"github.com/mcaci/msdb5/app/game/track"
	"github.com/mcaci/msdb5/dom/phase"
)

// Process func
func (g *Game) Process(inputRequest, origin string) Round {
	// verify phase
	phInfo := phase.NewInfo(g.Phase(), input.Parse(inputRequest, input.Com))
	err := phase.Check(phInfo)
	if err != nil {
		return Round{Game: g, req: inputRequest, rErr: err}
	}

	// verify player
	err = CheckOrigin(g.Players(), origin, g.CurrentPlayer())
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
	nextPhInfo := action.NewPhInfo(startPhase, g.Players(), g.Briscola(), g.IsSideUsed(),
		g.Caller(), g.Companion(), len(*g.PlayedCards()) == 5, input.Parse(inputRequest, input.Val))
	g.setPhase(action.Phase(nextPhInfo))

	// next player
	plInfo := action.NewPlInfo(startPhase, g.Players(), g.Briscola(), g.PlayedCards(), origin)
	nextPl := action.Player(plInfo)
	track.Player(g.LastPlaying(), nextPl)

	// collect cards
	cardToCollect := action.Collector(g, g.Players(), g.SideDeck(), g.PlayedCards())
	set.Move(cardToCollect(), g.CurrentPlayer().Pile())
	return Round{Game: g, req: inputRequest}
}
