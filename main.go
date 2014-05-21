package main

import (
	"github.com/liuzhe0223/remember/dt"
	"github.com/liuzhe0223/remember/utils"
	"io"
	"log"
	"net/http"
)

var Db dt.Rmap

func main() {
	//init db
	Db = dt.Rmap{}

	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {

	method, key, op, params, err := utils.ParseRequest(r)
	if err != nil {
		io.WriteString(w, "bad request")
	}

	value, ok := Db[key]

	//this key is not in the db
	if !ok && method == "Get" {
		io.WriteString(w, "")
		return
	}

	resValues := utils.DoOp(&value, method, op, params)

	if res, ok := resValues[0].Interface().([]dt.Robj); !ok {
		if res, ok = resValues[0].Interface().(dt.Robj); !ok {
			io.WriteString(w, "server err!")
		}
	}

	jsonStr := utils.ParseRes(res)
	io.WriteString(w, jsonStr)
}
