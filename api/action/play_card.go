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
	Finder
}

func NewPlay(request, origin string, playerInTurn *player.Player,
	players playerset.Players, playedCards *deck.Cards, sideDeck *deck.Cards, briscolaSeed card.Seed) Action {
	return &PlayCardStruct{request, origin, players,
		playedCards, sideDeck, briscolaSeed, NewPlayerFinder(origin, playerInTurn)}
}

func (pcs PlayCardStruct) Phase() game.Phase { return game.PlayingCards }
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
		if len(*pcs.sideDeck) > 0 && pcs.NextPhase(pcs.players, pcs) == game.End {
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
func (pcs PlayCardStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) game.Phase {
	if players.All(predicate.NextPhasePlayerInfo) {
		return game.End
	}
	return game.PlayingCards
}
func (pcs PlayCardStruct) NextPhasePlayerInfo(p *player.Player) bool { return p.IsHandEmpty() }

var roundWinnerIndex = func(playerInTurn uint8, cardsPlayed deck.Cards, seed card.Seed) uint8 {
	winningCardIndex := briscola.IndexOfWinningCard(cardsPlayed, seed)
	return playersRoundRobin(playerInTurn + winningCardIndex)
}
