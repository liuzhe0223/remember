package pst

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoads(t *testing.T) {
	pster := Pster{}
	db, err := pster.Loads()

	fmt.Println(db)
	assert.Nil(t, err)
}
