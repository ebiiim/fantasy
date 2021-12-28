package base

type LayerName string

type Layer struct {
	Name      LayerName
	Dimension Vertex
	Objects   []Locatable
}

func NewLayer(name LayerName, size Vertex, objList []Locatable) *Layer {
	l := Layer{
		Name:      name,
		Dimension: size,
		Objects:   objList,
	}
	return &l
}

func (l *Layer) GetObject(loc Vertex) Locatable {
	if loc.IsOutside(l.Dimension) {
		return NewObject(ObjUndef, loc)
	}
	return l.Objects[loc.ToIndex(l.Dimension)]
}

func (l *Layer) GetObjectOrErr(loc Vertex) (Locatable, error) {
	if loc.IsOutside(l.Dimension) {
		return nil, ErrNoObjectFound
	}
	return l.Objects[loc.ToIndex(l.Dimension)], nil
}
