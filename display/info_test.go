package display

import "testing"

func TestPrintOfEmptyDataIsEmpty(t *testing.T) {
	info := NewInfo("", "", "", "")
	if info.PrintIt() != "" {
		t.Fatal("nothing should be in the output")
	}
}
