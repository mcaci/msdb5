package card

// ToNumber func
func (id ID) ToNumber() uint8 {
	return id.toZeroBased()%10 + 1
}

// ToSeed func
func (id ID) ToSeed() Seed {
	return Seed(id.toZeroBased() / 10)
}

func (id ID) toZeroBased() uint8 {
	return uint8(id) - 1
}
