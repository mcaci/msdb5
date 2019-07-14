package play

import (
	"github.com/mcaci/msdb5/dom/auction"
)

// SideCardsToDisplay func
func SideCardsToDisplay(score auction.Score) uint8 {
	return uint8(score)/90 + uint8(score)/100 + uint8(score)/110 + uint8(score)/120 + uint8(score)/120
}
