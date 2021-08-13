package misc

// Players is a slice of Players
type Players []Player

// NewPlayers creates new container for players
func NewPlayers(nPlayers int) *Players {
	players := make(Players, nPlayers)
	for i := range players {
		players[i] = New(&Options{})
	}
	return &players
}

func (players *Players) Add(p Player) {
	*players = append(*players, p)
}

func (players Players) SelectIndex(prd Predicate) (uint8, error) {
	for i, p := range players {
		if !prd(p) {
			continue
		}
		return uint8(i), nil
	}
	return 0, ErrPlayerNotFound
}

// Part partition players in two groups according to a predicate
func (players Players) Part(predicate Predicate) (t1, t2 Players) {
	for _, p := range players {
		if predicate(p) {
			t1.Add(p)
			continue
		}
		t2.Add(p)
	}
	return
}

// Count counts the number of players satisfying the predicate
func Count(players Players, predicate Predicate) (count uint8) {
	for _, p := range players {
		if predicate(p) {
			count++
		}
	}
	return
}