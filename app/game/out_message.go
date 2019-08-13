package game

import "io"

type PlMsg struct {
	dest io.Writer
	msg  string
}

func (m PlMsg) Dest() io.Writer { return m.dest }
func (m PlMsg) Msg() string     { return m.msg }
