package gamelog

import (
	"log"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// ToConsole func
func ToConsole(gameInfo informer, sender *player.Player, request string, err error) {
	log.Printf("New Action by %s\n", sender.Name())
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %+v\n", err)
	log.Printf("Player info after action: %+v\n", sender)
	log.Printf("Game info after action: %+v\n", gameInfo)
}
