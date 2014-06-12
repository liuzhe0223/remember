package pst

import (
	"github.com/liuzhe0223/remember/dt"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
)

func (p *Pster) Dumps() (err error) {
	db, err := leveldb.OpenFile("/home/liuzhe/db", nil)
	if err != nil {
		return
	}
	defer db.Close()

	for k, v := range *p {

		//string store k,v
		if rString, ok := v.(*dt.Rstring); ok {
			db.Put([]byte(k), []byte(*v), nil)

			//rset store k~nu el
		} else if rSet, ok := v.(*dt.Rset); ok {
			db.Put([]byte(k+"~"), []byte("rset"), nil)
			//the nu
			i := 1
			for el, _ := range *rSet {
				salt := strconv.FormatInt(i, 10)
				db.Put([]byte(k+"~"+i), []byte(v), nil)
				i += 1
			}
			db.Put([]byte(k+"~end"), []byte(""), nil)

			//store rMap k~el_k v
		} else if rmap, ok := v.(*dt.Rmap); ok {
			db.Put([]byte(k+"~"), []byte("rmap"), nil)
			for elK, elV := range *rmap {
				db.Put([]byte(k+"~"+elK), []byte(elV), nil)
			}
			db.Put([]byte(k+"~end"), []byte(""), nil)

			//store rList  k~nu v
		} else if rList, ok := v.(*dt.Rlist); ok {
			db.Put([]byte(k+"~"), []byte("rlist"), nil)
			i := 1
			for e := rList.Front(); e != nil; e = e.Next() {
				salt := strconv.FormatInt(i, 10)
				db.Put([]byte(k+"~"+i), []byte(e.Value.(string)), nil)
				i += 1
			}
			db.Put([]byte(k+"~end"), []byte(""), nil)
		}
	}
}
