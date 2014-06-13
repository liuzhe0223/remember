package utils

import (
	"encoding/json"
)

func ParseRes(res interface{}) string {

	resByte, _ := json.Marshal(res)
	resStr := string(resByte)

	return `{"data":` + resStr + "}"
}
