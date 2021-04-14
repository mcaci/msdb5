package srv

import (
	"html/template"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/game"
)

var (
	files     = []string{"assets/home.html", "assets/start.html"}
	templates = template.Must(template.ParseFiles(files...))

	n        string
	g        *game.Game
	register func(string)

	cards       = set.Deck()
	currentBody = [][]byte{}
)
