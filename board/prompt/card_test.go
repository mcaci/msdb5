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

type RoundData struct {
	winningCardIndex uint8
	winningCard      card.ID
}

func Round() RoundData {
	briscola := card.Coin
	pChans := make([]chan card.ID, 5)
	for i := range pChans {
		pChans[i] = make(chan card.ID)
	}
	var playedCards set.Cards
	var roundData RoundData
	for i, prompt := range cardPrompts {
		nextCard := Card(prompt, pChans[i])
		playedCards.Add(nextCard)
		roundData.winningCard = rule.WinningCard(roundData.winningCard, nextCard, briscola)
		roundData.winningCardIndex = rule.IndexOfWinningCard(playedCards, briscola)
	}
	return roundData
}

func TestBoardRoundExecutionOneShot(t *testing.T) {
	var expectedWinningCardIndex uint8 = 2
	data := Round()
	if expectedWinningCardIndex != data.winningCardIndex {
		t.Fatal("Unexpected winner")
	}
}

func TestBoardRoundExecutionStepByStep(t *testing.T) {
	data := Round()
	expectedWinningCard := card.ID(3)
	if expectedWinningCard != data.winningCard {
		t.Fatal("Unexpected winner")
	}
}
