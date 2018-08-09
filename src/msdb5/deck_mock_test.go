package msdb5

import "msdb5/card"

type MockDeck struct {
}

func (d *MockDeck) RemoveTop() CardPtr {
	return card.ByID(0)
}
