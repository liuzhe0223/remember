package pst

import (
	"fmt"
	"github.com/liuzhe0223/remember/dt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"strings"
)

func (p *Pster) Loads() (rdb map[string]interface{}, err error) {
	db, err := leveldb.OpenFile("/home/zhe/db", nil)
	if err != nil {
		return
	}
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	rdb = map[string]interface{}{}

	for iter.Next() {
		k := string(iter.Key())
		v := string(iter.Value())

		dT := valType(k, v)

		splitK := strings.Split(k, "~")
		realK := splitK[0]

		switch dT {
		case "rstring":
			rstring := dt.Rstring(v)
			rdb[k] = &rstring
		case "rset":
			fmt.Println("in rset")
			rdb[realK] = loadRset(k, iter)
		case "rmap":
			fmt.Println("in rmap")
			rdb[realK] = loadRmap(k, iter)
		case "rlist":
			fmt.Println("in rlist")
			rdb[realK] = loadRlist(k, iter)
		default:
			rdb[k] = v
		}
	}

	return
}

func valType(k, v string) string {
	if strings.HasSuffix(k, "~") {
		return v
	}

	return "rstring"
}

func loadRset(k string, iter iterator.Iterator) (rset *dt.Rset) {

	rset = dt.NewRset()
	for iter.Next() {
		if string(iter.Key()) == k+"end" {
			return
		}

		rset.Sadd(string(iter.Value()))
	}

	return
}

func loadRmap(k string, iter iterator.Iterator) (rmap *dt.Rmap) {
	rmap = dt.NewRmap()
	for iter.Next() {
		if string(iter.Key()) == k+"~end" {
			return
		}

		splitElK := strings.Split(string(iter.Key()), "~")
		elK := splitElK[1]
		(*rmap)[elK] = string(iter.Value())
	}

	return
}

func loadRlist(k string, iter iterator.Iterator) (rlist *dt.Rlist) {

	rlist = dt.NewRlist()
	for iter.Next() {
		if string(iter.Key()) == k+"end" {
			return
		}

		rlist.Rpush(string(iter.Value()))
	}

	return
}
