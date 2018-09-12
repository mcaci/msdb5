package score

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// Nominate func
func Nominate(id uint8) (*card.Card, error) {
	return card.ByID(id)
}
