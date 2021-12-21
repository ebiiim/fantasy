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

func (m *Map) GetObjects(loc Vertex) []*Object {
	objs := make([]*Object, len(m.Layers))
	for idx, layer := range m.Layers {
		objs[idx] = layer.GetObject(loc)
	}
	return objs
}
