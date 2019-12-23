package phase

import (
	"testing"
)

func TestIDCreationMustWithNoErr(t *testing.T) {
	testID := MustID("Card")
	if testID != PlayingCards {
		t.Fatalf("Unexpected phase: %s", testID)
	}
}

func TestIDCreationWithNoErr(t *testing.T) {
	_, err := ToID("Card")
	if err != nil {
		t.Fatal("Unexpected error")
	}
}

func TestIDValueCreation_Join(t *testing.T) {
	p, _ := ToID("Join")
	if p != Joining {
		t.Fatal("Unexpected error")
	}
}

func TestIDValueCreation_Companion(t *testing.T) {
	p, _ := ToID("Companion")
	if p != ChoosingCompanion {
		t.Fatal("Unexpected error")
	}
}

func TestIDValueCreation_Auction(t *testing.T) {
	p, _ := ToID("Auction")
	if p != InsideAuction {
		t.Fatal("Unexpected error")
	}
}

func TestIDCreationWithErr(t *testing.T) {
	_, err := ToID("Budget")
	if err == nil {
		t.Fatal("Budget is not a valid phase")
	}
}
