package utils

import (
	"fmt"
	"github.com/liuzhe0223/remember/dt"
)

var op2ObjType = map[string]dt.RobjType{
	"iset": dt.RintType,
	"rpush": dt.RlistType,
	"lpush": dt.RlistType,
}

func CreateObjAccordingOp(op string) (robj dt.Robj) {
	fmt.Println("creating obj, op= ", op)
	switch op2ObjType[op] {
	case dt.RlistType:
		robj = dt.Robj{
			Type: dt.RlistType,
			Obj:  dt.NewRlist(),
		}
	case dt.RintType:
		robj = dt.Robj{
			Type: dt.RintType,
			Obj:  dt.NewRint(),
		}
	}
	return
}
