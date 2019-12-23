package input

import (
	"strings"
)

type act uint8

const (
	Com act = iota
	Val
)

func Parse(request string, a act) (parsed string) {
	fields := strings.Split(request, "#")
	if len(fields) > int(a) {
		parsed = fields[a]
	}
	return
}
