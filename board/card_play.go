package board

import "github.com/nikiforosFreespirit/msdb5/card"

// Play func
func (b *Board) Play(number, seed, origin string) {
	p, _ := b.Players().Find(origin)
	c, _ := p.Play(number, seed)
	b.PlayedCards().Add(c)
}

func playCard(b *Board) card.ID {
	h := b.Players()[0].Hand()
	card := (*h)[0]
	removeCardFromHand(card, h)
	b.PlayedCards().Add(card)
	return card
}

func removeCardFromHand(c card.ID, h *card.Cards) {
	index := 0
	for i, card := range *h {
		if card == c {
			index = i
			break
		}
	}
	*h = append((*h)[:index], (*h)[index+1])
}
