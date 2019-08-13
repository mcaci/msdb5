package action

import (
	"container/list"
	"testing"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type fakeGame struct {
	current *player.Player
	phase   phase.ID
	players team.Players
}

func newTestGame(ph phase.ID) fakeGame {
	f := fakeGame{}
	t := team.Players{}
	p := player.New()
	p.Join("127.0.0.51")
	p.RegisterAs("A")
	t.Add(p)
	f.players = t
	f.current = p
	f.phase = ph
	return f
}

func (g fakeGame) CurrentPlayer() *player.Player { return g.current }
func (g fakeGame) Players() team.Players         { return g.players }
func (g fakeGame) LastPlaying() *list.List       { return list.New() }
func (g fakeGame) Phase() phase.ID               { return g.phase }

type fakerequester struct{ action, from string }

func (rq fakerequester) From() string   { return rq.from }
func (rq fakerequester) Action() string { return rq.action }

func TestVerifyPlayerWithNoErr(t *testing.T) {
	err := VerifyPlayer(newTestGame(0), fakerequester{"Auction", "127.0.0.51"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPlayerWithErr(t *testing.T) {
	err := VerifyPlayer(newTestGame(0), fakerequester{"Auction", "127.0.0.52"})
	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestVerifyPhaseWithNoErr(t *testing.T) {
	err := VerifyPhase(newTestGame(0), fakerequester{"Join", "127.0.0.51"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPhaseWithErr(t *testing.T) {
	err := VerifyPhase(newTestGame(4), fakerequester{"Join", "127.0.0.51"})
	if err == nil {
		t.Fatal("Error was expected")
	}
}
