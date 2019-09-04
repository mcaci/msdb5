package input

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

const full = "Action#1#Coin"
const part = "Action#1"
const name = "Action"
const empty = ""

func TestParseCommand(t *testing.T) {
	res := Command(name)
	if res != name {
		t.Fatal("Unexpected value")
	}
}

func TestParseEmptyCommand(t *testing.T) {
	res := Command(empty)
	if res != empty {
		t.Fatal("Unexpected value")
	}
}

func TestParsePartCommand(t *testing.T) {
	res := Value(part)
	if res != "1" {
		t.Fatal("Unexpected value")
	}
}

func TestParseEmptyPartCommand(t *testing.T) {
	res := Value(name)
	if res != empty {
		t.Fatal("Unexpected value")
	}
}

func TestParseCardCommand(t *testing.T) {
	c, _ := Card(full)
	if c.Number() != 1 && c.Seed() != card.Coin {
		t.Fatal("Unexpected value")
	}
}

func TestParseEmptyCardCommand(t *testing.T) {
	_, err := Card(part)
	if err == nil {
		t.Fatal("Unexpected value")
	}
}
