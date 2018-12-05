package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// Prompt func
func Prompt(prompt func(chan<- card.ID), cardChan chan card.ID) card.ID {
	go prompt(cardChan)
	return <-cardChan
}
