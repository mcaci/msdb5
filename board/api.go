package board

import (
	"fmt"
	"strings"
)

// API interface
type API interface {
	Action(request, origin string)
	fmt.Stringer
}

// Action interface
func (b *Board) Action(request, origin string) {
	data := strings.Split(string(request), "#")
	switch data[0] {
	case "Join":
		b.Join(data[1], origin)
	case "Auction":
		b.RaiseAuction(data[1], origin)
	case "Play":
		b.Nominate(data[1], data[2], origin)
	}
}
