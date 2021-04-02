package srv

import (
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/pb"
)

func toScorers(cs pb.Cards) ([]interface{ Number() uint8 }, error) {
	scorers := make([]interface{ Number() uint8 }, len(cs.Cards))
	for i := range scorers {
		c, err := card.FromID(uint8(cs.Cards[i].Id))
		if err != nil {
			return nil, fmt.Errorf("error found: %w; could not convert %d to card, inside the set %v", err, cs.Cards[i].Id, cs.Cards)
		}
		scorers[i] = c
	}
	return scorers, nil
}
