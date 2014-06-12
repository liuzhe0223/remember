package utils

import (
	"fmt"
	"github.com/liuzhe0223/remember/dt"
)

var op2ObjType = map[string]dt.Rtype{
	"set":   dt.RstringType,
	"rpush": dt.RlistType,
	"lpush": dt.RlistType,
	"sadd":  dt.RsetType,
	"mset":  dt.RmapType,
}

func CreateObjAccordingOp(op string) (robj interface{}) {
	fmt.Println("creating obj, op= ", op)
	switch op2ObjType[op] {
	case dt.RlistType:
		robj = dt.NewRlist()
		fmt.Println("created a new list")
	case dt.RsetType:
		robj = dt.NewRset()
	case dt.RstringType:
		robj = dt.NewString()
	case dt.RmapType:
		robj = dt.NewRmap()
	}
	return
}
