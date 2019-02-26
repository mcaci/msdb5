package api

import "fmt"

// Action interface
type Action interface {
	Action(request, origin string)
	fmt.Stringer
}

// Info interface
type Info interface {
	Print() string
}
