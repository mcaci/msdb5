package game

import (
	"github.com/mcaci/msdb5/app/game/action"
	"github.com/mcaci/msdb5/app/game/end"
	"github.com/mcaci/msdb5/app/game/next"
	"github.com/mcaci/msdb5/app/game/track"
	"github.com/mcaci/msdb5/dom/phase"
)

// Process func
func (g *Game) Process(inputRequest, origin string) Round {
	// verify phase step
	phInfo := phase.NewInfo(g.Phase(), parse(inputRequest, com))
	err := phase.Check(phInfo)
	if err != nil {
		return Round{Game: g, req: inputRequest, rErr: err}
	}

	// verify player step
	es := action.NewExpectedSender(g.Players(), origin, g.CurrentPlayer())
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
		g.Caller(), g.Companion(), len(*g.PlayedCards()) == 5, parse(inputRequest, val))
	g.setPhase(next.Phase(nextPhInfo))

	// end round: next player
	plInfo := next.NewPlInfo(startPhase, g.Players(), g.Briscola(),
		g.PlayedCards(), len(*g.PlayedCards()) < 5, origin)
	nextPl := next.Player(plInfo)
	track.Player(g.LastPlaying(), nextPl)
	if g.Phase() == phase.PlayingCards && len(*g.PlayedCards()) == 5 {
		end.Collect(end.NewCollectInfo(g.CurrentPlayer(), g.PlayedCards()))
	}

	// end game: last round winner collects all cards
	if g.phase == phase.End {
		lastPl := end.LastPlayer(end.NewCollectInfo(g.CurrentPlayer(), g.PlayedCards()), g.Players())
		track.Player(g.LastPlaying(), lastPl)
		end.Collect(end.NewCollectInfo(g.CurrentPlayer(), g.PlayedCards()))
		end.Collect(end.NewCollectInfo(g.CurrentPlayer(), g.SideDeck()))
		for _, p := range g.Players() {
			end.Collect(end.NewCollectInfo(g.CurrentPlayer(), p.Hand()))
		}
	}
	return Round{Game: g, req: inputRequest}
}
