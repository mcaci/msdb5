package input

import (
	"fmt"
	"strings"

	"github.com/mcaci/ita-cards/card"
)

func Command(request string) string {
	fields := strings.Split(request, "#")
	if fields[0] != "" {
		return fields[0]
	}
	return ""
}

func Value(request string) string {
	fields := strings.Split(request, "#")
	if len(fields) > 1 {
		return fields[1]
	}
	return ""
}

func Card(request string) (*card.Item, error) {
	fields := strings.Split(request, "#")
	if len(fields) > 2 {
		return card.New(fields[1], fields[2])
	}
	return nil, fmt.Errorf("not enough data to make a card: %s", request)
}

