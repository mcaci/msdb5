package end

import (
	"container/list"
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/card"

	"github.com/nikiforosFreespirit/msdb5/app/track"

	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

type playersInformer interface {
	Caller() *player.Player
	Companion() *player.Player
	Players() team.Players
	Briscola() card.Seed
	LastPlaying() *list.List
}

func Check(g playersInformer, notify func(*player.Player, string)) bool {
	roundsLeft := g.Players()[0].HandSize()
	if roundsLeft <= 3 {
		highbriscolaCard := deck.Highest(g.Briscola())
		var callers, others bool
		var roundsChecked int
		for _, card := range highbriscolaCard {
			if roundsChecked == roundsLeft {
				break
			}
			_, p, err := g.Players().Find(func(p *player.Player) bool {
				return p.Has(card)
			})
			if err != nil { // no one has card
				continue
			}
			if p == g.Caller() || p == g.Companion() {
				callers = true
			} else {
				others = true
			}
			if callers == others {
				break
			}
			roundsChecked++
		}
		if callers != others {
			p := g.Caller()
			if others {
				_, p, _ = g.Players().Find(func(p *player.Player) bool {
					return p == g.Caller() || p == g.Companion()
				})
			}
			collect(g, p, notify)
			return true
		}
	}
	return team.Count(g.Players(), func(p *player.Player) bool { return p.IsHandEmpty() }) == 5
}

func collect(g playersInformer, p *player.Player, notify func(*player.Player, string)) {
	for _, pl := range g.Players() {
		p.Collect(pl.Hand())
		notify(pl, fmt.Sprint("Others team has all briscola, ending game"))
	}
	track.Player(g.LastPlaying(), p)
}

func Process(g playersInformer, notify func(*player.Player, string)) error {
	scorers := make([]team.Scorer, 0)
	for _, p := range g.Players() {
		scorers = append(scorers, p)
	}
	scoreTeam1, scoreTeam2 := team.Score(g.Caller(), g.Companion(), scorers...)
	for _, pl := range g.Players() {
		notify(pl, fmt.Sprintf("The end - Callers: %+v; Others: %+v", scoreTeam1, scoreTeam2))
	}
	return nil
}
