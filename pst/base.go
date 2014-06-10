package pst

import (
	"time"
)

type Pster struct {
	Db *map[string]interface{}
}

func timer(c chan int) {
	for {
		time.Sleep(time.Second * 20)
		c <- 0
	}
}

func (p *Pster) Go() {
	c := make(chan int)
	timer(c)

	for {
		<-c
		p.Dumps()
	}
}
