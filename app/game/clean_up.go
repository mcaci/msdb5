package game

import (
	"container/list"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/language"
)

type roundInformer interface {
	AuctionScore() *auction.Score
	Caller() *player.Player
	Companion() *player.Player
	CurrentPlayer() *player.Player
	LastPlayer() *player.Player
	Players() team.Players
	PlayedCards() *deck.Cards
	Phase() phase.ID
	Briscola() card.ID
	Lang() language.Tag
	LastPlaying() *list.List
	IsSideUsed() bool
	SideDeck() *deck.Cards
	IsRoundOngoing() bool
}

type requestInformer interface {
	From() string
	Action() string
	Value() string
}

func cleanUp(g roundInformer, rq requestInformer) {
	current := g.Phase()
	switch current {
	case phase.PlayingCards:
		if g.IsRoundOngoing() {
			break
		}
		playerIndex, _ := g.Players().Find(func(pl *player.Player) bool { return pl == g.CurrentPlayer() })
		winningCardIndex := briscola.IndexOfWinningCard(*g.PlayedCards(), g.Briscola())
		winnerIndex := (playerIndex + int(winningCardIndex) + 1) % 5
		p := g.Players()[winnerIndex]
		move(g.PlayedCards(), p.Pile())
		if !(team.Count(g.Players(), player.IsHandEmpty) == 5 && g.IsSideUsed()) {
			break
		}
		move(g.SideDeck(), p.Pile())
	}
}

func move(from, to *deck.Cards) {
	to.Add(*from...)
	from.Clear()
}
