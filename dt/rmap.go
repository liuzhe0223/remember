package dt

type Rmap map[string]Robj

func (m *Rmap) Get() (robj Robj) {
  return Robj{
    Type: RmapType,
    Obj: m,
  }
}

func (m *Rmap) Mset(robj Robj) bool{
  mValue := robj.Obj.(Rmap)
  var k string
  var v Robj
  for k, v = range mValue {
  }
  (*m)[k] = v
  return true
}

func (m *Rmap) Mget(robj Robj) Robj {
  return (*m)[robj.Obj.(string)]
}

func (m *Rmap) Mdelete(robj Robj) bool {
  delete(*m, robj.Obj.(string))
  return true
}

func (m *Rmap) Mkeys() []Robj{
  resList := []Robj{}
  for k, _ := range *m {
    kObj := Robj {
      Type: RstringType,
      Obj: k,
    }
    resList = append(resList, kObj)
  }
  return resList
}
