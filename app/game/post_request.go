package game

import (
	"container/list"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func postRequest(g roundInformer, rq requestInformer) {
	current := g.Phase()
	switch current {
	case phase.PlayingCards:
		roundHasEnded := len(*g.PlayedCards()) == 5
		if !roundHasEnded {
			break
		}
		winnerIndex := roundWinnerIndex(g)
		collectingPlayer := g.Players()[winnerIndex]
		collectingPlayer.Collect(g.PlayedCards())
		if !(team.Count(g.Players(), player.IsHandEmpty) == 5 && g.IsSideUsed()) {
			break
		}
		side := g.SideDeck()
		collectingPlayer.Collect(side)
		side.Clear()
	}
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
	collectingPlayer := g.Caller()
	printer := message.NewPrinter(g.Lang())
	team := printer.Sprintf("Callers")
	if others {
		_, collectingPlayer = g.Players().Find(func(p *player.Player) bool { return p == g.Caller() || p == g.Companion() })
		team = printer.Sprintf("Others")
	}
	for _, pl := range g.Players() {
		collectingPlayer.Collect(pl.Hand())
		printer.Fprintf(pl, "The end - %s team has all briscola cards", team)
	}
	track.Player(g.LastPlaying(), collectingPlayer)
	if !g.IsSideUsed() {
		return
	}
	collectingPlayer.Collect(g.SideDeck())
	g.SideDeck().Clear()
}

func roundWinnerIndex(g roundInformer) uint8 {
	playerIndex, _ := g.Players().Find(func(pl *player.Player) bool { return pl == g.CurrentPlayer() })
	winningCardIndex := briscola.IndexOfWinningCard(*g.PlayedCards(), g.Briscola())
	winnerIndex := (playerIndex + int(winningCardIndex) + 1) % 5
	return uint8(winnerIndex)
}
