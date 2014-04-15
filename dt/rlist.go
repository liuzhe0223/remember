package dt

import (
	"container/list"
)

type Rlist list.List

func NewRlist() {
	return Rlist.New()
}

func (rl *Rlist) Rpush(robj Robj) []Robj {
	el = Rlist.PushBack(robj)
	return el.Value
}
