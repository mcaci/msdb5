package player

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/display"
)

// Name func
func (player *Player) Name() display.Info {
	return display.NewInfo("Name", ":", player.name, ";")
}

// Info function
func (player Player) Info() []display.Info {
	hand := display.NewInfo("Hand", ":", player.hand.String(), ";")
	return display.Wrap("Player", player.Name(), hand)
}

func (player Player) String() string {
	host := display.NewInfo("Host", ":", player.host, ";")
	hand := display.NewInfo("Hand", ":", player.hand.String(), ";")
	pile := display.NewInfo("Pile", ":", player.pile.String(), ";")
	fold := display.NewInfo("Folded", ":", strconv.FormatBool(player.Folded()), ";")
	return display.All(display.Wrap("Player", player.Name(), host, hand, pile, fold)...)
}
