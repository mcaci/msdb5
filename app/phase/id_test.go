package phase

import (
	"testing"
)

func TestIDCreation(t *testing.T) {
	_, err := ToID("Card")
	if err != nil {
		t.Fatal("Unexpected error")
	}
}
