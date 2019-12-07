package team

import "github.com/mcaci/msdb5/dom/player"

// Callers interface
type Callers interface {
	Caller() *player.Player
	Companion() *player.Player
}

// IsInCallers func
func IsInCallers(g Callers) player.Predicate {
	matchingCaller := player.Matching(g.Caller())
	matchingCompanion := player.Matching(g.Companion())
	return func(p *player.Player) bool { return matchingCaller(p) || matchingCompanion(p) }
}
