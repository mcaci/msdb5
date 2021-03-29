package briscola5

import (
	"fmt"
	"log"
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
	return Phase(0), fmt.Errorf("Request %s not valid", phase)
}

// MustPhase func
func MustPhase(phase string) Phase {
	id, err := ToPhase(phase)
	if err != nil {
		log.Fatalf("error found: %v", err)
	}
	return id
}

func (id Phase) String() string {
	return phases[id]
}
