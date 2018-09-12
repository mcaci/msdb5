package board

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/deck"
)

type Board struct {
}

func (b *Board) Deck() *deck.Deck {
	return nil
}

func TestBoardHasADeck(t *testing.T) {
	b := Board{}
	if b.Deck() == nil {
		t.Fatal("The board has no Deck")
	}

}
