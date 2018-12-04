package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/card/set"
	"github.com/nikiforosFreespirit/msdb5/rule"
)

// PlayedCards func
func (b *Board) PlayedCards() *set.Cards {
	return &b.playedCards
}

func PromptNext(actualWinningCard card.ID, briscola card.Seed, prompt func(chan<- card.ID), cardChan chan card.ID) card.ID {
	nextID := Prompt(prompt, cardChan)
	if &actualWinningCard == nil || rule.DoesOtherCardWin(actualWinningCard, nextID, briscola) {
		actualWinningCard = nextID
	}
	return actualWinningCard
}

func Prompt(prompt func(chan<- card.ID), cardChan chan card.ID) card.ID {
	go prompt(cardChan)
	return <-cardChan
}
