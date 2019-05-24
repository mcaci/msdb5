package phase

import (
	"fmt"
	"strings"
)

// ID type
type ID uint8

const (
	Joining ID = iota
	InsideAuction
	ExchangingCards
	ChoosingCompanion
	PlayingCards
	End
)

// ToID func
func ToID(request string) (ID, error) {
	phase := strings.Split(request, "#")[0]
	var id ID
	var err error
	switch phase {
	case "Join":
		id = Joining
	case "Auction":
		id = InsideAuction
	case "Exchange":
		id = ExchangingCards
	case "Companion":
		id = ChoosingCompanion
	case "Card":
		id = PlayingCards
	default:
		err = fmt.Errorf("Request %s not valid", phase)
	}
	return id, err
}
