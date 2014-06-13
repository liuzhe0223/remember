package dt

type Rstring string

func NewString() *Rstring {
	var value Rstring
	return &value
}

func (s *Rstring) Get() string {
	return string(*s)
}

func (s *Rstring) Set(v string) bool {
	*s = Rstring(v)
	return true
}
