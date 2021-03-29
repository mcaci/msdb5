package briscola

// Points returns the value of each card number according to Briscola rules
func Points(valuer interface{ Number() uint8 }) uint8 {
	var points = map[uint8]uint8{1: 11, 3: 10, 8: 2, 9: 3, 10: 4}
	return points[valuer.Number()]
}
