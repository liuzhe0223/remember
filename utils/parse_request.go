package utils

import (
	"encoding/json"
	"errors"
	"github.com/liuzhe0223/remember/dt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
  "fmt"
)

type RequestData struct {
	Data interface{} `json:"data"`
}

func ParseRequest(r *http.Request) (method, key, op string, params map[string]interface{}, err error) {
	//for the root path (or index)
	method = r.Method

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

	//parameters
	params = map[string]interface{}{}
	r.ParseForm()
	for k, v := range r.Form {
		params[k] = v[0]
	}

	//data
	requestData := new(RequestData)
	body, _ := ioutil.ReadAll(r.Body)
	if method != "GET" && len(body) != 0 {
		err = json.Unmarshal(body, &requestData)
		if err != nil {
      fmt.Println("json unmarshal__ err", err)
			return
		}
		params["data"] = requestData.Data
	}

	return
}

func getRealOpAndParams(opAndParamsStr string, inParams map[string]interface{}) (realOp string, outParams []reflect.Value, err error) {
	splitList := strings.Split(opAndParamsStr, ":")
	outParams = make([]reflect.Value, 0, 2)
	realOp = splitList[0]
	length := len(splitList)
	for i := 1; i < length; i += 2 {
		var value reflect.Value
		if splitList[i] == "data" {
			data := inParams["data"]
			robj, err := parseData(data)
      if err != nil {
        fmt.Println(err)
      }
			value = reflect.ValueOf(robj)
		} else {
			strValue := inParams[splitList[i]].(string)
			if splitList[i+1] == "int" {
				int64Value, _ := strconv.ParseInt(strValue, 10, 0)
				intValue := int(int64Value)
				value = reflect.ValueOf(intValue)
			} else {
				value = reflect.ValueOf(strValue)
			}
		}
		outParams = append(outParams, value)
	}

	return
}

func parseData(data interface{}) (robj dt.Robj, err error) {
  fmt.Println("parseData________-data= ", data)
  fmt.Println("parseData________-data type= ", reflect.TypeOf(data))
	if intData, ok := data.(float64); ok {
    fmt.Println("parseData________- intData = ", intData)
		robj = dt.Robj{
			Type: dt.RintType,
			Obj:  dt.Rint(intData),
		}
	} else if strData, ok := data.(string); ok {
    fmt.Println("parseData________- strData = ", strData)
		robj = dt.Robj{
			Type: dt.RstringType,
			Obj:  strData,
		}
	} else if mapData, ok := data.(map[string]interface{}); ok {
		robj = dt.Robj{
			Type: dt.RmapType,
			Obj:  mapData,
		}
	} else {
		err = errors.New("wrong data type")
	}
	return
}
