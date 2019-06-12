package game

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
)

type req struct {
	action       string
	origin       string
	data1, data2 string
}

func newReq(request, origin string) *req {
	var rq req
	fields := strings.Split(request, "#")
	if len(fields) > 0 {
		rq.action = fields[0]
	}
	if len(fields) > 1 {
		rq.data1 = fields[1]
	}
	if len(fields) > 2 {
		rq.data2 = fields[2]
	}
	rq.origin = origin
	return &rq
}

func (r *req) Action() string {
	return r.action
}

func (r *req) From() string {
	return r.origin
}

func (r *req) Card() (card.ID, error) {
	return card.Create(r.data1, r.data2)
}

func (r *req) Value() string {
	return r.data1
}

func (r *req) EndExchange() bool {
	return r.data1 == "0"
}
