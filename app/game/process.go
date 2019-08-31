package game

import (
	"github.com/mcaci/msdb5/app/action"
	"github.com/mcaci/msdb5/app/action/collect"
	"github.com/mcaci/msdb5/app/input"
	"github.com/mcaci/msdb5/app/next"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/team"
)

// Process func
func (g *Game) Process(inputRequest, origin string) Round {
	// verify phase step
	phInfo := phaseInfo{g.Phase(), input.Command(inputRequest)}
	err := phase.Check(phInfo)
	if err != nil {
		return Round{Game: g, rErr: err}
	}

	// verify player step
	es := expectedSenderInfo{g.Players(), origin, g.CurrentPlayer()}
	err = team.CheckOrigin(es)
	if err != nil {
		return Round{Game: g, rErr: err}
	}

	// play step
	c, cerr := input.Card(inputRequest)
	gInfo := Round{Game: g, c: c, cErr: cerr, val: input.Value(inputRequest)}
	err = action.Play(gInfo)
	if err != nil {
		return Round{Game: g, rErr: err}
	}

	// end round: next player
	plInfo := next.NewPlInfo(g.Phase(), g.Players(), g.PlayedCards(), g.Briscola(),
		len(*g.SideDeck()) > 0, len(*g.PlayedCards()) < 5, origin)
	nextPl := next.Player(plInfo)
	track.Player(g.LastPlaying(), nextPl)
	if g.Phase() == phase.PlayingCards {
		collect.Played(collect.NewInfo(g.CurrentPlayer(), g.PlayedCards()))
	}

	// end round: next phase
	nextPhInfo := next.NewPhInfo(g.Phase(), g.Players(), g.Caller(), g.Companion(), g.Briscola(),
		len(*g.SideDeck()) > 0, len(*g.PlayedCards()) == 0, input.Value(inputRequest))
	g.setPhase(next.Phase(nextPhInfo))

	if g.phase != phase.End {
		return Round{Game: g, c: c, cErr: cerr, val: input.Value(inputRequest)}
	}

	// process end game
	// last round winner collects all cards
	collect.All(collect.NewAllInfo(g.CurrentPlayer(), g.SideDeck(), g.Players()))

	// compute score (output data)
	// pilers := make([]score.Piler, len(g.Players()))
	// for i, p := range g.Players() {
	// 	pilers[i] = p
	// }
	// scoreTeam1, scoreTeam2 := score.Calc(g.Caller(), g.Companion(), pilers, briscola.Points)

	return Round{Game: g, c: c, cErr: cerr, val: input.Value(inputRequest)}
}
