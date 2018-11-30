package board

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/rule"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestBoardRoundExecution(t *testing.T) {
	// Create board
	b := New()

	cardChan := make(chan card.ID)
	var id1 card.ID

	// Prompt card to player 1
	go PromptCard1(cardChan)
	id1 = <-cardChan
	b.PlayedCards().Add(id1)

	go PromptCard2(cardChan)
	id1 = <-cardChan
	b.PlayedCards().Add(id1)

	go PromptCard3(cardChan)
	id1 = <-cardChan
	b.PlayedCards().Add(id1)

	go PromptCard4(cardChan)
	id1 = <-cardChan
	b.PlayedCards().Add(id1)

	go PromptCard5(cardChan)
	id1 = <-cardChan
	b.PlayedCards().Add(id1)

	if 2 != IndexOfWinningCard(*b.PlayedCards(), card.Coin, rule.DoesOtherCardWin) {
		t.Fatal("Unexpected winner")
	}

}

func PromptCard1(cardChan chan<- card.ID) {
	cardChan <- 5
}

func PromptCard2(cardChan chan<- card.ID) {
	cardChan <- 14
}

func PromptCard3(cardChan chan<- card.ID) {
	cardChan <- 3
}

func PromptCard4(cardChan chan<- card.ID) {
	cardChan <- 35
}

func PromptCard5(cardChan chan<- card.ID) {
	cardChan <- 27
}
