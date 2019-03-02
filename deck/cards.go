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

// Move func
func (cards *Cards) Move(destination *Cards) {
	destination.Add(*cards...)
	*cards = Cards{}
}

// Find func
func (cards *Cards) Find(isInfoPresent func(c card.ID) bool) (int, error) {
	for index, c := range *cards {
		if isInfoPresent(c) {
			return index, nil
		}
	}
	return -1, errors.New("Card not found")
}

// Has func
func (cards Cards) Has(id card.ID) bool {
	_, err := cards.Find(func(c card.ID) bool { return c == id })
	return err == nil
}

// Supply func
func (cards *Cards) Supply() card.ID {
	card := (*cards)[0]
	cards.Remove(0)
	return card
}

func (cards Cards) String() (str string) {
	for _, card := range cards {
		str += card.String() + " "
	}
	return
}
