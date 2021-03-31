package briscola5

import "github.com/mcaci/ita-cards/set"

type Side struct{ set.Cards }

type PlayedCards struct{ *set.Cards }

func (c PlayedCards) Pile() *set.Cards {
	if len(*c.Cards) == 5 {
		return (*set.Cards)(c.Cards)
	}
	return &set.Cards{}
}
