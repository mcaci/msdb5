package input

import (
	"testing"

	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type fakeGameForOrigin struct {
	current *player.Player
	players team.Players
	origin  string
}

func newTestGame(origin string) fakeGameForOrigin {
	f := fakeGameForOrigin{}
	t := team.Players{}
	p := player.New()
	p.Join("127.0.0.51")
	p.RegisterAs("A")
	t.Add(p)
	f.players = t
	f.current = p
	f.origin = origin
	return f
}

func TestVerifyPlayerWithNoErr(t *testing.T) {
	g := newTestGame("127.0.0.51")
	err := CheckOrigin(g.players, g.origin, g.current)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPlayerWithErr(t *testing.T) {
	g := newTestGame("127.0.0.52")
	err := CheckOrigin(g.players, g.origin, g.current)
	if err == nil {
		t.Fatal("Error was expected")
	}
}
