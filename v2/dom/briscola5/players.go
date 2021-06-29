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
	pls      []*player.B5Player
	cal, cmp int
}

// NewPlayers creates new container for briscola5 players
func NewPlayers() *Players {
	players := Players{pls: make([]*player.B5Player, 5)}
	for i := range players.pls {
		players.pls[i] = player.New(&player.Options{For5P: true}).(*player.B5Player)
	}
	return &players
}

func ToGeneralPlayers(bp5 Players) team.Players {
	pls := make(team.Players, 0)
	for i := range bp5.pls {
		pls = append(pls, &bp5.pls[i].Player)
	}
	return pls
}

// Add adds a player to the team
func (playerSet *Players) Add(p *player.B5Player) {
	playerSet.pls = append(playerSet.pls, p)
}

// ErrPlayerNotFound error
var ErrPlayerNotFound = errors.New("Player not found")

// Index returns the index of the player that satisfies the predicate
// or an error if not found
func (playerSet Players) Index(predicate player.B5Predicate) (uint8, error) {
	for i, p := range playerSet.pls {
		if predicate(p) {
			return uint8(i), nil
		}
	}
	return 0, ErrPlayerNotFound
}

// MustIndex works as Index but exits fatally in case of error
func (playerSet Players) MustIndex(predicate player.B5Predicate) uint8 {
	i, err := playerSet.Index(predicate)
	if err != nil {
		log.Fatalf("error found: %v", err)
	}
	return i
}

// Count counts the number of players satisfying the predicate
func Count(players Players, predicate player.B5Predicate) (count uint8) {
	for _, p := range players.pls {
		if predicate(p) {
			count++
		}
	}
	return
}

// Part partition players in two groups according to a predicate
func (playerSet Players) Part(predicate player.B5Predicate) (t1, t2 Players) {
	for _, p := range playerSet.pls {
		if predicate(p) {
			t1.Add(p)
			continue
		}
		t2.Add(p)
	}
	return
}

func (playerSet *Players) Registration() func(string) error {
	var i int
	return func(s string) error {
		if i >= 5 {
			return errors.New("noop: max players reached")
		}
		log.Printf("registering player %d with name %q", i, s)
		playerSet.At(i).RegisterAs(s)
		i++
		return nil
	}
}

func (playerSet *Players) List() []*player.B5Player    { return playerSet.pls }
func (playerSet *Players) At(i int) *player.B5Player   { return playerSet.pls[i] }
func (playerSet *Players) Player(i int) *player.Player { return &playerSet.At(i).Player }
func (playerSet *Players) Caller() *player.Player      { return playerSet.Player(playerSet.cal) }
func (playerSet *Players) Companion() *player.Player   { return playerSet.Player(playerSet.cmp) }
func (playerSet *Players) SetCaller(i uint8)           { playerSet.cal = int(i) }
func (playerSet *Players) SetCompanion(i uint8)        { playerSet.cmp = int(i) }

func (playerSet Players) String() string {
	return fmt.Sprintf("Players: %v, caller's index: %d, companion's index: %d", playerSet.pls, playerSet.cal, playerSet.cmp)
}
