package base

import (
	"github.com/ebiiim/fantasy/log"
)

var lg = log.NewLogger("Field")

type Field struct {
	Intelligents []Intelligent
	Map          *Map
	landMovable  []bool
}

func NewField(m *Map) *Field {
	f := Field{Map: m}
	f.landMovable = make([]bool, f.Map.Dimension.X*f.Map.Dimension.Y)

	f.updateLandMovableAll()
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
				f.updateLandMovable(oldLoc) // set default land movable for now
				f.landMovable[newLoc.ToIndex(f.Map.Dimension)] = false
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
