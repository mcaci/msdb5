package gamelog

import (
	"log"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type senderInformer interface {
	Sender(string) *player.Player
}

// ToConsole func
func ToConsole(gameInfo senderInformer, rq requester) {
	sender := gameInfo.Sender(rq.From())
	log.Printf("New Action by %s: %s\n", sender.Name(), rq.Action())
	log.Printf("Sender info: %+v\n", sender)
	log.Printf("Game info: %+v\n", gameInfo)
}

// ErrToConsole func
func ErrToConsole(senderName, request string, err error) {
	log.Printf("New Action by %s: %s\n", senderName, request)
	log.Printf("Error raised: %+v\n", err)
}
