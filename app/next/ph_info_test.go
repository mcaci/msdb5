package next

import (
	"fmt"
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

func TestNextPhase(t *testing.T) {
	testObj := NewPhInfo(phase.ExchangingCards, testPlayers(), *card.MustID(1), true, player.New(), player.New(), true, "0")
	next := Phase(testObj)
	expected := phase.ChoosingCompanion
	if next != expected {
		t.Fatalf("Expecting %s, found %s", expected, next)
	}
}

func testPredictionPlayersToEnd() team.Players {
	pls := make(team.Players, 5)
	for i := range pls {
		pls[i] = player.New()
		pls[i].Join(fmt.Sprintf("127.0.0.%d", i))
		pls[i].Hand().Add(*card.MustID(uint8(i) + 9), *card.MustID(uint8(i) + 20))
	}
	return pls
}

func TestNextPhaseToEnd(t *testing.T) {
	pls := testPredictionPlayersToEnd()
	testObj := NewPhInfo(phase.PlayingCards, pls, *card.MustID(1), true, pls[0], pls[1], true, "0")
	next := Phase(testObj)
	expected := phase.End
	if next != expected {
		t.Fatalf("Expecting %s, found %s", expected, next)
	}
}

func testPredictionPlayersToStayStill() team.Players {
	pls := make(team.Players, 5)
	for i := range pls {
		pls[i] = player.New()
		pls[i].Join(fmt.Sprintf("127.0.0.%d", i))
		pls[i].Hand().Add(*card.MustID(uint8(i) + 1), *card.MustID(uint8(i) + 10))
	}
	return pls
}

func TestNextPhaseToStayStill(t *testing.T) {
	pls := testPredictionPlayersToStayStill()
	testObj := NewPhInfo(phase.PlayingCards, pls, *card.MustID(1), true, pls[1], pls[0], true, "0")
	next := Phase(testObj)
	expected := phase.PlayingCards
	if next != expected {
		t.Fatalf("Expecting %s, found %s", expected, next)
	}
}

func testPredictionPlayersToStay() team.Players {
	pls := make(team.Players, 5)
	for i := range pls {
		pls[i] = player.New()
		pls[i].Join(fmt.Sprintf("127.0.0.%d", i))
		pls[i].Hand().Add(*card.MustID(uint8(i) + 10), *card.MustID(uint8(i) + 20))
	}
	return pls
}

func TestNextPhaseToStayInPlay(t *testing.T) {
	pls := testPredictionPlayersToStay()
	testObj := NewPhInfo(phase.PlayingCards, pls, *card.MustID(1), true, pls[0], pls[1], true, "0")
	next := Phase(testObj)
	expected := phase.PlayingCards
	if next != expected {
		t.Fatalf("Expecting %s, found %s", expected, next)
	}
}
