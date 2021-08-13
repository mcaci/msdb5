package briscola5

import (
	"fmt"
)

// Phase type
type Phase uint8

const (
	InsideAuction Phase = iota
	ExchangingCards
	ChoosingCompanion
	PlayingCards
	End
)

var phases = []string{"Auction", "Exchange", "Companion", "Card", "End"}

// ToPhase func
func ToPhase(phase string) (Phase, error) {
	for i := range phases {
		if phases[i] != phase {
			continue
		}
		return Phase(i), nil
	}
	return Phase(0), fmt.Errorf("request %s not valid", phase)
}

func (id Phase) String() string {
	return phases[id]
}
