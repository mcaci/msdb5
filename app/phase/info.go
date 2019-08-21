package phase

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

// Info struct
type Info struct {
	players            team.Players
	caller, companion  *player.Player
	briscolaCard       card.Item
	side, roundOngoing bool
	exchInput          string
}

func NewInfo(pls team.Players, call, comp *player.Player, briscola card.Item,
	isSide, isRound bool, exchVal string) *Info {
	return &Info{players: pls, caller: call, companion: comp, exchInput: exchVal,
		briscolaCard: briscola, side: isSide, roundOngoing: isRound}
}

func (ph Info) Briscola() card.Item       { return ph.briscolaCard }
func (ph Info) Caller() *player.Player    { return ph.caller }
func (ph Info) Companion() *player.Player { return ph.companion }
func (ph Info) IsSideUsed() bool          { return ph.side }
func (ph Info) Players() team.Players     { return ph.players }
func (ph Info) IsRoundOngoing() bool      { return ph.roundOngoing }
func (ph Info) ExchangeInput() string     { return ph.exchInput }
