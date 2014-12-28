package remember

func (r *Remember) KVEncodeKey(key []byte) (encodeKey []byte) {
	encodeKeyLen := len(key) + 1
	encodeKey = make([]byte, encodeKeyLen, encodeKeyLen)

	encodeKey[0] = byte(kvType)
	copy(encodeKey[1:], key)
	return
}

func (r *Remember) Get(key []byte) (value []byte, err error) {
	value, err = r.Store.Get(r.KVEncodeKey(key))
	return
}

func (r *Remember) Put(key, value []byte) (err error) {
	err = r.Store.Put(r.KVEncodeKey(key), value)
	return
}

func (r *Remember) Delete(key []byte) (err error) {
	err = r.Store.Delete(r.KVEncodeKey(key))
	return
}
