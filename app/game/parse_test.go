package game

import (
	"testing"
)

const full = "Action#1#Coin"
const part = "Action#1"
const name = "Action"
const empty = ""

func TestParseCommand(t *testing.T) {
	res := command(name)
	if res != name {
		t.Fatal("Unexpected value")
	}
}

func TestParseEmptyCommand(t *testing.T) {
	res := command(empty)
	if res != empty {
		t.Fatal("Unexpected value")
	}
}

func TestParsePartCommand(t *testing.T) {
	res := value(part)
	if res != "1" {
		t.Fatal("Unexpected value")
	}
}

func TestParseEmptyPartCommand(t *testing.T) {
	res := value(name)
	if res != empty {
		t.Fatal("Unexpected value")
	}
}
