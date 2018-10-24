package player

import (
	"testing"
)

var testPlayers = []*Player{&Player{name: "A"}, &Player{name: "B"}}

func TestPlayerPresentInList(t *testing.T) {
	name := "A"
	if player, _ := ByName(name, testPlayers); name != player.Name() {
		t.Fatalf("%v and %v are expected to be the same player", name, player)
	}
}

func errorCheck(t *testing.T, name string, errorPredicate func(error) bool) {
	if _, err := ByName(name, testPlayers); errorPredicate(err) {
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
// 	// if player.Name() != player.Name() {
// 	// 	t.Fatalf("%v and %v are expected to be the same player", player, player)
// 	// }
// }
