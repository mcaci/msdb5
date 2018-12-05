package board

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/rule"
)

var prompts = []func(chan<- card.ID){
	func(cardChan chan<- card.ID) {
		cardChan <- 5
	},
	func(cardChan chan<- card.ID) {
		cardChan <- 14
	},
	func(cardChan chan<- card.ID) {
		cardChan <- 3
	},
	func(cardChan chan<- card.ID) {
		cardChan <- 35
	},
	func(cardChan chan<- card.ID) {
		cardChan <- 27
	}}

func TestBoardRoundExecutionOneShot(t *testing.T) {
	b := New()
	var expectedWinningCardIndex uint8 = 2
	briscola := card.Coin
	for i, prompt := range prompts {
		nextCard := Prompt(prompt, b.pChans[i])
		b.PlayedCards().Add(nextCard)
	}
	if expectedWinningCardIndex != rule.IndexOfWinningCard(*b.PlayedCards(), briscola) {
		t.Fatal("Unexpected winner")
	}
}

func TestBoardRoundExecutionStepByStep(t *testing.T) {
	b := New()
	expectedWinningCard := card.ID(3)
	briscola := card.Coin
	var winningCard card.ID
	for i, prompt := range prompts {
		nextCard := Prompt(prompt, b.pChans[i])
		winningCard = rule.WinningCard(winningCard, nextCard, briscola)
	}
	if expectedWinningCard != winningCard {
		t.Fatal("Unexpected winner")
	}
}
