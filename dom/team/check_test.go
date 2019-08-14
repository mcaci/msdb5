package team

import (
	"testing"

	"github.com/mcaci/msdb5/dom/player"
)

type fakeGame struct {
	current *player.Player
	players Players
}

func newTestGame() fakeGame {
	f := fakeGame{}
	t := Players{}
	p := player.New()
	p.Join("127.0.0.51")
	p.RegisterAs("A")
	t.Add(p)
	f.players = t
	f.current = p
	return f
}

func (g fakeGame) CurrentPlayer() *player.Player { return g.current }
func (g fakeGame) Players() Players              { return g.players }

type fakeFrom string

func (rq fakeFrom) From() string { return string(rq) }

func TestVerifyPlayerWithNoErr(t *testing.T) {
	err := CheckOrigin(newTestGame(), fakeFrom("127.0.0.51"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPlayerWithErr(t *testing.T) {
	err := CheckOrigin(newTestGame(), fakeFrom("127.0.0.52"))
	if err == nil {
		t.Fatal("Error was expected")
	}
}
