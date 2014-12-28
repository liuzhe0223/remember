package remember

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLEncodeMetaKey(t *testing.T) {
	store := NewLeveldb("testdb")
	r := NewRemember(store)
	defer r.Close()

	key := []byte("test_key")
	metaKey := r.lEncodeMetakey(key)

	expectMetakey := append([]byte{1}, []byte(key)...)
	assert.Equal(t, expectMetakey, metaKey)
}

func TestLEncodeKey(t *testing.T) {
	store := NewLeveldb("testdb")
	r := NewRemember(store)
	defer r.Close()

	key := []byte("test_key")
	seq := int32(5)
	metaKey := r.lEncodeMetakey(key)

	buf := make([]byte, 4, 4)
	binary.BigEndian.PutUint32(buf, uint32(seq))

	expectListKey := append(metaKey, buf...)
	listKey := r.lEncodeKey(key, seq)

	assert.Equal(t, expectListKey, listKey)
}

func TestLSetMetaAndLGetMeta(t *testing.T) {
	store := NewLeveldb("testdb")
	r := NewRemember(store)
	defer r.Close()

	lKey := []byte("testListKey")
	metaKey := r.lEncodeMetakey(lKey)

	size, err := r.lSetMeta(metaKey, 500, 600)
	assert.Nil(t, err)
	assert.Equal(t, 600-500+1, size)

	hp, tp, size, err := r.lGetMeta(metaKey)
	assert.Nil(t, err)
	assert.Equal(t, 500, hp)
	assert.Equal(t, 600, tp)
	assert.Equal(t, 600-500+1, size)
}

func TestLpushAndLPop(t *testing.T) {
	store := NewLeveldb("testdb")
	r := NewRemember(store)
	defer r.Close()

	key := []byte("testListLpushAndPopKey")
	value := []byte("testListLpushAndPopValue")

	err := r.Lpush(key, value)
	assert.Nil(t, err)

	testValue, err := r.Lpop(key)
	assert.Nil(t, err)
	assert.Equal(t, value, testValue)
}
