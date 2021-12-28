package base

type Locatable interface {
	ObjectType() ObjectType

	Loc() Vertex
	SetLoc(Vertex)

	Flag() Flag
	SetFlag(Flag)
}
