package msdb5

import "testing"

func TestPlayerDrawsOneCard(t *testing.T) {
	var d Deck
	p := ConcretePlayer{}
	playedCard := p.Draw(&d)
	if p.Hasnt(playedCard) {
		t.Fatalf("Expecting player to have drawn %v", playedCard)
	}
}