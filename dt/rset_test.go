package dt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRset(t *testing.T) {
	rset := NewRset()
	newRset := new(Rset)
	newRset.Value = map[string]string{}
	assert.Equal(t, rset, newRset)
}

func TestRsetOp(t *testing.T) {
	rset := NewRset()

	v1 := "1"
	v2 := "2"
	v3 := "3"

	rset.Sadd(v1)
	rset.Sadd(v2)
	rset.Sadd(v3)

	rset.Spop(v1)
	rset.Spop(v3)
	assert.Equal(t, rset.Get(), []string{v2})
}
