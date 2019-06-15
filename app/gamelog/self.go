package gamelog

import (
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type selfInformer interface {
	CurrentPlayer() *player.Player
	LastPlayer() *player.Player
	Phase() phase.ID
	SideDeck() deck.Cards
}

// ToCurrent func
func ToCurrent(gameInfo selfInformer) string {
	return createInGameMsg(gameInfo, gameInfo.CurrentPlayer())
}

// ToLast func
func ToLast(gameInfo selfInformer) string {
	return createInGameMsg(gameInfo, gameInfo.LastPlayer())
}
