package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type PlayWithSideCardStruct struct {
	request, origin string
	playerInTurn    *player.Player
	players         playerset.Players
	board           *board.Board
	briscolaSeed    card.Seed
}

func NewPlayWithSide(request, origin string, playerInTurn *player.Player,
	players playerset.Players, board *board.Board, briscolaSeed card.Seed) Action {
	return &PlayWithSideCardStruct{request, origin, playerInTurn,
		players, board, briscolaSeed}
}

func (pcs PlayWithSideCardStruct) Phase() game.Phase { return game.PlayingCards }
func (pcs PlayWithSideCardStruct) Find(p *player.Player) bool {
	return p.IsExpectedPlayer(pcs.playerInTurn, pcs.origin)
}
func (pcs PlayWithSideCardStruct) Do(p *player.Player) error {
	data := strings.Split(pcs.request, "#")
	number := data[1]
	seed := data[2]
	c, err := card.Create(number, seed)
	err = p.Play(c)
	if err != nil {
		return err
	}
	pcs.board.PlayedCards().Add(c)
	roundHasEnded := len(*pcs.board.PlayedCards()) == 5
	if roundHasEnded {
		index, _ := pcs.players.FindIndex(func(pl *player.Player) bool { return pl == p })
		next := roundWinnerIndex(uint8(index), *pcs.board.PlayedCards(), pcs.briscolaSeed)
		pcs.players[next].Collect(pcs.board.PlayedCards())
	}
	return err
}
func (pcs PlayWithSideCardStruct) NextPlayer(playerInTurn uint8) uint8 {
	next := playersRoundRobin(playerInTurn)
	roundHasEnded := len(*pcs.board.PlayedCards()) == 5
	if roundHasEnded {
		next = roundWinnerIndex(playerInTurn, *pcs.board.PlayedCards(), pcs.briscolaSeed)
		pcs.board.PlayedCards().Clear()
	}
	return next
}
func (pcs PlayWithSideCardStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) game.Phase {
	if players.All(predicate.NextPhasePlayerInfo) {
		return game.End
	}
	return game.PlayingCards
}
func (pcs PlayWithSideCardStruct) NextPhasePlayerInfo(p *player.Player) bool { return p.IsHandEmpty() }
