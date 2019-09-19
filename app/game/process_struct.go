package game

import (
	"fmt"
	"strings"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type phaseInfo struct {
	phase  phase.ID
	action string
}

func (s phaseInfo) Action() string  { return s.action }
func (s phaseInfo) Phase() phase.ID { return s.phase }

type expectedSenderInfo struct {
	players team.Players
	origin  string
	p       *player.Player
}

func (s expectedSenderInfo) CurrentPlayer() *player.Player { return s.p }
func (s expectedSenderInfo) From() string                  { return s.origin }
func (s expectedSenderInfo) Players() team.Players         { return s.players }

type Round struct {
	*Game
	req  string
	rErr error
}

func (g Round) Card() (*card.Item, error) {
	fields := strings.Split(g.req, "#")
	if len(fields) > 2 {
		return card.New(fields[1], fields[2])
	}
	return nil, fmt.Errorf("not enough data to make a card: %s", g.req)
}
func (g Round) Value() string     { return value(g.req) }
func (g Round) RoundError() error { return g.rErr }
func (g Round) PlayedCard() card.Item {
	c, err := g.Card()
	if err != nil {
		return card.Item{}
	}
	return *c
}
