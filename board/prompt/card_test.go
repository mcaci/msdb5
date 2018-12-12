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

func RoundData() (winningCard card.ID, winningCardIndex uint8) {
	briscola := card.Coin
	pChans := make([]chan card.ID, 5)
	for i := range pChans {
		pChans[i] = make(chan card.ID)
	}
	var playedCards set.Cards
	for i, prompt := range cardPrompts {
		nextCard := Card(prompt, pChans[i])
		playedCards.Add(nextCard)
		winningCard = rule.WinningCard(winningCard, nextCard, briscola)
		winningCardIndex = rule.IndexOfWinningCard(playedCards, briscola)
	}
	return
}

func TestBoardRoundExecutionOneShot(t *testing.T) {
	var expectedWinningCardIndex uint8 = 2
	_, winningCardIndex := RoundData()
	if expectedWinningCardIndex != winningCardIndex {
		t.Fatal("Unexpected winner")
	}
}

func TestBoardRoundExecutionStepByStep(t *testing.T) {
	winningCard, _ := RoundData()
	expectedWinningCard := card.ID(3)
	if expectedWinningCard != winningCard {
		t.Fatal("Unexpected winner")
	}
}
