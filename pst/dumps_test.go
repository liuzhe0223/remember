package pst

import (
	"encoding/json"
	"fmt"
	"github.com/liuzhe0223/remember/dt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDumps(t *testing.T) {
	rlistObj := dt.NewRlistObj(dt.NewRlist())
	robj := dt.Robj{
		Type: dt.RlistType,
		Obj:  rlistObj,
	}
	m := dt.Rmap{"test": robj}

	jsonByte, _ := json.Marshal(m)
	fmt.Println(string(jsonByte))

	var m2 dt.Rmap
	json.Unmarshal(jsonByte, &m2)

	fmt.Println(m2)
	assert.Equal(t, m, m2)
}
