package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type PlayWithSideCardStruct struct {
	request, origin string
	playerInTurn    *player.Player
	players         playerset.Players
	playedCards     *deck.Cards
	sideDeck        *deck.Cards
	briscolaSeed    card.Seed
}

func NewPlayWithSide(request, origin string, playerInTurn *player.Player,
	players playerset.Players, playedCards *deck.Cards, sideDeck *deck.Cards, briscolaSeed card.Seed) Action {
	return &PlayWithSideCardStruct{request, origin, playerInTurn,
		players, playedCards, sideDeck, briscolaSeed}
}

func (pcs PlayWithSideCardStruct) Phase() game.Phase { return game.PlayingCards }
func (pcs PlayWithSideCardStruct) Find(p *player.Player) bool {
	return p.IsExpectedPlayer(pcs.playerInTurn, pcs.origin)
}
func (pcs PlayWithSideCardStruct) Do(p *player.Player) error {
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
		if pcs.NextPhase(pcs.players, pcs) == game.End {
			pcs.players[next].Collect(pcs.sideDeck)
			pcs.sideDeck.Clear()
		}
	}
	return err
}
func (pcs PlayWithSideCardStruct) NextPlayer(playerInTurn uint8) uint8 {
	next := playersRoundRobin(playerInTurn)
	roundHasEnded := len(*pcs.playedCards) == 5
	if roundHasEnded {
		next = roundWinnerIndex(playerInTurn, *pcs.playedCards, pcs.briscolaSeed)
		pcs.playedCards.Clear()
	}
	return next
}
func (pcs PlayWithSideCardStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) game.Phase {
	if players.All(predicate.NextPhasePlayerInfo) {
		return game.End
	}
	return game.PlayingCards
}
func (pcs PlayWithSideCardStruct) NextPhasePlayerInfo(p *player.Player) bool { return p.IsHandEmpty() }
