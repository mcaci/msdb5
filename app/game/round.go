package game

import (
	"container/list"

	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/deck"
	"golang.org/x/text/language"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
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
	CardsOnTheBoard() int
}

type requestInformer interface {
	From() string
	Action() string
	Value() string
}

// Round func
func Round(g roundInformer, rq requestInformer, setCaller func(*player.Player), setPhase func(phase.ID)) {
	// actions to do post request
	postRequest(g, rq)
	// next player
	nextPlayer(g, rq)
	// next phase
	nextPhase(g, rq, setCaller, setPhase)
	// clean phase
	cleanPhase(g, rq)
}

func senderIndex(g roundInformer, rq requestInformer) int {
	index, _ := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(rq.From()) })
	return index
}

func cleanPhase(g roundInformer, rq requestInformer) {
	if g.CardsOnTheBoard() < 5 {
		return
	}
	g.PlayedCards().Clear()
}
