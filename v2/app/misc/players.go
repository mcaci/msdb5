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
