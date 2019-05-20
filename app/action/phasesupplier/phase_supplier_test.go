package phasesupplier

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
)

func evaluate(t *testing.T, input InputAction, expected phase.ID) {
	if input.ID() != expected {
		t.Fatalf("expecting %v phase to be returned", expected)
	}
}

func TestSupplyJoiningPhase(t *testing.T)   { evaluate(t, "Join", phase.Joining) }
func TestSupplyAuctionPhase(t *testing.T)   { evaluate(t, "Auction", phase.InsideAuction) }
func TestSupplyExchangePhase(t *testing.T)  { evaluate(t, "Exchange", phase.ExchangingCards) }
func TestSupplyCompanionPhase(t *testing.T) { evaluate(t, "Companion", phase.ChosingCompanion) }
func TestSupplyPlayPhase(t *testing.T)      { evaluate(t, "Card", phase.PlayingCards) }
func TestSupplyEndPhase(t *testing.T)       { evaluate(t, "End", phase.End) }
