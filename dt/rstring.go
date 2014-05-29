package dt

type Rstring string

func NewString() *Rstring {
  var value Rstring 
  return &value
}

func (s *Rstring) Get() Robj {
  return Robj {
    Type: RintType,
    Obj: *s,
  }
}

func (s *Rstring) Sset(robj Robj) bool {
  if value, ok := robj.Obj.(Rstring); ok {
    *s = value
    return true
  }
  return false
}
