package gamelog

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
)

// ToAll func
func ToAll(gameInfo informer) string {
	all := fmt.Sprintf("Game: %+v", gameInfo)
	sideDeck := gameInfo.SideDeck()
	if gameInfo.Phase() == phase.InsideAuction && gameInfo.IsSideUsed() {
		score := gameInfo.AuctionScore()
		if score >= 90 {
			all += fmt.Sprintf("First card: %+v", sideDeck[0])
		}
		if score >= 100 {
			all += fmt.Sprintf("Second card: %+v", sideDeck[1])
		}
		if score >= 110 {
			all += fmt.Sprintf("Third card: %+v", sideDeck[2])
		}
		if score >= 120 {
			all += fmt.Sprintf("Fourth card: %+v", sideDeck[3])
			all += fmt.Sprintf("Fifth card: %+v", sideDeck[4])
		}
	}
	return all
}
