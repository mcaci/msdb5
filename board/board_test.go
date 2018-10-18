package board

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func TestBoardHasADeck(t *testing.T) {
	b := New()
	if b.Deck() == nil {
		t.Fatal("The board has no Deck")
	}
}

func TestBoardHas5Player(t *testing.T) {
	b := New()
	if b.Players() == nil {
		t.Fatal("The board has no Player")
	}
}

func Test5PlayersDrawUntilDeckIsEmpty(t *testing.T) { // not a Unit test
	d := deck.Deck()

	var players [5]player.Player
	for i := range players {
		players[i] = player.New()
	}

	for i := 0; i < deck.Size; i++ {
		players[i%5].Draw(&d)
	}

	if !d.IsEmpty() {
		t.Fatal("All players should have drawn all cards")
		for _, player := range players {
			t.Log(player)
		}
	}
}
