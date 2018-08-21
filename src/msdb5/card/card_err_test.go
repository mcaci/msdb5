package card

import "testing"

func TestId0IsInvalid(t *testing.T) {
	verifyInvalidID(t, 0)
}
func TestId41IsInvalid(t *testing.T) {
	verifyInvalidID(t, 41)
}

func verifyInvalidID(t *testing.T, id int) {
	_, err := ByID(id)
	if err == nil {
		t.Fatalf("%d is not valid id", id)
	}
}
