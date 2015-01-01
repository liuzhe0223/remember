package remember

import (
	"sync"
)

type Remember struct {
	Store KV
	sync.Mutex
}

type KV interface {
	Get([]byte) ([]byte, error)
	Put([]byte, []byte) error
	Delete([]byte) error
	Close() error
}

func NewRemember(store KV) *Remember {
	return &Remember{
		Store: store,
	}
}

func (r *Remember) Close() (err error) {
	err = r.Store.Close()
	return
}
