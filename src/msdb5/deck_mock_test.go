package msdb5

type MockDeck struct {
}

func (d *MockDeck) First() *Card {
	return &Card{number: 1, seed: Coin}
}
