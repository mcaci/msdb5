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

var phases = []string{"Join", "Auction", "Exchange", "Companion", "Card", "End"}

type requester interface {
	Action() string
}

// ToID func
func ToID(rq requester) (ID, error) {
	phase := rq.Action()
	for i := range phases {
		if phases[i] != phase {
			continue
		}
		return ID(i), nil
	}
	return ID(0), fmt.Errorf("Request %s not valid", phase)
}

// MustID func
func MustID(rq requester) ID {
	phase, err := ToID(rq)
	if err != nil {
		panic(err)
	}
	return phase
}

func (id ID) String() string {
	return phases[id]
}
