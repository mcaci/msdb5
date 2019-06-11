package game

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func nextPlayer(g *Game, request, origin string, notify func(*player.Player, string)) {
	current := g.phase
	actingPlayerIndex := g.senderIndex(origin)
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
		if nextPlayer == playerIndex {
			g.caller = g.players[playerIndex]
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
}

func nextPhase(g *Game, request string) phase.ID {
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
	case phase.ExchangingCards:
		predicateToNextPhase = func() bool {
			data := strings.Split(request, "#")
			return len(data) > 1 && data[1] == "0"
		}
	case phase.ChoosingCompanion:
		nextPhase = phase.PlayingCards
	case phase.PlayingCards:
		predicateToNextPhase = func() bool {
			return team.Count(g.players, func(p *player.Player) bool { return p.IsHandEmpty() }) == 5
		}
	}
	if predicateToNextPhase() {
		g.phase = nextPhase
		return nextPhase
	}
	return current
}

func cleanPhase(g *Game, request, origin string, notify func(*player.Player, string)) error {
	if g.cardsOnTheBoard() >= 5 {
		g.playedCards.Clear()
	}
	return nil
}
