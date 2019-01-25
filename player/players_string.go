package player

func (players Players) String() string {
	var str string
	for _, player := range players {
		str += player.String() + " "
	}
	return str
}
