package player

import (
	"github.com/nikiforosFreespirit/msdb5/deck"
	"testing"
)

func TestPlayerDrawsOneCard(t *testing.T) {
	var d MockDeck
	player := New()
	playedCard := player.Draw(&d)
	if !player.has(playedCard) {
		t.Fatalf("Expecting player to have drawn %v", playedCard)
	}
}

func Test5PlayersDrawUntilDeckIsEmpty(t *testing.T) { // not a Unit test
	deck := deck.New()

	var players [5]Player
	for i := range players {
		players[i] = New()
	}

	for i := 0; i < 40; i++ {
		players[i%5].Draw(deck)
	}

	if !deck.IsEmpty() {
		t.Fatal("All players should have drawn all cards")
		for _, player := range players {
			t.Log(player)
		}
	}
}
