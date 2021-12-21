package flag

import "github.com/ebiiim/fantasy/base"

func init() {
	m = make(map[base.Object]Flag)

	m[base.ObjUndef] = BlockAll
	m[base.ObjMe] = Me
	m[base.ObjBG] = BlockAll
	m[base.ObjNone] = None
	m[base.ObjBase] = None
	m[base.ObjGrass] = None
	m[base.ObjTree] = LandObject
	m[base.ObjBox] = Box
}
