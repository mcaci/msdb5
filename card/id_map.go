package card

func mapIDToNumber(index ID) uint8 {
	return (uint8(index)-1)%10 + 1
}

func mapIDToSeed(index ID) Seed {
	return Seed((uint8(index) - 1) / 10)
}
