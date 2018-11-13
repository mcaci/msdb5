package rule

import "github.com/nikiforosFreespirit/msdb5/card/set"

// Points func
func Points(cards set.Cards) (sum uint8) {
	for _, card := range cards {
		sum += card.Points()
	}
	return sum
}
