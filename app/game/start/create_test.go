package start

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type testStr struct {
	t team.Players
	s *set.Cards
}

func (t testStr) Players() team.Players { return t.t }
func (t testStr) SideDeck() *set.Cards  { return t.s }

func TestDistributeAllWithEmptySide(t *testing.T) {
	testObj := testStr{t: team.Players{player.New(), player.New(), player.New(), player.New(), player.New()}, s: &set.Cards{}}
	DistributeCards(testObj, false)
	if len(*testObj.SideDeck()) != 0 {
		t.Fatal("Expecting empty side deck")
	}
}

func TestDistributeAllWithFiveCardsAside(t *testing.T) {
	testObj := testStr{t: team.Players{player.New(), player.New(), player.New(), player.New(), player.New()}, s: &set.Cards{}}
	DistributeCards(testObj, true)
	if len(*testObj.SideDeck()) != 5 {
		t.Log(testObj)
		t.Fatal("Expecting side deck with five cards")
	}
}

func TestDistributeWithSide(t *testing.T) {
	testObj := testStr{t: team.Players{player.New(), player.New(), player.New(), player.New(), player.New()}, s: &set.Cards{}}
	DistributeCards(testObj, true)
	if len(*testObj.t[0].Hand()) != 7 {
		t.Fatal("Expecting 7 cards in hand")
	}
}

func TestDistributeAllCards(t *testing.T) {
	testObj := testStr{t: team.Players{player.New(), player.New(), player.New(), player.New(), player.New()}, s: &set.Cards{}}
	DistributeCards(testObj, false)
	if len(*testObj.t[0].Hand()) != 8 {
		t.Fatal("Expecting 8 cards in hand")
	}
}
