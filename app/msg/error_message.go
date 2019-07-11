package msg

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Error func
func Error(text string, lang language.Tag) error {
	printer := message.NewPrinter(lang)
	msg := printer.Sprintf(text)
	return fmt.Errorf(msg)
}
