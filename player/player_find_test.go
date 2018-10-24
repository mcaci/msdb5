package player

import (
	"testing"
)

var testPlayers = []*Player{&Player{name: "A"}, &Player{name: "B"}}
var testPlayersWithHost = []*Player{&Player{host: "A"}, &Player{host: "B"}}

func TestPlayerPresentInListByName(t *testing.T) {
	name := "A"
	if player, _ := Find(name, testPlayers); name != player.Name() {
		t.Fatalf("%v and %v are expected to be the same player", name, player)
	}
}

func TestPlayerPresentInListByHost(t *testing.T) {
	host := "A"
	if player, _ := Find(host, testPlayersWithHost); host != player.Host() {
		t.Fatalf("%v and %v are expected to be the same player", host, player)
	}
}

func errorCheck(t *testing.T, nameOrHost string, players []*Player, errorPredicate func(error) bool) {
	if _, err := Find(nameOrHost, players); errorPredicate(err) {
		t.Fatal(err)
	}
}
func TestPlayerPresentInListNoErr(t *testing.T) {
	name := "A"
	errorCheck(t, name, testPlayers, func(e error) bool { return e != nil })
}

func TestPlayerNotPresentToReturnErr(t *testing.T) {
	name := "C"
	errorCheck(t, name, testPlayers, func(e error) bool { return e == nil })
}

func TestPlayerPresentInListByHostNoErr(t *testing.T) {
	host := "A"
	errorCheck(t, host, testPlayersWithHost, func(e error) bool { return e != nil })
	// if player.Name() != player.Name() {
	// 	t.Fatalf("%v and %v are expected to be the same player", player, player)
	// }
}

func TestPlayerPresentInListNotPresentByHost(t *testing.T) {
	host := "C"
	errorCheck(t, host, testPlayersWithHost, func(e error) bool { return e == nil })
	// if player.Name() != player.Name() {
	// 	t.Fatalf("%v and %v are expected to be the same player", player, player)
	// }
}
