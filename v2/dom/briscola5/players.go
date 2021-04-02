package briscola5

import (
	"errors"
	"fmt"
	"log"

	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

// Players is a slice of pointers to players
type Players struct {
	pls      []*Player
	cal, cmp int
}

// NewPlayers creates new container for briscola5 players
func NewPlayers() *Players {
	players := Players{pls: make([]*Player, 5)}
	for i := range players.pls {
		players.pls[i] = NewPlayer()
	}
	return &players
}

func ToGeneralPlayers(bp5 Players) team.Players {
	pls := make(team.Players, 5)
	for i := range bp5.pls {
		pls[i] = &bp5.pls[i].Player
	}
	return pls
}

// Add adds a player to the team
func (playerSet *Players) Add(p *Player) {
	playerSet.pls = append(playerSet.pls, p)
}

// ErrPlayerNotFound error
var ErrPlayerNotFound = errors.New("Player not found")

// Index returns the index of the player that satisfies the predicate
// or an error if not found
func (playerSet Players) Index(predicate Predicate) (uint8, error) {
	for i, p := range playerSet.pls {
		if predicate(p) {
			return uint8(i), nil
		}
	}
	return 0, ErrPlayerNotFound
}

// MustIndex works as Index but exits fatally in case of error
func (playerSet Players) MustIndex(predicate Predicate) uint8 {
	i, err := playerSet.Index(predicate)
	if err != nil {
		log.Fatalf("error found: %v", err)
	}
	return i
}

// Count counts the number of players satisfying the predicate
func Count(players Players, predicate Predicate) (count uint8) {
	for _, p := range players.pls {
		if predicate(p) {
			count++
		}
	}
	return
}

func (playerSet *Players) List() []*Player             { return playerSet.pls }
func (playerSet *Players) At(i int) *Player            { return playerSet.pls[i] }
func (playerSet *Players) Player(i int) *player.Player { return &playerSet.At(i).Player }
func (playerSet *Players) Caller() *player.Player      { return playerSet.Player(playerSet.cal) }
func (playerSet *Players) Companion() *player.Player   { return playerSet.Player(playerSet.cmp) }
func (playerSet *Players) SetCaller(i uint8)           { playerSet.cal = int(i) }
func (playerSet *Players) SetCompanion(i uint8)        { playerSet.cmp = int(i) }

func (playerSet Players) String() string {
	return fmt.Sprintf("Players: %v, caller's index: %d, companion's index: %d", playerSet.pls, playerSet.cal, playerSet.cmp)
}