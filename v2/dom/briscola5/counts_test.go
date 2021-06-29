package briscola5

import (
	"testing"

	"github.com/mcaci/msdb5/v2/dom/player"
)

func TestCount(t *testing.T) {
	p := player.New(&player.Options{For5P: true}).(*player.B5Player)
	if count := Count(Players{pls: []*player.B5Player{p, p}}, func(pl *player.B5Player) bool { return true }); count != 2 {
		t.Fatal("Count should be 2")
	}
}
