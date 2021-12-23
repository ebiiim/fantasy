package base

type ObjectType uint

type Object struct {
	Type ObjectType
	Loc  Vertex
	Flag Flag
}

func NewObject(t ObjectType, loc Vertex) *Object {

	o := Object{
		Type: t,
		Loc:  loc,
		Flag: GetDefaultFlags(t),
	}
	return &o
}
