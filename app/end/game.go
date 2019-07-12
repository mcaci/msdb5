package end

import (
	"container/list"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type playersInformer interface {
	Caller() *player.Player
	Companion() *player.Player
	Players() team.Players
	Briscola() card.Seed
	LastPlaying() *list.List
	Lang() language.Tag
	CardsOnTheBoard() int
	IsSideUsed() bool
	SideDeck() *deck.Cards
}

func check(g playersInformer) bool {
	roundsLeft := len(*g.Players()[0].Hand())
	if g.CardsOnTheBoard() >= 5 && roundsLeft <= 3 {
		highbriscolaCard := deck.BriscolaSerie(g.Briscola())
		var callers, others bool
		var roundsChecked int
		for _, card := range highbriscolaCard {
			if roundsChecked == roundsLeft {
				break
			}
			_, p := g.Players().Find(func(p *player.Player) bool { return p.Has(card) })
			if p == nil { // no one has card
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
			printer := message.NewPrinter(g.Lang())
			team := printer.Sprintf("Callers")
			if others {
				_, p = g.Players().Find(func(p *player.Player) bool { return p == g.Caller() || p == g.Companion() })
				team = printer.Sprintf("Others")
			}
			collect(g, p, team)

			return true
		}
	}
	return team.Count(g.Players(), player.IsHandEmpty) == 5
}

func collect(g playersInformer, p *player.Player, team string) {
	printer := message.NewPrinter(g.Lang())
	for _, pl := range g.Players() {
		p.Collect(pl.Hand())
		printer.Fprintf(pl, "The end - %s team has all briscola cards", team)
	}
	if g.IsSideUsed() {
		p.Collect(g.SideDeck())
	}
	track.Player(g.LastPlaying(), p)
}
