package action

import (
	"errors"
	"strconv"

	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

// Play func
func Play(g gamePlayer) error {
	var err error
	switch g.Phase() {
	case phase.Joining:
		g.CurrentPlayer().RegisterAs(g.Value())
	case phase.InsideAuction:
		score, err := strconv.Atoi(g.Value())
		toFold := player.Folded(g.CurrentPlayer()) || err != nil || !auction.CheckScores(*g.AuctionScore(), auction.Score(score))
		if toFold {
			g.CurrentPlayer().Fold()
		}
		newScore := auction.Update(*g.AuctionScore(), auction.Score(score))
		g.SetAuction(newScore)
		if len(*g.SideDeck()) > 0 {
			quantity := uint8(newScore/90 + newScore/100 + newScore/110 + newScore/120 + newScore/120)
			g.SetShowSide(quantity)
		}
		if newScore >= 120 {
			for _, p := range g.Players() {
				if p == g.CurrentPlayer() {
					continue
				}
				p.Fold()
			}
		}
		notFolded := func(p *player.Player) bool { return !player.Folded(p) }
		if team.Count(g.Players(), notFolded) == 1 {
			g.SetCaller(notFolded)
		}
	case phase.ExchangingCards:
		if g.Value() == "0" {
			return nil
		}
		c, err := g.Card()
		if err != nil {
			return err
		}
		idx, err := g.Players().Index(player.IsCardInHand(*c))
		if err != nil {
			return err
		}
		pl := g.Players().At(idx)
		cards := pl.Hand()
		index := cards.Find(*c)
		toCards := g.SideDeck()
		awayCard := (*cards)[index]
		(*cards)[index] = (*toCards)[0]
		*toCards = append((*toCards)[1:], awayCard)
		return nil
	case phase.ChoosingCompanion:
		if g.Value() == "0" {
			return errors.New("Value 0 for card allowed only for ExchangingCard phase")
		}
		c, err := g.Card()
		if err != nil {
			return err
		}
		idx, err := g.Players().Index(player.IsCardInHand(*c))
		if err != nil {
			return err
		}
		pl := g.Players().At(idx)
		g.SetBriscola(c)
		g.SetCompanion(pl)
		return nil
	case phase.PlayingCards:
		if g.Value() == "0" {
			return errors.New("Value 0 for card allowed only for ExchangingCard phase")
		}
		c, err := g.Card()
		if err != nil {
			return err
		}
		idx, err := g.Players().Index(player.IsCardInHand(*c))
		if err != nil {
			return err
		}
		pl := g.Players().At(idx)
		cards := pl.Hand()
		index := cards.Find(*c)
		g.PlayedCards().Add((*cards)[index])
		*cards = append((*cards)[:index], (*cards)[index+1:]...)
		return nil
	}
	return err
}
