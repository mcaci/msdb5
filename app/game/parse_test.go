package game

import (
	"testing"
)

const part = "Action#1"
const name = "Action"
const empty = ""

func TestParseCommand(t *testing.T) {
	res := parse(name, 0)
	if res != name {
		t.Fatal("Unexpected value")
	}
}

func TestParseEmptyCommand(t *testing.T) {
	res := parse(empty, 0)
	if res != empty {
		t.Fatal("Unexpected value")
	}
}

func TestParsePartCommand(t *testing.T) {
	res := parse(part, 1)
	if res != "1" {
		t.Fatal("Unexpected value")
	}
}

func TestParseEmptyPartCommand(t *testing.T) {
	res := parse(name, 1)
	if res != empty {
		t.Fatal("Unexpected value")
	}
}
