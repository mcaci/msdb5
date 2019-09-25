package pl

import (
	"golang.org/x/text/message"
)

var sideDeckRef = func(p *message.Printer) string { return p.Sprintf("Side deck") }
var undefinedCard = func(p *message.Printer) string { return p.Sprintf("(Undefined card)") }

var seeds = func(p *message.Printer, idx uint8) string {
	return []string{p.Sprintf("Coin"), p.Sprintf("Cup"), p.Sprintf("Sword"), p.Sprintf("Cudgel")}[idx]
}

var errRef = func(p *message.Printer) string { return p.Sprintf("Error") }
var errMsgRef = func(p *message.Printer, exp, act string) string {
	return p.Sprintf("Expecting %s, but found %s", exp, act)
}

var phases = func(p *message.Printer, idx uint8) string {
	return []string{p.Sprintf("Join"), p.Sprintf("Auction"), p.Sprintf("Exchange"),
		p.Sprintf("Companion"), p.Sprintf("Card"), p.Sprintf("End")}[idx]
}

var gameRef = func(p *message.Printer) string { return p.Sprintf("Game") }
var gameElemRef = func(p *message.Printer, idx uint8) string {
	return []string{p.Sprintf("Turn of"), p.Sprintf("Phase"), p.Sprintf("Companion"),
		p.Sprintf("Played cards"), p.Sprintf("Last card"), p.Sprintf("Auction score")}[idx]
}

var teams = func(p *message.Printer, idx uint8) string {
	return []string{p.Sprintf("Callers"), p.Sprintf("Others")}[idx]
}

var endRef = func(p *message.Printer) string { return p.Sprintf("The end") }
var allBriscolaRef = func(p *message.Printer) string { return p.Sprintf("have all briscola cards") }

var plRef = func(p *message.Printer) string { return p.Sprintf("Player") }
var plElemRef = func(p *message.Printer, idx uint8) string {
	return []string{p.Sprintf("Name"), p.Sprintf("Cards"), p.Sprintf("Pile"),
		p.Sprintf("Played cards"), p.Sprintf("Has folded")}[idx]
}

var yesNoRef = func(p *message.Printer, idx uint8) string {
	return []string{p.Sprintf("Yes"), p.Sprintf("No")}[idx]
}
