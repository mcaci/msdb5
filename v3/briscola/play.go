package briscola

import (
	"log"

	"github.com/mcaci/ita-cards/set"
)

func Play(g interface {
	Players() *Players
	Board() *PlayedCards
}, b interface {
	Name() string
	Idx() uint8
}) error {
	i, err := g.Players().Index(func(p Player) bool { return p.Name() == b.Name() })
	if err != nil {
		return err
	}
	hand := (*g.Players())[i].Hand()
	log.Print(g.Players())
	card := (*hand)[b.Idx()]
	log.Println("playing card: ", card)
	err = set.MoveOne(&card, hand, g.Board().Cards)
	if err != nil {
		return err
	}
	return nil
}
