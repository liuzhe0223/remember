package dt

type RobjType uint

const (
	RintType    RobjType = 0
	RstringType RobjType = 1
	RlistType   RobjType = 2
	RmapType    RobjType = 3
)

func (robjType RobjType) IsInt() bool {
	return robjType == RintType
}

func (robjType RobjType) IsString() bool {
	return robjType == RstringType
}

//robj
type Robj struct {
	Type RobjType
	Obj  interface{}
}

func NewRintObj(value int) *Robj {
	return &Robj{
		Type: RintType,
		Obj:  Rint(value),
	}
}

func NewRstringObj(value string) *Robj {
	return &Robj{
		Type: RstringType,
		Obj:  Rstring(value),
	}
}

func NewRlistObj(value *Rlist) *Robj {
	return &Robj{
		Type: RlistType,
		Obj:  value,
	}
}
