package chantest

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestCardNominatedIsStoredOnBoard(t *testing.T) {
	board := board.New()
	pChan := make(chan card.ID)
	promptFunc := func(cardChan chan<- card.ID) { cardChan <- 21 }
	actualCard := AskNominatedCard(board, promptFunc, pChan)
	expectedCard := board.NominatedCard()
	if *expectedCard != actualCard {
		t.Fatalf("Card nominated should be %v but %v was computed", expectedCard, actualCard)
	}
}
