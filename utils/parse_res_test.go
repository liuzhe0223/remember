package utils

import (
	"github.com/liuzhe0223/remember/dt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRes(t *testing.T) {
	rString := dt.Rstring("test")

	assert.Equal(t, ParseRes(rString), "{\"data\":\"test\"}")
}
