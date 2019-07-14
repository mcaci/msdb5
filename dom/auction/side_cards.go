package auction

// SideCardsToDisplay func
func SideCardsToDisplay(score Score) uint8 {
	return uint8(score/90 + score/100 + score/110 + score/120 + score/120)
}
