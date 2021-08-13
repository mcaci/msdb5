package misc

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func testPredicateA(p Player) bool { return p.Name() == "A" }

func TestPartitionT1(t *testing.T) {
	t1, _ := testPlayers.Part(testPredicateA)
	if t1.None(testPredicateA) {
		t.Fatal("t1 should contain only players named A")
	}
}

func TestPartitionT2(t *testing.T) {
	_, t2 := testPlayers.Part(testPredicateA)
	if !t2.None(testPredicateA) {
		t.Fatal("t2 should not contain players named A")
	}
}

func TestUnsuccessfulFind(t *testing.T) {
	if _, err := testPlayers.Index(IsCardInHand(*card.MustID(8))); err == nil {
		t.Fatal("Player should not be found")
	}
}

func TestSuccessfulFindNoErr(t *testing.T) {
	if _, err := testPlayers.Index(IsCardInHand(*card.MustID(33))); err != nil {
		t.Fatal("Player not found with criteria misc.IsCardInHand(33)")
	}
}

func TestSuccessfulFindWithNone(t *testing.T) {
	if testPlayers.None(IsCardInHand(*card.MustID(33))) {
		t.Fatal("Player not found with criteria misc.IsCardInHand(33)")
	}
}

func TestUnsuccessfulFindWithNone(t *testing.T) {
	if !testPlayers.None(IsCardInHand(*card.MustID(24))) {
		t.Fatal("Player should not be found")
	}
}

func TestUnsuccessfulFindWithAll(t *testing.T) {
	if testPlayers.All(IsCardInHand(*card.MustID(24))) {
		t.Fatal("Player should not be found")
	}
}

func TestSuccessfulFindWithAll(t *testing.T) {
	if !testPlayers.All(IsCardInHand(*card.MustID(34))) {
		t.Fatal("All players have the 4 of Cudgel")
	}
}
