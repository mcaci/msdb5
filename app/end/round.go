package end

import (
	"container/list"
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/deck"

	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/app/track"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

type roundInformer interface {
	Caller() *player.Player
	Companion() *player.Player
	CurrentPlayer() *player.Player
	LastPlayer() *player.Player
	Players() team.Players
	PlayedCards() *deck.Cards
	Phase() phase.ID
	Briscola() card.Seed
	LastPlaying() *list.List
	SenderIndex(string) int

	IsSideUsed() bool
	SideDeck() deck.Cards
	CardsOnTheBoard() int
}

type requestInformer interface {
	From() string
	EndExchange() bool
}

func Round(g roundInformer, rq requestInformer, setCaller func(*player.Player), setPhase func(phase.ID), notify func(*player.Player, string)) {
	// next player step
	nextPlayer(g, rq, notify)

	// next phase
	nextPhase(g, rq, setCaller, setPhase, notify)

	// clean phase
	cleanPhase(g, rq, notify)
}

func nextPlayer(g roundInformer, rq requestInformer, notify func(*player.Player, string)) error {
	current := g.Phase()
	actingPlayerIndex := g.SenderIndex(rq.From())
	var playersRoundRobin = func(playerIndex uint8) uint8 { return (playerIndex + 1) % 5 }
	playerIndex := uint8(actingPlayerIndex)
	nextPlayer := playersRoundRobin(playerIndex)
	switch current {
	case phase.ChoosingCompanion, phase.ExchangingCards:
		nextPlayer = playerIndex
	case phase.InsideAuction:
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

func nextPhase(g roundInformer, rq requestInformer, setCaller func(*player.Player), setPhase func(phase.ID), notify func(*player.Player, string)) error {
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
			_, p, _ := g.Players().Find(func(p *player.Player) bool { return !p.Folded() })
			setCaller(p)
		}
	case phase.ExchangingCards:
		predicateToNextPhase = rq.EndExchange
	case phase.ChoosingCompanion:
		nextPhase = phase.PlayingCards
	case phase.PlayingCards:
		predicateToNextPhase = func() bool {
			return Check(g, notify)
		}
	}
	if predicateToNextPhase() {
		setPhase(nextPhase)
	}
	notify(g.LastPlayer(), gamelog.ToLast(g))
	for _, pl := range g.Players() {
		notify(pl, fmt.Sprintf("Game: %+v", g))
	}
	notify(g.CurrentPlayer(), gamelog.ToCurrent(g))
	return nil
}

func cleanPhase(g roundInformer, rq requestInformer, notify func(*player.Player, string)) error {
	if g.CardsOnTheBoard() >= 5 {
		g.PlayedCards().Clear()
	}
	return nil
}
