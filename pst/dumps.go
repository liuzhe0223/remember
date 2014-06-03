package pst

import (
	"encoding/json"
)

func (p *Pster) Dumps() string {
	jsonStr, _ := json.Marshal(p.Db)
	return string(jsonStr)
}
