package briscola5

import (
	"testing"
)

func TestIDCreationWithNoErr(t *testing.T) {
	_, err := ToPhase("Card")
	if err != nil {
		t.Fatal("Unexpected error")
	}
}

func TestIDValueCreation_Exchange(t *testing.T) {
	p, _ := ToPhase("Exchange")
	if p != ExchangingCards {
		t.Fatal("Unexpected error")
	}
}

func TestIDValueCreation_Companion(t *testing.T) {
	p, _ := ToPhase("Companion")
	if p != ChoosingCompanion {
		t.Fatal("Unexpected error")
	}
}

func TestIDValueCreation_Auction(t *testing.T) {
	p, _ := ToPhase("Auction")
	if p != InsideAuction {
		t.Fatal("Unexpected error")
	}
}

func TestIDCreationWithErr(t *testing.T) {
	_, err := ToPhase("Budget")
	if err == nil {
		t.Fatal("Budget is not a valid phase")
	}
}
