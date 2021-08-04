package team

import (
	"testing"

	"github.com/mcaci/msdb5/v2/dom/player"
)

func testPredicateA(p player.Player) bool { return p.Name() == "A" }

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
