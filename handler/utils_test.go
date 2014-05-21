package handler

import (
	"github.com/liuzhe0223/remember/dt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRealOPAndParams(t *testing.T) {
	opAndParamsStr := "Lrange:start:end"
	inParams := map[string]interface{}{
		"start": 0,
		"end":   6,
	}

	realOp, outParmas := getRealOpAndParams(opAndParamsStr, inParams)
	assert.Equal(t, "Lrange", realOp)

	testOutParmas := []interface{}{0, 6}
	outStrParmas := make([]interface{}, 0, 2)
	for _, value := range outParmas {
		valueInterface := value.Interface()
		outStrParmas = append(outStrParmas, valueInterface)
	}

	assert.Equal(t, testOutParmas, outStrParmas)
}

func TestParseBaseRobj(t *testing.T) {
	rintObj := dt.NewRintObj(3)
	rstringObj := dt.NewRstringObj("test")

	assert.Equal(t, string(ParseBaseRobj(rintObj)), "{\"data\":3}")
	assert.Equal(t, string(ParseBaseRobj(rstringObj)), "{\"data\":\"test\"}")
}

func TestDoOP(t *testing.T) {
	rlist := dt.NewRlist()

	robj1 := dt.Robj{dt.RintType, 1}
	robj2 := dt.Robj{dt.RintType, 2}
	robj3 := dt.Robj{dt.RintType, 3}

	rlist.Rpush(robj1)
	rlist.Rpush(robj2)
	rlist.Rpush(robj3)

	params := map[string]interface{}{
		"start": 1,
		"end":   2,
	}

	rlistObj := dt.Robj{dt.RlistType, rlist}
	opRes := DoOp(&rlistObj, "GET", "lrange", params)
	res, _ := opRes[0].Interface().([]dt.Robj)
	assert.Equal(t, []dt.Robj{robj2, robj3}, res)
}
