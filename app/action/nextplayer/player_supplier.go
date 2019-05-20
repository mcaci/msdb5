package nextplayer

import (
	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

type NextPlayerStruct struct {
	current      phase.ID
	players      team.Players
	playedCards  *deck.Cards
	briscolaSeed card.Seed
}

func NewPlayerChanger(current phase.ID, players team.Players,
	playedCards *deck.Cards, briscolaSeed card.Seed) action.NextPlayerSelector {
	return &NextPlayerStruct{current, players, playedCards, briscolaSeed}
}

var playersRoundRobin = func(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }

func (nps NextPlayerStruct) NextPlayer(playerInTurn uint8) uint8 {
	switch nps.current {
	case phase.Joining:
		return playersRoundRobin(playerInTurn)
	case phase.InsideAuction:
		return nps.NextPlayerAuction(playerInTurn)
	case phase.PlayingCards:
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
