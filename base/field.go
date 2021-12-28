package base

import (
	"github.com/ebiiim/fantasy/log"
)

var lg = log.NewLogger("Field")

type Field struct {
	Map         *Map
	landMovable []bool
}

func NewField(m *Map) *Field {
	f := Field{Map: m}
	f.landMovable = make([]bool, f.Map.Dimension.X*f.Map.Dimension.Y)

	ints := make([]Locatable, m.Dimension.X*m.Dimension.Y)
	for idx := range ints {
		i := NewNopIntelligent(NewObject(ObjNone, VertexFromIndex(m.Dimension, idx)))
		go i.Born(&f)
		ints[idx] = i
	}
	f.Map.Layers = append(f.Map.Layers, NewLayer("intelligents", m.Dimension, ints))

	f.updateLandMovableAll()
	return &f
}

func (f *Field) layerIntelligents() *Layer {
	return f.Map.Layers[len(f.Map.Layers)-1]
}

func (f *Field) MoveIntelligent(from, to Vertex) {
	if from.IsOutside(f.Map.Dimension) || to.IsOutside(f.Map.Dimension) {
		lg.Error(log.TypeValidation, "field", "move intelligent to wrong location")
		return
	}
	x1 := f.layerIntelligents().Objects[from.ToIndex(f.Map.Dimension)]
	x2 := f.layerIntelligents().Objects[to.ToIndex(f.Map.Dimension)]
	x1.SetLoc(to)
	x2.SetLoc(from)
	f.layerIntelligents().Objects[to.ToIndex(f.Map.Dimension)] = x1
	f.layerIntelligents().Objects[from.ToIndex(f.Map.Dimension)] = x2
	f.updateLandMovable(from)
	f.updateLandMovable(to)
}

func (f *Field) ReplaceIntelligent(i Intelligent, loc Vertex) {
	i.SetLoc(loc)
	f.layerIntelligents().Objects[loc.ToIndex(f.Map.Dimension)] = i
	f.updateLandMovable(loc)
	go i.Born(f)
}

func (f *Field) Update() error {
	for _, i := range f.layerIntelligents().Objects {
		// TODO: might use generics
		intelli := i.(Intelligent)
		select {
		default:
			// do nothing
		case act := <-intelli.RecvCh():
			switch act.Type {
			case ActMove:
				oldLoc := intelli.Loc()
				newLoc := oldLoc.Add(act.MoveAmount)
				// only move {+-1,0} or {0,+-1} for now
				if newLoc.L1Norm(oldLoc) > 1 {
					lg.Info(log.TypeValidation, "sheep", "wrong move norm")
					continue
				}
				if !f.IsMovable(newLoc) {
					continue
				}
				f.MoveIntelligent(oldLoc, newLoc)
				// TODO: might block for now
				intelli.SendCh() <- Action{
					Type:     ActMoved,
					MovedLoc: newLoc,
				}
			}
		}
	}
	return nil
}

func (f *Field) updateLandMovableAll() {
	for idx := range f.landMovable {
		f.landMovable[idx] = f.calcLandMovable(idx)
	}
}

func (f *Field) updateLandMovable(loc Vertex) {
	idx := loc.ToIndex(f.Map.Dimension)
	f.landMovable[idx] = f.calcLandMovable(idx)
}

func (f *Field) calcLandMovable(idx int) bool {
	objs := f.Map.GetObjects(VertexFromIndex(f.Map.Dimension, idx))
	fs := None
	for _, obj := range objs {
		fs |= obj.Flag()
	}
	return fs.Has(TerrainLand) && fs.Excepts(IsBlockingObject)
}

func (f *Field) IsMovable(loc Vertex) bool {
	if loc.IsOutside(f.Map.Dimension) {
		return false
	}
	idx := loc.ToIndex(f.Map.Dimension)
	return f.landMovable[idx]
}
