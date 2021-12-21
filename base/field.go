package base

type Field struct {
	Map         *Map
	landMovable []bool
}

func NewField(m *Map) *Field {
	f := Field{Map: m}
	f.landMovable = make([]bool, f.Map.Dimension.X*f.Map.Dimension.Y)
	return &f
}

func (f *Field) Update() error {
	f.updateLandMovable()
	return nil
}

func (f *Field) updateLandMovable() {
	for idx := range f.landMovable {
		objs := f.Map.GetObjects(VertexFromIndex(f.Map.Dimension, idx))
		movable := FNone
		for _, obj := range objs {
			movable |= obj.FlagSet
		}
		f.landMovable[idx] = (movable & FLandObject) != FLandObject
	}
}

func (f *Field) IsMovable(loc Vertex) bool {
	if loc.IsOutside(f.Map.Dimension) {
		return false
	}
	idx := loc.ToIndex(f.Map.Dimension)
	return f.landMovable[idx]
}
