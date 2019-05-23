package gamelog

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
)

// ToMe func
func ToMe(gameInfo informer) string {
	me := fmt.Sprintf("%+v", gameInfo.CurrentPlayer())
	if gameInfo.Phase() == phase.ExchangingCards {
		me += fmt.Sprintf("Side deck: %+v", gameInfo.SideDeck())
	}
	return me
}
