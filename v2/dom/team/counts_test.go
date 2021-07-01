package team

import (
	"testing"

	"github.com/mcaci/msdb5/v2/dom/player"
)

func TestCount(t *testing.T) {
	p := player.New(&player.Options{For5P: true}).(*player.B5Player)
	if count := Count(Players{p, p}, func(pl player.Player) bool { return true }); count != 2 {
		t.Fatal("Count should be 2")
	}
}
