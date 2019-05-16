package team

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
)

type mockCounter int

func (m mockCounter) Count(scorer func(card.ID) uint8) uint8 {
	return 1
}

func TestTeam1(t *testing.T) {
	fakeCaller := new(mockCounter)
	if score1, _ := Score(fakeCaller, nil, fakeCaller); score1 != 1 {
		t.Fatal("Count string should contain the total of 1")
	}
}

func TestTeam2(t *testing.T) {
	if _, score2 := Score(nil, nil, new(mockCounter)); score2 != 1 {
		t.Fatal("Count string should contain the total of 1")
	}
}
