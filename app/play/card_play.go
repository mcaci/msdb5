package play

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/team"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Play func
func Play(c card.ID, cards *deck.Cards) error {
	index := cards.Find(c)
	if index == -1 {
		return fmt.Errorf("Card is not in players hand")
	}
	func(cards, to *deck.Cards, index, toIndex int) {
		*cards = append((*cards)[:index], (*cards)[index+1:]...)
	}(cards, nil, index, 0)
	return nil
}

// Exchange func
func Exchange(c card.ID, cards, to *deck.Cards) error {
	index := cards.Find(c)
	if index == -1 {
		return fmt.Errorf("Card is not in players hand")
	}
	func(cards, to *deck.Cards, index, toIndex int) {
		(*cards)[index], (*to)[toIndex] = (*to)[index], (*cards)[toIndex]
	}(cards, to, index, 0)
	return nil
}

// Companion func
func Companion(c card.ID, players team.Players, setCompanion func(*player.Player), setBriscolaCard func(card.ID)) error {
	setBriscolaCard(c)
	_, pl := players.Find(func(p *player.Player) bool { return p.Has(c) })
	setCompanion(pl)
	return nil
}
