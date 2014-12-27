package remember

import (
	"testing"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/stretchr/testify/assert"
)

func TestDbOriginOPs(t *testing.T) {
	store := NewLeveldb("testdb")
	r := NewRemember(store)
	defer r.Close()

	key := []byte("testkey")
	value := []byte("testvalue")

	err := r.Put(key, value)
	assert.NoError(t, err)

	valueGet, err := r.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, value, valueGet)

	err = r.Delete(key)
	assert.NoError(t, err)

	_, err = r.Get(key)
	assert.Equal(t, err, leveldb.ErrNotFound)
}
