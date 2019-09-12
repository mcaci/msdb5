package game

import (
	"github.com/mcaci/msdb5/app/action"
	"github.com/mcaci/msdb5/app/action/collect"
	"github.com/mcaci/msdb5/app/action/end"
	"github.com/mcaci/msdb5/app/input"
	"github.com/mcaci/msdb5/app/next"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/track"
)

// Process func
func (g *Game) Process(inputRequest, origin string) Round {
	// verify phase step
	phInfo := phaseInfo{g.Phase(), input.Command(inputRequest)}
	err := phase.Check(phInfo)
	if err != nil {
		return Round{Game: g, req: inputRequest, rErr: err}
	}

	// verify player step
	es := expectedSenderInfo{g.Players(), origin, g.CurrentPlayer()}
	err = action.CheckOrigin(es)
	if err != nil {
		return Round{Game: g, req: inputRequest, rErr: err}
	}

	// play step
	gInfo := Round{Game: g, req: inputRequest}
	err = action.Play(gInfo)
	if err != nil {
		return Round{Game: g, req: inputRequest, rErr: err}
	}

	// end round: next phase
	startPhase := g.Phase()
	nextPhInfo := next.NewPhInfo(startPhase, g.Players(), g.Briscola(), g.IsSideUsed(),
		g.Caller(), g.Companion(), len(*g.PlayedCards()) == 5, input.Value(inputRequest))
	g.setPhase(next.Phase(nextPhInfo))

	// end round: next player
	plInfo := next.NewPlInfo(startPhase, g.Players(), g.Briscola(),
		g.PlayedCards(), len(*g.PlayedCards()) < 5, origin)
	nextPl := next.Player(plInfo)
	track.Player(g.LastPlaying(), nextPl)
	if g.Phase() == phase.PlayingCards {
		collect.Played(collect.NewInfo(g.CurrentPlayer(), g.PlayedCards()))
	}

	// end game: last round winner collects all cards
	if g.phase == phase.End {
		lastPl := end.LastPlayer(g)
		track.Player(g.LastPlaying(), lastPl)
		collect.Played(collect.NewInfo(g.CurrentPlayer(), g.PlayedCards()))
		collect.All(collect.NewAllInfo(g.CurrentPlayer(), g.SideDeck(), g.Players()))
	}
	return Round{Game: g, req: inputRequest}
}
