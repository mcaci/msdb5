package orchestrator

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func (g *Game) play(request, origin string) (all []display.Info, me []display.Info, err error) {
	_, err = cardAction(request)
	if err != nil {
		return
	}
	playerInTurn := g.playerInTurn
	roundMayEnd := len(*g.info.PlayedCards()) >= 4
	if roundMayEnd {
		info := g.playEndRoundData(request, origin)
		err = g.playPhase(info)
	} else {
		info := g.playData(request, origin)
		err = g.playPhase(info)
	}
	if g.phase == end {
		return g.endGame()
	}
	return g.Info(), g.players[playerInTurn].Info(), err
}

func (g *Game) playData(request, origin string) dataPhase {
	c, _ := cardAction(request)
	phase := playBriscola
	find := func(p *player.Player) bool { return isExpectedPlayer(p, g, origin) }
	do := func(p *player.Player) (err error) {
		p.Play(c)
		g.info.PlayedCards().Add(c)
		return
	}
	nextPlayerOperator := nextPlayer
	nextPhasePredicate := g.endGameCondition
	playerPredicate := func(p *player.Player) bool { return p.IsHandEmpty() }
	return dataPhase{phase, find, do, nextPlayerOperator, nextPhasePredicate, playerPredicate}
}

func (g *Game) playEndRoundData(request, origin string) dataPhase {
	c, _ := cardAction(request)
	phase := playBriscola
	find := func(p *player.Player) bool { return isExpectedPlayer(p, g, origin) }
	do := func(p *player.Player) (err error) {
		p.Play(c)
		g.info.PlayedCards().Add(c)
		roundWinnerIndex := roundWinner(g)
		g.players[roundWinnerIndex].Collect(g.info.PlayedCards())
		return
	}
	nextPlayerOperator := func(uint8) uint8 {
		roundWinnerIndex := roundWinner(g)
		g.info.PlayedCards().Clear()
		return roundWinnerIndex
	}
	nextPhasePredicate := g.endGameCondition
	playerPredicate := func(p *player.Player) bool { return p.IsHandEmpty() }
	return dataPhase{phase, find, do, nextPlayerOperator, nextPhasePredicate, playerPredicate}
}

func roundWinner(g *Game) uint8 {
	return (g.playerInTurn + briscola.IndexOfWinningCard(*g.info.PlayedCards(), g.companion.Card().Seed()) + 1) % 5
}

func (g *Game) endGameCondition(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return players.All(searchCriteria)
}

func (g *Game) endGame() ([]display.Info, []display.Info, error) {
	caller, _ := g.players.Find(func(p *player.Player) bool { return p.NotFolded() })
	score1 := caller.Count() + g.companion.Ref().Count()
	score2 := uint8(0)
	for _, pl := range g.players {
		if pl != caller && pl != g.companion.Ref() {
			score2 += pl.Count()
		}
	}
	score1info := display.NewInfo("Callers", ":", strconv.Itoa(int(score1)), ";")
	score2info := display.NewInfo("Others", ":", strconv.Itoa(int(score2)), ";")
	return display.Wrap("Final Score", score1info, score2info), nil, nil
}
