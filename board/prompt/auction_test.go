package prompt

import (
	"testing"
)

var playerSays61 = func(cardChan chan<- uint8) {
	cardChan <- 61
}

func TestAuctionPromptRound_NoChecks(t *testing.T) {
	scoreChan := make(chan uint8)
	expectedScore := uint8(61)
	actualScore := PromptScore(playerSays61, scoreChan)
	if expectedScore != actualScore {
		t.Fatal("Unexpected score")
	}
}

