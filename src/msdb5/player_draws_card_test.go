package msdb5

import (
	"testing"
)

func TestPlayerDrawsOneCard(t *testing.T) {
	var d MockDeck
	var p Player
	p = &ConcretePlayer{}
	p.Init()
	playedCard := p.Draw(&d)
	if p.Hasnt(playedCard) {
		t.Fatalf("Expecting player to have drawn %v", playedCard)
	}
}
