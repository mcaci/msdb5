package phase

import (
	"testing"
)

type fakeRequester string

func (fr fakeRequester) Action() string {
	return string(fr)
}

func TestIDCreationWithNoErr(t *testing.T) {
	_, err := ToID(fakeRequester("Card"))
	if err != nil {
		t.Fatal("Unexpected error")
	}
}

func TestIDValueCreation_Join(t *testing.T) {
	p, _ := ToID(fakeRequester("Join"))
	if p != Joining {
		t.Fatal("Unexpected error")
	}
}

func TestIDValueCreation_Exchange(t *testing.T) {
	p, _ := ToID(fakeRequester("Exchange"))
	if p != ExchangingCards {
		t.Fatal("Unexpected error")
	}
}

func TestIDValueCreation_Companion(t *testing.T) {
	p, _ := ToID(fakeRequester("Companion"))
	if p != ChoosingCompanion {
		t.Fatal("Unexpected error")
	}
}

func TestIDValueCreation_Auction(t *testing.T) {
	p, _ := ToID(fakeRequester("Auction"))
	if p != InsideAuction {
		t.Fatal("Unexpected error")
	}
}

func TestIDCreationWithErr(t *testing.T) {
	_, err := ToID(fakeRequester("Budget"))
	if err == nil {
		t.Fatal("Budget is not a valid phase")
	}
}
