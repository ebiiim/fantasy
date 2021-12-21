package base

func initRegisterFlags() map[ObjectType]FlagSet {
	m := make(map[ObjectType]FlagSet)

	m[ObjUndef] = FBlockAll
	m[ObjMe] = FMe
	m[ObjBG] = FBlockAll
	m[ObjNone] = FNone
	m[ObjBase] = FNone
	m[ObjGrass] = FNone
	m[ObjTree] = FLandObject
	m[ObjBox] = FBox

	return m
}
