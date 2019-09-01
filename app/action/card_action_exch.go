package action

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/team"
)

type exchangeData struct {
	side    *set.Cards
	players team.Players
}

func (e exchangeData) act(rq data) {
	cards := rq.pl.Hand()
	index := cards.Find(*rq.card)
	toCards := e.side
	awayCard := (*cards)[index]
	(*cards)[index] = (*toCards)[0]
	*toCards = append((*toCards)[1:], awayCard)
}

func (e exchangeData) notAcceptedZeroErr() error {
	return nil
}

func (e exchangeData) pls() team.Players {
	return e.players
}
