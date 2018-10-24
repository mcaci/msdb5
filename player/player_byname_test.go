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

func errorCheck(t *testing.T, name string, errorPredicate func(error) bool) {
	if _, err := ByName(name, []*Player{&Player{name: "A"}, &Player{name: "B"}}); errorPredicate(err) {
		t.Fatal(err)
	}
}
func TestPlayerPresentInListNoErr(t *testing.T) {
	errorCheck(t, "A", func(e error) bool { return e != nil })
}

func TestPlayerNotPresentToReturnErr(t *testing.T) {
	errorCheck(t, "C", func(e error) bool { return e == nil })
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
