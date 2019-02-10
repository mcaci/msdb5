package card

func toZeroBased(index ID) uint8 {
	return uint8(index) - 1
}

func mapToID(number uint8, seed Seed) uint8 {
	return number + (uint8)(seed)*10
}

// FromZeroBased func
func FromZeroBased(index int) uint8 {
	return uint8(index) + 1
}
