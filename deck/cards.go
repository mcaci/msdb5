package deck

import (
	"errors"

	"github.com/nikiforosFreespirit/msdb5/card"
)

// Cards type
type Cards []card.ID

// Add func
func (cards *Cards) Add(ids ...card.ID) {
	*cards = append(*cards, ids...)
}

// Remove func
func (cards *Cards) Remove(index int) {
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}

// Clear func
func (cards *Cards) Clear() {
	*cards = Cards{}
}

// Sum func
func (cards *Cards) Sum(point func(card.ID) uint8) (sum uint8) {
	for _, c := range *cards {
		sum += point(c)
	}
	return
}

// Find func
func (cards *Cards) Find(id card.ID) (int, error) {
	f := func(c card.ID) bool { return c == id }
	return cards.find(f)
}

func (cards *Cards) find(isInfoPresent func(c card.ID) bool) (int, error) {
	for index, c := range *cards {
		if isInfoPresent(c) {
			return index, nil
		}
	}
	return -1, errors.New("Card not found")
}

// Has func
func (cards Cards) Has(id card.ID) bool {
	_, err := cards.Find(id)
	return err == nil
}

// Supply func
func (cards *Cards) Supply() card.ID {
	card := (*cards)[0]
	cards.Remove(0)
	return card
}