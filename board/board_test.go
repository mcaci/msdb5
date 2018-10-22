package board

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func TestBoardHasADeck(t *testing.T) {
	if b := New(); b.Deck() == nil {
		t.Fatal("The board has no Deck")
	}
}

func TestBoardHas5Player(t *testing.T) {
	if b := New(); b.Players() == nil {
		t.Fatal("The board has no Player")
	}
}

func Test5PlayersDrawUntilDeckIsEmpty(t *testing.T) { // not a Unit test
	d := card.Deck()

	var players [5]player.Player
	for i := range players {
		players[i] = player.New()
	}

	for i := 0; i < card.DeckSize; i++ {
		players[i%5].Draw(&d)
	}

	if !d.IsEmpty() {
		t.Fatal("All players should have drawn all cards")
		for _, player := range players {
			t.Log(player)
		}
	}
}
