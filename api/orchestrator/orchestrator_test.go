package orchestrator

import (
	"testing"
)

func TestActionCreationAndAuctionUsage(t *testing.T) {
	b := NewAction()
	b.Action("Join#A", "100.1.1.1")
	_, pInfo, _ := b.Action("Auction#102", "100.1.1.1")
	if pInfo == nil {
		t.Fatal("Auction action was not properly performed")
	}
}

func TestActionCreationAndPCompanionUsage(t *testing.T) {
	b := NewAction()
	b.Action("Join#A", "100.1.1.1")
	_, pInfo, _ := b.Action("Companion#3#Cup", "100.1.1.1")
	if pInfo == nil {
		t.Fatal("Companion action was not properly performed")
	}
}

func TestActionCreationAndPlayCardUsage(t *testing.T) {
	b := NewAction()
	b.Action("Join#A", "100.1.1.1")
	_, pInfo, _ := b.Action("Card#6#Cudgel", "100.1.1.1")
	if pInfo == nil {
		t.Fatal("Card action was not properly performed")
	}
}
