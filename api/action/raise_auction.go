package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type AuctionStruct struct {
	request, origin string
	playerInTurn    *player.Player
	players         playerset.Players
	board           *board.Board
}

func NewAuction(request, origin string, playerInTurn *player.Player,
	players playerset.Players, board *board.Board) Action {
	return &AuctionStruct{request, origin,
		playerInTurn, players, board}
}
func (as AuctionStruct) Phase() game.Phase { return game.InsideAuction }
func (as AuctionStruct) Find(p *player.Player) bool {
	return p.IsExpectedPlayer(as.playerInTurn, as.origin)
}
func (as AuctionStruct) Do(p *player.Player) error {
	data := strings.Split(as.request, "#")
	score := data[1]
	auction.CheckAndUpdate(score, p.Folded, p.Fold, as.board.AuctionScore, as.board.SetAuctionScore)
	return nil
}
func (as AuctionStruct) NextPlayer(playerInTurn uint8) uint8 {
	winnerIndex := playersRoundRobin(playerInTurn)
	for as.NextPhasePlayerInfo(as.players[winnerIndex]) {
		winnerIndex = playersRoundRobin(winnerIndex)
	}
	return winnerIndex
}
func (as AuctionStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) bool {
	return players.Count(predicate.NextPhasePlayerInfo) == 4
}
func (as AuctionStruct) NextPhasePlayerInfo(p *player.Player) bool { return p.Folded() }
