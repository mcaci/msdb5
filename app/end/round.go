package end

import (
	"container/list"

	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/nikiforosFreespirit/msdb5/app/msg"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/app/track"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
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
	Briscola() card.Seed
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

func nextPlayer(g roundInformer, rq requestInformer) error {
	current := g.Phase()
	actingPlayerIndex := senderIndex(g, rq)
	var playersRoundRobin = func(playerIndex uint8) uint8 { return (playerIndex + 1) % 5 }
	playerIndex := uint8(actingPlayerIndex)
	nextPlayer := playersRoundRobin(playerIndex)
	switch current {
	case phase.ChoosingCompanion, phase.ExchangingCards:
		nextPlayer = playerIndex
	case phase.InsideAuction:
		if *g.AuctionScore() >= 120 {
			for i := range g.Players() {
				if i == actingPlayerIndex {
					continue
				}
				g.Players()[i].Fold()
			}
		}
		for player.Folded(g.Players()[nextPlayer]) {
			nextPlayer = playersRoundRobin(nextPlayer)
		}
	case phase.PlayingCards:
		roundHasEnded := len(*g.PlayedCards()) == 5
		if roundHasEnded {
			winningCardIndex := briscola.IndexOfWinningCard(*g.PlayedCards(), g.Briscola())
			nextPlayer = playersRoundRobin(playerIndex + winningCardIndex)
		}
	default:
	}
	track.Player(g.LastPlaying(), g.Players()[nextPlayer])
	return nil
}

func nextPhase(g roundInformer, rq requestInformer, setCaller func(*player.Player), setPhase func(phase.ID)) {
	current := g.Phase()
	nextPhase := current + 1
	switch {
	case current == phase.InsideAuction && !g.IsSideUsed():
		nextPhase = current + 2
	default:
		nextPhase = current + 1
	}
	isNext := true
	switch current {
	case phase.Joining:
		isNext = team.Count(g.Players(), player.IsNameEmpty) == 0
	case phase.InsideAuction:
		isNext = team.Count(g.Players(), player.Folded) == 4
	case phase.ExchangingCards:
		isNext = rq.Value() == "0"
	case phase.PlayingCards:
		isNext = check(g)
	}
	if isNext && current == phase.InsideAuction {
		_, p := g.Players().Find(func(p *player.Player) bool { return !player.Folded(p) })
		setCaller(p)
		setPhase(nextPhase)
	}
	if isNext && current != phase.InsideAuction {
		setPhase(nextPhase)
	}
	printer := message.NewPrinter(g.Lang())
	printer.Fprintf(g.LastPlayer(), msg.CreateInGameMsg(g, g.LastPlayer()))
	for _, pl := range g.Players() {
		printer.Fprintf(pl, "Game: %+v", g)
	}
	printer.Fprintf(g.CurrentPlayer(), msg.CreateInGameMsg(g, g.CurrentPlayer()))
}

func cleanPhase(g roundInformer, rq requestInformer) {
	if g.CardsOnTheBoard() < 5 {
		return
	}
	g.PlayedCards().Clear()
}
