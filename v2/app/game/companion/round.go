package companion

import (
	"log"

	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Round(c briscola.Card, players team.Players) struct {
	Briscola  briscola.Card
	Companion uint8
} {
	idx, err := players.Index(player.IsCardInHand(c.Item))
	if err != nil {
		log.Printf("error: %v. Card %v chosen from the side deck", err, c)
		log.Printf("not expecting at this point to have issues. Fallback to Player1 for now.")
	}

	return struct {
		Briscola  briscola.Card
		Companion uint8
	}{
		Briscola:  c,
		Companion: idx,
	}
}
