package main

import (
	"fmt"
	"github.com/liuzhe0223/remember/dt"
	"github.com/liuzhe0223/remember/pst"
	"github.com/liuzhe0223/remember/utils"
	"io"
	"log"
	"net/http"
)

var Db dt.Rmap

func main() {
	//init db
	pster := pst.Pster{Db: &Db}
	go pster.Go()

	http.HandleFunc("/", Handler)
  fmt.Println("This is a remember server, listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		io.WriteString(w, "Hello, this is a remember server !")
		return
	}
	method, key, op, params, err := utils.ParseRequest(r)
	fmt.Println("r parsed______", r)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "bad request")
		return
	}

	//this key is not in the db
	if _, ok := Db[key]; !ok {
		if method == "GET" {
			io.WriteString(w, "not found")
			return
		} else {
			Db[key] = utils.CreateObjAccordingOp(op)
		}
	}
	value := Db[key]
	fmt.Println("found or create key___", value)

	resValues := utils.DoOp(&value, method, op, params)

	fmt.Println("final value =   ", value)

  w.Header().Set("Content-Type", "application/json")
	var jsonStr string
	if objRes, ok := resValues[0].Interface().([]dt.Robj); ok {
		jsonStr = utils.ParseRes(objRes)
	} else if objListRes, ok := resValues[0].Interface().(dt.Robj); ok {
		jsonStr = utils.ParseRes(objListRes)
  } else if boolVlaue, ok := resValues[0].Interface().(bool); ok{
    jsonStr = utils.ParseRes(boolVlaue)
	} else {
		io.WriteString(w, "server err!")
	}

	io.WriteString(w, jsonStr)
	return
}
