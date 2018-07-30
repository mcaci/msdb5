package msdb5

import (
	"container/list"
	"testing"
)

func TestPlayerDrawsOneCard(t *testing.T) {
	var d Deck
	var p Player
	p = &ConcretePlayer{cards: list.New()}
	playedCard := p.Draw(&d)
	if p.Hasnt(playedCard) {
		t.Fatalf("Expecting player to have drawn %v", playedCard)
	}
}