package end

import (
	"container/list"

	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"golang.org/x/text/language"

	"github.com/nikiforosFreespirit/msdb5/app/notify"
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
	SenderIndex(string) int

	IsSideUsed() bool
	SideDeck() *deck.Cards
	CardsOnTheBoard() int
}

type requestInformer interface {
	From() string
	Value() string
}

// Round func
func Round(g roundInformer, rq requestInformer, setCaller func(*player.Player), setPhase func(phase.ID), sendMsg func(*player.Player, string)) {
	// next player step
	nextPlayer(g, rq, sendMsg)

	// next phase
	nextPhase(g, rq, setCaller, setPhase, sendMsg)

	// clean phase
	cleanPhase(g, rq, sendMsg)
}

func nextPlayer(g roundInformer, rq requestInformer, sendMsg func(*player.Player, string)) error {
	current := g.Phase()
	actingPlayerIndex := g.SenderIndex(rq.From())
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
		for g.Players()[nextPlayer].Folded() {
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

func nextPhase(g roundInformer, rq requestInformer, setCaller func(*player.Player), setPhase func(phase.ID), sendMsg func(*player.Player, string)) error {
	current, nextPhase := g.Phase(), g.Phase()+1
	predicateToNextPhase := func() bool { return true }
	switch current {
	case phase.Joining:
		predicateToNextPhase = func() bool {
			return team.Count(g.Players(), func(p *player.Player) bool { return p.IsNameEmpty() }) == 0
		}
	case phase.InsideAuction:
		predicateToNextPhase = func() bool {
			return team.Count(g.Players(), func(p *player.Player) bool { return p.Folded() }) == 4
		}
		if !g.IsSideUsed() {
			nextPhase = current + 2
		}
		if predicateToNextPhase() {
			_, p := g.Players().Find(func(p *player.Player) bool { return !p.Folded() })
			setCaller(p)
		}
	case phase.ExchangingCards:
		predicateToNextPhase = func() bool { return rq.Value() == "0" }
	case phase.ChoosingCompanion:
		nextPhase = phase.PlayingCards
	case phase.PlayingCards:
		predicateToNextPhase = func() bool {
			return Check(g, sendMsg)
		}
	}
	if predicateToNextPhase() {
		setPhase(nextPhase)
	}
	sendMsg(g.LastPlayer(), notify.ToLast(g))
	for _, pl := range g.Players() {
		sendMsg(pl, notify.GameInfoMsg(g))
	}
	sendMsg(g.CurrentPlayer(), notify.ToCurrent(g))
	return nil
}

func cleanPhase(g roundInformer, rq requestInformer, sendMsg func(*player.Player, string)) error {
	if g.CardsOnTheBoard() >= 5 {
		g.PlayedCards().Clear()
	}
	return nil
}
