package remember

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Leveldb struct {
	*leveldb.DB
}

func NewLeveldb(DbPath string) *Leveldb {
	db, err := leveldb.OpenFile(DbPath, nil)
	if err != nil {
		panic(err)
	}

	StoreErrorNotFound = leveldb.ErrNotFound

	return &Leveldb{
		DB: db,
	}
}

func (db *Leveldb) Get(key []byte) (value []byte, err error) {
	value, err = db.DB.Get(key, nil)
	return
}

func (db *Leveldb) Put(key, value []byte) (err error) {
	err = db.DB.Put(key, value, nil)
	return
}

func (db *Leveldb) Delete(key []byte) (err error) {
	err = db.DB.Delete(key, nil)
	return
}
