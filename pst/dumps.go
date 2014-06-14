package pst

import (
	"github.com/liuzhe0223/remember/dt"
	"strconv"
)

func (p *Pster) Dumps() (err error) {
	iter := p.Leveldb.NewIterator(nil, nil)
	for k, v := range *(p.Db) {

		var kTemp string
		for ok := iter.Seek([]byte(k + "~")); ok; ok = iter.Next() {
			p.Leveldb.Delete([]byte(kTemp), nil)
			kTemp = string(iter.Key())

			if string(iter.Key()) == k+"~end" || string(iter.Key()) == k+"~~end" {
				kTemp = string(iter.Key())
				break
			}
		}
		p.Leveldb.Delete([]byte(kTemp), nil)

		//string store k,v
		if rString, ok := v.(*dt.Rstring); ok {
			p.Leveldb.Put([]byte(k), []byte(*rString), nil)

			//rset store k~nu el
		} else if rSet, ok := v.(*dt.Rset); ok {
			p.Leveldb.Put([]byte(k+"~"), []byte("rset"), nil)
			//the nu
			var i int64 = 1
			for el, _ := range rSet.Value {
				salt := strconv.FormatInt(i, 10)
				p.Leveldb.Put([]byte(k+"~"+salt), []byte(el), nil)
				i += 1
			}
			p.Leveldb.Put([]byte(k+"~end"), []byte(""), nil)

			//store rMap k~el_k v
		} else if rmap, ok := v.(*dt.Rmap); ok {
			p.Leveldb.Put([]byte(k+"~"), []byte("rmap"), nil)
			for elK, elV := range *rmap {
				p.Leveldb.Put([]byte(k+"~"+elK), []byte(elV), nil)
			}
			p.Leveldb.Put([]byte(k+"~~end"), []byte(""), nil)

			//store rList  k~nu v
		} else if rList, ok := v.(*dt.Rlist); ok {
			p.Leveldb.Put([]byte(k+"~"), []byte("rlist"), nil)
			var i int64 = 1
			for e := rList.Front(); e != nil; e = e.Next() {
				salt := strconv.FormatInt(i, 10)
				p.Leveldb.Put([]byte(k+"~"+salt), []byte(e.Value.(string)), nil)
				i += 1
			}
			p.Leveldb.Put([]byte(k+"~end"), []byte(""), nil)
		}
	}

	return
}
