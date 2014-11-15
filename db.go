package remember

import (
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	Db     *db
	DbPath string
)

func init() {
	DbPath = "testdb"

	leveldbP, err := leveldb.OpenFile(DbPath, nil)
	if err != nil {
		panic(err)
	}

	Db = &db{
		DB: leveldbP,
	}
}

type db struct {
	*leveldb.DB
	sync.Mutex
}

func (d *db) Get(key []byte) (value []byte, err error) {
	value, err = d.DB.Get(key, nil)
	return
}

func (d *db) Put(key, value []byte) (err error) {
	err = d.DB.Put(key, value, nil)
	return
}

func (d *db) Delete(key []byte) (err error) {
	err = d.DB.Delete(key, nil)
	return
}
