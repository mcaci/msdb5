package game

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/app/input"
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

func (g Round) Card() (*card.Item, error) { return input.Card(g.req) }
func (g Round) PlayedCard() *card.Item    { c, _ := input.Card(g.req); return c }
func (g Round) Value() string             { return input.Value(g.req) }
func (g Round) RoundError() error         { return g.rErr }
