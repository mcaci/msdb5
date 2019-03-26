package action

import (
	"errors"
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/deck"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type ExchangeCardsStruct struct {
	request, origin string
	sideDeck        *deck.Cards
}

func NewExchangeCards(request, origin string, sideDeck *deck.Cards) Action {
	return &ExchangeCardsStruct{request, origin, sideDeck}
}

func (ecs ExchangeCardsStruct) Do(p *player.Player) error {
	data := strings.Split(ecs.request, "#")
	number := data[1]
	if number == "0" {
		return nil
	}
	seed := data[2]
	c, err := card.Create(number, seed)
	if err != nil {
		return err
	}
	if !p.Has(c) {
		return errors.New("card is not in players hand")
	}
	index, err := p.Hand().Find(c)
	if err != nil {
		return err
	}
	p.Hand().Add((*ecs.sideDeck)[0])
	ecs.sideDeck.Remove(0)
	ecs.sideDeck.Add(c)
	p.Hand().Remove(index)
	return nil
}
func (ecs ExchangeCardsStruct) NextPlayer(playerInTurn uint8) uint8 { return playerInTurn }
func (ecs ExchangeCardsStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) game.Phase {
	data := strings.Split(ecs.request, "#")
	number, err := strconv.Atoi(data[1])
	if number == 0 || err != nil {
		return game.ChosingCompanion
	}
	return game.ExchangingCards
}
func (ecs ExchangeCardsStruct) NextPhasePlayerInfo(p *player.Player) bool { return true }
