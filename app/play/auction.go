package play

import (
	"github.com/mcaci/msdb5/dom/auction"
)

// SideCardsToDisplay func
func SideCardsToDisplay(score auction.Score) uint8 {
	var cardsToShow uint8
	switch {
	case score < 90:
		cardsToShow = 0
	case score < 100:
		cardsToShow = 1
	case score < 110:
		cardsToShow = 2
	case score < 120:
		cardsToShow = 3
	default:
		cardsToShow = 5
	}
	return cardsToShow
}
