package game

import (
	"strings"
)

func command(request string) string {
	fields := strings.Split(request, "#")
	if fields[0] != "" {
		return fields[0]
	}
	return ""
}

func value(request string) string {
	fields := strings.Split(request, "#")
	if len(fields) > 1 {
		return fields[1]
	}
	return ""
}
