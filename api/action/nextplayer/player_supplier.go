package nextplayer

import (
	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type NextPlayerStruct struct {
	current      game.Phase
	players      playerset.Players
	playedCards  *deck.Cards
	briscolaSeed card.Seed
}

func NewPlayerChanger(current game.Phase, players playerset.Players,
	playedCards *deck.Cards, briscolaSeed card.Seed) action.NextPlayerSelector {
	return &NextPlayerStruct{current, players, playedCards, briscolaSeed}
}

var playersRoundRobin = func(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }

func (nps NextPlayerStruct) NextPlayer(playerInTurn uint8) uint8 {
	switch nps.current {
	case game.Joining:
		return playersRoundRobin(playerInTurn)
	case game.InsideAuction:
		return nps.NextPlayerAuction(playerInTurn)
	case game.PlayingCards:
		return nps.NextPhasePlay(playerInTurn)
	default:
		return playerInTurn
	}
}

func (nps NextPlayerStruct) NextPhasePlay(playerInTurn uint8) uint8 {
	next := playersRoundRobin(playerInTurn)
	roundHasEnded := len(*nps.playedCards) == 5
	if roundHasEnded {
		winningCardIndex := briscola.IndexOfWinningCard(*nps.playedCards, nps.briscolaSeed)
		next = playersRoundRobin(playerInTurn + winningCardIndex)
		nps.playedCards.Clear()
	}
	return next
}

func (nps NextPlayerStruct) NextPlayerAuction(playerInTurn uint8) uint8 {
	winnerIndex := playersRoundRobin(playerInTurn)
	for nps.players[winnerIndex].Folded() {
		winnerIndex = playersRoundRobin(winnerIndex)
	}
	return winnerIndex
}
