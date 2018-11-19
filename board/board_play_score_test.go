package board

import (
	"testing"
)

func TestAllPlayersHave0ScoreWhenBoardIsCreated(t *testing.T) {
	b := New()
	actualScore := b.Players()[0].Score()
	if 0 != actualScore {
		t.Fatal("Score should be 0 at the beginning")
	}
}
