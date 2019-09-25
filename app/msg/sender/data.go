package sender

import "github.com/mcaci/msdb5/dom/team"

type Data struct {
	players team.Players
	origin  string
}

func New(origin string, players team.Players) Data {
	return Data{players, origin}
}

func (s Data) From() string          { return s.origin }
func (s Data) Players() team.Players { return s.players }
