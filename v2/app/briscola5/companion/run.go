package companion

import (
	"log"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola"
)

type companionIn struct {
	Player  misc.Player
	Players misc.Players
}
type companionOut struct {
	Briscola  briscola.Card
	Companion uint8
}

func Run(s companionIn) companionOut {
	selectedCard := selectCardFromSeedThatHasMostCardsInHand(s.Player)
	serie := briscola.Serie(selectedCard)
	for _, c := range serie {
		i, err := s.Players.Index(misc.IsCardInHand(c))
		switch {
		case err != nil:
			log.Printf("error: %v. Card %v is inside the side deck", err, c)
		case s.Players[i] == s.Player:
			log.Print("Player is self")
		default:
			return companionOut{
				Briscola:  briscola.Card{Item: c},
				Companion: i,
			}
		}
	}
	log.Printf("not expecting at this point to have issues. Fallback to player1 for now.")
	return companionOut{}
}

func selectCardFromSeedThatHasMostCardsInHand(p misc.Player) card.Item {
	// count how many cards per seed
	count := make(map[card.Seed]uint8)
	for _, c := range *p.Hand() {
		count[c.Seed()]++
	}
	// check seed with highest card
	var max uint8
	var seed card.Seed
	for k, v := range count {
		if v < max {
			continue
		}
		max = v
		seed = k
	}
	// select card from seed
	for _, c := range *p.Hand() {
		if c.Seed() != seed {
			continue
		}
		return c
	}
	return card.Item{}
}
