package start

import (
	"testing"

	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type testStr team.Players

func (t testStr) Players() team.Players { return team.Players(t) }

func TestDistributeAllWithEmptySide(t *testing.T) {
	testObj := testStr{player.New(), player.New(), player.New(), player.New(), player.New()}
	sideP := DistributeAll(testObj, false)
	if len(*sideP) != 0 {
		t.Fatal("Expecting empty side deck")
	}
}

func TestDistributeAllWithFiveCardsAside(t *testing.T) {
	testObj := testStr{player.New(), player.New(), player.New(), player.New(), player.New()}
	sideP := DistributeAll(testObj, true)
	if len(*sideP) != 5 {
		t.Fatal("Expecting empty side deck")
	}
}

func TestDistributeWithSide(t *testing.T) {
	testObj := testStr{player.New(), player.New(), player.New(), player.New(), player.New()}
	DistributeAll(testObj, true)
	if len(*testObj[0].Hand()) != 7 {
		t.Fatal("Expecting 7 cards in hand")
	}
}

func TestDistributeAllCards(t *testing.T) {
	testObj := testStr{player.New(), player.New(), player.New(), player.New(), player.New()}
	DistributeAll(testObj, false)
	if len(*testObj[0].Hand()) != 8 {
		t.Fatal("Expecting 8 cards in hand")
	}
}
