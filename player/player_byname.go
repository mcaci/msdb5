package player

// ByName func
func ByName(name string, players []*Player) *Player {
	var p1 *Player
	for _, p := range players {
		if p.Name() == name {
			p1 = p
			break
		}
	}
	return p1
}
