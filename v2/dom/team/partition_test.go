package team

import (
	"testing"

	"github.com/mcaci/msdb5/v2/dom/player"
)

func TestPartitionT1(t *testing.T) {
	pred := func(p *player.Player) bool { return p.Name() == "A" }
	t1, _ := testPlayers.Part(pred)
	if t1.None(pred) {
		t.Fatal("t1 should contain only players named A")
	}
}

func TestPartitionT2(t *testing.T) {
	pred := func(p *player.Player) bool { return p.Name() == "A" }
	_, t2 := testPlayers.Part(pred)
	if !t2.None(pred) {
		t.Fatal("t2 should not contain players named A")
	}
}
