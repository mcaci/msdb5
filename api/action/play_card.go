package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type PlayCardStruct struct {
	request, origin   string
	playerInTurnIndex uint8
	playerInTurn      *player.Player
	players           playerset.Players
	board             *board.Board
	briscolaSeed      card.Seed
}

func NewPlay(request, origin string, playerInTurnIndex uint8,
	playerInTurn *player.Player, players playerset.Players,
	board *board.Board, briscolaSeed card.Seed) Action {
	return &PlayCardStruct{request, origin, playerInTurnIndex,
		playerInTurn, players, board, briscolaSeed}
}

func (pcs PlayCardStruct) Phase() game.Phase { return game.PlayingCards }
func (pcs PlayCardStruct) Find(p *player.Player) bool {
	return p.IsExpectedPlayer(pcs.playerInTurn, pcs.origin)
}
func (pcs PlayCardStruct) Do(p *player.Player) error {
	data := strings.Split(pcs.request, "#")
	number := data[1]
	seed := data[2]
	c, err := card.Create(number, seed)
	p.Play(c)
	pcs.board.PlayedCards().Add(c)
	roundHasEnded := len(*pcs.board.PlayedCards()) == 5
	if roundHasEnded {
		roundWinner := briscola.IndexOfWinningCard(*pcs.board.PlayedCards(), pcs.briscolaSeed)
		roundWinnerIndex := (pcs.playerInTurnIndex + roundWinner + 1) % 5
		pcs.players[roundWinnerIndex].Collect(pcs.board.PlayedCards())
	}
	return err
}
func (pcs PlayCardStruct) NextPlayer(playerInTurn uint8) uint8 {
	next := nextPlayerInTurn(playerInTurn)
	roundHasEnded := len(*pcs.board.PlayedCards()) == 5
	if roundHasEnded {
		roundWinner := briscola.IndexOfWinningCard(*pcs.board.PlayedCards(), pcs.briscolaSeed)
		next = (pcs.playerInTurnIndex + roundWinner + 1) % 5
		pcs.board.PlayedCards().Clear()
	}
	return next
}
func (pcs PlayCardStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) bool {
	return players.All(predicate.NextPhasePlayerInfo)
}
func (pcs PlayCardStruct) NextPhasePlayerInfo(p *player.Player) bool { return p.IsHandEmpty() }
