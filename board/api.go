package board

import (
	"strings"
)

// Action interface
func (b *Board) Action(request, origin string) {
	data := strings.Split(string(request), "#")
	switch data[0] {
	case "Join":
		b.Join(data[1], origin)
	case "Auction":
		b.RaiseAuction(data[1], origin)
	case "Companion":
		b.Nominate(data[1], data[2], origin)
	case "Card":
		b.Play(data[1], data[2], origin)
	}
}
