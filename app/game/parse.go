package game

import (
	"strings"
)

type act uint8

const (
	com act = iota
	val
)

func parse(request string, a act) (parsed string) {
	fields := strings.Split(request, "#")
	if len(fields) > int(a) {
		parsed = fields[a]
	}
	return
}
