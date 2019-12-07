package team

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
)

type fakeGame struct {
	current *player.Player
	players Players
	origin  string
}

func newTestGame(origin string) fakeGame {
	f := fakeGame{}
	t := Players{}
	p := player.New()
	p.Join("127.0.0.51")
	p.RegisterAs("A")
	t.Add(p)
	f.players = t
	f.current = p
	f.origin = origin
	return f
}

func TestVerifyPlayerWithNoErr(t *testing.T) {
	g := newTestGame("127.0.0.51")
	err := CheckOrigin(g.players, g.origin, g.current)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPlayerWithErr(t *testing.T) {
	g := newTestGame("127.0.0.52")
	err := CheckOrigin(g.players, g.origin, g.current)
	if err == nil {
		t.Fatal("Error was expected")
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
