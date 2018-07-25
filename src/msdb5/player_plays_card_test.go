package msdb5

import "testing"

type Player struct {
	card *Card
}
type Card struct{}

func (p *Player) assignCard(card *Card) { p.card = card }
func (p *Player) playCard() *Card       { return p.card }

func TestPlayerThatPlaysCardHasNotItsCardAnymore(t *testing.T) {
	p := new(Player)
	card := new(Card)
	p.assignCard(card)
	playedCard := p.playCard()
	if playedCard == nil {
		t.Fatal("Expecting player to have played a card")
	}
}
