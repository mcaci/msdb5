package next

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

// PhInfo struct
type PhInfo struct {
	phase              phase.ID
	players            team.Players
	caller, companion  *player.Player
	briscolaCard       card.Item
	side, roundOngoing bool
	exchInput          string
}

func NewPhInfo(ph phase.ID, pls team.Players, call, comp *player.Player, briscola card.Item,
	isSide, isRound bool, exchVal string) *PhInfo {
	return &PhInfo{phase: ph, players: pls, caller: call, companion: comp, exchInput: exchVal,
		briscolaCard: briscola, side: isSide, roundOngoing: isRound}
}

func (ph PhInfo) Briscola() card.Item       { return ph.briscolaCard }
func (ph PhInfo) Caller() *player.Player    { return ph.caller }
func (ph PhInfo) Companion() *player.Player { return ph.companion }
func (ph PhInfo) IsSideUsed() bool          { return ph.side }
func (ph PhInfo) IsRoundOngoing() bool      { return ph.roundOngoing }
func (ph PhInfo) Phase() phase.ID           { return ph.phase }
func (ph PhInfo) Players() team.Players     { return ph.players }
func (ph PhInfo) ExchangeInput() string     { return ph.exchInput }
