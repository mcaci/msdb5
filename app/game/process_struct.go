package game

import (
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
	c    *card.Item
	cErr error
	val  string
	rErr error
}

func (g Round) Card() (*card.Item, error) { return g.c, g.cErr }
func (g Round) Value() string             { return g.val }
func (g Round) RoundError() error         { return g.rErr }
