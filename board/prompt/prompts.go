package prompt

import (
	"errors"

	"github.com/nikiforosFreespirit/msdb5/card"
)

// PromptCard func
func PromptCard(prompt func(chan<- card.ID), cardChan chan card.ID) card.ID {
	go prompt(cardChan)
	return <-cardChan
}

// PromptScore func
func PromptScore(prompt func(chan<- uint8), cardChan chan uint8) uint8 {
	go prompt(cardChan)
	return <-cardChan
}

func EvaluateScore(before, after uint8) error {
	return errors.New("Score is low")
}
