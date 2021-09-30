package briscola

import (
	"fmt"
	"log"

	"github.com/mcaci/ita-cards/set"
)

func Play(g interface {
	Players() *Players
	InTurn() *Player
	nextPlayer(func() int)
	BriscolaCard() *Card
	board() *PlayedCards
	deckCards() *Deck
	roundrobin() int
}, b interface {
	Name() string
	Idx() uint8
}) (*struct {
	Pl  *Player    `json:"player"`
	Brd *set.Cards `json:"board"`
}, error) {
	pl := g.InTurn()
	if pl.Name() != b.Name() {
		return nil, fmt.Errorf("player %q is not expected to play`, turn of player %q", b.Name(), pl.Name())
	}
	hand := pl.Hand()
	if int(b.Idx()) >= len(*hand) {
		return nil, fmt.Errorf("card number %d is present to play, choose between 0 and %d", b.Idx(), len(*hand))
	}
	card := (*hand)[b.Idx()]
	log.Println("player", pl.Name(), "playing card:", card)
	err := set.MoveOne(&card, hand, g.board().Cards)
	if err != nil {
		return nil, err
	}
	g.nextPlayer(g.roundrobin)
	isRoundOngoing := func(playedCards set.Cards) bool { return len(playedCards) < 2 }
	if isRoundOngoing(*g.board().Cards) {
		return &struct {
			Pl  *Player    `json:"player"`
			Brd *set.Cards `json:"board"`
		}{Pl: pl, Brd: g.board().Cards}, nil
	}
	winId := Winner(*g.board().Cards, g.BriscolaCard().Seed())
	g.nextPlayer(func() int { return int(winId) })
	plWin := (*g.Players())[winId]
	plLos := (*g.Players())[(winId+1)%2]
	set.Move(g.board().Pile(), plWin.Pile())

	switch len(g.deckCards().Cards) {
	case 0:
	case 1:
		plWin.Hand().Add(g.deckCards().Top())
		plLos.Hand().Add(g.BriscolaCard().Item)
	default:
		plWin.Hand().Add(g.deckCards().Top())
		plLos.Hand().Add(g.deckCards().Top())
	}
	return &struct {
		Pl  *Player    `json:"player"`
		Brd *set.Cards `json:"board"`
	}{Pl: pl, Brd: g.board().Cards}, nil
}
