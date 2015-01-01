package remember

import (
	"encoding/binary"
	"errors"
)

const (
	currentHeadPos int32 = 1
	currentTailPos int32 = 2

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

func (r *Remember) lEncodeKey(key []byte, pos int32) []byte {
	metaKey := r.lEncodeMetakey(key)

	bufLength := len(metaKey) + 4
	buf := make([]byte, bufLength, bufLength)

	copy(buf, metaKey)

	binary.BigEndian.PutUint32(buf[len(metaKey):], uint32(pos))

	return buf
}

func (r *Remember) lGetMeta(metaKey []byte) (headPos, tailPos, size int32, err error) {
	value, err := r.Store.Get(metaKey)

	if err != nil && err != ErrNotFound {
		return
	}

	if err == ErrNotFound {
		err = nil
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

func (r *Remember) Lpush(key, elVlaue []byte) error {
	return r.push(key, elVlaue, currentHeadPos)
}

func (r *Remember) Rpush(key, elVlaue []byte) error {
	return r.push(key, elVlaue, currentTailPos)
}

func (r *Remember) Lpop(key []byte) ([]byte, error) {
	return r.pop(key, currentHeadPos)
}

func (r *Remember) Rpop(key []byte) ([]byte, error) {
	return r.pop(key, currentTailPos)
}

func (r *Remember) pop(key []byte, where int32) (value []byte, err error) {
	var (
		encodeKey []byte
		pos       int32
	)

	metaKey := r.lEncodeMetakey(key)
	headPos, tailPos, _, err := r.lGetMeta(metaKey)
	if err != nil {
		return
	}

	if where == currentHeadPos {
		pos = headPos
		headPos += 1
	} else if where == currentTailPos {
		pos = tailPos
		tailPos -= 1
	} else {
		err = errors.New("Invalid pop position")
	}

	encodeKey = r.lEncodeKey(key, pos)
	value, err = r.Store.Get(encodeKey)
	if err != nil {
		return
	}

	err = r.Store.Delete(encodeKey)
	if err != nil {
		return
	}

	_, err = r.lSetMeta(metaKey, headPos, tailPos)
	return
}

func (r *Remember) push(key, elVlaue []byte, where int32) (err error) {
	var (
		encodeKey []byte
		pos       int32
	)

	metaKey := r.lEncodeMetakey(key)
	headPos, tailPos, size, err := r.lGetMeta(metaKey)
	if err != nil {
		return
	}

	if size == 0 {
		pos = headPos
	} else if where == currentHeadPos {
		pos = headPos - 1
		headPos = pos
	} else if where == currentTailPos {
		pos = tailPos + 1
		tailPos = pos
	} else {
		err = errors.New("Invalid push position")
	}

	encodeKey = r.lEncodeKey(key, pos)
	err = r.Store.Put(encodeKey, elVlaue)
	if err != nil {
		return
	}

	_, err = r.lSetMeta(metaKey, headPos, tailPos)
	return
}
