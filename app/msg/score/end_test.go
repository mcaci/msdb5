package score

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type mockTeam struct {
	call *player.Player
}

func (m mockTeam) Caller() *player.Player  { return m.call }
func (mockTeam) Companion() *player.Player { return player.New() }

func TestTeam1(t *testing.T) {
	fakePlayer := player.New()
	fakePlayer.Pile().Add(*set.NewMust(1)...)
	if score1, _ := Calc(mockTeam{fakePlayer}, team.Players{fakePlayer}); score1 != 11 {
		t.Fatal("Points string should contain the total of 11")
	}
}

func TestTeam2(t *testing.T) {
	fakePlayer := player.New()
	fakePlayer.Pile().Add(*set.NewMust(1)...)
	if _, score2 := Calc(mockTeam{}, team.Players{fakePlayer}); score2 != 11 {
		t.Fatal("Points string should contain the total of 11")
	}
}
