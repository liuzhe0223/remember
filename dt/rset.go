package dt

type Rset struct {
	Len    uint
	Value  map[string]string
}

func NewRset() (rset *Rset) {
	rset = new(Rset)
	rset.Value = map[string]string{}
	return
}

func (rset *Rset) All() []string{
	res := make([]string, 0, rset.Len)
	for k, _ := range rset.Value {
		res = append(res, k)
	}
	return res
}

//return ok
func (rset *Rset) Sadd(v string) bool {
	if _, ok := rset.Value[v]; ok {
		return true
	}

	rset.Value[v] = ""
	rset.Len += 1
	return true
}

func (rset *Rset) Spop(v string) string{
	if _, ok := rset.Value[v]; ok {
		delete(rset.Value, v)
		rset.Len -= 1
		return v
	} else {
		return ""
	}
}
