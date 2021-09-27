package briscola

// Players is a slice of Players
type Players []*Player

// NewPlayers creates new container for players
func NewPlayers(nPlayers int) *Players {
	players := make(Players, nPlayers)
	for i := range players {
		players[i] = NewPlayer("")
	}
	return &players
}
