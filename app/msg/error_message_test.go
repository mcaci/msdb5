package msg

import (
	"testing"

	"golang.org/x/text/language"
)

func TestValidErrorCreation(t *testing.T) {
	err := Error("Hello", language.English)
	if err == nil {
		t.Fatal("Expecting error")
	}
}
