package dt

import (
	"container/list"
)

type Rlist struct {
	*list.List
}

func NewRlist() *Rlist {
	return &Rlist{list.New()}
}

//return ok
func (rl *Rlist) Rpush(robj Robj) bool {
	rl.PushBack(robj)
	return true
}

func (rl *Rlist) Lpush(robj Robj) bool {
	return rl.PushFront(robj) != nil
}

func (rl *Rlist) Rpop() (value Robj) {
	el := rl.Back()
	value, _ = el.Value.(Robj)
	return
}

func (rl *Rlist) Lpop() (value Robj) {
	el := rl.Front()
	value, _ = el.Value.(Robj)
	return
}

func (rl *Rlist) Lrange(start, end int) (values []Robj) {
	if start+1 > rl.Len() || start > end {
		values = make([]Robj, 0)
		return
	}

	stop := 0
	if end+1 <= rl.Len() {
		stop = end
	} else {
		stop = rl.Len() - 1
	}

	values = make([]Robj, 0, stop-start+1)

	el := rl.Front()
	for i := 0; i <= stop; i++ {
		if i >= start {
			value := el.Value.(Robj)
			values = append(values, value)
		}
		el = el.Next()
	}

	return
}
