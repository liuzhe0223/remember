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
