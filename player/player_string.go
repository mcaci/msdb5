package player

func (player Player) String() string {
	str := "Player["
	str += "Name:" + player.name + ";"
	str += "Host:" + player.host + ";"
	str += "Hand:" + player.hand.String() + ";"
	str += "Pile:" + player.pile.String() + ";"
	str += "]"
	return str
}
