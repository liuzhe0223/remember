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

//waiting for test
func (rl *Rlist) Get() (resList []string) {
	resList = make([]string, 0, rl.Len())
	for el := rl.Front(); el != nil; el = el.Next() {
		resList = append(resList, el.Value.(string))
	}
	return
}

//return ok
func (rl *Rlist) Rpush(v string) bool {
	rl.PushBack(v)
	return true
}

func (rl *Rlist) Lpush(v string) bool {
	return rl.PushFront(v) != nil
}

func (rl *Rlist) Rpop() (v string) {
	el := rl.Back()
	v, _ = el.Value.(string)
	rl.Remove(el)
	return
}

func (rl *Rlist) Lpop() (v string) {
	el := rl.Front()
	v, _ = el.Value.(string)
	rl.Remove(el)
	return
}

func (rl *Rlist) Lrange(start, end int) (values []string) {
	if start+1 > rl.Len() || start > end {
		values = make([]string, 0)
		return
	}

	stop := 0
	if end+1 <= rl.Len() {
		stop = end
	} else {
		stop = rl.Len() - 1
	}

	values = make([]string, 0, stop-start+1)

	el := rl.Front()
	for i := 0; i <= stop; i++ {
		if i >= start {
			value := el.Value.(string)
			values = append(values, value)
		}
		el = el.Next()
	}

	return
}
