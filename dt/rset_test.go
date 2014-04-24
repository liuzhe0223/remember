package dt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRset(t *testing.T) {
	rset := NewRset()
	newRset := new(Rset)
	newRset.Value = map[Robj]string{}
	assert.Equal(t, rset, newRset)
}

func TestRsetOp(t *testing.T) {
	rset := NewRset()

	robj1 := Robj{RintType, 1}
	robj2 := Robj{RintType, 2}
	robj3 := Robj{RintType, 3}

	rset.Sadd(robj1)
	rset.Sadd(robj2)
	rset.Sadd(robj3)

	rset.Spop(robj1)
	rset.Spop(robj3)
	assert.Equal(t, rset.All(), []Robj{robj2})
}
