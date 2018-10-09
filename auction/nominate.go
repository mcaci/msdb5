package score

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// Nominate func
func Nominate(number, seed string) (*card.Card, error) {
	return card.ByName(number, seed)
}
