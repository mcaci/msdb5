package team

import (
	"testing"
)

type mockCounter struct{}

func (mockCounter) Folded() bool      { return true }
func (mockCounter) IsNameEmpty() bool { return true }
func (mockCounter) IsHandEmpty() bool { return true }

type mockCounterFalse struct{}

func (mockCounterFalse) IsHandEmpty() bool { return false }

func TestCountFolded(t *testing.T) {
	if count := CountFolded(new(mockCounter), new(mockCounter)); count != 2 {
		t.Fatal("Count should be 2")
	}
}

func TestCountEmptyNames(t *testing.T) {
	if count := CountEmptyNames(new(mockCounter)); count != 1 {
		t.Fatal("Count should be 1")
	}
}

func TestCountEmptyHands(t *testing.T) {
	if count := CountEmptyHands(new(mockCounter), new(mockCounter), new(mockCounterFalse)); count != 2 {
		t.Fatal("Count should be 2")
	}
}
