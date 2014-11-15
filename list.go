package remember

import (
	"encoding/binary"
)

func (d *db) lEncodeMetakey(key []byte) []byte {
	bufLength := len(key) + 1
	buf := make([]byte, bufLength, bufLength)
	buf[0] = byte(listType)
	copy(buf[1:], key)

	return buf
}

func (d *db) lEncodeKey(key []byte, seq int32) []byte {
	metaKey := d.lEncodeMetakey(key)

	bufLength := len(metaKey) + 4
	buf := make([]byte, bufLength, bufLength)

	copy(buf, metaKey)

	binary.BigEndian.PutUint32(buf[len(metaKey):], uint32(seq))

	return buf
}
