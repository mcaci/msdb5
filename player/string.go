package player

func (player Player) String() string {
	str := "Player["
	str += print("Name", player.name)
	str += print("Host", player.host)
	str += print("Hand", player.hand.String())
	str += print("Pile", player.pile.String())
	str += "]"
	return str
}

func print(info, field string) string {
	return info + ":" + field + ";"
}
