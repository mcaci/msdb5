package auction

// SideCards func
func SideCards(score Score) uint8 {
	return uint8(score/90 + score/100 + score/110 + score/120 + score/120)
}
