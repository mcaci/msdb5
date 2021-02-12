package phase

import "fmt"

// ID type
type ID uint8

const (
	InsideAuction ID = iota
	ExchangingCards
	ChoosingCompanion
	PlayingCards
	End
)

var phases = []string{"Auction", "Exchange", "Companion", "Card", "End"}

// ToID func
func ToID(phase string) (ID, error) {
	for i := range phases {
		if phases[i] != phase {
			continue
		}
		return ID(i), nil
	}
	return ID(0), fmt.Errorf("Request %s not valid", phase)
}

// MustID func
func MustID(phase string) ID {
	id, err := ToID(phase)
	if err != nil {
		panic(err)
	}
	return id
}

func (id ID) String() string {
	return phases[id]
}
