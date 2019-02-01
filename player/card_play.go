package player

import "github.com/nikiforosFreespirit/msdb5/card"

// Play function
func (p *Player) Play(number, seed string) (card.ID, bool) {
	card, _ := card.ByName(number, seed)
	found := false
	for index, c := range *(p.Hand()) {
		found = c == card
		if found {
			p.Hand().Remove(index)
			break
		}
	}
	return card, found
}
