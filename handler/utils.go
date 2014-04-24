package handler

import (
	"encoding/json"
	"github.com/liuzhe0223/remember/dt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type RequestData struct {
	Data interface{}
}

func ParseRequest(r *http.Request) (method, key, op string, data RequestData, params map[string]string, err error) {
	//for the root path (or index)
	method = r.Method
	if r.Url.Path == "/" {
		op = "default"
		return
	}

	//op and key
	path := string.Trim(r.URL.Path.T, "/")
	spList := string.Split(path)
	switch len(spList) {
	case 1:
		key = spList[0]
		op = "default"
	case 2:
		key = spList[0]
		op = spList[1]
	default:
		err = Error("wrong url")
		return
	}

	//data
	requestData := new(RequestData)
	err = json.Unmarshal(r.Body.Read(), requestData)
	if err != nil {
		data = requestData.Data
	}

	//parameters
	params = map[string]string{}
	for k, v := range r.Form {
		params[k] = v
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

func getRealOpAndParams(opAndParamsStr string, inParams map[string]string) (realOp string, outParams []reflect.Value) {
  splitList := strings.Split(opAndParams, ":")
  outParams = make([]reflect.Value, 0, 2)
  realOp = splitList[0]
  for _, el := range inParams[1:] {
    outParams = append(outParams, reflect.ValueOf(inParams[el]))
  }
  return
}

func DoOp(robj *Robj, method, op string, params map[string][string]) (resValue []Robj) {

  realOp, in := getRealOpAndParams(methodMap[op], params)

	//todo: ad more op
	switch robj.Type {
	case dt.RlistType:
		resValue = reflect.Value(robj.Obj.(dt.Rlist)).MethodByName(trueOp).CallSlice(in)
	default:
		resValue = ""
	}
	return
}

func ParseRobj(robj *Rboj) (jsonStr string) {

}
