package briscola5

import "github.com/mcaci/msdb5/v2/dom/player"

// IsInCallers func
func IsInCallers(g interface {
	Caller() *player.Player
	Companion() *player.Player
}) player.Predicate {
	matchingCaller := player.Matching(g.Caller())
	matchingCompanion := player.Matching(g.Companion())
	return func(p *player.Player) bool { return matchingCaller(p) || matchingCompanion(p) }
}
