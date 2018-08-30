package player

import (
	"msdb5/card"
	"msdb5/deck"
	"testing"
)

type MockDeck struct {
}

func (d *MockDeck) RemoveTop() *card.Card {
	mockCard, _ := card.ByID(0)
	return mockCard
}

func (d *MockDeck) IsEmpty() bool {
	return false
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

func Test5PlayersDrawUntilDeckIsEmpty(t *testing.T) { // not a Unit test
	var deck deck.ConcreteDeck
	deck.Create()

	var players [5]Player
	for i := range players {
		players[i] = &ConcretePlayer{}
		players[i].Init()
	}

	for i := 0; i < 40; i++ {
		players[i%5].Draw(&deck)
	}

	if !deck.IsEmpty() {
		t.Fatal("All players should have drawn all cards")
		for _, player := range players {
			t.Log(player)
		}
	}
}
