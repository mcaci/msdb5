package phase

import (
	"fmt"
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

type requester interface {
	Action() string
}

// ToID func
func ToID(rq requester) (ID, error) {
	phase := rq.Action()
	set := []string{"Join", "Auction", "Exchange", "Companion", "Card"}
	for i := range set {
		if set[i] != phase {
			continue
		}
		return ID(i), nil
	}
	return ID(0), fmt.Errorf("Request %s not valid", phase)
}
