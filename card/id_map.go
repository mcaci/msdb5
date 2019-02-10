package card

func mapToID(number uint8, seed Seed) uint8 {
	return number + (uint8)(seed)*10
}

func mapIDToNumber(index ID) uint8 {
	return toZeroBased(index)%10 + 1
}

func mapIDToSeed(index ID) Seed {
	return Seed(toZeroBased(index) / 10)
}

func toZeroBased(index ID) uint8 {
	return uint8(index) - 1
}
