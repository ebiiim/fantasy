package base

type ObjectName string

type Object interface {
	ObjectName() ObjectName
	ObjectType() ObjectType

	Loc() Vertex
	SetLoc(Vertex)

	Flag() Flag
	SetFlag(Flag)
}

type BaseObject struct {
	objName ObjectName
	objType ObjectType
	loc     Vertex
	flag    Flag
}

var _ Object = (*BaseObject)(nil)

func NewObject(n ObjectName, t ObjectType, loc Vertex) *BaseObject {
	o := BaseObject{
		objName: n,
		objType: t,
		loc:     loc,
		flag:    GetDefaultFlags(t),
	}
	return &o
}

func (o *BaseObject) ObjectName() ObjectName { return o.objName }

func (o *BaseObject) ObjectType() ObjectType { return o.objType }

func (o *BaseObject) Loc() Vertex { return o.loc }

func (o *BaseObject) SetLoc(loc Vertex) { o.loc = loc }

func (o *BaseObject) Flag() Flag { return o.flag }

func (o *BaseObject) SetFlag(flag Flag) { o.flag = flag }
