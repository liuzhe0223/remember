package utils

import (
	"github.com/liuzhe0223/remember/dt"
	"reflect"
)

func DoOp(robj *dt.Robj, method, op string, params map[string]interface{}) (resReflectValues []reflect.Value) {

	realOp, in, _ := getRealOpAndParams(methodMap[method][op], params)

	//todo: ad more op
	switch robj.Type {
	case dt.RlistType:
		resReflectValues = reflect.ValueOf(robj.Obj.(*dt.Rlist)).MethodByName(realOp).Call(in)
	default:
		resReflectValues = make([]reflect.Value, 0, 0)
	}

	//convent the res to []dt.Robj form []relect.Value
	// resValue = make([]dt.Robj, 0, len(resReflectValues))
	// for _, value := range resReflectValues {
	// 	valueInterface := value.Interface()
	// 	typedValue, _ := valueInterface.(dt.Robj)
	// 	resValue = append(resValue, typedValue)
	// }
	return
}
