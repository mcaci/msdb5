package action

import (
	"errors"

	"github.com/mcaci/msdb5/dom/team"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
)

type companionData struct {
	setCompanion  func(*card.Item)
	setCompPlayer func(*player.Player)
	players       team.Players
}

func (c companionData) act(rq data) {
	c.setCompanion(rq.card)
	c.setCompPlayer(rq.pl)
}

func (c companionData) notAcceptedZeroErr() error {
	return errors.New("Value 0 for card allowed only for ExchangingCard phase")
}

func (c companionData) pls() team.Players {
	return c.players
}
