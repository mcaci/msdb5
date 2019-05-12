package auction

import (
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type AuctionStruct struct {
	request, origin string
	score           *auction.Score
}

func NewAuction(request, origin string, score *auction.Score) action.Executer {
	return &AuctionStruct{request, origin, score}
}
func (as AuctionStruct) Do(p *player.Player) error {
	if p.Folded() {
		return nil
	}
	data := strings.Split(as.request, "#")
	score := data[1]
	currentScore, err := strconv.Atoi(score)
	if err == nil && as.score.CheckWith(auction.Score(currentScore)) {
		as.score.Update(auction.Score(currentScore))
	} else {
		p.Fold()
	}
	return nil
}
