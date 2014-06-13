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
		"slength": "Slength",
		"mget":    "Mget:data:string",
		"mkeys":   "Mkeys",
	},
	"POST": map[string]string{
		"set":   "Set:data:string",
		"rpush": "Rpush:data:string",
		"lpush": "Lpush:data:string",
		"sadd":  "Sadd:data:string",
		"mset":  "Mset:data:map",
	},
	"PUT": map[string]string{
		"rpop":    "Rpop",
		"lpop":    "Lpop",
		"spop":    "Spop:data:string",
		"mdelete": "Mdelete:data:string",
	},
}

func DoOp(robj interface{}, method, op string, params map[string]interface{}) (resReflectValues []reflect.Value) {

	fmt.Println("do oping, ___params =", params)

	realOp, in, _ := getRealOpAndParams(methodMap[method][op], params)

	fmt.Println("do oping, ___in =", in)

	//todo: ad more op
	fmt.Println("robj type: ", reflect.TypeOf(robj).String())

	if pRlist, ok := robj.(*dt.Rlist); ok {
		fmt.Println("do op rlist ing _______")
		resReflectValues = reflect.ValueOf(pRlist).MethodByName(realOp).Call(in)

	} else if pRset, ok := robj.(*dt.Rset); ok {
		fmt.Println("do op rset ing _______")
		fmt.Println("*rset _______: ", pRset)
		resReflectValues = reflect.ValueOf(pRset).MethodByName(realOp).Call(in)

	} else if pRstring, ok := robj.(*dt.Rstring); ok {
		fmt.Println("do op rstring ________-")
		resReflectValues = reflect.ValueOf(pRstring).MethodByName(realOp).Call(in)

	} else if pRmap, ok := robj.(*dt.Rmap); ok {
		fmt.Println("do op rmap________-")
		resReflectValues = reflect.ValueOf(pRmap).MethodByName(realOp).Call(in)
	} else {
		fmt.Println("do op default _______")
		resReflectValues = make([]reflect.Value, 0, 0)
	}

	fmt.Println("done op, ___value =", resReflectValues)
	return
}
