package player

import (
	"testing"
)

func TestPlayerPresentInList(t *testing.T) {
	player := Player{name: "B"}
	playerByName := ByName("A", []*Player{&Player{name: "A"}, &Player{name: "B"}})
	if player.Name() != playerByName.Name() {
		t.Fatalf("%v and %v are expected to be the same player", player, playerByName)
	}
}
