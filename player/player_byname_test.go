package player

import (
	"testing"
)

func TestPlayerPresentInList(t *testing.T) {
	player := Player{name: "A"}
	playerByName := ByName("A", []*Player{&Player{name: "A"}, &Player{name: "B"}})
	if player.Name() != playerByName.Name() {
		t.Fatalf("%v and %v are expected to be the same player", player, playerByName)
	}
}

func TestPlayerPresentInListByHost(t *testing.T) {
	player := Player{host: "A"}
	playerByName := ByName("A", []*Player{&Player{host: "A"}, &Player{host: "B"}})
	if player.Name() != playerByName.Name() {
		t.Fatalf("%v and %v are expected to be the same player", player, playerByName)
	}
}
