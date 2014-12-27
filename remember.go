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

func (r *Remember) Get(key []byte) (value []byte, err error) {
	value, err = r.Store.Get(key)
	return
}

func (r *Remember) Put(key, value []byte) (err error) {
	err = r.Store.Put(key, value)
	return
}

func (r *Remember) Delete(key []byte) (err error) {
	err = r.Store.Delete(key)
	return
}

func (r *Remember) Close() (err error) {
	err = r.Store.Close()
	return
}
