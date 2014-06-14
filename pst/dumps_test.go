package pst

import (
	"github.com/liuzhe0223/remember/dt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDumps(t *testing.T) {
	pRstring := dt.NewString()
	pRstring.Set("test_string")

	pRset := dt.NewRset()
	pRset.Sadd("test_set_string")
	pRset.Sadd("test_set_string2")

	pRlist := dt.NewRlist()
	pRlist.Rpush("hello_list")
	pRlist.Rpush("hello2_list")

	pRmap := dt.NewRmap()
	(*pRmap)["test_k"] = "test_key"
	(*pRmap)["test_k1"] = "test_key1"

	Db := map[string]interface{}{
		"stringV": pRstring,
		"setV":    pRset,
		"listV":   pRlist,
		"mapV":    pRmap,
	}

	p := Pster{Db: &Db}
	p.Init()
	defer p.Leveldb.Close()

	assert.Nil(t, p.Dumps())
}
