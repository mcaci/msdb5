package srv

import (
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/pb"
)

func toCards(cs *pb.Cards) (*set.Cards, error) {
	cards := make(set.Cards, len(cs.Cards))
	for i := range cards {
		c, err := card.FromID(uint8(cs.Cards[i].Id))
		if err != nil {
			return nil, fmt.Errorf("error found: %w; could not convert %d to card, inside the set %v", err, cs.Cards[i].Id, cs.Cards)
		}
		cards[i] = *c
	}
	return &cards, nil
}
