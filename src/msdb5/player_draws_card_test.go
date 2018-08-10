package msdb5

import (
	"msdb5/card"
	"testing"
)

type MockDeck struct {
}

func (d *MockDeck) RemoveTop() *(card.Card) {
	return card.ByID(0)
}
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
