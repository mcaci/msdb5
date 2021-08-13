package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/briscola/player"
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

func TestCount(t *testing.T) {
	p := player.New(&player.Options{For5P: true}).(*player.B5Player)
	if count := Count(Players{p, p}, func(pl player.Player) bool { return true }); count != 2 {
		t.Fatal("Count should be 2")
	}
}

func TestUnsuccessfulFind(t *testing.T) {
	if _, err := testPlayers.Index(player.IsCardInHand(*card.MustID(8))); err == nil {
		t.Fatal("Player should not be found")
	}
}

func TestSuccessfulFindNoErr(t *testing.T) {
	if _, err := testPlayers.Index(player.IsCardInHand(*card.MustID(33))); err != nil {
		t.Fatal("Player not found with criteria player.IsCardInHand(33)")
	}
}

func TestSuccessfulFindWithNone(t *testing.T) {
	if testPlayers.None(player.IsCardInHand(*card.MustID(33))) {
		t.Fatal("Player not found with criteria player.IsCardInHand(33)")
	}
}

func TestUnsuccessfulFindWithNone(t *testing.T) {
	if !testPlayers.None(player.IsCardInHand(*card.MustID(24))) {
		t.Fatal("Player should not be found")
	}
}

func TestUnsuccessfulFindWithAll(t *testing.T) {
	if testPlayers.All(player.IsCardInHand(*card.MustID(24))) {
		t.Fatal("Player should not be found")
	}
}

func TestSuccessfulFindWithAll(t *testing.T) {
	if !testPlayers.All(player.IsCardInHand(*card.MustID(34))) {
		t.Fatal("All players have the 4 of Cudgel")
	}
}
