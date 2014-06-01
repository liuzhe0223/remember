package pst

import (
	"encoding/json"
	"github.com/liuzhe0223/remember/dt"
)

func (p *Pster) Dumps() string {
	return json.Marshal(p.Db)
}
