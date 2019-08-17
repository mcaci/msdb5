package frw

import "github.com/mcaci/msdb5/app/game"

// Action interface
type Action interface {
	Join(origin string, playerChannel chan []byte)
	Process(request, origin string) []game.PlMsg
}
