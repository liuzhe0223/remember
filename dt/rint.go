package dt

type Rint int

func NewRint() *Rint {
  var value Rint
  return &value
}

func (i *Rint) Get() Robj {
  return Robj {
    Type: RintType,
    Obj: *i,
  }
}

func (i *Rint) Iset(robj Robj) bool {
  if value, ok := robj.Obj.(Rint); ok {
    *i = value
    return true
  }
  return false
}
