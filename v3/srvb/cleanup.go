package srvb

const CleanupURL = "/cln"

func (g *Game) Cleanup() { g.Game = nil }
