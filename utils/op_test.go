package utils

import (
	"github.com/liuzhe0223/remember/dt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoOP(t *testing.T) {
	rlist := dt.NewRlist()

	robj1 := dt.Robj{dt.RintType, 1}
	robj2 := dt.Robj{dt.RintType, 2}
	robj3 := dt.Robj{dt.RintType, 3}

	rlist.Rpush(robj1)
	rlist.Rpush(robj2)
	rlist.Rpush(robj3)

	params := map[string]interface{}{
		"start": "1",
		"end":   "2",
	}

	rlistObj := dt.Robj{dt.RlistType, rlist}
	opRes := DoOp(&rlistObj, "GET", "lrange", params)
	res, _ := opRes[0].Interface().([]dt.Robj)
	assert.Equal(t, []dt.Robj{robj2, robj3}, res)
}
