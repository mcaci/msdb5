package input

import (
	"testing"
)

const part = "Action#1"
const name = "Action"
const empty = ""

func TestParseCommand(t *testing.T) {
	res := Parse(name, 0)
	if res != name {
		t.Fatal("Unexpected value")
	}
}

func TestParseEmptyCommand(t *testing.T) {
	res := Parse(empty, 0)
	if res != empty {
		t.Fatal("Unexpected value")
	}
}

func TestParsePartCommand(t *testing.T) {
	res := Parse(part, 1)
	if res != "1" {
		t.Fatal("Unexpected value")
	}
}

func TestParseEmptyPartCommand(t *testing.T) {
	res := Parse(name, 1)
	if res != empty {
		t.Fatal("Unexpected value")
	}
}
