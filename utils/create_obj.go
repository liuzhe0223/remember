package utils

import (
	"fmt"
	"github.com/liuzhe0223/remember/dt"
)

var op2ObjType = map[string]dt.Rtype{
	"rpush": dt.RlistType,
	"lpush": dt.RlistType,
}

func CreateObjAccordingOp(op string) (robj interface{}) {
	fmt.Println("creating obj, op= ", op)
	switch op2ObjType[op] {
	case dt.RlistType:
		robj = dt.NewRlist()
		fmt.Println("created a new list")
	}
	return
}
