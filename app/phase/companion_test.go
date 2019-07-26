package phase

import (
	"testing"

	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/player"
)

type companiontest string

func (c companiontest) Find(player.Predicate) (int, *player.Player) {
	return 1, nil
}

type fakeInput card.ID

func (rq fakeInput) Card() (card.ID, error) {
	return card.ID(rq), nil
}

func (c companiontest) Value() string { return string(c) }

func TestCompanionIndex(t *testing.T) {
	data := Companion(fakeInput(1), companiontest("A"))
	if data.PlIdx() != 1 {
		t.Fatal("Unexpected player")
	}
}

func TestCompanionCard(t *testing.T) {
	data := Companion(fakeInput(11), companiontest("A"))
	if data.Card() != 11 {
		t.Fatal("Unexpected briscola")
	}
}

type errortest struct{}

func (e errortest) Find(player.Predicate) (int, *player.Player) {
	return -1, nil
}

func TestCompanionErr(t *testing.T) {
	data := Companion(fakeInput(11), errortest{})
	if data.CardNotFound() != nil {
		t.Fatal("Error is expected")
	}
}
