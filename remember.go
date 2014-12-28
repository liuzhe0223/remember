package remember

import (
	"sync"
)

var (
	StoreErrorNotFound error
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
	if !isInitValuesSet() {
		panic("all init vars need to set first")
	}

	return &Remember{
		Store: store,
	}
}

func isInitValuesSet() (res bool) {
	if StoreErrorNotFound.Error() == "" {
		return
	}

	return true
}

func (r *Remember) Close() (err error) {
	err = r.Store.Close()
	return
}
