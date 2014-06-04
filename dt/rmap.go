package dt

type Rmap map[string]string

func (m *Rmap) Get() Rmap {
  return *m
}

func (m *Rmap) Mset(k, v string) bool{
  (*m)[k] = v
  return true
}

func (m *Rmap) Mget(k string) string {
  return (*m)[k]
}

func (m *Rmap) Mdelete(k string) bool {
  delete(*m, k)
  return true
}

func (m *Rmap) Mkeys() []string{
  resList := make([]string, 0, len(*m))
  for k, _ := range *m {
    resList = append(resList, k)
  }
  return resList
}
