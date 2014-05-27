package utils

import (
	"github.com/liuzhe0223/remember/dt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseBaseRobj(t *testing.T) {
	rintObj := dt.NewRintObj(3)
	rstringObj := dt.NewRstringObj("test")

	assert.Equal(t, parseBaseRobj(rintObj), "3")
	assert.Equal(t, parseBaseRobj(rstringObj), "\"test\"")
}

func TestParseRobjList(t *testing.T) {
	rintObj := dt.NewRintObj(3)
	rstringObj := dt.NewRstringObj("test")
	robjList := []dt.Robj{*rintObj, *rstringObj}

	assert.Equal(t, parseRobjList(robjList), "[3,\"test\"]")
}

func TestParseRes(t *testing.T) {

	rintObj := dt.NewRintObj(3)
	rstringObj := dt.NewRstringObj("test")
	robjList := []dt.Robj{*rintObj, *rstringObj}

	assert.Equal(t, ParseRes(robjList), "{\"data\":[3,\"test\"]}")
	assert.Equal(t, ParseRes(*rintObj), "{\"data\":3}")
}
