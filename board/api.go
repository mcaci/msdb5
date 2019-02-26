package board

import (
	"log"
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api"
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Action interface
func (b *Board) Action(request, origin string) (api.Info, api.Info, error) {
	data := strings.Split(string(request), "#")
	var err error
	switch data[0] {
	case "Join":
		err = b.Join(data[1], origin)
	case "Auction":
		err = b.RaiseAuction(data[1], origin)
	case "Companion":
		_, err = b.Nominate(data[1], data[2], origin)
	case "Card":
		err = b.Play(data[1], data[2], origin)
	}
	pInfo, err := b.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
	return b, pInfo, err
}

// RaiseAuction func
func (b *Board) RaiseAuction(score, origin string) error {
	p, err := b.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
	if err == nil {
		prevScore := b.AuctionScore()
		currentScore, err := strconv.Atoi(score)
		if err != nil {
			log.Printf("Error was raised during auction: %v\n", err)
		}
		updateAuction(0, prevScore, uint8(currentScore), p.SetAuctionScore)
		updateAuction(prevScore, prevScore, uint8(currentScore), b.SetAuctionScore)
	}
	return err
}

func updateAuction(baseScore, prevScore, currentScore uint8, set func(uint8)) {
	const minScore = 61
	const maxScore = 120
	actualScore := currentScore
	if currentScore < prevScore {
		actualScore = baseScore
	} else if currentScore < minScore {
		actualScore = minScore
	} else if currentScore > maxScore {
		actualScore = maxScore
	}
	set(actualScore)
}

// Play func
func (b *Board) Play(number, seed, origin string) error {
	p, err := b.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
	if err == nil {
		c, _ := p.Play(number, seed)
		// c, err := p.Play(number, seed)
		// if err == nil { // TODO: FOR SOME CHECKS IT'S TRUE
		b.PlayedCards().Add(c)
		if len(*b.PlayedCards()) >= 5 {
			playerIndex := briscola.IndexOfWinningCard(*b.PlayedCards(), card.Coin)
			b.PlayedCards().Move(b.Players()[playerIndex].Pile())
		}
		// }
	}
	return err
}

// Nominate func
func (b *Board) Nominate(number, seed, origin string) (card.ID, error) {
	card, err := card.Create(number, seed)
	if err == nil {
		p, err := b.Players().Find(func(p *player.Player) bool { return p.Has(card) })
		if err == nil {
			b.selectedCard = card
			b.selectedPlayer = *p
		}
	}
	return card, err
}

// Join func
func (b *Board) Join(name, origin string) error {
	p, err := b.Players().Find(func(p *player.Player) bool { return p.Name() == "" })
	if err == nil {
		p.SetName(name)
		p.MyHostIs(origin)
	} else {
		log.Println("All players have joined, no further players are expected: " + err.Error())
		log.Println(b.Players())
	}
	return err
}
