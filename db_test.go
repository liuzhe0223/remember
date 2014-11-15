package remember

import (
	"testing"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/stretchr/testify/assert"
)

func TestDbInit(t *testing.T) {
	assert.NotNil(t, Db)
	assert.NotNil(t, Db.DB)
}

func TestDbOriginOPs(t *testing.T) {
	key := []byte("testkey")
	value := []byte("testvalue")

	err := Db.Put(key, value)
	assert.NoError(t, err)

	valueGet, err := Db.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, value, valueGet)

	err = Db.Delete(key)
	assert.NoError(t, err)

	_, err = Db.Get(key)
	assert.Equal(t, err, leveldb.ErrNotFound)
}
