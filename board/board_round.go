package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/card/set"
	"github.com/nikiforosFreespirit/msdb5/rule"
)

// PromptNext func
func PromptNext(actualWinningCard card.ID, briscola card.Seed, prompt func(chan<- card.ID), cardChan chan card.ID) card.ID {
	nextID := Prompt(prompt, cardChan)
	if &actualWinningCard == nil || rule.DoesOtherCardWin(actualWinningCard, nextID, briscola) {
		actualWinningCard = nextID
	}
	return actualWinningCard
}

// Prompt func
func Prompt(prompt func(chan<- card.ID), cardChan chan card.ID) card.ID {
	go prompt(cardChan)
	return <-cardChan
}

// IndexOfWinningCard func
func IndexOfWinningCard(cardsOnTheTable set.Cards, briscola card.Seed, rule func(card.ID, card.ID, card.Seed) bool) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if rule(base, other, briscola) {
			base = other
			max = i
		}
	}
	return uint8(max)
}
