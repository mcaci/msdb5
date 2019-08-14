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

// ToID func
func ToID(rq interface{ Action() string }) (ID, error) {
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
func MustID(rq interface{ Action() string }) ID {
	phase, err := ToID(rq)
	if err != nil {
		panic(err)
	}
	return phase
}

func (id ID) String() string {
	return phases[id]
}
