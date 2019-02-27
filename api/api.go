package api

// Action interface
type Action interface {
	Action(request, origin string) (Info, Info, error)
}

// Info interface
type Info interface {
	Print() string
}
