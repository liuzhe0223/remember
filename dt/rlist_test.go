package dt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRlist(t *testing.T) {
	rlist := NewRlist()

	assert.NotNil(t, rlist)
}

func TestGet(t *testing.T) {
	rlist := NewRlist()

  v1 := "test1"
  v2 := "test2"

	rlist.Rpush(v1)
	rlist.Rpush(v2)

	assert.Equal(t, rlist.Get(), []string{v1, v2})
}

func TestRpush(t *testing.T) {
	rlist := NewRlist()
  v1 := "test1"

	assert.True(t, rlist.Rpush(v1))
}

func TestLpush(t *testing.T) {
	rlist := NewRlist()
  v1 := "test1"

	assert.True(t, rlist.Lpush(v1))
}

func TestPop(t *testing.T) {
	rlist := NewRlist()
	v1 := "1"
	v2 := "2"
	v3 := "3"

	rlist.Rpush(v1)
	rlist.Rpush(v2)
	rlist.Rpush(v3)

	assert.Equal(t, rlist.Lpop(), v1)
	assert.Equal(t, rlist.Rpop(), v3)
}

func TestLrange(t *testing.T) {
	rlist := NewRlist()
	v1 := "1"
	v2 := "2"
	v3 := "3"

	rlist.Rpush(v1)
	rlist.Rpush(v2)
	rlist.Rpush(v3)

	assert.Equal(t, rlist.Lrange(0, 1), []string{v1, v2})
}
