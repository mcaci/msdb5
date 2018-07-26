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

func createPlayer(numberOfCards uint8) *Player {
	p := new(Player)
	for i := 0; i < int(numberOfCards); i++ {
		p.assignCard(new(Card))
	}
	return p
}

func TestPlayerThatPlaysCardActuallyChoosesACard(t *testing.T) {
	p := createPlayer(1)
	playedCard := p.playCard()
	if playedCard == nil {
		t.Fatal("Expecting player to have chosen a card")
	}
}

func TestPlayerThatPlaysCardActuallyGivesTheCard(t *testing.T) {
	p := createPlayer(1)
	firstCard := p.card
	p.playCard()
	newFirstCard := p.card
	if firstCard == newFirstCard {
		t.Fatal("Expecting player to have given the card")
	}
}
