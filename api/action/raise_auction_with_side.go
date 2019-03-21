package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type AuctionWithSideStruct struct {
	request, origin string
	playerInTurn    *player.Player
	players         playerset.Players
	board           *board.Board
}

func NewAuctionWithSide(request, origin string, playerInTurn *player.Player,
	players playerset.Players, board *board.Board) Action {
	return &AuctionWithSideStruct{request, origin,
		playerInTurn, players, board}
}
func (as AuctionWithSideStruct) Phase() game.Phase { return game.InsideAuction }
func (as AuctionWithSideStruct) Find(p *player.Player) bool {
	return p.IsExpectedPlayer(as.playerInTurn, as.origin)
}
func (as AuctionWithSideStruct) Do(p *player.Player) error {
	data := strings.Split(as.request, "#")
	score := data[1]
	auction.CheckAndUpdate(score, p.Folded, p.Fold, as.board.AuctionScore, as.board.SetAuctionScore)
	return nil
}
func (as AuctionWithSideStruct) NextPlayer(playerInTurn uint8) uint8 {
	winnerIndex := playersRoundRobin(playerInTurn)
	for as.NextPhasePlayerInfo(as.players[winnerIndex]) {
		winnerIndex = playersRoundRobin(winnerIndex)
	}
	return winnerIndex
}
func (as AuctionWithSideStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) game.Phase {
	if players.Count(predicate.NextPhasePlayerInfo) == 4 {
		return game.ExchangingCards
	}
	return game.InsideAuction
}
func (as AuctionWithSideStruct) NextPhasePlayerInfo(p *player.Player) bool { return p.Folded() }
