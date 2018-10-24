package player

import (
	"testing"
)

func TestPlayerPresentInList(t *testing.T) {
	player := Player{name: "A"}
	playerByName, _ := ByName("A", []*Player{&Player{name: "A"}, &Player{name: "B"}})
	if player.Name() != playerByName.Name() {
		t.Fatalf("%v and %v are expected to be the same player", player, playerByName)
	}
}

func TestPlayerPresentInListNoErr(t *testing.T) {
	if _, err := ByName("A", []*Player{&Player{name: "A"}, &Player{name: "B"}}); err != nil {
		t.Fatal(err)
	}
}

func TestPlayerNotPresentToReturnErr(t *testing.T) {
	if _, err := ByName("C", []*Player{&Player{name: "A"}, &Player{name: "B"}}); err == nil {
		t.Fatal(err)
	}
}

// func TestPlayerPresentInListByHost(t *testing.T) {
// 	// player := Player{host: "A"}
// 	_, err := ByName("A", []*Player{&Player{host: "A"}, &Player{host: "B"}})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// if player.Name() != playerByName.Name() {
// 	// 	t.Fatalf("%v and %v are expected to be the same player", player, playerByName)
// 	// }
// }
