package base

type MapName string

type Map struct {
	Name      MapName
	Dimension Vertex
	Layers    []*Layer
}

func NewMap(name MapName, size Vertex, layers []*Layer) *Map {
	m := Map{Name: name, Dimension: size, Layers: layers}
	return &m
}

func (m *Map) GetObjects(loc Vertex) []Object {
	objs := make([]Object, len(m.Layers))
	for idx, layer := range m.Layers {
		objs[idx] = layer.GetObject(loc)
	}
	return objs
}

func (m *Map) GetObjectsOrErr(loc Vertex) ([]Object, error) {
	objs := make([]Object, len(m.Layers))
	for idx, layer := range m.Layers {
		obj, err := layer.GetObjectOrErr(loc)
		if err != nil {
			return nil, ErrNoObjectFound
		}
		objs[idx] = obj
	}
	return objs, nil
}
