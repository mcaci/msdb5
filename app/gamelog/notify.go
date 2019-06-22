package gamelog

import (
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type selfInformer interface {
	CurrentPlayer() *player.Player
	LastPlayer() *player.Player
	Phase() phase.ID
	SideDeck() *deck.Cards
}

// ToCurrent func
func ToCurrent(gameInfo selfInformer) string {
	return createInGameMsg(gameInfo, gameInfo.CurrentPlayer())
}

// ToLast func
func ToLast(gameInfo selfInformer) string {
	return createInGameMsg(gameInfo, gameInfo.LastPlayer())
}

type sidedeckInformer interface {
	SideDeck() *deck.Cards
}

// SideDeckContent func
func SideDeckContent(gameInfo sidedeckInformer, quantity uint8) string {
	return createSideGameMsg(gameInfo, quantity)
}

// GameInfoMsg func
func GameInfoMsg(gameInfo sidedeckInformer) string {
	printer := message.NewPrinter(language.English)
	return printer.Sprintf("Game: %+v", gameInfo)
}

// NotifyAnticipatedEnding func
func NotifyAnticipatedEnding(team string) string {
	printer := message.NewPrinter(language.English)
	return printer.Sprintf("The end - %s team has all briscola, ending game", team)
}

// NotifyScore func
func NotifyScore(scoreTeam1, scoreTeam2 uint8) string {
	printer := message.NewPrinter(language.English)
	return printer.Sprintf("The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
}
