package utils

import (
	"fmt"
	"github.com/liuzhe0223/remember/dt"
	"reflect"
)

//method op map for do op
var methodMap = map[string]map[string]string{
	"GET": map[string]string{
		"default": "Get",
		"lrange":  "Lrange:start:int:end:int",
	},
	"POST": map[string]string{
		"iset":  "Iset:data:Robj",
		"rpush": "Rpush:data:Robj",
		"lpush": "Lpush:data:Robj",
	},
	"PUT": map[string]string{
		"rpop": "Rpop",
		"lpop": "Lpop",
	},
}

func DoOp(robj interface{}, method, op string, params map[string]interface{}) (resReflectValues []reflect.Value) {

	fmt.Println("do oping, ___params =", params)
	realOp, in, _ := getRealOpAndParams(methodMap[method][op], params)

	fmt.Println("do oping, ___in =", in)
	//todo: ad more op
	switch robj.Type {
	case dt.RlistType:
		resReflectValues = reflect.ValueOf(robj.Obj.(*dt.Rlist)).MethodByName(realOp).Call(in)
	default:
		resReflectValues = make([]reflect.Value, 0, 0)
	}

	fmt.Println("done op, ___value =", resReflectValues)
	return
}
