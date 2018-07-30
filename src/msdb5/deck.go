package msdb5

type Deck struct {
}

func (d *Deck) First() *Card {
	return &Card{number: 1, seed: Coin}
}
