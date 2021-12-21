package base

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
		return ObjUndef
	}
	return l.Objects[loc.ToIndex(l.Dimension)]
}
