package phasesupplier

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/api/game"
)

func evaluate(t *testing.T, input InputAction, expected game.Phase) {
	if input.Phase() != expected {
		t.Fatalf("expecting %v phase to be returned", expected)
	}
}

func TestSupplyJoiningPhase(t *testing.T)   { evaluate(t, "Join", game.Joining) }
func TestSupplyAuctionPhase(t *testing.T)   { evaluate(t, "Auction", game.InsideAuction) }
func TestSupplyExchangePhase(t *testing.T)  { evaluate(t, "Exchange", game.ExchangingCards) }
func TestSupplyCompanionPhase(t *testing.T) { evaluate(t, "Companion", game.ChosingCompanion) }
func TestSupplyPlayPhase(t *testing.T)      { evaluate(t, "Card", game.PlayingCards) }
func TestSupplyEndPhase(t *testing.T)       { evaluate(t, "End", game.End) }
