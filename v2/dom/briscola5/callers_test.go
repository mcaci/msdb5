package briscola5

import (
	"testing"
)

func TestTeamCallers(t *testing.T) {
	fakePlayer := NewPlayer()
	if !IsInCallers(NewCallersTeam(&fakePlayer.Player, &NewPlayer().Player))(&fakePlayer.Player) {
		t.Fatal("Player should be in Callers")
	}
}

func TestTeamOthers(t *testing.T) {
	fakePlayer := NewPlayer()
	if IsInCallers(NewCallersTeam(&NewPlayer().Player, &NewPlayer().Player))(&fakePlayer.Player) {
		t.Fatal("Player should be in Others")
	}
}
