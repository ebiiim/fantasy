package base

import (
	"math/rand"
	"time"

	"github.com/ebiiim/fantasy/log"
)

type Me struct{ *BaseIntelligent }

func NewMe() *Me {
	return &Me{BaseIntelligent: NewIntelligent(NewObject(ObjMe, NewVertex(-1, -1)), MeBornFunc, MeDieFunc, MeActFunc)}
}

var MeBornFunc = NopBornFunc
var MeDieFunc = NopDieFunc
var MeActFunc = NopActionFunc

type Sheep struct {
	*BaseIntelligent
}

func NewSheep() *Sheep {
	return &Sheep{BaseIntelligent: NewIntelligent(NewObject(ObjSheep, NewVertex(-1, -1)), SheepBornFunc, SheepDieFunc, SheepActFunc)}
}

var SheepBornFunc = NopBornFunc
var SheepDieFunc = NopDieFunc
var SheepActFunc = func(self0 Intelligent) {
	self := self0.(*Sheep)
	select {
	case <-time.After(time.Millisecond * time.Duration(500+rand.Intn(2000))):
		lg.Debug(log.TypeIntelligent, "SheepActFunc", "try to move")
		axis := rand.Intn(10)     // X:Y=7:3
		value := rand.Intn(3) - 1 // -1,0,1
		var moveAmount Vertex
		if axis < 3 {
			moveAmount = NewVertex(value, 0)
		} else {
			moveAmount = NewVertex(0, value)
		}
		act := Action{
			Type:       ActMove,
			MoveAmount: moveAmount,
		}
		self.ToFieldCh() <- act
	case act, ok := <-self.FromFieldCh():
		if !ok { // died
			return
		}
		switch act.Type {
		case ActMoved:
			lg.Debug(log.TypeIntelligent, "SheepActFunc", "baa")
		}
	}
}
