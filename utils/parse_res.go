package utils

import (
	"github.com/liuzhe0223/remember/dt"
	"strconv"
	"strings"
  "fmt"
  "reflect"
)

//parse base robj, int and string
func parseBaseRobj(robj *dt.Robj) (stringValue string) {
	switch robj.Type {
	case dt.RintType:
    fmt.Println("res obj type ==", reflect.TypeOf(robj.Obj))
		intValue, ok := robj.Obj.(dt.Rint)
    fmt.Println("parse data ...ok= ", ok)
		stringValue = strconv.FormatInt(int64(intValue), 10)
	case dt.RstringType:
		stringValue, _ = robj.Obj.(string)
		stringValue = "\"" + stringValue + "\""
	default:
		stringValue = ""
	}

	return
}

func parseRobjList(res []dt.Robj) string {
	resStrList := make([]string, 0, len(res))

	for _, robj := range res {
		resStrList = append(resStrList, parseBaseRobj(&robj))
	}

	return "[" + strings.Join(resStrList, ",") + "]"
}

func ParseRes(res interface{}) string {
	var resStr string
	if robj, ok := res.(dt.Robj); ok {
		resStr = parseBaseRobj(&robj)
	} else if robjList, ok := res.([]dt.Robj); ok {
		resStr = parseRobjList(robjList)
	}
	return `{"data":` + resStr + "}"
}
