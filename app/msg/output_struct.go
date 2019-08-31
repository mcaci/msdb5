package msg

import (
	"github.com/mcaci/msdb5/dom/team"
)

type senderInfo struct {
	players team.Players
	origin  string
}

func (s senderInfo) From() string          { return s.origin }
func (s senderInfo) Players() team.Players { return s.players }
