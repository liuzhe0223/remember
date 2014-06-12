package dt

type Rmap map[string]string

func NewRmap() *Rmap {
	rmap := Rmap{}
	return &rmap
}

func (m *Rmap) Get() map[string]string {
	return map[string]string(*m)
}

func (m *Rmap) Mset(data map[string]string) bool {
	for k, v := range data {
		(*m)[k] = v
	}
	return true
}

func (m *Rmap) Mget(k string) string {
	return (*m)[k]
}

func (m *Rmap) Mdelete(k string) bool {
	delete(*m, k)
	return true
}

func (m *Rmap) Mkeys() []string {
	resList := make([]string, 0, len(*m))
	for k, _ := range *m {
		resList = append(resList, k)
	}
	return resList
}
