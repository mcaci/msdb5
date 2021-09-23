package briscola

import (
	"log"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola"
)

func Play(a interface {
	Players() *misc.Players
	Board() *briscola.PlayedCards
}, b interface {
	Name() string
	Idx() uint8
}) error {
	i, err := a.Players().Index(func(p misc.Player) bool { return p.Name() == b.Name() })
	if err != nil {
		return err
	}
	hand := (*a.Players())[i].Hand()
	card := (*hand)[b.Idx()]
	log.Println("playing card: ", card)
	err = set.MoveOne(&card, hand, a.Board().Cards)
	if err != nil {
		return err
	}
	return nil
}
