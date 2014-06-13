package utils

import (
	"fmt"
	"github.com/liuzhe0223/remember/dt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoOP(t *testing.T) {
	rlist := dt.NewRlist()

	robj1 := "1"
	robj2 := "2"
	robj3 := "3"

	rlist.Rpush(robj1)
	rlist.Rpush(robj2)
	rlist.Rpush(robj3)

	params := map[string]interface{}{
		"start": "1",
		"end":   "2",
	}
	fmt.Println("________rlist: ", rlist)
	opRes := DoOp(rlist, "GET", "lrange", params)
	fmt.Println(opRes)

	res, _ := opRes[0].Interface().([]string)
	assert.Equal(t, []string{"2", "3"}, res)
}
