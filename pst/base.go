package pst

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"time"
)

type Pster struct {
	Db      *map[string]interface{}
	Leveldb *leveldb.DB
}

func (p *Pster) Init() {
	p.Leveldb, _ = leveldb.OpenFile("/home/zhe/db", nil)
	return
}

func timer(c chan int) {
	for {
		time.Sleep(time.Second * 2)
		c <- 0
	}
}

func (p *Pster) Go() {
	c := make(chan int)
	go timer(c)

	for {
		<-c
		fmt.Println("_________sync_________--")
		p.Dumps()
	}
}
