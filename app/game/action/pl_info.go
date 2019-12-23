package action

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/team"
)

// PlInfo struct
type PlInfo struct {
	phase        phase.ID
	briscolaCard card.Item
	players      team.Players
	playedCards  *set.Cards
	fromInput    string
}

func NewPlInfo(ph phase.ID, pls team.Players, briscola card.Item, plCards *set.Cards, frInput string) *PlInfo {
	return &PlInfo{phase: ph, players: pls, playedCards: plCards, fromInput: frInput, briscolaCard: briscola}
}

func (nx PlInfo) Briscola() card.Item     { return nx.briscolaCard }
func (nx PlInfo) IsRoundOngoing() bool    { return len(*nx.playedCards) < 5 }
func (nx PlInfo) Phase() phase.ID         { return nx.phase }
func (nx PlInfo) PlayedCards() *set.Cards { return nx.playedCards }
func (nx PlInfo) Players() team.Players   { return nx.players }
func (nx PlInfo) FromInput() string       { return nx.fromInput }
