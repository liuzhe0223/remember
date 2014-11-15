package remember

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLEncodeMetaKey(t *testing.T) {
	key := []byte("test_key")
	metaKey := Db.lEncodeMetakey(key)

	expectMetakey := append([]byte{1}, []byte(key)...)
	assert.Equal(t, expectMetakey, metaKey)
}

func TestLEncodeKey(t *testing.T) {
	key := []byte("test_key")
	seq := int32(5)
	metaKey := Db.lEncodeMetakey(key)

	buf := make([]byte, 4, 4)
	binary.BigEndian.PutUint32(buf, uint32(seq))

	expectListKey := append(metaKey, buf...)
	listKey := Db.lEncodeKey(key, seq)

	assert.Equal(t, expectListKey, listKey)
}
