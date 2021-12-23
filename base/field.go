package base

type Field struct {
	Intelligents []Intelligent
	Map          *Map
	landMovable  []bool
}

func NewField(m *Map) *Field {
	f := Field{Map: m}
	f.landMovable = make([]bool, f.Map.Dimension.X*f.Map.Dimension.Y)

	f.updateLandMovable()
	return &f
}

func (f *Field) AddIntelligent(i Intelligent) {
	f.Intelligents = append(f.Intelligents, i)
	go i.Born(f)
}

func (f *Field) Update() error {
	for _, intelli := range f.Intelligents {
		select {
		default:
			// do nothing
		case act := <-intelli.RecvCh():
			switch act.Type {
			case ActMove:
				if f.IsMovable(act.MoveLoc) {
					// TODO: might block for now
					intelli.SendCh() <- Action{
						Type:     ActMoved,
						MovedLoc: act.MoveLoc,
					}
				}
			}
		}
	}
	return nil
}

func (f *Field) updateLandMovable() {
	for idx := range f.landMovable {
		objs := f.Map.GetObjects(VertexFromIndex(f.Map.Dimension, idx))
		fs := None
		for _, obj := range objs {
			fs |= obj.Flag
		}
		f.landMovable[idx] = fs.Has(TerrainLand) && fs.Excepts(IsBlockingObject)
	}
}

func (f *Field) IsMovable(loc Vertex) bool {
	if loc.IsOutside(f.Map.Dimension) {
		return false
	}
	idx := loc.ToIndex(f.Map.Dimension)
	return f.landMovable[idx]
}
