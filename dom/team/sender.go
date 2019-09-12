package team

import "github.com/mcaci/msdb5/dom/player"

type senderInformer interface {
	Players() Players
	From() string
}

func Sender(s senderInformer) *player.Player {
	_, p := s.Players().Find(player.MatchingHost(s.From()))
	return p
}
