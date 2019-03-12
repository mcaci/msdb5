package team

import (
	"strings"
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

type mockCounter int

func (m mockCounter) Count(scorer func(card.ID) uint8) uint8 {
	return 1
}

func TestTeam(t *testing.T) {
	team := new(BriscolaTeam)
	team.Add(mockCounter(1))
	if !strings.Contains(team.Info("Test").Display(), "1") {
		t.Fatal("Count string should contain the total of 1")
	}
}
