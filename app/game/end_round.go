package game

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func nextPlayer(g *Game, rq *req, notify func(*player.Player, string)) error {
	current := g.phase
	actingPlayerIndex := g.senderIndex(rq.From())
	var playersRoundRobin = func(playerIndex uint8) uint8 { return (playerIndex + 1) % 5 }
	playerIndex := uint8(actingPlayerIndex)
	nextPlayer := playersRoundRobin(playerIndex)
	switch current {
	case phase.ChoosingCompanion, phase.ExchangingCards:
		nextPlayer = playerIndex
	case phase.InsideAuction:
		for g.players[nextPlayer].Folded() {
			nextPlayer = playersRoundRobin(nextPlayer)
		}
	case phase.PlayingCards:
		roundHasEnded := len(g.playedCards) == 5
		if roundHasEnded {
			winningCardIndex := briscola.IndexOfWinningCard(g.playedCards, g.briscola())
			nextPlayer = playersRoundRobin(playerIndex + winningCardIndex)
		}
	default:
	}
	trackActing(&g.lastPlaying, g.players[nextPlayer])
	return nil
}

func nextPhase(g *Game, rq *req, notify func(*player.Player, string)) error {
	current, nextPhase := g.phase, g.phase+1
	predicateToNextPhase := func() bool { return true }
	switch current {
	case phase.Joining:
		predicateToNextPhase = func() bool {
			return team.Count(g.players, func(p *player.Player) bool { return p.IsNameEmpty() }) == 0
		}
	case phase.InsideAuction:
		predicateToNextPhase = func() bool {
			return team.Count(g.players, func(p *player.Player) bool { return p.Folded() }) == 4
		}
		if !g.IsSideUsed() {
			nextPhase = current + 2
		}
		if predicateToNextPhase() {
			_, p, _ := g.players.Find(func(p *player.Player) bool { return !p.Folded() })
			g.caller = p
		}
	case phase.ExchangingCards:
		predicateToNextPhase = rq.EndExchange
	case phase.ChoosingCompanion:
		nextPhase = phase.PlayingCards
	case phase.PlayingCards:
		predicateToNextPhase = func() bool {
			return team.Count(g.players, func(p *player.Player) bool { return p.IsHandEmpty() }) == 5
		}
	}
	if predicateToNextPhase() {
		g.phase = nextPhase
	}
	notify(g.LastPlayer(), gamelog.ToLast(g))
	for _, pl := range g.players {
		notify(pl, fmt.Sprintf("Game: %+v", g))
	}
	notify(g.CurrentPlayer(), gamelog.ToCurrent(g))
	return nil
}

func cleanPhase(g *Game, rq *req, notify func(*player.Player, string)) error {
	if g.cardsOnTheBoard() >= 5 {
		g.playedCards.Clear()
	}
	return nil
}
