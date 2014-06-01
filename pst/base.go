package pst

import (
	"github.com/liuzhe0223/remember/dt"
	"io"
	"time"
)

type Pster struct {
	Db *dt.Rmap
}

func timer(c chan int) {
	for {
		time.Sleep(time.Second * 20)
		c <- 0
	}
}

func (p *Pster) Go() {
	c = make(chan int)
	timer(c)

	for {
		<-c
		p.Dumps(p.Db)
	}
}
