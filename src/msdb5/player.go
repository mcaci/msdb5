package msdb5

type Player interface {
	Draw(d *Deck) *Card
	Has(c *Card) bool
}

// func (p *Player) assignCard(card *Card) { p.card = card }

// func (p *Player) playCard() *Card {
// 	chosenCard := p.card
// 	p.card = nil
// 	return chosenCard
// }

// func createPlayer(numberOfCards uint8) *Player {
// 	p := new(Player)
// 	for i := 0; i < int(numberOfCards); i++ {
// 		p.assignCard(new(Card))
// 	}
// 	return p
// }
