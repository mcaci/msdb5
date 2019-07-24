package phase

import "testing"

type jointest string

func (j jointest) Value() string { return string(j) }

func TestProcessRequestWithNoErr(t *testing.T) {
	data := Join(jointest("A"))
	if data.Name() != "A" {
		t.Fatal("Unexpected name")
	}
}
