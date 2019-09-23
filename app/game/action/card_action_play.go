package action

import (
	"errors"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/team"
)

type playCardData struct {
	playedCards *set.Cards
	players     team.Players
}

func (pc playCardData) act(rq data) {
	cards := rq.pl.Hand()
	index := cards.Find(*rq.card)
	pc.playedCards.Add((*cards)[index])
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}

func (pc playCardData) notAcceptedZeroErr() error {
	return errors.New("Value 0 for card allowed only for ExchangingCard phase")
}

func (pc playCardData) pls() team.Players {
	return pc.players
}
