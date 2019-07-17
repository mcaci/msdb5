package game

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

type checker interface {
	Players() team.Players
}

func check(g checker) bool {
	return team.Count(g.Players(), player.IsHandEmpty) == 5
}

type allInformer interface {
	Caller() *player.Player
	Companion() *player.Player
	Players() team.Players
	Briscola() card.ID
	LastPlaying() *list.List
	Lang() language.Tag
	CardsOnTheBoard() int
	IsSideUsed() bool
	SideDeck() *deck.Cards
}

func predict(g allInformer, roundsBefore, limit uint8) bool {
	if g.CardsOnTheBoard() < 5 || roundsBefore > limit {
		return false
	}
	callers, others := anyoneHasAllBriscola(g, limit)
	if callers == others {
		return false
	}
	collect(g, others)
	return true
}

type collector interface {
	Caller() *player.Player
	Companion() *player.Player
	Players() team.Players
	LastPlaying() *list.List
	Lang() language.Tag
	IsSideUsed() bool
	SideDeck() *deck.Cards
}

func collect(g collector, others bool) {
	p := g.Caller()
	printer := message.NewPrinter(g.Lang())
	team := printer.Sprintf("Callers")
	if others {
		_, p = g.Players().Find(func(p *player.Player) bool { return p == g.Caller() || p == g.Companion() })
		team = printer.Sprintf("Others")
	}
	for _, pl := range g.Players() {
		move(pl.Hand(), p.Pile())
		printer.Fprintf(pl, "The end - %s team has all briscola cards", team)
	}
	if g.IsSideUsed() {
		move(g.SideDeck(), p.Pile())
	}
	track.Player(g.LastPlaying(), p)
}
