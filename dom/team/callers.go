package team

import "github.com/mcaci/msdb5/dom/player"

// CallerTeam struct
type CallerTeam struct {
	call *player.Player
}

func NewEmptyCallers() Callers               { return CallerTeam{player.New()} }
func NewCallers(call *player.Player) Callers { return CallerTeam{call} }
func (t CallerTeam) Caller() *player.Player  { return t.call }
func (CallerTeam) Companion() *player.Player { return player.New() }

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
