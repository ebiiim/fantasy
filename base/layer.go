package base

import (
	"fmt"
)

type LayerName string

type Layer struct {
	Name      LayerName
	Dimension Vertex
	Objects   []Object
}

func NewLayer(name LayerName, size Vertex, objList []Object) *Layer {
	l := Layer{
		Name:      name,
		Dimension: size,
		Objects:   objList,
	}
	return &l
}

func (l *Layer) GetObject(loc Vertex) Object {
	if loc.IsOutside(l.Dimension) {
		return NewObject(
			ObjectName(fmt.Sprintf("MapOutside-%d_%d", loc.X, loc.Y)), ObjUndef, loc)
	}
	return l.Objects[loc.ToIndex(l.Dimension)]
}

func (l *Layer) GetObjectOrErr(loc Vertex) (Object, error) {
	if loc.IsOutside(l.Dimension) {
		return nil, ErrNoObjectFound
	}
	return l.Objects[loc.ToIndex(l.Dimension)], nil
}
