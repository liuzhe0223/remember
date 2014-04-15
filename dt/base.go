package dt

type RobjType uint

const (
	RintType    RobjType = 0
	RstringType RobjType = 1
	RlistType   RobjType = 2
	RmapType    RobjType = 3
)

type Robj struct {
	Type RobjType
	Obj  interface{}
}
