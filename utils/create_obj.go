package utils

import (
	"github.com/liuzhe0223/remember/dt"
)

var op2ObjType = map[string]dt.RobjType{
	"Rpush": dt.RlistType,
}

func createObjAccordingOp(op string) (robj dt.Robj) {
	switch op2ObjType[op] {
	case dt.RlistType:
		robj = dt.Robj{
			Type: dt.RlistType,
			Obj:  dt.NewRlist(),
		}
	}
}
