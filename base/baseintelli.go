package base

import (
	"fmt"
	"sync"

	"github.com/ebiiim/fantasy/log"
)

type Intelligent interface {
	Born(self Intelligent, f *Field, loc Vertex)
	Die(self Intelligent)
	ToFieldCh() chan Action
	FromFieldCh() chan Action

	Locatable
}

type BornFunc func(self0 Intelligent)
type DieFunc func(self0 Intelligent)
type ActionFunc func(self0 Intelligent)

type BaseIntelligent struct {
	*BaseObject

	toFieldCh   chan Action
	fromFieldCh chan Action
	field       *Field

	bornFunc BornFunc
	dieFunc  DieFunc
	actFunc  ActionFunc

	onceBorn sync.Once
	onceDie  sync.Once
}

func NewIntelligent(obj *BaseObject, bornFunc BornFunc, dieFunc DieFunc, actFunc ActionFunc) *BaseIntelligent {
	x := BaseIntelligent{
		BaseObject:  obj,
		toFieldCh:   make(chan Action),
		fromFieldCh: make(chan Action),

		bornFunc: NopBornFunc,
		dieFunc:  NopDieFunc,
		actFunc:  NopActionFunc,
	}
	if bornFunc != nil {
		x.bornFunc = bornFunc
	}
	if dieFunc != nil {
		x.dieFunc = dieFunc
	}
	if actFunc != nil {
		x.actFunc = actFunc
	}
	return &x
}

func (x *BaseIntelligent) ToFieldCh() chan Action { return x.toFieldCh }

func (x *BaseIntelligent) FromFieldCh() chan Action { return x.fromFieldCh }

func (x *BaseIntelligent) Born(self Intelligent, f *Field, loc Vertex) {
	x.onceBorn.Do(
		func() {
			x.field = f
			x.SetLoc(loc)
			lg.Debug(log.TypeIntelligent, "BaseIntelligent.Born", fmt.Sprintf("ObjectType %v, Loc %v", self.ObjectType(), self.Loc()))
			x.bornFunc(self)
			go func() {
				for {
					x.actFunc(self)
				}
			}()
		})
}

func (x *BaseIntelligent) Die(self Intelligent) {
	x.onceDie.Do(func() {
		lg.Debug(log.TypeIntelligent, "BaseIntelligent.Die", fmt.Sprintf("ObjectType %v, Loc %v", self.ObjectType(), self.Loc()))
		x.dieFunc(self)

		// FIXME: super slow
		// close(x.toFieldCh)
		// close(x.fromFieldCh)
	})
}

var NopBornFunc = func(self Intelligent) {}

var NopDieFunc = func(self Intelligent) {}

var NopActionFunc = func(self Intelligent) {
	for {
		_, ok := <-self.FromFieldCh()
		if !ok {
			// died
			return
		}
		// do nothing
	}
}

type NopIntelligent struct{ *BaseIntelligent }

func NewNopIntelligent() *NopIntelligent {
	return &NopIntelligent{BaseIntelligent: NewIntelligent(NewObject(ObjNone, NewVertex(-1, -1)), nil, nil, nil)}
}
