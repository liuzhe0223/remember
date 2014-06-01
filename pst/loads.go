package pst

import (
	"encoding/json"
	"github.com/liuzhe0223/remember/dt"
	"io"
)

func (p *Pster) Loads() {

	var jsonStr string
	r.Read(&jsonStr)
	json.Unmarshal(jsonStr, p.Db)

	return
}
