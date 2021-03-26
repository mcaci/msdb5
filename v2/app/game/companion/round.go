package companion

import (
	"log"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Round(c *card.Item, players team.Players) struct {
	Briscola  *card.Item
	Companion *player.Player
} {
	idx, err := players.Index(player.IsCardInHand(*c))
	if err != nil {
		log.Printf("error: %v. Card %v chosen from the side deck", err, *c)
		log.Printf("not expecting at this point to have issues. Fallback to Player1 for now.")
	}

	return struct {
		Briscola  *card.Item
		Companion *player.Player
	}{
		Briscola:  c,
		Companion: players[idx],
	}
}
