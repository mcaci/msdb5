package player

import "msdb5/card"

type MockDeck struct {
}

func (d *MockDeck) RemoveTop() *card.Card {
	mockCard, _ := card.ByID(0)
	return mockCard
}

func (d *MockDeck) IsEmpty() bool {
	return false
}
