package briscola5

import (
	"github.com/mcaci/msdb5/v2/app/misc"
)

// IsInCallers func
func IsInCallers(t interface {
	Caller() misc.Player
	Companion() misc.Player
}) func(p misc.Player) bool {
	return func(p misc.Player) bool { return eq(p, t.Caller()) || eq(p, t.Companion()) }
}

func eq(p, q misc.Player) bool {
	return p.Name() == q.Name() && p.Hand() == q.Hand() && p.Pile() == q.Pile()
}
