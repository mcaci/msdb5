package prompt

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// PromptCard func
func PromptCard(prompt func(chan<- card.ID), cardChan chan card.ID) card.ID {
	go prompt(cardChan)
	return <-cardChan
}
