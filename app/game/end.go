package game

import (
	"container/list"

	"golang.org/x/text/language"

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
