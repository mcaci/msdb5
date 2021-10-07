package srvb

import "github.com/mcaci/msdb5/v3/briscola"

var g *briscola.Game

type Game struct{ *briscola.Game }

func (g *Game) Cleanup() { g.Game = nil }
