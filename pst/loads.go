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

		fmt.Println(k)

		dT := valType(k, v)

		splitK := strings.Split(k, "~")
		realK := splitK[0]

		switch dT {
		case "rstring":
			rdb[k] = v
		case "rset":
			rdb[realK] = loadRset(k, iter)
		case "rmap":
			rdb[realK] = loadRmap(k, iter)
		case "rlist":
			rdb[realK] = loadRlist(k, iter)
		default:
			fmt.Println("in default")
			rdb[k] = v
		}
		fmt.Println("one  return , rdb:", rdb)
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

		fmt.Println("in load rset, key", string(iter.Key()))
		rset.Sadd(string(iter.Value()))
		fmt.Println("in load rset, rdb", rset)
	}

	return
}

func loadRmap(k string, iter iterator.Iterator) (rmap *dt.Rmap) {
	fmt.Println("in load rmap")
	rmap = dt.NewRmap()
	for iter.Next() {
		if string(iter.Key()) == k+"end" {
			return
		}

		splitElK := strings.Split(string(iter.Key()), "~")
		elK := splitElK[1]
		(*rmap)[elK] = string(iter.Value())
	}

	return
}

func loadRlist(k string, iter iterator.Iterator) (rlist *dt.Rlist) {
	fmt.Println("in load list")

	rlist = dt.NewRlist()
	for iter.Next() {
		if string(iter.Key()) == k+"end" {
			return
		}

		rlist.Rpush(string(iter.Value()))
	}

	return
}
