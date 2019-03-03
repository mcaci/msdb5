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
	isInfoPresent := func(p *player.Player) bool { return p.Has(33) }
	if _, err := testPlayers.Find(isInfoPresent); err != nil {
		t.Fatal("Player not found with criteria p.Has(33)")
	}
}

func TestSuccessfulFindDataCorresponds(t *testing.T) {
	isInfoPresent := func(p *player.Player) bool { return p.Name() == "A" }
	if player, _ := testPlayers.Find(isInfoPresent); "A" != player.Name() {
		t.Fatalf("%s and %v are expected to be the same player", "A", player)
	}
}

func TestUnsuccessfulFind(t *testing.T) {
	isInfoPresent := func(p *player.Player) bool { return p.Has(24) }
	if _, err := testPlayers.Find(isInfoPresent); err == nil {
		t.Fatal(err)
	}
}
