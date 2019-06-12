package gamelog

import (
	"fmt"

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

// requester
type requester interface {
	From() string
	Action() string
}

// ToCurrent func
func ToCurrent(gameInfo selfInformer) string {
	return to(gameInfo, gameInfo.CurrentPlayer())
}

// ToLast func
func ToLast(gameInfo selfInformer) string {
	return to(gameInfo, gameInfo.LastPlayer())
}

func to(gameInfo selfInformer, pl *player.Player) string {
	me := fmt.Sprintf("Player: %+v\n", pl)
	if gameInfo.Phase() == phase.ExchangingCards {
		me += fmt.Sprintf("Side deck: %+v\n", gameInfo.SideDeck())
	}
	return me
}

// SendErrToSender func
func SendErrToSender(err error, gameInfo senderInformer, rq requester, notify func(*player.Player, string)) {
	sender := gameInfo.Sender(rq.From())
	ErrToConsole(sender.Name(), rq.Action(), err)
	notify(sender, err.Error())
}
