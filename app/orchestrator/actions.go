package orchestrator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app/phase"

	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

// Play func
func Play(g *Game, p *player.Player, requestname, request, origin string) error {
	switch requestname {
	case "Join":
		name := strings.Split(request, "#")[1]
		p.Join(name, origin)
		return nil
	case "Auction":
		if p.Folded() {
			return nil
		}
		data := strings.Split(request, "#")
		score := data[1]
		currentScore, err := strconv.Atoi(score)
		if err == nil && g.auctionScore.CheckWith(auction.Score(currentScore)) {
			g.auctionScore.Update(auction.Score(currentScore))
		} else {
			p.Fold()
		}
		return nil
	case "Exchange":
		data := strings.Split(request, "#")
		number := data[1]
		if number == "0" {
			return nil
		}
		seed := data[2]
		c, err := card.Create(number, seed)
		if err != nil {
			return err
		}
		if !p.Has(c) {
			return errors.New("card is not in players hand")
		}
		index, err := p.Hand().Find(c)
		if err != nil {
			return err
		}
		p.Hand().Add(g.side[0])
		g.side.Remove(0)
		g.side.Add(c)
		p.Hand().Remove(index)
		return nil
	case "Companion":
		data := strings.Split(request, "#")
		number := data[1]
		seed := data[2]
		c, err := card.Create(number, seed)
		if err != nil {
			return err
		}
		_, pl, err := g.players.Find(func(p *player.Player) bool { return p.Has(c) })
		if err != nil {
			return err
		}
		g.SetCompanion(c, pl)
		return nil
	case "Card":
		data := strings.Split(request, "#")
		number := data[1]
		seed := data[2]
		c, err := card.Create(number, seed)
		err = p.Play(c)
		if err != nil {
			return err
		}
		g.playedCards.Add(c)
		roundHasEnded := len(g.playedCards) == 5
		if roundHasEnded {
			playerInTurn, _, _ := g.players.Find(func(pl *player.Player) bool { return pl == p })
			winningCardIndex := briscola.IndexOfWinningCard(g.playedCards, g.BriscolaSeed())
			var playersRoundRobin = func(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }
			next := playersRoundRobin(uint8(playerInTurn) + winningCardIndex)
			g.players[next].Collect(g.PlayedCards())
			a := make([]player.EmptyHandChecker, 0)
			for _, p := range g.players {
				a = append(a, p)
			}
			if team.Count(g.players, func(p *player.Player) bool { return p.IsHandEmpty() }) == 5 &&
				len(g.side) > 0 {
				g.players[next].Collect(g.SideDeck())
				g.SideDeck().Clear()
			}
		}
		return err
	default:
		return fmt.Errorf("Action %s not valid", requestname)
	}
}

// NextPhase func
func NextPhase(g *Game, request string) phase.ID {
	isSideDeckUsed := len(*g.SideDeck()) > 0
	current, nextPhase := g.phase, g.phase+1
	predicateToNextPhase := func() bool { return true }
	switch current {
	case phase.Joining:
		predicateToNextPhase = func() bool {
			return team.Count(g.players, func(p *player.Player) bool { return p.IsNameEmpty() }) == 0
		}
	case phase.InsideAuction:
		predicateToNextPhase = func() bool {
			return team.Count(g.players, func(p *player.Player) bool { return p.Folded() }) == 4
		}
		if !isSideDeckUsed {
			nextPhase = current + 2
		}
	case phase.ExchangingCards:
		predicateToNextPhase = func() bool {
			data := strings.Split(request, "#")
			if len(data) > 1 {
				number, err := strconv.Atoi(data[1])
				return number == 0 || err != nil
			}
			return false
		}
	case phase.ChosingCompanion:
		nextPhase = phase.PlayingCards
	case phase.PlayingCards:
		predicateToNextPhase = func() bool {
			return team.Count(g.players, func(p *player.Player) bool { return p.IsHandEmpty() }) == 5
		}
	default:
		nextPhase = phase.End
	}
	if predicateToNextPhase() {
		return nextPhase
	}
	return current
}

// NextPlayer func
func NextPlayer(g *Game, current phase.ID, playerInTurn uint8) uint8 {
	var playersRoundRobin = func(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }
	nextPlayer := playersRoundRobin(playerInTurn)
	switch current {
	case phase.Joining: // nothing
	case phase.InsideAuction:
		for g.players[nextPlayer].Folded() {
			nextPlayer = playersRoundRobin(nextPlayer)
		}
	case phase.PlayingCards:
		roundHasEnded := len(g.playedCards) == 5
		if roundHasEnded {
			winningCardIndex := briscola.IndexOfWinningCard(g.playedCards, g.BriscolaSeed())
			nextPlayer = playersRoundRobin(playerInTurn + winningCardIndex)
		}
	default:
		nextPlayer = playerInTurn
	}
	return nextPlayer
}
