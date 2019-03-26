package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type PlayCardStruct struct {
	request, origin string
	players         playerset.Players
	playedCards     *deck.Cards
	sideDeck        *deck.Cards
	briscolaSeed    card.Seed
}

func NewPlay(request, origin string, players playerset.Players,
	playedCards *deck.Cards, sideDeck *deck.Cards, briscolaSeed card.Seed) Action {
	return &PlayCardStruct{request, origin, players,
		playedCards, sideDeck, briscolaSeed}
}

func (pcs PlayCardStruct) Do(p *player.Player) error {
	data := strings.Split(pcs.request, "#")
	number := data[1]
	seed := data[2]
	c, err := card.Create(number, seed)
	err = p.Play(c)
	if err != nil {
		return err
	}
	pcs.playedCards.Add(c)
	roundHasEnded := len(*pcs.playedCards) == 5
	if roundHasEnded {
		index, _ := pcs.players.FindIndex(func(pl *player.Player) bool { return pl == p })
		next := roundWinnerIndex(uint8(index), *pcs.playedCards, pcs.briscolaSeed)
		pcs.players[next].Collect(pcs.playedCards)
		if len(*pcs.sideDeck) > 0 && pcs.NextPhase() == game.End {
			pcs.players[next].Collect(pcs.sideDeck)
			pcs.sideDeck.Clear()
		}
	}
	return err
}
func (pcs PlayCardStruct) NextPlayer(playerInTurn uint8) uint8 {
	next := playersRoundRobin(playerInTurn)
	roundHasEnded := len(*pcs.playedCards) == 5
	if roundHasEnded {
		next = roundWinnerIndex(playerInTurn, *pcs.playedCards, pcs.briscolaSeed)
		pcs.playedCards.Clear()
	}
	return next
}
func (pcs PlayCardStruct) NextPhase() game.Phase {
	var isHandEmpty = func(p *player.Player) bool { return p.IsHandEmpty() }
	if pcs.players.All(isHandEmpty) {
		return game.End
	}
	return game.PlayingCards
}

var roundWinnerIndex = func(playerInTurn uint8, cardsPlayed deck.Cards, seed card.Seed) uint8 {
	winningCardIndex := briscola.IndexOfWinningCard(cardsPlayed, seed)
	return playersRoundRobin(playerInTurn + winningCardIndex)
}
