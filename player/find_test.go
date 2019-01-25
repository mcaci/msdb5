package player

import (
	"testing"
)

var (
	testPlayers         Players
	testPlayersWithHost Players
)

func init() {
	// &player.Player{name: "A"}, &player.Player{name: "B"}
	var a Player
	a.SetName("A")
	testPlayers.Add(a)
	var b Player
	b.SetName("B")
	testPlayers.Add(b)
	//  Players{&player.Player{host: "A"}, &player.Player{host: "B"}}
	var a1 Player
	a1.MyHostIs("A")
	testPlayersWithHost.Add(a1)
	var b1 Player
	b1.MyHostIs("B")
	testPlayersWithHost.Add(b1)
}

func TestPlayerPresentInListByName(t *testing.T) {
	name := "A"
	if player, _ := testPlayers.Find(name); name != player.Name() {
		t.Fatalf("%v and %v are expected to be the same player", name, player)
	}
}

func TestPlayerPresentInListByHost(t *testing.T) {
	host := "A"
	if player, _ := testPlayersWithHost.Find(host); host != player.Host() {
		t.Fatalf("%v and %v are expected to be the same player", host, player)
	}
}

func errorCheck(t *testing.T, nameOrHost string, players Players, errorPredicate func(error) bool) {
	if _, err := players.Find(nameOrHost); errorPredicate(err) {
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
}

func TestPlayerPresentInListNotPresentByHost(t *testing.T) {
	host := "C"
	errorCheck(t, host, testPlayersWithHost, func(e error) bool { return e == nil })
}
