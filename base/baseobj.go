package base

type BaseObject struct {
	objType ObjectType
	loc     Vertex
	flag    Flag
}

func NewObject(t ObjectType, loc Vertex) *BaseObject {

	o := BaseObject{
		objType: t,
		loc:     loc,
		flag:    GetDefaultFlags(t),
	}
	return &o
}

var _ Locatable = (*BaseObject)(nil)

func (o *BaseObject) ObjectType() ObjectType {
	return o.objType
}
func (o *BaseObject) Loc() Vertex {
	return o.loc
}

func (o *BaseObject) SetLoc(loc Vertex) {
	o.loc = loc
}

func (o *BaseObject) Flag() Flag {
	return o.flag
}

func (o *BaseObject) SetFlag(flag Flag) {
	o.flag = flag
}
