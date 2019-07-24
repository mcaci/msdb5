package phase

import (
	"testing"

	"github.com/mcaci/msdb5/dom/player"
)

type companiontest string

func (c companiontest) Find(player.Predicate) (int, *player.Player) {
	return 1, nil
}

func (c companiontest) Value() string { return string(c) }

func TestCompanionIndex(t *testing.T) {
	data := Companion(fakeInput(1), companiontest("A"))
	if data.CompIdx() != 1 {
		t.Fatal("Unexpected player")
	}
}

func TestCompanionCard(t *testing.T) {
	data := Companion(fakeInput(11), companiontest("A"))
	if data.Briscola() != 11 {
		t.Fatal("Unexpected briscola")
	}
}
