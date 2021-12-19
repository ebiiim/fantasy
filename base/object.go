package base

type FlagSet uint

type ObjectType int

type Object struct {
	Type    ObjectType
	Loc     Vertex
	FlagSet FlagSet
}

func NewObject(objType ObjectType, v Vertex) *Object {
	obj := Object{objType, v, GetFlagSet[objType]}
	return &obj
}
