package dt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRlist(t *testing.T) {
	rlist := NewRlist()

	assert.NotNil(t, rlist)
}

func TestRpush(t *testing.T) {
	rlist := NewRlist()
	robj := Robj{RintType, 1}

	assert.True(t, rlist.Rpush(robj))
}

func TestLpush(t *testing.T) {
	rlist := NewRlist()
	robj := Robj{RintType, 1}

	assert.True(t, rlist.Lpush(robj))
}

func TestPop(t *testing.T) {
	rlist := NewRlist()
	robj1 := Robj{RintType, 1}
	robj2 := Robj{RintType, 2}
	robj3 := Robj{RintType, 3}

	rlist.Rpush(robj1)
	rlist.Rpush(robj2)
	rlist.Rpush(robj3)

	assert.Equal(t, rlist.Lpop(), robj1)
	assert.Equal(t, rlist.Rpop(), robj3)
}

func TestLrange(t *testing.T) {
	rlist := NewRlist()
	robj1 := Robj{RintType, 1}
	robj2 := Robj{RintType, 2}
	robj3 := Robj{RintType, 3}

	rlist.Rpush(robj1)
	rlist.Rpush(robj2)
	rlist.Rpush(robj3)

	assert.Equal(t, rlist.Lrange(0, 1), []Robj{robj1, robj2})
}
