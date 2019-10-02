package score

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

func TestTeam1(t *testing.T) {
	fakePlayer := player.New()
	fakePlayer.Pile().Add(*set.NewMust(1)...)
	if score1, _ := Calc(team.NewCallers(fakePlayer), team.Players{fakePlayer}); score1 != 11 {
		t.Fatal("Points string should contain the total of 11")
	}
}

func TestTeam2(t *testing.T) {
	fakePlayer := player.New()
	fakePlayer.Pile().Add(*set.NewMust(1)...)
	if _, score2 := Calc(team.NewEmptyCallers(), team.Players{fakePlayer}); score2 != 11 {
		t.Fatal("Points string should contain the total of 11")
	}
}
