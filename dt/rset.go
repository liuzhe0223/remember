package dt

type Rset struct {
	ElType RobjType
	Len    uint
	Value  map[Robj]string
}

func NewRset() (rset *Rset) {
	rset = new(Rset)
	rset.Value = map[Robj]string{}
	return
}

func (rset *Rset) All() []Robj {
	res := make([]Robj, 0, rset.Len)
	for k, _ := range rset.Value {
		res = append(res, k)
	}
	return res
}

//return ok
func (rset *Rset) Sadd(value interface{}) bool {
	data := value.(Robj)
	if _, ok := rset.Value[data]; ok {
		return true
	}
	rset.Value[data] = ""
	rset.Len += 1
	return true
}

func (rset *Rset) Spop(key Robj) Robj {
	if _, ok := rset.Value[key]; ok {
		delete(rset.Value, key)
		rset.Len -= 1
		return key
	} else {
		return Robj{}
	}
}
