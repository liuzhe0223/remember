package handler

import (
	"encoding/json"
	"errors"
	"github.com/liuzhe0223/remember/dt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type RequestData struct {
	Data interface{} `json:"data"`
}

func ParseRequest(r *http.Request) (method, key, op string, data interface{}, params map[string]string, err error) {
	//for the root path (or index)
	method = r.Method
	if r.URL.Path == "/" {
		op = "default"
		return
	}

	//op and key
	path := strings.Trim(r.URL.Path, "/")
	spList := strings.Split(path, "/")
	switch len(spList) {
	case 1:
		key = spList[0]
		op = "default"
	case 2:
		key = spList[0]
		op = spList[1]
	default:
		err = errors.New("wrong url")
		return
	}

	//data
	requestData := new(RequestData)
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, requestData)
	if err != nil {
		data = requestData.Data
	}

	//parameters
	params = map[string]string{}
	for k, v := range r.Form {
		params[k] = v[0]
	}

	return
}

//method op map for do op
var methodMap = map[string]map[string]string{
	"GET": map[string]string{
		"default": "Get",
		"lrange":  "Lrange:start:end",
	},
}

func getRealOpAndParams(opAndParamsStr string, inParams map[string]interface{}) (realOp string, outParams []reflect.Value) {
	splitList := strings.Split(opAndParamsStr, ":")
	outParams = make([]reflect.Value, 0, 2)
	realOp = splitList[0]
	for _, el := range splitList[1:] {
		outParams = append(outParams, reflect.ValueOf(inParams[el]))
	}
	return
}

func DoOp(robj *dt.Robj, method, op string, params map[string]interface{}) (resReflectValues []reflect.Value) {

	realOp, in := getRealOpAndParams(methodMap[method][op], params)

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

//parse base robj, int and string
func ParseBaseRobj(robj *dt.Robj) (jsonBytes []byte) {
	switch robj.Type {
	case dt.RintType:
		intValue, _ := robj.Obj.(int)
		data2Int := map[string]int{"data": intValue}
		jsonBytes, _ = json.Marshal(data2Int)
	case dt.RstringType:
		stringValue, _ := robj.Obj.(string)
		data2String := map[string]string{"data": stringValue}
		jsonBytes, _ = json.Marshal(data2String)
	}
	return
}
