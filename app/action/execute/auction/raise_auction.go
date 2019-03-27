package auction

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/board"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type AuctionStruct struct {
	request, origin string
	board           *board.Board
}

func NewAuction(request, origin string, board *board.Board) action.Executer {
	return &AuctionStruct{request, origin, board}
}
func (as AuctionStruct) Do(p *player.Player) error {
	data := strings.Split(as.request, "#")
	score := data[1]
	auction.CheckAndUpdate(score, p.Folded, p.Fold, as.board.AuctionScore, as.board.SetAuctionScore)
	return nil
}
