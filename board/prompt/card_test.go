package prompt

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/card/set"
	"github.com/nikiforosFreespirit/msdb5/rule"
)

var cardPrompts = []func(chan<- card.ID){
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
	briscola := card.Coin
	pChans := make([]chan card.ID, 5)
	for i := range pChans {
		pChans[i] = make(chan card.ID)
	}
	playedCards := set.Cards{}
	for i, prompt := range cardPrompts {
		nextCard := Card(prompt, pChans[i])
		playedCards.Add(nextCard)
	}
	var expectedWinningCardIndex uint8 = 2
	if expectedWinningCardIndex != rule.IndexOfWinningCard(playedCards, briscola) {
		t.Fatal("Unexpected winner")
	}
}

func TestBoardRoundExecutionStepByStep(t *testing.T) {
	briscola := card.Coin
	pChans := make([]chan card.ID, 5)
	for i := range pChans {
		pChans[i] = make(chan card.ID)
	}
	var winningCard card.ID
	for i, prompt := range cardPrompts {
		nextCard := Card(prompt, pChans[i])
		winningCard = rule.WinningCard(winningCard, nextCard, briscola)
	}
	expectedWinningCard := card.ID(3)
	if expectedWinningCard != winningCard {
		t.Fatal("Unexpected winner")
	}
}
