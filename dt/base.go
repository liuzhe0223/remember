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

type Robj struct {
	Type RobjType
	Obj  interface{}
}
