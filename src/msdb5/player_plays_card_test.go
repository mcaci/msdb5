package msdb5

import "testing"

type Player struct {
	card *Card
}
type Card struct{}

func (p *Player) assignCard(card *Card) { p.card = card }
func (p *Player) playCard() *Card {
	chosenCard := p.card
	p.card = nil
	return chosenCard
}

func TestPlayerThatPlaysCardActuallyChoosesACard(t *testing.T) {
	p := new(Player)
	card := new(Card)
	p.assignCard(card)
	playedCard := p.playCard()
	if playedCard == nil {
		t.Fatal("Expecting player to have chosen a card")
	}
}

func TestPlayerThatPlaysCardActuallyGivesTheCard(t *testing.T) {
	p := new(Player)
	card := new(Card)
	p.assignCard(card)
	p.playCard()
	if p.card != nil {
		t.Fatal("Expecting player to have given the card")
	}
}
