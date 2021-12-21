package base

type Object uint

const (
	ObjUndef Object = iota
	ObjNone
	ObjBG
	ObjBase
	ObjGrass
	ObjTree
	ObjBox

	ObjMe Object = 123456789
)
