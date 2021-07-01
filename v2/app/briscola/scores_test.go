package briscola

import (
	"errors"
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func TestScore(t *testing.T) {
	expected := "[: 0], [: 1]"
	actual := PrintScore(&struct {
		Players *team.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: team.New(2),
		Method:  func(i int) (interface{ GetPoints() uint32 }, error) { p := briscola.Pnts(i); return &p, nil },
	})
	if expected != actual {
		t.Errorf("Expecting %q, actual %q", expected, actual)
	}
}

func TestScoreWithErr(t *testing.T) {
	expected := ""
	actual := PrintScore(&struct {
		Players *team.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: team.New(2),
		Method: func(i int) (interface{ GetPoints() uint32 }, error) {
			p := briscola.Pnts(i)
			return &p, errors.New("error")
		},
	})
	if expected != actual {
		t.Errorf("Expecting errors but got %q", actual)
	}
}

func TestPlayerScore(t *testing.T) {
	players := team.New(2)
	(*players)[0] = player.New(&player.Options{Name: "Player 1", For2P: true})
	(*players)[0].Pile().Add(*card.MustID(1))
	(*players)[1] = player.New(&player.Options{Name: "Player 2", For2P: true})

	expected := "[Player 1: 11], [Player 2: 0]"
	actual := PrintScore(&struct {
		Players *team.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: players,
		Method: func(i int) (interface{ GetPoints() uint32 }, error) {
			p := briscola.Score(*(*players)[i].Pile())
			return p, nil
		},
	})
	if expected != actual {
		t.Errorf("Expecting %q, actual %q", expected, actual)
	}
}
