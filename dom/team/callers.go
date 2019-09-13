package team

import "github.com/mcaci/msdb5/dom/player"

// Callers interface
type Callers interface {
	Caller() *player.Player
	Companion() *player.Player
}

// IsInCallers func
func IsInCallers(g Callers, p *player.Player) bool {
	matching := player.Matching(p)
	return matching(g.Caller()) || matching(g.Companion())
}
