package base

type MapName int

type Map struct {
	Name        MapName
	Size        Vertex
	Layers      []*Layer
	landMovable []bool
}

func NewMap(name MapName, size Vertex, layers []*Layer) *Map {
	m := Map{Name: name, Size: size, Layers: layers}
	m.landMovable = make([]bool, m.Size.X*m.Size.Y)
	return &m
}

func (m *Map) GetObjects(loc Vertex) []*Object {
	objs := make([]*Object, len(m.Layers))
	for idx, layer := range m.Layers {
		objs[idx] = layer.GetObject(loc)
	}
	return objs
}

func (m *Map) Update() error {
	m.updateLandMovable()
	return nil
}

func (m *Map) updateLandMovable() {
	for idx := range m.landMovable {
		objs := m.GetObjects(VertexFromIndex(m.Size, idx))
		movable := FNone
		for _, obj := range objs {
			movable |= obj.FlagSet
		}
		m.landMovable[idx] = (movable & FLandObject) != FLandObject
	}
}

func (m *Map) IsMovable(loc Vertex) bool {
	if loc.IsOutside(m.Size) {
		return false
	}
	idx := loc.ToIndex(m.Size)
	return m.landMovable[idx]
}
