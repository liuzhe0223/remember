package handler

import (
	"github.com/liuzhe0223/remember/dt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRealOPAndParams(t *testing.T) {
	opAndParamsStr := "Lrange:start:end"
	inParams := map[string]string{
		"start": "0",
		"end":   "6",
	}

	realOp, outParmas := getRealOpAndParams(opAndParamsStr, inParams)
	assert.Equal(t, "Lrange", realOp)

	testOutParmas := []string{"0", "6"}
	outStrParmas := make([]string, 0, 2)
	for _, value := range outParmas {
		valueInterface := value.Interface()
		valueStr, _ := valueInterface.(string)
		outStrParmas = append(outStrParmas, valueStr)
	}

	assert.Equal(t, testOutParmas, outStrParmas)
}

func TestParseBaseRobj(t *testing.T) {
	rintObj := dt.NewRintObj(3)
	rstringObj := dt.NewRstringObj("test")

	assert.Equal(t, string(ParseBaseRobj(rintObj)), "{\"data\":3}")
	assert.Equal(t, string(ParseBaseRobj(rstringObj)), "{\"data\":\"test\"}")
}
