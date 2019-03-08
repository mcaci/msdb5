package playerset

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/player"
)

var testPlayers Players

func init() {
	var a player.Player
	a.Join("A", "")
	testPlayers.Add(a)
	var b player.Player
	b.Join("B", "")
	b.Hand().Add(33)
	testPlayers.Add(b)
}

func TestSuccessfulFindNoErr(t *testing.T) {
	if _, err := testPlayers.Find(func(p *player.Player) bool { return p.Has(33) }); err != nil {
		t.Fatal("Player not found with criteria p.Has(33)")
	}
}

func TestSuccessfulFindDataCorresponds(t *testing.T) {
	if player, _ := testPlayers.Find(func(p *player.Player) bool { return p.IsName("A") }); !player.IsName("A") {
		t.Fatalf("%s and %v are expected to be the same player", "A", player)
	}
}

func TestUnsuccessfulFind(t *testing.T) {
	if _, err := testPlayers.Find(func(p *player.Player) bool { return p.Has(24) }); err == nil {
		t.Fatal(err)
	}
}

func TestCount(t *testing.T) {
	if 2 != testPlayers.Count(func(p *player.Player) bool { return true }) {
		t.Fatal("Count should be 2")
	}
}
