package remember

import (
	"encoding/binary"
	"errors"
)

const (
	listHeadPos int32 = 1
	listTailPos int32 = 2

	listMinSeq     int32 = 1000
	listMaxSeq     int32 = 1<<31 - 1000
	listInitialPos int32 = listMinSeq + (listMaxSeq-listMinSeq)/2
)

func (r *Remember) lEncodeMetakey(key []byte) []byte {
	bufLength := len(key) + 1
	buf := make([]byte, bufLength, bufLength)
	buf[0] = byte(listType)
	copy(buf[1:], key)

	return buf
}

func (r *Remember) lEncodeKey(key []byte, seq int32) []byte {
	metaKey := r.lEncodeMetakey(key)

	bufLength := len(metaKey) + 4
	buf := make([]byte, bufLength, bufLength)

	copy(buf, metaKey)

	binary.BigEndian.PutUint32(buf[len(metaKey):], uint32(seq))

	return buf
}

func (r *Remember) lGetMeta(metaKey []byte) (headPos, tailPos, size int32, err error) {
	value, err := r.Store.Get(metaKey)
	if err != nil {
		return
	}

	if value == nil {
		headPos = listInitialPos
		tailPos = listInitialPos
		size = 0
		return
	}

	if len(value) != 8 {
		err = errors.New("Invalid list meta value")
		return
	}

	headPos = int32(binary.BigEndian.Uint32(value[0:4]))
	tailPos = int32(binary.BigEndian.Uint32(value[4:8]))
	size = tailPos - headPos + 1

	return
}

func (r *Remember) lSetMeta(metaKey []byte, headPos, tailPos int32) (size int32, err error) {
	size = tailPos - headPos + 1

	if size < 0 {
		err = errors.New("Invalid list size")
		panic(err)
		return
	}

	if size == 0 {
		err = r.Store.Delete(metaKey)
		return
	}

	metaValue := make([]byte, 8, 8)
	binary.BigEndian.PutUint32(metaValue[0:4], uint32(headPos))
	binary.BigEndian.PutUint32(metaValue[4:8], uint32(tailPos))

	err = r.Store.Put(metaKey, metaValue)
	return
}
