package handler

// import (
// 	"github.com/liuzhe0223/remember/dt"
// 	"io"
// 	"net/http"
// )

// func Handler(w http.ResponseWriter, r *http.Request) {

// 	method, key, op, data, params, err := ParseRquest(r)

// 	value, ok := Db[key]

// 	//this key is not in the db
// 	if !ok && method == "Get" {
// 		io.WriteString(w, "")
// 		return
// 	}

// 	resValues, err := DoOp(&value, method, op, params)

// 	resString := ParseRobj(&resValue)

// 	io.WriteString(w, resString)
// }
