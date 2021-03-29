package briscola5

import "github.com/mcaci/msdb5/v2/dom/player"

type Callerer interface {
	Caller() *player.Player
	Companion() *player.Player
}

// IsInCallers func
func IsInCallers(g Callerer) player.Predicate {
	matchingCaller := player.Matching(g.Caller())
	matchingCompanion := player.Matching(g.Companion())
	return func(p *player.Player) bool { return matchingCaller(p) || matchingCompanion(p) }
}

// Callers holds information about briscola's caller and companion
type Callers struct{ caller, companion *player.Player }

func NewCallersTeam(clr, cmp *player.Player) *Callers { return &Callers{caller: clr, companion: cmp} }
func (c Callers) Caller() *player.Player              { return c.caller }
func (c Callers) Companion() *player.Player           { return c.companion }
