package msg

import (
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type senderInformer interface {
	Players() team.Players
	From() string
}

func sender(s senderInformer) *player.Player {
	_, p := s.Players().Find(player.MatchingHost(s.From()))
	return p
}
