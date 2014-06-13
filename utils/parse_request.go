package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
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
		fmt.Println("origin data ____--:", params["data"])
	}

	return
}

func getRealOpAndParams(opAndParamsStr string, inParams map[string]interface{}) (realOp string, outParams []reflect.Value, err error) {
	splitList := strings.Split(opAndParamsStr, ":")
	realOp = splitList[0]

	outParams = make([]reflect.Value, 0, 2)
	for i := 1; i < len(splitList); i += 2 {
		switch splitList[i+1] {
		case "int":
			inValue, _ := strconv.ParseInt(inParams[splitList[i]].(string), 10, 64)
			outParams = append(outParams, reflect.ValueOf(int(inValue)))
		case "string":
			outParams = append(outParams, reflect.ValueOf(inParams[splitList[i]]))
		case "map":
			m := formatMapData(inParams[splitList[i]].(map[string]interface{}))
			outParams = append(outParams, reflect.ValueOf(m))
		}
	}

	return
}

func isValidData(data interface{}) (res bool) {
	fmt.Println("parseData________-data= ", data)
	fmt.Println("parseData________-data type= ", reflect.TypeOf(data))

	if _, ok := data.(string); ok {
		res = true
		return
	} else if _, ok := data.(map[string]interface{}); ok {
		res = true
		return
	}
	res = false
	return
}

func formatMapData(data map[string]interface{}) (resMap map[string]string) {
	resMap = map[string]string{}
	for k, v := range data {
		resMap[k] = v.(string)
	}
	return
}
