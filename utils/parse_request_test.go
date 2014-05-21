package utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestGetRealOPAndParams(t *testing.T) {
	opAndParamsStr := "Lrange:start:int:end:int"
	inParams := map[string]interface{}{
		"start": "0",
		"end":   "6",
	}

	realOp, outParmas, _ := getRealOpAndParams(opAndParamsStr, inParams)
	assert.Equal(t, "Lrange", realOp)

	testOutParmas := []interface{}{0, 6}
	outStrParmas := make([]interface{}, 0, 2)
	for _, value := range outParmas {
		valueInterface := value.Interface()
		outStrParmas = append(outStrParmas, valueInterface)
	}

	assert.Equal(t, testOutParmas, outStrParmas)
}

func TestParseRequest(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://test.com/key/lrange?start=1&end=2", strings.NewReader("{\"data\":\"hello\"}"))
	method, key, op, params, _ := ParseRequest(r)

	assert.Equal(t, "GET", method)
	assert.Equal(t, "key", key)
	assert.Equal(t, "lrange", op)
	expectParams := map[string]interface{}{
		"start": "1",
		"end":   "2",
		"data":  "hello",
	}
	assert.Equal(t, expectParams, params)
}
