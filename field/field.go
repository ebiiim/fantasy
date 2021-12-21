package field

import (
	"github.com/ebiiim/fantasy/base"
	"github.com/ebiiim/fantasy/flag"
)

type Field struct {
	Map         *base.Map
	landMovable []bool
}

func NewField(m *base.Map) *Field {
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
		objs := f.Map.GetObjects(base.VertexFromIndex(f.Map.Dimension, idx))
		movable := flag.None
		for _, obj := range objs {
			movable |= flag.Get(obj)
		}
		f.landMovable[idx] = (movable & flag.BlockLand) != flag.BlockLand
	}
}

func (f *Field) IsMovable(loc base.Vertex) bool {
	if loc.IsOutside(f.Map.Dimension) {
		return false
	}
	idx := loc.ToIndex(f.Map.Dimension)
	return f.landMovable[idx]
}
