package player

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

func (player Player) String() string {
	str := "Player["
	str += "Name:" + player.name + ";"
	str += "Host:" + player.host + ";"
	str += "Hand:"
	for _, cardID := range player.hand {
		c, _ := card.ByID(cardID)
		str += c.String() + " "
	}
	str += "Pile:"
	for _, cardID := range player.pile {
		c, _ := card.ByID(cardID)
		str += c.String() + " "
	}
	str += "]"
	return str
}
